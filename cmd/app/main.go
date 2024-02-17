package main

import (
	"context"
	"fmt"
	"os"

	"github.com/HOCKNAS/demo-app/internal/app"
	"github.com/HOCKNAS/demo-app/internal/core/domain"
	"github.com/HOCKNAS/demo-app/pkg/database/mongodb"
)

const uri = "mongodb://172.16.1.100:32120"

func main() {

	mongoClient, err := mongodb.NewClient(uri, "", "")
	if err != nil {
		fmt.Println(err)
	}

	db := mongoClient.Database("demo-app")

	repositories := app.NewRepositories(db)

	services := app.NewServices(app.Deps{
		Repos: repositories,
	})

	fmt.Println(banner())

	user, err := services.Users.Register(context.Background(), &domain.User{
		Name:     "Santiago",
		LastName: "Chacon",
		Username: "hocknas",
		Email:    "santiago.chacon99@gmail.com",
		Password: "Hola123@",
		IsAdmin:  false,
		IsActive: true,
	})

	if err != nil {
		fmt.Println(err)
	}

	if user != nil {
		fmt.Println(user)
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
