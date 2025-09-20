package main

import (
	"fmt"
	"main/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", handler.Hello)

	fmt.Println("Server is running on :8080...")
	r.Run()
}
