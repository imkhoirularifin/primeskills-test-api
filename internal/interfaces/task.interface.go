package interfaces

import (
	"primeskills-test-api/internal/dto"
	"primeskills-test-api/internal/entity"
)

type TaskRepository interface {
	Create(task *entity.Task) error
	FindById(id string) (*entity.Task, error)
	FindByTaskListIds(taskListIds []string) (*[]entity.Task, error)
	Update(task *entity.Task) error
	Delete(task *entity.Task) error
}

type TaskService interface {
	Create(req *dto.CreateTaskDto) (*dto.TaskDto, error)
	FindMapByTaskListIds(taskListIds []string) (map[string][]dto.TaskDto, error)
	Update(id string, req *dto.UpdateTaskDto) error
	Delete(id string) error
}
