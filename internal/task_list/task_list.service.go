package tasklist

import (
	"net/http"
	"primeskills-test-api/internal/dto"
	"primeskills-test-api/internal/entity"
	"primeskills-test-api/internal/interfaces"
	"primeskills-test-api/internal/utilities"
	"primeskills-test-api/pkg/xerrors"

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
	claims := utilities.ExtractClaims(ctx)

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
	claims := utilities.ExtractClaims(ctx)

	taskLists, err := s.taskListRepository.FindByUserId(claims.Subject)
	if err != nil {
		return nil, xerrors.Throw(http.StatusNotFound, "task list not found")
	}

	mapRelation := s.findTaskListsRelation(taskLists)
	taskListDtos := make([]dto.TaskListDto, 0)

	for _, taskList := range *taskLists {
		relation := mapRelation[taskList.ID]
		dto := taskListEntityToDto(&taskList, relation)

		taskListDtos = append(taskListDtos, *dto)
	}

	return &taskListDtos, nil
}

func (s *service) Update(id string, req *dto.UpdateTaskListDto) error {
	taskList, err := s.taskListRepository.FindById(id)
	if err != nil {
		return xerrors.Throw(http.StatusNotFound, "task list not found")
	}

	taskList.Title = req.Title
	return s.taskListRepository.Update(taskList)
}

func (s *service) Delete(id string) error {
	taskList, err := s.taskListRepository.FindById(id)
	if err != nil {
		return xerrors.Throw(http.StatusNotFound, "task list not found")
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
