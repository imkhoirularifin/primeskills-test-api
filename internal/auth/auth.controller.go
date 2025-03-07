package auth

import (
	"net/http"
	"primeskills-test-api/internal/domain/dto"
	"primeskills-test-api/internal/domain/interfaces"
	"primeskills-test-api/pkg/middleware"
	"primeskills-test-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

type controller struct {
	authService interfaces.AuthService
}

func NewController(router *gin.RouterGroup, authService interfaces.AuthService) {
	controller := &controller{
		authService: authService,
	}

	router.POST("/login", middleware.Validate[dto.LoginDto](), controller.login)
	router.POST("/register", middleware.Validate[dto.RegisterDto](), controller.register)
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login
//	@Tags			auth
//	@Param			body	body		dto.LoginDto	true	"login request"
//	@Success		200		{object}	dto.TokenDto
//	@Router			/auth/login [post]
func (c *controller) login(ctx *gin.Context) {
	req := utils.ExtractStructFromValidator[dto.LoginDto](ctx)

	token, err := c.authService.Login(req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.ResponseDto{
		Message: "Login successful",
		Data:    token,
	})
}

// Register godoc
//
//	@Summary		Register
//	@Description	Register
//	@Tags			auth
//	@Param			body	body		dto.RegisterDto	true	"register request"
//	@Success		200		{object}	dto.TokenDto
//	@Router			/auth/register [post]
func (c *controller) register(ctx *gin.Context) {
	req := utils.ExtractStructFromValidator[dto.RegisterDto](ctx)

	token, err := c.authService.Register(req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &dto.ResponseDto{
		Message: "Registration successful",
		Data:    token,
	})
}
