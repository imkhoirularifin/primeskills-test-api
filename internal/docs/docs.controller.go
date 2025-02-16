package docs

import (
	"primeskills-test-api/docs"
	"primeskills-test-api/internal/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewController(router *gin.RouterGroup) {
	docs.SwaggerInfo.Host = config.Cfg.Swagger.Host

	router.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))
}
