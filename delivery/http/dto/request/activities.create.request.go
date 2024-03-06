package request

import "todolist-api/domain/entity"

type CreateActivityGroupRequest struct {
	Title string `json:"title" validate:"required"`
	Email string `json:"email"`
}

func (a CreateActivityGroupRequest) ToEntity() *entity.Activity {
	return &entity.Activity{
		Title: a.Title,
		Email: a.Email,
	}
}
