package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/szygendaborys/social-media-agent/internal/ai"
	"github.com/szygendaborys/social-media-agent/internal/todo"
)

func init() {
	fmt.Println("Initializing server...")

	// Load .env file
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	fmt.Println("Starting server...")

	router := gin.Default()
	v1 := router.Group("/api/v1")

	initTodosRoutes(v1)
	initAiRoutes(v1)

	router.Run(":8080")
}

func initTodosRoutes(r *gin.RouterGroup) {
	r.GET("/todos", todo.GetTodosRoute)
	r.POST("/todos", todo.CreateTodoRoute)
}

func initAiRoutes(r *gin.RouterGroup) {
	r.POST("/ai/send-chat-request", ai.SendChatRequestRoute)
}
