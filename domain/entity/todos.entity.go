package entity

import (
	"gorm.io/gorm"
	"todolist-api/delivery/http/dto/response"
)

type Todos struct {
	gorm.Model
	ActivityGroupId uint     `gorm:"column:activity_group_id"`
	Activity        Activity `gorm:"foreignKey:ActivityGroupId;references:ID" json:"activity_group_id"`
	Title           string   `json:"title" gorm:"column:title"`
	IsActive        bool     `json:"is_active" gorm:"column:is_active"`
	Priority        string   `json:"priority" gorm:"column:priority"`
}

func (t Todos) ToResponse() *response.CreateTodoResponse {
	return &response.CreateTodoResponse{
		ID:              int(t.ID),
		Title:           t.Title,
		ActivityGroupId: int(t.ActivityGroupId),
		IsActive:        t.IsActive,
		Priority:        t.Priority,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
	}
}
