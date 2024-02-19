package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/HOCKNAS/demo-app/internal/adapters/identity_provider"
	"github.com/HOCKNAS/demo-app/internal/adapters/repository"
	"github.com/HOCKNAS/demo-app/internal/adapters/server"
	"github.com/HOCKNAS/demo-app/internal/core/domain"
	"github.com/HOCKNAS/demo-app/internal/core/use_cases"
	"github.com/HOCKNAS/demo-app/pkg/auth/firebase_auth"
	"github.com/HOCKNAS/demo-app/pkg/db/mongo_db"
	logger "github.com/HOCKNAS/demo-app/pkg/logger"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

const uri = "mongodb://172.16.1.100:32120"

func Run() {

	// Mongo
	mongoClient, err := mongo_db.NewClient(uri, "", "")
	if err != nil {
		fmt.Println(err)
	}
	db := mongoClient.Database("demo-app")

	// firebase
	firebaseApp, err := firebase_auth.NewApp(firebaseConfigFile())
	if err != nil {
		fmt.Println(err)
	}
	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	// logger
	configlogrus := &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		DisableQuote:    true,
		TimestampFormat: "2006-01-02T15:04:05.000000000",
	}

	configzap := `{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "/tmp/logs"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase",
		  "timeKey": "ts",
		  "callerKey": "caller"
		}
	  }`

	config := &logger.LoggerConfig{
		UseLogger:    "zap",
		ZapConfig:    configzap,
		LogrusConfig: configlogrus,
	}

	identity_provider := identity_provider.NewIdentityProvider(authClient)

	repositories := repository.NewRepositories(db)

	logger := logger.NewLogger(*config)

	type Options struct {
		Debug bool   `doc:"Enable debug logging"`
		Host  string `doc:"Hostname to listen on."`
		Port  int    `doc:"Port to listen on." short:"p" default:"8888"`
	}

	type GreetingInput struct {
		Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	}

	type GreetingOutput struct {
		Body struct {
			Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
		}
	}

	cli := huma.NewCLI(func(hooks huma.Hooks, opts *Options) {

		handler := chi.NewRouter()

		srvConfig := &http.Server{
			Addr:           ":" + "8080",
			Handler:        handler,
			ReadTimeout:    15 * time.Second,
			WriteTimeout:   15 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		api := humachi.New(handler, huma.DefaultConfig("DEMO - APP", "1.0.0"))

		huma.Register(api, huma.Operation{
			OperationID: "get-greeting",
			Summary:     "Get a greeting",
			Method:      http.MethodGet,
			Path:        "/greeting/{name}",
		}, func(ctx context.Context, input *GreetingInput) (*GreetingOutput, error) {
			resp := &GreetingOutput{}
			resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
			return resp, nil
		})

		srv := server.NewServer(srvConfig).HTTPServer

		hooks.OnStart(func() {
			if err := srv.Start(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		})

		hooks.OnStop(func() {
			srv.Stop(context.Background())
		})
	})

	services := use_cases.NewServices(use_cases.Deps{
		Repos:            repositories,
		IdentityProvider: identity_provider,
		Logs:             logger,
	})

	fmt.Println(banner())

	cli.Run()

	user, _ := services.Users.CreateUser(context.Background(), &domain.User{
		Name:     "Santiago",
		LastName: "Chacon",
		Username: "hocknas",
		Email:    "santiago.chacon99@gmail.com",
		Password: "Hola123@",
		IsAdmin:  false,
		IsActive: true,
	})
	if user != nil {
		fmt.Println(user.Email)
	}

	deactivate, _ := services.Users.DeactivateUser(context.Background(), user.ID)
	if deactivate != nil {
		fmt.Println(deactivate.IsActive)
	}

	eliminar := services.Users.DeleteUser(context.Background(), user.ID)
	if eliminar != nil {
		fmt.Println(eliminar)
	}
}

func banner() string {
	bannerPath := os.Getenv("BANNER_PATH")
	if bannerPath == "" {
		bannerPath = "resources/banner.txt"
	}

	b, err := os.ReadFile(bannerPath)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func firebaseConfigFile() string {
	firebaseConfigFilePath := os.Getenv("FIREBASE_PATH")
	if firebaseConfigFilePath == "" {
		firebaseConfigFilePath = "firebase.json"
	}
	return string(firebaseConfigFilePath)
}
