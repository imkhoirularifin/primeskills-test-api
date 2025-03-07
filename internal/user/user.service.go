package user

import (
	"primeskills-test-api/internal/domain/dto"
	"primeskills-test-api/internal/domain/entity"
	"primeskills-test-api/internal/domain/interfaces"
	"primeskills-test-api/pkg/exception"
	utils2 "primeskills-test-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

type service struct {
	userRepository interfaces.UserRepository
}

func (s *service) Create(req dto.CreateUserDto) error {
	existingUser, _ := s.userRepository.FindByEmail(req.Email)
	if existingUser != nil {
		return exception.Conflict("Email already exists")
	}

	hashedPassword, err := utils2.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	return s.userRepository.Create(user)
}

func (s *service) FindUserProfile(ctx *gin.Context) (*dto.UserProfileDto, error) {
	claims := utils2.ExtractClaims(ctx)

	user, err := s.userRepository.FindById(claims.Subject)
	if err != nil {
		return nil, exception.Unauthorized("")
	}

	return &dto.UserProfileDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *service) Update(ctx *gin.Context, req *dto.UpdateUserDto) error {
	claims := utils2.ExtractClaims(ctx)

	existingUser, err := s.userRepository.FindById(claims.Subject)
	if err != nil {
		return exception.Unauthorized("")
	}

	existingUserEmail, _ := s.userRepository.FindByEmail(req.Email)
	if existingUserEmail != nil && existingUser.Email != req.Email {
		return exception.Conflict("Email already exists")
	}

	user := &entity.User{
		Name:  req.Name,
		Email: req.Email,
	}
	user.ID = existingUser.ID

	return s.userRepository.Update(user)
}

func (s *service) UpdatePassword(ctx *gin.Context, req *dto.UpdateUserPasswordDto) error {
	claims := utils2.ExtractClaims(ctx)

	user, err := s.userRepository.FindById(claims.Subject)
	if err != nil {
		return exception.Unauthorized("")
	}

	err = utils2.ComparePassword(req.OldPassword, user.Password)
	if err != nil {
		return exception.Unauthorized("Invalid old password")
	}

	hashedPassword, err := utils2.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return s.userRepository.Update(user)
}

func NewService(repository interfaces.UserRepository) interfaces.UserService {
	return &service{
		userRepository: repository,
	}
}
