package repository

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"todolist-api/domain/entity"
)

type TodosRepository interface {
	CreateTodoItem(tx *gorm.DB, value *entity.Todos) error
	FindAllTodoItem(tx *gorm.DB) ([]*entity.Todos, error)
	FindTodoItemById(tx *gorm.DB, id int) (*entity.Todos, error)
	UpdateTodoItem(tx *gorm.DB, value *entity.Todos, id int) error
	DeleteTodoItem(tx *gorm.DB, value *entity.Todos, id int) error
}

type todoRepositoryImpl struct {
}

func NewTodosRepository() TodosRepository {
	return &todoRepositoryImpl{}
}

func (t todoRepositoryImpl) CreateTodoItem(tx *gorm.DB, value *entity.Todos) error {
	create := tx.Create(value)
	if create.Error != nil {
		log.Println(fmt.Sprintf("Error when create activity group: %v", create.Error))
		return create.Error
	}
	return nil
}

func (t todoRepositoryImpl) FindAllTodoItem(tx *gorm.DB) ([]*entity.Todos, error) {
	var todos []*entity.Todos
	find := tx.Find(&todos)
	if find.Error != nil {
		log.Println(fmt.Sprintf("Error when find all activity group: %v", find.Error))
		return nil, find.Error
	}
	return todos, nil
}

func (t todoRepositoryImpl) FindTodoItemById(tx *gorm.DB, id int) (*entity.Todos, error) {
	var todo entity.Todos
	find := tx.Take(&todo)
	if find.Error != nil {
		log.Println(fmt.Sprintf("Error when find activity group by id: %v", find.Error))
		return nil, find.Error
	}

	return &todo, nil
}

func (t todoRepositoryImpl) UpdateTodoItem(tx *gorm.DB, value *entity.Todos, id int) error {
	update := tx.Model(value).Where("id = ?", id).Updates(value)
	if update.Error != nil {
		log.Println(fmt.Sprintf("Error when update activity group: %v", update.Error))
		return update.Error
	}
	return nil
}

func (t todoRepositoryImpl) DeleteTodoItem(tx *gorm.DB, value *entity.Todos, id int) error {
	result := tx.Where("id = ?", id).Delete(value)
	if result.Error != nil {
		log.Println(fmt.Sprintf("Error when delete activity group: %v", result.Error))
		return result.Error
	}
	return nil
}
