package request

import "todolist-api/domain/entity"

type CreateTodoRequest struct {
	ActivityGroupById uint   `json:"activity_group_id" validate:"required"`
	Title             string `json:"title" validate:"required"`
	Priority          string `json:"priority"`
	IsActive          bool   `json:"is_active"`
}

func (t CreateTodoRequest) ToEntity() *entity.Todos {
	return &entity.Todos{
		ActivityGroupId: t.ActivityGroupById,
		Title:           t.Title,
		Priority:        t.Priority,
		IsActive:        t.IsActive,
	}
}
