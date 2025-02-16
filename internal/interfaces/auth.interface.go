package interfaces

import "primeskills-test-api/internal/dto"

type AuthService interface {
	Register(req *dto.RegisterDto) (*dto.TokenDto, error)
	Login(req *dto.LoginDto) (*dto.TokenDto, error)
}
