package user

import (
	"net/http"
	"primeskills-test-api/internal/domain/dto"
	"primeskills-test-api/internal/domain/interfaces"
	"primeskills-test-api/internal/middleware"
	"primeskills-test-api/internal/utilities"

	"github.com/gin-gonic/gin"
)

type controller struct {
	userService interfaces.UserService
}

func NewController(router *gin.RouterGroup, userService interfaces.UserService) {
	controller := &controller{
		userService: userService,
	}

	protected := router.Group("/", middleware.RequireToken())

	protected.GET("/my-profile", controller.getMyProfile)
	protected.PUT("/", middleware.Validate[dto.UpdateUserDto](), controller.update)
	protected.PUT("/password", middleware.Validate[dto.UpdateUserPasswordDto](), controller.updatePassword)
}

// Get my profile godoc
//
//	@Summary		Get my profile
//	@Description	Get current login user profile
//	@Tags			user
//	@Security		Bearer
//	@Success		200	{object}	dto.UserProfileDto
//	@Router			/users/my-profile [get]
func (c *controller) getMyProfile(ctx *gin.Context) {
	userProfile, err := c.userService.FindUserProfile(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.ResponseDto{
		Message: "Successfully get user profile",
		Data:    userProfile,
	})
}

// Update godoc
//
//	@Summary		Update user profile
//	@Description	Update current login user profile
//	@Tags			user
//	@Security		Bearer
//	@Param			body	body		dto.UpdateUserDto	true	"Update user profile"
//	@Success		200		{object}	dto.ResponseDto
//	@Router			/users [put]
func (c *controller) update(ctx *gin.Context) {
	req := utilities.ExtractStructFromValidator[dto.UpdateUserDto](ctx)

	err := c.userService.Update(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.ResponseDto{
		Message: "Successfully update user profile",
	})
}

// Update password godoc
//
//	@Summary		Update user password
//	@Description	Update current login user password
//	@Tags			user
//	@Security		Bearer
//	@Param			body	body		dto.UpdateUserPasswordDto	true	"Update user password"
//	@Success		200		{object}	dto.ResponseDto
//	@Router			/users/password [put]
func (c *controller) updatePassword(ctx *gin.Context) {
	req := utilities.ExtractStructFromValidator[dto.UpdateUserPasswordDto](ctx)

	err := c.userService.UpdatePassword(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.ResponseDto{
		Message: "Successfully update user password",
	})
}
