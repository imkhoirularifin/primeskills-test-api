package interfaces

import (
	"primeskills-test-api/internal/dto"
	"primeskills-test-api/internal/entity"

	"github.com/gin-gonic/gin"
)

type TaskListRepository interface {
	Create(taskList *entity.TaskList) error
	FindById(id string) (*entity.TaskList, error)
	FindByUserId(userId string) (*[]entity.TaskList, error)
	Update(taskList *entity.TaskList) error
	Delete(taskList *entity.TaskList) error
}

type TaskListService interface {
	Create(ctx *gin.Context, req *dto.CreateTaskListDto) (*dto.TaskListDto, error)
	FindByUserId(ctx *gin.Context) (*[]dto.TaskListDto, error)
	Update(id string, req *dto.UpdateTaskListDto) error
	Delete(id string) error
}
