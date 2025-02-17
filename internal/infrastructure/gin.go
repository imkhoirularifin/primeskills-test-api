package infrastructure

import (
	"fmt"
	"io"
	"net/http"
	swaggerDocs "primeskills-test-api/docs"
	"primeskills-test-api/internal/auth"
	"primeskills-test-api/internal/docs"
	"primeskills-test-api/internal/middleware"
	"primeskills-test-api/internal/task"
	tasklist "primeskills-test-api/internal/task_list"
	"primeskills-test-api/internal/user"
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

	// disable gin's startup message
	gin.DefaultWriter = io.Discard

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowCredentials = true

	app := gin.New()

	app.Use(gin.Recovery())
	app.Use(cors.New(corsConfig))
	app.Use(middleware.Zerolog(logger))
	app.Use(middleware.HandleError())

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := app.Group("/api/v1")
	docs.NewController(api.Group("/docs"))
	auth.NewController(api.Group("/auth"), authService)
	user.NewController(api.Group("/users"), userService)
	tasklist.NewController(api.Group("/task-lists"), taskListService)
	task.NewController(api.Group("/tasks"), taskService)

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	logger.Info().Msgf("Server is running on address: %s", addr)

	err := app.Run(addr)
	if err != nil {
		logger.Fatal().Err(err).Msg("Server failed to start")
	}
}
