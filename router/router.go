package router

import (
	"AccountManage/controller"
	utils "AccountManage/utils"

	"github.com/gin-gonic/gin"
)

// InitRouter make router
func InitRouter(SECRET_KEY, CurrentPath string) *gin.Engine {
	router := gin.Default()
	router.Use(utils.CORSMiddleware())
	api := router.Group("/api")
	api.Use(utils.SetConfigMiddleWare(SECRET_KEY, CurrentPath))
	{
		router.GET("/", controller.Index)
	}

	return router
}
