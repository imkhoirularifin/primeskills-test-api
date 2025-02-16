package docs

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewController(router *gin.RouterGroup) {
	router.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))
}
