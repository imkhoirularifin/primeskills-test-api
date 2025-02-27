package user

import (
	"net/http"
	"primeskills-test-api/internal/domain/dto"
	"primeskills-test-api/internal/domain/entity"
	"primeskills-test-api/internal/domain/interfaces"
	"primeskills-test-api/internal/utilities"
	"primeskills-test-api/pkg/xerrors"

	"github.com/gin-gonic/gin"
)

type service struct {
	userRepository interfaces.UserRepository
}

func (s *service) Create(req dto.CreateUserDto) error {
	existingUser, _ := s.userRepository.FindByEmail(req.Email)
	if existingUser != nil {
		return xerrors.Throw(http.StatusConflict, "Email already exists")
	}

	hashedPassword, err := utilities.HashPassword(req.Password)
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
	claims := utilities.ExtractClaims(ctx)

	user, err := s.userRepository.FindById(claims.Subject)
	if err != nil {
		return nil, xerrors.Throw(http.StatusUnauthorized, "Unauthorized")
	}

	return &dto.UserProfileDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *service) Update(ctx *gin.Context, req *dto.UpdateUserDto) error {
	claims := utilities.ExtractClaims(ctx)

	existingUser, err := s.userRepository.FindById(claims.Subject)
	if err != nil {
		return xerrors.Throw(http.StatusUnauthorized, "Unauthorized")
	}

	existingUserEmail, _ := s.userRepository.FindByEmail(req.Email)
	if existingUserEmail != nil && existingUser.Email != req.Email {
		return xerrors.Throw(http.StatusConflict, "Email already exists")
	}

	user := &entity.User{
		Name:  req.Name,
		Email: req.Email,
	}
	user.ID = existingUser.ID

	return s.userRepository.Update(user)
}

func (s *service) UpdatePassword(ctx *gin.Context, req *dto.UpdateUserPasswordDto) error {
	claims := utilities.ExtractClaims(ctx)

	user, err := s.userRepository.FindById(claims.Subject)
	if err != nil {
		return xerrors.Throw(http.StatusUnauthorized, "Unauthorized")
	}

	err = utilities.ComparePassword(req.OldPassword, user.Password)
	if err != nil {
		return xerrors.Throw(http.StatusUnauthorized, "Invalid old password")
	}

	hashedPassword, err := utilities.HashPassword(req.NewPassword)
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
