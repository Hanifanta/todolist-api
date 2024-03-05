package request

import "todolist-api/domain/entity"

type CreateTodoRequest struct {
	ActivityGroupById uint   `json:"activity_group_id"`
	Title             string `json:"title"`
}

func (t CreateTodoRequest) ToEntity() *entity.Todos {
	return &entity.Todos{
		ActivityGroupId: t.ActivityGroupById,
		Title:           t.Title,
	}
}
