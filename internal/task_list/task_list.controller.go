package tasklist

import (
	"net/http"
	"primeskills-test-api/internal/domain/dto"
	"primeskills-test-api/internal/domain/interfaces"
	"primeskills-test-api/internal/middleware"
	"primeskills-test-api/internal/utilities"

	"github.com/gin-gonic/gin"
)

type controller struct {
	taskListService interfaces.TaskListService
}

func NewController(router *gin.RouterGroup, taskListService interfaces.TaskListService) {
	controller := &controller{
		taskListService: taskListService,
	}

	protected := router.Group("/", middleware.RequireToken())

	protected.POST("/", middleware.Validate[dto.CreateTaskListDto](), controller.create)
	protected.GET("/", controller.findByUserId)
	protected.PUT("/:id", middleware.Validate[dto.UpdateTaskListDto](), controller.update)
	protected.DELETE("/:id", controller.delete)
}

// Create godoc
//
//	@Summary		Create task list
//	@Description	Create task list
//	@Tags			task-list
//	@Security		Bearer
//	@Param			body	body		dto.CreateTaskListDto	true	"Create task list"
//	@Success		200		{object}	dto.TaskListDto
//	@Router			/task-lists [post]
func (c *controller) create(ctx *gin.Context) {
	req := utilities.ExtractStructFromValidator[dto.CreateTaskListDto](ctx)

	taskListDto, err := c.taskListService.Create(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.ResponseDto{
		Message: "Successfully create task list",
		Data:    taskListDto,
	})
}

// Find by user id godoc
//
//	@Summary		Find my task lists
//	@Description	Find my task lists
//	@Tags			task-list
//	@Security		Bearer
//	@Success		200	{array}	dto.TaskListDto
//	@Router			/task-lists [get]
func (c *controller) findByUserId(ctx *gin.Context) {
	taskLists, err := c.taskListService.FindByUserId(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.ResponseDto{
		Message: "Successfully get task lists",
		Data:    taskLists,
	})
}

// Update godoc
//
//	@Summary		Update task list
//	@Description	Update task list
//	@Tags			task-list
//	@Security		Bearer
//	@Param			id		path		string					true	"Task list ID"
//	@Param			body	body		dto.UpdateTaskListDto	true	"Update task list"
//	@Success		200		{object}	dto.ResponseDto
//	@Router			/task-lists/{id} [put]
func (c *controller) update(ctx *gin.Context) {
	id := ctx.Param("id")
	req := utilities.ExtractStructFromValidator[dto.UpdateTaskListDto](ctx)

	err := c.taskListService.Update(id, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.ResponseDto{
		Message: "Successfully update task list",
	})
}

// Delete godoc
//
//	@Summary		Delete task list
//	@Description	Delete task list
//	@Tags			task-list
//	@Security		Bearer
//	@Param			id	path		string	true	"Task list ID"
//	@Success		200	{object}	dto.ResponseDto
//	@Router			/task-lists/{id} [delete]
func (c *controller) delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.taskListService.Delete(id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.ResponseDto{
		Message: "Successfully delete task list",
	})
}
