package dto

import "time"

type TaskListDto struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Tasks     []TaskDto `json:"tasks"`
}

type CreateTaskListDto struct {
	Title string `json:"title" validate:"required,min=3,max=255"`
}

type UpdateTaskListDto struct {
	Title string `json:"title" validate:"required,min=3,max=255"`
}
