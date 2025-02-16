package infrastructure

import (
	"primeskills-test-api/internal/auth"
	"primeskills-test-api/internal/config"
	"primeskills-test-api/internal/interfaces"
	"primeskills-test-api/internal/task"
	tasklist "primeskills-test-api/internal/task_list"
	"primeskills-test-api/internal/user"
	"primeskills-test-api/pkg/xlogger"

	_ "github.com/joho/godotenv/autoload"
)

var (
	cfg config.Config

	userRepository     interfaces.UserRepository
	taskListRepository interfaces.TaskListRepository
	taskRepository     interfaces.TaskRepository

	userService     interfaces.UserService
	authService     interfaces.AuthService
	taskListService interfaces.TaskListService
	taskService     interfaces.TaskService
)

func init() {
	config.Setup()
	cfg = config.Cfg

	xlogger.Setup()
	dbSetup()

	userRepository = user.NewMysqlRepository(db)
	taskListRepository = tasklist.NewMysqlRepository(db)
	taskRepository = task.NewMysqlRepository(db)

	userService = user.NewService(userRepository)
	authService = auth.NewService(userRepository, userService)
	taskService = task.NewService(taskRepository, taskListRepository)
	taskListService = tasklist.NewService(taskListRepository, taskService)
}
