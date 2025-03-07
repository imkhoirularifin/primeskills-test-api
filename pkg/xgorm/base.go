package xgorm

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base is a struct that includes common fields for all database models.
type Base struct {
	ID        string         `sql:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// BeforeCreate is a GORM hook that is called before a new record is created.
// It sets the ID field to a new UUID.
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.New().String()

	return nil
}
