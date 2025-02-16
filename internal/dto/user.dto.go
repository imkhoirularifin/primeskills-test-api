package dto

type UserProfileDto struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserDto struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type UpdateUserDto struct {
	Name  string `json:"name" validate:"min=3,max=100"`
	Email string `json:"email" validate:"email"`
}

type UpdateUserPasswordDto struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required,min=6,max=100"`
}
