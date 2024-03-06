package usecase

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
	"strings"
	"todolist-api/delivery/http/dto/request"
	"todolist-api/delivery/http/dto/response"
	"todolist-api/domain/entity"
	"todolist-api/domain/repository"
	"todolist-api/shared/util"
)

type TodosUsecase interface {
	CreateTodoItem(ctx context.Context, requestData *request.CreateTodoRequest) (*response.CreateTodoResponse, error)
	FindAllTodoItem(ctx context.Context, id int) ([]*response.CreateTodoResponse, error)
	FindTodoItemById(ctx context.Context, id int) (*response.CreateTodoResponse, error)
	UpdateTodoItem(ctx context.Context, id int, requestData *request.CreateTodoRequest) (*response.CreateTodoResponse, error)
	DeleteTodoItem(ctx context.Context, id int) error
}

type todosUsecaseImpl struct {
	DB                   *gorm.DB
	Validate             *validator.Validate
	TodosRepository      repository.TodosRepository
	ActivitiesRepository repository.ActivitiesRepository
}

func NewTodosUsecase(DB *gorm.DB, validate *validator.Validate, todosRepository repository.TodosRepository, activitiesRepository repository.ActivitiesRepository) TodosUsecase {
	return &todosUsecaseImpl{DB: DB, Validate: validate, TodosRepository: todosRepository, ActivitiesRepository: activitiesRepository}
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
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, e := range validationErrors {
			log.Printf("Field: %s, Tag: %s, Param: %s", e.Field(), e.Tag(), e.Param())
			switch {
			case e.Field() == "ActivityGroupById" && e.Tag() == "required":
				errorMessages = append(errorMessages, util.ActivityGroupId.Error())
			case e.Field() == "Title" && e.Tag() == "required":
				errorMessages = append(errorMessages, util.TitleIsRequired.Error())
			}
		}
		return nil, errors.New(strings.Join(errorMessages, "; "))
	}

	activityGroup, err := t.ActivitiesRepository.FindActivityGroupById(tx, int(requestData.ActivityGroupById))
	if err != nil || activityGroup == nil {
		errorMessages := []string{"Activity Group ID Not Found"}
		return nil, errors.New(strings.Join(errorMessages, "; "))
	}

	requestDataEntity := requestData.ToEntity()

	err = t.TodosRepository.CreateTodoItem(tx, requestDataEntity)
	if err != nil {
		tx.Rollback()
		if strings.Contains(err.Error(), "Error 1452") {
			return nil, errors.New("ActivityGroupId Not Found")
		}
		log.Printf("Error when create todo item: %v", err)
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return requestDataEntity.ToResponse(), nil
}

func (t todosUsecaseImpl) FindAllTodoItem(ctx context.Context, id int) ([]*response.CreateTodoResponse, error) {
	tx := t.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	todos, err := t.TodosRepository.FindAllTodoItem(tx, id)
	if err != nil {
		tx.Rollback()
		log.Printf("Error when find all todo items: %v", err)
		return nil, errors.New("Failed to retrieve all todo items")
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, errors.New("Transaction commit failed")
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
		if err == gorm.ErrRecordNotFound {
			return nil, util.TodoIdNotFound
		}
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

	err := t.Validate.Struct(requestData)
	if err != nil {
		log.Printf("Invalid Request Body: %v", err)
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, e := range validationErrors {
			log.Printf("Field: %s, Tag: %s, Param: %s", e.Field(), e.Tag(), e.Param())
			switch {
			case e.Field() == "Title" && e.Tag() == "required":
				errorMessages = append(errorMessages, "Title is required")
			case e.Field() == "ActivityGroupById" && e.Tag() == "required":
				errorMessages = append(errorMessages, "Activity Group ID is required")
			}
		}
		tx.Rollback()
		return nil, errors.New(strings.Join(errorMessages, "; "))
	}

	activityGroup, err := t.ActivitiesRepository.FindActivityGroupById(tx, int(requestData.ActivityGroupById))
	if err != nil || activityGroup == nil {
		errorMessages := []string{"Activity Group ID Not Found"}
		return nil, errors.New(strings.Join(errorMessages, "; "))
	}

	todo, err := t.TodosRepository.FindTodoItemById(tx, id)
	if err != nil || todo == nil {
		tx.Rollback()
		log.Printf("Error when find todo item by id: %v", err)
		return nil, util.TodoIdNotFound
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

	todo, err := t.TodosRepository.FindTodoItemById(tx, id)
	if err != nil || todo == nil {
		tx.Rollback()
		log.Printf("Error when find todo item by id: %v", err)
		return util.TodoIdNotFound
	}

	err = t.TodosRepository.DeleteTodoItem(tx, &entity.Todos{}, id)
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
