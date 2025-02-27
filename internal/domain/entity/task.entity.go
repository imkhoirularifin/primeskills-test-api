package entity

import "primeskills-test-api/pkg/xgorm"

type Task struct {
	xgorm.Base
	TaskListID  string   `json:"taskListId" gorm:"type:uuid;not null"`
	Title       string   `json:"title" gorm:"type:varchar(255);not null"`
	Description *string  `json:"description" gorm:"type:text"`
	IsCompleted bool     `json:"isCompleted" gorm:"type:boolean;not null;default:false"`
	TaskList    TaskList `json:"taskList" gorm:"foreignKey:TaskListID"`
}
