package interfaces

import "primeskills-test-api/internal/domain/dto"

type AuthService interface {
	Register(req *dto.RegisterDto) (*dto.TokenDto, error)
	Login(req *dto.LoginDto) (*dto.TokenDto, error)
}
