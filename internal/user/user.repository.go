package user

import (
	"primeskills-test-api/internal/entity"
	"primeskills-test-api/internal/interfaces"

	"gorm.io/gorm"
)

type mysqlRepository struct {
	db *gorm.DB
}

func (m *mysqlRepository) Create(user *entity.User) error {
	return m.db.Create(&user).Error
}

func (m *mysqlRepository) FindById(id string) (*entity.User, error) {
	var user *entity.User

	if err := m.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (m *mysqlRepository) FindByEmail(email string) (*entity.User, error) {
	var user *entity.User

	if err := m.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (m *mysqlRepository) Update(user *entity.User) error {
	return m.db.Updates(&user).Error
}

func NewMysqlRepository(db *gorm.DB) interfaces.UserRepository {
	return &mysqlRepository{
		db: db,
	}
}
