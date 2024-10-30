package services

import (
	"fkedweb/models"
	"fkedweb/repositories"
)

type TodoService interface {
	Create(todo *models.Todo) error
	GetAll() ([]models.Todo, error)
	GetByID(id uint) (*models.Todo, error)
	Update(todo *models.Todo) error
	Delete(id uint) error
}

type todoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{repo}
}

func (s *todoService) Create(todo *models.Todo) error {
	return s.repo.Create(todo)
}

func (s *todoService) GetAll() ([]models.Todo, error) {
	return s.repo.GetAll()
}

func (s *todoService) GetByID(id uint) (*models.Todo, error) {
	return s.repo.GetByID(id)
}

func (s *todoService) Update(todo *models.Todo) error {
	return s.repo.Update(todo)
}

func (s *todoService) Delete(id uint) error {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(todo)
}
