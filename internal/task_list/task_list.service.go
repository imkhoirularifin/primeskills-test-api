package tasklist

import (
	"primeskills-test-api/internal/domain/dto"
	"primeskills-test-api/internal/domain/entity"
	"primeskills-test-api/internal/domain/interfaces"
	"primeskills-test-api/pkg/exception"
	"primeskills-test-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

type service struct {
	taskListRepository interfaces.TaskListRepository
	taskService        interfaces.TaskService
}

type taskListRelation struct {
	tasks []dto.TaskDto
}

func (s *service) Create(ctx *gin.Context, req *dto.CreateTaskListDto) (*dto.TaskListDto, error) {
	claims := utils.ExtractClaims(ctx)

	taskList := &entity.TaskList{
		UserID: claims.Subject,
		Title:  req.Title,
	}

	err := s.taskListRepository.Create(taskList)
	if err != nil {
		return nil, err
	}

	return &dto.TaskListDto{
		ID:        taskList.ID,
		UserID:    taskList.UserID,
		Title:     taskList.Title,
		CreatedAt: taskList.CreatedAt,
		UpdatedAt: taskList.UpdatedAt,
	}, nil
}

func (s *service) FindByUserId(ctx *gin.Context) (*[]dto.TaskListDto, error) {
	claims := utils.ExtractClaims(ctx)
	taskListDtos := make([]dto.TaskListDto, 0)

	taskLists, _ := s.taskListRepository.FindByUserId(claims.Subject)

	mapRelation := s.findTaskListsRelation(taskLists)

	for _, taskList := range *taskLists {
		relation := mapRelation[taskList.ID]
		taskListDto := taskListEntityToDto(&taskList, relation)

		taskListDtos = append(taskListDtos, *taskListDto)
	}

	return &taskListDtos, nil
}

func (s *service) Update(id string, req *dto.UpdateTaskListDto) error {
	taskList, err := s.taskListRepository.FindById(id)
	if err != nil {
		return exception.NotFound("")
	}

	taskList.Title = req.Title
	return s.taskListRepository.Update(taskList)
}

func (s *service) Delete(id string) error {
	taskList, err := s.taskListRepository.FindById(id)
	if err != nil {
		return exception.NotFound("")
	}

	return s.taskListRepository.Delete(taskList)
}

func (s *service) findTaskListsRelation(taskLists *[]entity.TaskList) map[string]taskListRelation {
	mapRelation := make(map[string]taskListRelation)
	if taskLists == nil {
		return mapRelation
	}

	taskListIds := make([]string, 0)

	for _, taskList := range *taskLists {
		taskListIds = append(taskListIds, taskList.ID)
	}

	mapTasks, err := s.taskService.FindMapByTaskListIds(taskListIds)
	if err != nil {
		return mapRelation
	}

	for _, taskList := range *taskLists {
		mapRelation[taskList.ID] = taskListRelation{
			tasks: mapTasks[taskList.ID],
		}
	}

	return mapRelation
}

func taskListEntityToDto(taskList *entity.TaskList, relation taskListRelation) *dto.TaskListDto {
	return &dto.TaskListDto{
		ID:        taskList.ID,
		UserID:    taskList.UserID,
		Title:     taskList.Title,
		CreatedAt: taskList.CreatedAt,
		UpdatedAt: taskList.UpdatedAt,
		Tasks:     relation.tasks,
	}
}

func NewService(taskListRepository interfaces.TaskListRepository, taskService interfaces.TaskService) interfaces.TaskListService {
	return &service{
		taskListRepository: taskListRepository,
		taskService:        taskService,
	}
}
