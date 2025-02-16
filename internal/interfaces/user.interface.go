package interfaces

import (
	"primeskills-test-api/internal/dto"
	"primeskills-test-api/internal/entity"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindById(id string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
}

type UserService interface {
	Create(req dto.CreateUserDto) error
	FindUserProfile(ctx *gin.Context) (*dto.UserProfileDto, error)
	Update(ctx *gin.Context, req *dto.UpdateUserDto) error
	UpdatePassword(ctx *gin.Context, req *dto.UpdateUserPasswordDto) error
}
