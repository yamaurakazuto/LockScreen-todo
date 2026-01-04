package storage

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"

	"lockscreen-todo/backend/models"
)

var ErrTodoNotFound = errors.New("todo not found")

type TodoRepository interface {
	List() ([]models.Todo, error)
	Create(todo models.Todo) (models.Todo, error)
	Update(id string, updates models.Todo) (models.Todo, error)
	Delete(id string) error
	Toggle(id string) (models.Todo, error)
}

type MemoryTodoRepository struct {
	mu    sync.RWMutex
	todos map[string]models.Todo
}

func NewMemoryTodoRepository() *MemoryTodoRepository {
	return &MemoryTodoRepository{
		todos: make(map[string]models.Todo),
	}
}

func (repo *MemoryTodoRepository) List() ([]models.Todo, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	items := make([]models.Todo, 0, len(repo.todos))
	for _, todo := range repo.todos {
		items = append(items, todo)
	}

	return items, nil
}

func (repo *MemoryTodoRepository) Create(todo models.Todo) (models.Todo, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	id := uuid.NewString()
	now := time.Now()
	todo.ID = id
	todo.CreatedAt = now
	todo.UpdatedAt = now

	repo.todos[id] = todo
	return todo, nil
}

func (repo *MemoryTodoRepository) Update(id string, updates models.Todo) (models.Todo, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	todo, ok := repo.todos[id]
	if !ok {
		return models.Todo{}, ErrTodoNotFound
	}

	todo.Title = updates.Title
	todo.Description = updates.Description
	todo.Priority = updates.Priority
	todo.Completed = updates.Completed
	todo.UpdatedAt = time.Now()

	repo.todos[id] = todo
	return todo, nil
}

func (repo *MemoryTodoRepository) Delete(id string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, ok := repo.todos[id]; !ok {
		return ErrTodoNotFound
	}

	delete(repo.todos, id)
	return nil
}

func (repo *MemoryTodoRepository) Toggle(id string) (models.Todo, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	todo, ok := repo.todos[id]
	if !ok {
		return models.Todo{}, ErrTodoNotFound
	}

	todo.Completed = !todo.Completed
	todo.UpdatedAt = time.Now()
	repo.todos[id] = todo

	return todo, nil
}
