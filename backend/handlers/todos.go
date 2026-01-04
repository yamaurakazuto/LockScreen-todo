package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"lockscreen-todo/backend/models"
	"lockscreen-todo/backend/storage"
)

type TodoHandler struct {
	repo storage.TodoRepository
}

type TodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Completed   bool   `json:"completed"`
}

func NewTodoHandler(repo storage.TodoRepository) *TodoHandler {
	return &TodoHandler{repo: repo}
}

func (handler *TodoHandler) ListTodos(c *gin.Context) {
	todos, err := handler.repo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (handler *TodoHandler) CreateTodo(c *gin.Context) {
	var req TodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	created, err := handler.repo.Create(models.Todo{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		Completed:   false,
		UserID:      "default",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (handler *TodoHandler) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var req TodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	updated, err := handler.repo.Update(id, models.Todo{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		Completed:   req.Completed,
	})
	if err != nil {
		status := http.StatusInternalServerError
		if err == storage.ErrTodoNotFound {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (handler *TodoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if err := handler.repo.Delete(id); err != nil {
		status := http.StatusInternalServerError
		if err == storage.ErrTodoNotFound {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (handler *TodoHandler) ToggleTodo(c *gin.Context) {
	id := c.Param("id")
	updated, err := handler.repo.Toggle(id)
	if err != nil {
		status := http.StatusInternalServerError
		if err == storage.ErrTodoNotFound {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}
