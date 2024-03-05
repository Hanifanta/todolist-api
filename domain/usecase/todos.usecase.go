package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
	"todolist-api/delivery/http/dto/request"
	"todolist-api/delivery/http/dto/response"
	"todolist-api/domain/entity"
	"todolist-api/domain/repository"
)

type TodosUsecase interface {
	CreateTodoItem(ctx context.Context, requestData *request.CreateTodoRequest) (*response.CreateTodoResponse, error)
	FindAllTodoItem(ctx context.Context) ([]*response.CreateTodoResponse, error)
	FindTodoItemById(ctx context.Context, id int) (*response.CreateTodoResponse, error)
	UpdateTodoItem(ctx context.Context, id int, requestData *request.CreateTodoRequest) (*response.CreateTodoResponse, error)
	DeleteTodoItem(ctx context.Context, id int) error
}

type todosUsecaseImpl struct {
	DB              *gorm.DB
	Validate        *validator.Validate
	TodosRepository repository.TodosRepository
}

func NewTodosUsecase(DB *gorm.DB, validate *validator.Validate, todosRepository repository.TodosRepository) TodosUsecase {
	return &todosUsecaseImpl{DB: DB, Validate: validate, TodosRepository: todosRepository}
}

func (t todosUsecaseImpl) CreateTodoItem(ctx context.Context, requestData *request.CreateTodoRequest) (*response.CreateTodoResponse, error) {
	tx := t.DB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := t.Validate.Struct(requestData)
	if err != nil {
		log.Printf("Invalid Request Body: %v", err)
		return nil, err
	}

	requestDataEntity := requestData.ToEntity()

	err = t.TodosRepository.CreateTodoItem(tx, requestDataEntity)
	if err != nil {
		tx.Rollback()
		log.Printf("Error when create todo item: %v", err)
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return requestDataEntity.ToResponse(), nil
}

func (t todosUsecaseImpl) FindAllTodoItem(ctx context.Context) ([]*response.CreateTodoResponse, error) {
	tx := t.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	todos, err := t.TodosRepository.FindAllTodoItem(tx)
	if err != nil {
		tx.Rollback()
		log.Printf("Error when find all todo items: %v", err)
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	var responses []*response.CreateTodoResponse
	for _, todo := range todos {
		responses = append(responses, todo.ToResponse())
	}
	return responses, nil
}

func (t todosUsecaseImpl) FindTodoItemById(ctx context.Context, id int) (*response.CreateTodoResponse, error) {
	tx := t.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	todo, err := t.TodosRepository.FindTodoItemById(tx, id)
	if err != nil {
		tx.Rollback()
		log.Printf("Error when find todo item by id: %v", err)
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return todo.ToResponse(), nil
}

func (t todosUsecaseImpl) UpdateTodoItem(ctx context.Context, id int, requestData *request.CreateTodoRequest) (*response.CreateTodoResponse, error) {
	tx := t.DB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := t.Validate.Struct(requestData)
	if err != nil {
		log.Printf("Invalid Request Body: %v", err)
		return nil, err
	}

	requestDataEntity := requestData.ToEntity()

	err = t.TodosRepository.UpdateTodoItem(tx, requestDataEntity, id)
	if err != nil {
		tx.Rollback()
		log.Printf("Error when update todo item: %v", err)
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return requestDataEntity.ToResponse(), nil
}

func (t todosUsecaseImpl) DeleteTodoItem(ctx context.Context, id int) error {
	tx := t.DB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := t.TodosRepository.DeleteTodoItem(tx, &entity.Todos{}, id)
	if err != nil {
		tx.Rollback()
		log.Printf("Error when delete todo item: %v", err)
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}
