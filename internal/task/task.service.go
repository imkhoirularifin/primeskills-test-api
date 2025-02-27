package task

import (
	"net/http"
	"primeskills-test-api/internal/domain/dto"
	"primeskills-test-api/internal/domain/entity"
	"primeskills-test-api/internal/domain/interfaces"
	"primeskills-test-api/pkg/xerrors"
)

type service struct {
	taskRepository     interfaces.TaskRepository
	taskListRepository interfaces.TaskListRepository
}

func (s *service) Create(req *dto.CreateTaskDto) (*dto.TaskDto, error) {
	_, err := s.taskListRepository.FindById(req.TaskListID)
	if err != nil {
		return nil, xerrors.Throw(http.StatusNotFound, "Task list not found")
	}

	task := &entity.Task{
		TaskListID:  req.TaskListID,
		Title:       req.Title,
		Description: req.Description,
	}

	err = s.taskRepository.Create(task)
	if err != nil {
		return nil, err
	}

	taskDto := taskEntityToDto(task)

	return taskDto, nil
}

func (s *service) FindMapByTaskListIds(taskListIds []string) (map[string][]dto.TaskDto, error) {
	mapTasksByTaskListId := make(map[string][]dto.TaskDto)

	if len(taskListIds) == 0 {
		return mapTasksByTaskListId, nil
	}

	tasks, err := s.taskRepository.FindByTaskListIds(taskListIds)
	if err != nil {
		return nil, err
	}

	for _, task := range *tasks {
		taskDto := taskEntityToDto(&task)
		mapTasksByTaskListId[task.TaskListID] = append(mapTasksByTaskListId[task.TaskListID], *taskDto)
	}

	return mapTasksByTaskListId, nil
}

func (s *service) Update(id string, req *dto.UpdateTaskDto) error {
	task, err := s.taskRepository.FindById(id)
	if err != nil {
		return xerrors.Throw(http.StatusNotFound, "task not found")
	}

	task.Title = req.Title
	task.Description = req.Description
	task.IsCompleted = req.IsCompleted

	err = s.taskRepository.Update(task)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(id string) error {
	task, err := s.taskRepository.FindById(id)
	if err != nil {
		return xerrors.Throw(http.StatusNotFound, "task not found")
	}

	err = s.taskRepository.Delete(task)
	if err != nil {
		return err
	}

	return nil
}

func taskEntityToDto(task *entity.Task) *dto.TaskDto {
	return &dto.TaskDto{
		ID:          task.ID,
		TaskListID:  task.TaskListID,
		Title:       task.Title,
		Description: task.Description,
		IsCompleted: task.IsCompleted,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func NewService(taskRepository interfaces.TaskRepository, taskListRepository interfaces.TaskListRepository) interfaces.TaskService {
	return &service{
		taskRepository:     taskRepository,
		taskListRepository: taskListRepository,
	}
}
