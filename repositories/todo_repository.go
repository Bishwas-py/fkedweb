package repositories

import (
	"fkedweb/models"
	"github.com/jinzhu/gorm"
)

type TodoRepository interface {
	Create(todo *models.Todo) error
	GetAll() ([]models.Todo, error)
	GetByID(id uint) (*models.Todo, error)
	Update(todo *models.Todo) error
	Delete(todo *models.Todo) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) Create(todo *models.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *todoRepository) GetByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

func (r *todoRepository) Update(todo *models.Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepository) Delete(todo *models.Todo) error {
	return r.db.Delete(todo).Error
}
