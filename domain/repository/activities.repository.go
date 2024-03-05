package repository

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"todolist-api/domain/entity"
)

type ActivitiesRepository interface {
	CreateActivityGroup(tx *gorm.DB, value *entity.Activity) error
	FindAllActivityGroup(tx *gorm.DB) ([]*entity.Activity, error)
	FindActivityGroupById(tx *gorm.DB, id int) (*entity.Activity, error)
	UpdateActivityGroup(tx *gorm.DB, value *entity.Activity, id int) error
	DeleteActivityGroup(tx *gorm.DB, id int) error
}

type activitiesRepositoryImpl struct {
}

func NewActivitiesRepository() ActivitiesRepository {
	return &activitiesRepositoryImpl{}
}

func (a activitiesRepositoryImpl) CreateActivityGroup(tx *gorm.DB, value *entity.Activity) error {
	create := tx.Create(value)
	if create.Error != nil {
		log.Println(fmt.Sprintf("Error when create activity group: %v", create.Error))
		return create.Error
	}
	return nil
}

func (a activitiesRepositoryImpl) FindAllActivityGroup(tx *gorm.DB) ([]*entity.Activity, error) {
	var activities []*entity.Activity
	find := tx.Find(&activities)
	if find.Error != nil {
		log.Println(fmt.Sprintf("Error when find all activity group: %v", find.Error))
		return nil, find.Error
	}
	return activities, nil
}

func (a activitiesRepositoryImpl) FindActivityGroupById(tx *gorm.DB, id int) (*entity.Activity, error) {
	var activity entity.Activity
	find := tx.Take(&activity)
	if find.Error != nil {
		log.Println(fmt.Sprintf("Error when find activity group by id: %v", find.Error))
		return nil, find.Error
	}
	return &activity, nil
}

func (a activitiesRepositoryImpl) UpdateActivityGroup(tx *gorm.DB, value *entity.Activity, id int) error {
	update := tx.Model(value).Where("id = ?", id).Updates(value)
	if update.Error != nil {
		log.Println(fmt.Sprintf("Error when update activity group: %v", update.Error))
		return update.Error
	}
	return nil
}

func (a activitiesRepositoryImpl) DeleteActivityGroup(tx *gorm.DB, id int) error {
	deletes := tx.Delete(&entity.Activity{}, id)
	if deletes.Error != nil {
		log.Println(fmt.Sprintf("Error when delete activity group: %v", deletes.Error))
		return deletes.Error
	}
	return nil
}
