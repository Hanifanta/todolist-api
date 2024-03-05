package response

import "time"

type CreateTodoResponse struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	ActivityGroupId int       `json:"activity_group_id"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
