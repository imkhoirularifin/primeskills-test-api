package tasklist

import (
	"primeskills-test-api/internal/domain/entity"
	"primeskills-test-api/internal/domain/interfaces"

	"gorm.io/gorm"
)

type mysqlRepository struct {
	db *gorm.DB
}

func (m *mysqlRepository) Create(taskList *entity.TaskList) error {
	return m.db.Create(&taskList).Error
}

func (m *mysqlRepository) FindById(id string) (*entity.TaskList, error) {
	var taskList *entity.TaskList

	if err := m.db.First(&taskList, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return taskList, nil
}

func (m *mysqlRepository) FindByUserId(userId string) (*[]entity.TaskList, error) {
	var taskLists *[]entity.TaskList

	if err := m.db.Find(&taskLists, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}

	return taskLists, nil
}

func (m *mysqlRepository) Update(taskList *entity.TaskList) error {
	return m.db.Updates(taskList).Error
}

func (m *mysqlRepository) Delete(taskList *entity.TaskList) error {
	return m.db.Delete(taskList).Error
}

func NewMysqlRepository(db *gorm.DB) interfaces.TaskListRepository {
	return &mysqlRepository{
		db: db,
	}
}
