package dto

import "time"

type TaskDto struct {
	ID          string    `json:"id"`
	TaskListID  string    `json:"taskListId"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	IsCompleted bool      `json:"isCompleted"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateTaskDto struct {
	TaskListID  string  `json:"taskListId" validate:"required,uuid4"`
	Title       string  `json:"title" validate:"required,min=3,max=255"`
	Description *string `json:"description" validate:"omitempty,max=255"`
}

type UpdateTaskDto struct {
	Title       string  `json:"title" validate:"required,min=3,max=255"`
	Description *string `json:"description" validate:"omitempty,max=255"`
	IsCompleted bool    `json:"isCompleted" validate:"required,boolean"`
}
