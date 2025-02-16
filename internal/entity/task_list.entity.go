package entity

import "primeskills-test-api/pkg/xgorm"

type TaskList struct {
	xgorm.Base
	UserID string `json:"userId" gorm:"type:uuid;not null"`
	Title  string `json:"title" gorm:"type:varchar(255);not null"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
	Tasks  []Task `json:"tasks"`
}
