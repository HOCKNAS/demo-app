package main

import (
	"context"
	"fmt"
	"os"

	"github.com/HOCKNAS/demo-app/internal/app"
	"github.com/HOCKNAS/demo-app/internal/core/domain"
	"github.com/HOCKNAS/demo-app/pkg/auth/firebase_auth"
	"github.com/HOCKNAS/demo-app/pkg/database/mongodb"
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
	firebaseApp, err := firebase_auth.NewApp(firebaseConfigFile())
	if err != nil {
		fmt.Println(err)
	}
	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	identity_provider := app.NewIdentityProvider(authClient)

	repositories := app.NewRepositories(db)

	services := app.NewServices(app.Deps{
		Repos:            repositories,
		IdentityProvider: identity_provider,
	})

	fmt.Println(banner())

	user, err := services.Users.CreateAccount(context.Background(), &domain.User{
		Name:     "Santiago",
		LastName: "Chacon",
		Username: "hocknas",
		Email:    "santiago.chacon99@gmail.com",
		Password: "Hola123@",
		IsAdmin:  false,
		IsActive: true,
	})

	deactivate, err := services.Users.DeactivateAccount(context.Background(), user.ID)

	eliminar := services.Users.DeleteAccount(context.Background(), user.ID)

	if err != nil {
		fmt.Println(err)
	}

	if eliminar != nil {
		fmt.Println(eliminar)
	}

	if user != nil {
		fmt.Println(user.Email)
	}

	if deactivate != nil {
		fmt.Println(deactivate.IsActive)
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
