package todo

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTodosRoute(c *gin.Context) {
	todos, err := GetTodos()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func CreateTodoRoute(c *gin.Context) {
	todo := Task{
		ID:        1,
		Title:     "Buy groceries" + time.Now().Format(time.RFC3339),
		Completed: false,
	}

	err := AddTodo(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}
