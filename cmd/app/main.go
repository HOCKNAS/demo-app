package main

import (
	"context"
	"fmt"
	"os"

	"github.com/HOCKNAS/demo-app/internal/app"
	"github.com/HOCKNAS/demo-app/internal/core/domain"
	firebaseauth "github.com/HOCKNAS/demo-app/pkg/auth/firebase_auth"
	"github.com/HOCKNAS/demo-app/pkg/database/mongodb"
	"github.com/sirupsen/logrus"
)

const uri = "mongodb://172.16.1.100:32120"

func main() {
	// Mongo
	mongoClient, err := mongodb.NewClient(uri, "", "")
	if err != nil {
		fmt.Println(err)
	}
	db := mongoClient.Database("demo-app")

	//firebase
	firebaseApp, err := firebaseauth.NewApp(firebaseConfigFile())
	if err != nil {
		fmt.Println(err)
	}
	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	//logger
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
			"initialFields": {"app": "demo-app"},
			"encoderConfig": {
			  "messageKey": "message",
			  "levelKey": "level",
			  "levelEncoder": "lowercase"
			}
		  }`

	config := &app.LoggerConfig{
		UseLogger:    "zap",
		ZapConfig:    configzap,
		LogrusConfig: configlogrus,
	}

	identity_provider := app.NewIdentityProviders(authClient)

	repositories := app.NewRepositories(db)

	logger := app.NewLoggers(*config)

	services := app.NewServices(app.Deps{
		Repos:            repositories,
		IdentityProvider: identity_provider,
		Logger:           logger,
	})

	fmt.Println(banner())

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
