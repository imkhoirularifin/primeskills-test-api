package task

import (
	"primeskills-test-api/internal/domain/entity"
	"primeskills-test-api/internal/domain/interfaces"

	"gorm.io/gorm"
)

type mysqlRepository struct {
	db *gorm.DB
}

func (m *mysqlRepository) Create(task *entity.Task) error {
	return m.db.Create(&task).Error
}

func (m *mysqlRepository) FindById(id string) (*entity.Task, error) {
	var task *entity.Task

	if err := m.db.First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (m *mysqlRepository) FindByTaskListIds(taskListIds []string) (*[]entity.Task, error) {
	var tasks *[]entity.Task

	if err := m.db.Find(&tasks, "task_list_id IN ?", taskListIds).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (m *mysqlRepository) Update(task *entity.Task) error {
	return m.db.Updates(&task).Error
}

func (m *mysqlRepository) Delete(task *entity.Task) error {
	return m.db.Delete(&task).Error
}

func NewMysqlRepository(db *gorm.DB) interfaces.TaskRepository {
	return &mysqlRepository{
		db: db,
	}
}
