package entity

import (
	"primeskills-test-api/pkg/xgorm"
)

type User struct {
	xgorm.Base
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password string `json:"-" gorm:"not null"`
}
