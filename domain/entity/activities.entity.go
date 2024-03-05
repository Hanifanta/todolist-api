package entity

import (
	"gorm.io/gorm"
	"todolist-api/delivery/http/dto/response"
)

type Activity struct {
	gorm.Model
	Email string `json:"email" gorm:"column:email"`
	Title string `json:"title" gorm:"column:title"`
}

func (a Activity) ToResponse() *response.CreateActivityResponse {
	return &response.CreateActivityResponse{
		ID:        int(a.ID),
		Title:     a.Title,
		Email:     a.Email,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}
