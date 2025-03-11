package infrastructure

import (
	"fmt"
	"io"
	"net/http"
	swaggerDocs "primeskills-test-api/docs"
	"primeskills-test-api/internal/auth"
	"primeskills-test-api/internal/docs"
	"primeskills-test-api/internal/domain/dto"
	"primeskills-test-api/internal/task"
	tasklist "primeskills-test-api/internal/task_list"
	"primeskills-test-api/internal/user"
	"primeskills-test-api/pkg/config"
	middleware2 "primeskills-test-api/pkg/middleware"
	"primeskills-test-api/pkg/xlogger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	logger := xlogger.Logger

	if !cfg.IsDevelopment {
		gin.SetMode(gin.ReleaseMode)
		swaggerDocs.SwaggerInfo.Schemes = []string{"https"}
	}
	swaggerDocs.SwaggerInfo.Host = cfg.Swagger.Host

	// disable gin startup message
	gin.DefaultWriter = io.Discard

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(cors.New(config.CorsConfig))
	engine.Use(middleware2.Logger(logger))
	engine.Use(middleware2.HandleError())

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, &dto.ResponseDto{
			Message: "pong",
		})
	})

	api := engine.Group("/api/v1")
	docs.NewController(api.Group("/docs"))
	auth.NewController(api.Group("/auth"), authService)
	user.NewController(api.Group("/users"), userService)
	tasklist.NewController(api.Group("/task-lists"), taskListService)
	task.NewController(api.Group("/tasks"), taskService)

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	logger.Info().Msgf("Server is running on address: %s", addr)

	err := engine.Run(addr)
	if err != nil {
		logger.Fatal().Err(err).Msg("Server failed to start")
	}
}
