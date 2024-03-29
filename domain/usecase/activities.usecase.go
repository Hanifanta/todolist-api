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
	"todolist-api/domain/repository"
	"todolist-api/shared/util"
)

type ActivitiesUsecase interface {
	CreateActivityGroup(ctx context.Context, requestData *request.CreateActivityGroupRequest) (*response.CreateActivityResponse, error)
	FindAllActivityGroup(ctx context.Context) ([]*response.CreateActivityResponse, error)
	FindActivityGroupById(ctx context.Context, id int) (*response.CreateActivityResponse, error)
	UpdateActivityGroup(ctx context.Context, id int, requestData *request.CreateActivityGroupRequest) (*response.CreateActivityResponse, error)
	DeleteActivityGroup(ctx context.Context, id int) error
}

type activitiesUsecaseImpl struct {
	DB                   *gorm.DB
	Validate             *validator.Validate
	ActivitiesRepository repository.ActivitiesRepository
}

func NewActivitiesUsecase(DB *gorm.DB, validate *validator.Validate, activitiesRepository repository.ActivitiesRepository) ActivitiesUsecase {
	return &activitiesUsecaseImpl{DB: DB, Validate: validate, ActivitiesRepository: activitiesRepository}
}

func (a activitiesUsecaseImpl) CreateActivityGroup(ctx context.Context, requestData *request.CreateActivityGroupRequest) (*response.CreateActivityResponse, error) {
	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := a.Validate.Struct(requestData)
	if err != nil {
		log.Printf("Invalid Request Body: %v", err)
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, e := range validationErrors {
			log.Printf("Field: %s, Tag: %s, Param: %s", e.Field(), e.Tag(), e.Param())
			switch {
			case e.Field() == "Title" && e.Tag() == "required":
				errorMessages = append(errorMessages, "Title is required")
			}
		}
		return nil, errors.New(strings.Join(errorMessages, "; "))
	}

	requestDataEntity := requestData.ToEntity()

	err = a.ActivitiesRepository.CreateActivityGroup(tx, requestDataEntity)
	if err != nil {
		log.Printf("Error when create activity group: %v", err)
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return requestDataEntity.ToResponse(), nil
}

func (a activitiesUsecaseImpl) FindAllActivityGroup(ctx context.Context) ([]*response.CreateActivityResponse, error) {
	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	activities, err := a.ActivitiesRepository.FindAllActivityGroup(tx)
	if err != nil {
		tx.Rollback()
		log.Printf("Error when find all activity groups: %v", err)
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	var responses []*response.CreateActivityResponse
	for _, activity := range activities {
		responses = append(responses, activity.ToResponse())
	}
	return responses, nil
}

func (a activitiesUsecaseImpl) FindActivityGroupById(ctx context.Context, id int) (*response.CreateActivityResponse, error) {
	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	activity, err := a.ActivitiesRepository.FindActivityGroupById(tx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.ActivityIdNotFound
		}
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return activity.ToResponse(), nil
}

func (a activitiesUsecaseImpl) UpdateActivityGroup(ctx context.Context, id int, requestData *request.CreateActivityGroupRequest) (*response.CreateActivityResponse, error) {
	tx := a.DB.WithContext(ctx).Begin()

	err := a.Validate.Struct(requestData)
	if err != nil {
		log.Printf("Invalid Request Body: %v", err)
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, e := range validationErrors {
			log.Printf("Field: %s, Tag: %s, Param: %s", e.Field(), e.Tag(), e.Param())
			switch {
			case e.Field() == "Title" && e.Tag() == "required":
				errorMessages = append(errorMessages, "Title is required")
			}
		}
		tx.Rollback()
		return nil, errors.New(strings.Join(errorMessages, "; "))
	}

	activity, err := a.ActivitiesRepository.FindActivityGroupById(tx, id)
	if err != nil || activity == nil {
		tx.Rollback()
		log.Printf("Error when find activity group by id: %v", err)
		return nil, util.ActivityIdNotFound
	}

	requestDataEntity := requestData.ToEntity()

	err = a.ActivitiesRepository.UpdateActivityGroup(tx, requestDataEntity, id)
	if err != nil {
		tx.Rollback()
		log.Printf("Error when update activity group: %v", err)
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return requestDataEntity.ToResponse(), nil
}

func (a activitiesUsecaseImpl) DeleteActivityGroup(ctx context.Context, id int) error {
	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	activity, err := a.ActivitiesRepository.FindActivityGroupById(tx, id)
	if err != nil || activity == nil {
		tx.Rollback()
		log.Printf("Error when find activity group by id: %v", err)
		return util.ActivityIdNotFound
	}

	err = a.ActivitiesRepository.DeleteActivityGroup(tx, id)
	if err != nil {
		tx.Rollback()
		log.Printf("Error when delete activity group: %v", err)
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}
