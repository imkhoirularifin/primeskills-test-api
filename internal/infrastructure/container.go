package infrastructure

import (
	"primeskills-test-api/internal/auth"
	"primeskills-test-api/internal/domain/interfaces"
	"primeskills-test-api/internal/gcloud_storage"
	"primeskills-test-api/internal/task"
	tasklist "primeskills-test-api/internal/task_list"
	"primeskills-test-api/internal/user"
	"primeskills-test-api/pkg/config"
	"primeskills-test-api/pkg/midtrans/midtrans_subscription"
	"primeskills-test-api/pkg/xlogger"

	_ "github.com/joho/godotenv/autoload"
)

var (
	cfg config.Config

	userRepository     interfaces.UserRepository
	taskListRepository interfaces.TaskListRepository
	taskRepository     interfaces.TaskRepository

	userService                 interfaces.UserService
	authService                 interfaces.AuthService
	taskListService             interfaces.TaskListService
	taskService                 interfaces.TaskService
	gcloudStorageService        interfaces.GcloudStorageService
	midtransSubscriptionService midtrans_subscription.IMidtransSubscriptionService
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
	gcloudStorageService = gcloud_storage.NewService()
	midtransSubscriptionService = midtrans_subscription.NewService()
}
