package task

import (
	"net/http"
	"primeskills-test-api/internal/dto"
	"primeskills-test-api/internal/interfaces"
	"primeskills-test-api/internal/middleware"
	"primeskills-test-api/internal/utilities"

	"github.com/gin-gonic/gin"
)

type controller struct {
	taskService interfaces.TaskService
}

func NewController(router *gin.RouterGroup, taskService interfaces.TaskService) {
	controller := &controller{
		taskService: taskService,
	}

	protected := router.Group("/", middleware.RequireToken())

	protected.POST("/", middleware.Validate[dto.CreateTaskDto](), controller.create)
	protected.PUT("/:id", middleware.Validate[dto.UpdateTaskDto](), controller.update)
	protected.DELETE("/:id", controller.delete)
}

// Create godoc
//
//	@Summary		Create task
//	@Description	Create new task
//	@Tags			task
//	@Security		Bearer
//	@Param			body	body		dto.CreateTaskDto	true	"Create task"
//	@Success		201		{object}	dto.TaskDto
//	@Router			/tasks [post]
func (c *controller) create(ctx *gin.Context) {
	req := utilities.ExtractStructFromValidator[dto.CreateTaskDto](ctx)

	task, err := c.taskService.Create(req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, &dto.ResponseDto{
		Message: "Task created successfully",
		Data:    task,
	})
}

// Update godoc
//
//	@Summary		Update task
//	@Description	Update task
//	@Tags			task
//	@Security		Bearer
//	@Param			id		path		string				true	"Task ID"
//	@Param			body	body		dto.UpdateTaskDto	true	"Update task"
//	@Success		200		{object}	dto.ResponseDto
//	@Router			/tasks/{id} [put]
func (c *controller) update(ctx *gin.Context) {
	id := ctx.Param("id")
	req := utilities.ExtractStructFromValidator[dto.UpdateTaskDto](ctx)

	err := c.taskService.Update(id, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.ResponseDto{
		Message: "Task updated successfully",
	})
}

// Delete godoc
//
//	@Summary		Delete task
//	@Description	Delete task
//	@Tags			task
//	@Security		Bearer
//	@Param			id	path		string	true	"Task ID"
//	@Success		200	{object}	dto.ResponseDto
//	@Router			/tasks/{id} [delete]
func (c *controller) delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.taskService.Delete(id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.ResponseDto{
		Message: "Task deleted successfully",
	})
}
