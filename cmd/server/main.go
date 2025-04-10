package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/szygendaborys/social-media-agent/internal/todo"
)

func init() {
	fmt.Println("Initializing server...")

}

func main() {
	fmt.Println("Starting server...")

	server := gin.Default()

	server.GET("/todos", todo.GetTodosRoute)
	server.POST("/todos", todo.CreateTodoRoute)

	server.Run(":8080")
}
