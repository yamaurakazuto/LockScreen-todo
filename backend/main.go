package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"lockscreen-todo/backend/handlers"
	"lockscreen-todo/backend/storage"
)

func main() {
	repo := storage.NewMemoryTodoRepository()
	handler := handlers.NewTodoHandler(repo)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/todos", handler.ListTodos)
		v1.POST("/todos", handler.CreateTodo)
		v1.PUT("/todos/:id", handler.UpdateTodo)
		v1.DELETE("/todos/:id", handler.DeleteTodo)
		v1.PATCH("/todos/:id/toggle", handler.ToggleTodo)
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	_ = router.Run(":8080")
}
