package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/HOCKNAS/demo-app/internal/adapters/delivery/http_delivery"
	"github.com/HOCKNAS/demo-app/internal/adapters/identity_provider"
	"github.com/HOCKNAS/demo-app/internal/adapters/repository"
	"github.com/HOCKNAS/demo-app/internal/core/use_cases"
	"github.com/HOCKNAS/demo-app/pkg/auth/firebase_auth"
	"github.com/HOCKNAS/demo-app/pkg/db/mongo_db"
	logger "github.com/HOCKNAS/demo-app/pkg/logger"
	"github.com/HOCKNAS/demo-app/pkg/server"
	"github.com/danielgtaylor/huma/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const uri = "mongodb://172.16.1.100:32120"

func Run() {

	fmt.Println(banner())

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

	services := use_cases.NewServices(use_cases.Deps{
		Repos:            repositories,
		IdentityProvider: identity_provider,
		Logs:             logger,
	})

	type Options struct {
		Host string `default:"localhost" doc:"Host to listen on"`
		Port int    `default:"8080" doc:"Port to listen on"`
	}

	var handler *http_delivery.Handler

	cli := huma.NewCLI(func(hooks huma.Hooks, opts *Options) {

		handler = http_delivery.NewHandlerHTTP(services)

		srvConfig := &http.Server{
			Addr:              fmt.Sprintf("%s:%d", opts.Host, opts.Port),
			Handler:           handler.Router,
			ReadTimeout:       5 * time.Second,
			ReadHeaderTimeout: 1 * time.Second,
			WriteTimeout:      10 * time.Second,
			IdleTimeout:       30 * time.Second,
		}

		srv := server.NewServer(srvConfig).HTTPServer

		hooks.OnStart(func() {

			fmt.Println("Starting server on http://" + srvConfig.Addr)

			err := srv.Start()
			if err != nil && err != http.ErrServerClosed {
				logger.Logger.Error("Error al iniciar el Servidor", "error", err)
			}

		})

		hooks.OnStop(func() {
			srv.Stop(context.Background())
		})
	})

	cli.Root().AddCommand(&cobra.Command{
		Use:   "openapi",
		Short: "Generate OpenAPI spec",
		Run: func(cmd *cobra.Command, args []string) {
			b, err := json.MarshalIndent(handler.GetApi().OpenAPI(), "", "  ")
			if err != nil {
				logger.Logger.Error("Error al generar la especificación OpenAPI", "error", err)
				return
			}

			fileName := "openapi-spec.json"

			file, err := os.Create(fileName)
			if err != nil {
				logger.Logger.Error("Error al crear el archivo", "error", err)
				return
			}
			defer file.Close()

			_, err = file.Write(b)
			if err != nil {
				logger.Logger.Error("Error al escribir en el archivo", "error", err)
				return
			}

			logger.Logger.Info("La especificación OpenAPI ha sido guardada", "fileName", fileName)
		},
	})

	cli.Run()

	// user, _ := services.Users.CreateUser(context.Background(), &domain.User{
	// 	Name:     "Santiago",
	// 	LastName: "Chacon",
	// 	Username: "hocknas",
	// 	Email:    "santiago.chacon99@gmail.com",
	// 	Password: "Hola123@",
	// 	IsAdmin:  false,
	// 	IsActive: true,
	// })
	// if user != nil {
	// 	fmt.Println(user.Email)
	// }

	// deactivate, _ := services.Users.DeactivateUser(context.Background(), user.ID)
	// if deactivate != nil {
	// 	fmt.Println(deactivate.IsActive)
	// }

	// eliminar := services.Users.DeleteUser(context.Background(), user.ID)
	// if eliminar != nil {
	// 	fmt.Println(eliminar)
	// }
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
