package main

import (
	"fmt"
	"main/db"
	"main/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// .env読み込み
	if err := godotenv.Load(); err != nil {
		fmt.Println("Failed to load .env")
	}

	// db接続
	if err := db.Init(); err != nil {
		panic(err)
	}

	userRepo := db.NewUserRepository(db.DB)
	userHandler := &handler.UserHandler{Repo: userRepo}

	r := gin.Default()
	r.GET("/", handler.Hello)
	r.POST("/user", userHandler.NewUser)

	fmt.Println("Server is running on :8080...")
	r.Run()
}
