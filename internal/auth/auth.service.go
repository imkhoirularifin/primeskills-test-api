package auth

import (
	"net/http"
	"primeskills-test-api/internal/domain/dto"
	"primeskills-test-api/internal/domain/interfaces"
	"primeskills-test-api/internal/utilities"
	"primeskills-test-api/pkg/xerrors"
)

type service struct {
	userRepository interfaces.UserRepository
	userService    interfaces.UserService
}

func (s *service) Register(req *dto.RegisterDto) (*dto.TokenDto, error) {
	err := s.userService.Create(dto.CreateUserDto{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	token, err := utilities.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.TokenDto{
		Token: *token,
	}, nil
}

func (s *service) Login(req *dto.LoginDto) (*dto.TokenDto, error) {
	user, err := s.userRepository.FindByEmail(req.Email)
	if err != nil {
		return nil, xerrors.Throw(http.StatusUnauthorized, "invalid credentials")
	}

	err = utilities.ComparePassword(req.Password, user.Password)
	if err != nil {
		return nil, xerrors.Throw(http.StatusUnauthorized, "invalid credentials")
	}

	token, err := utilities.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.TokenDto{
		Token: *token,
	}, nil
}

func NewService(userRepository interfaces.UserRepository, userService interfaces.UserService) interfaces.AuthService {
	return &service{
		userRepository: userRepository,
		userService:    userService,
	}
}
