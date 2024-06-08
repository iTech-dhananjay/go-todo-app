package services

import (
    "todo-app/internal/models"
    "todo-app/internal/repositories"
)

type TodoService struct {
    repo *repositories.TodoRepository
}

func NewTodoService(db *models.DB) *TodoService {
    return &TodoService{
        repo: repositories.NewTodoRepository(db),
    }
}

func (s *TodoService) Add(todo *models.Todo) error {
    return s.repo.Add(todo)
}

func (s *TodoService) GetById(id string) (*models.Todo, error) {
    return s.repo.GetById(id)
}

func (s *TodoService) GetAll() ([]models.Todo, error) {
    return s.repo.GetAll()
}

func (s *TodoService) UpdateById(id string, todo *models.Todo) error {
    return s.repo.UpdateById(id, todo)
}

func (s *TodoService) DeleteById(id string) error {
    return s.repo.DeleteById(id)
}
