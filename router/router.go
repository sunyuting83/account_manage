package router

import (
	"AccountManage/controller"
	Admin "AccountManage/controller/Admin"
	utils "AccountManage/utils"

	"github.com/gin-gonic/gin"
)

// InitRouter make router
func InitRouter(SECRET_KEY, CurrentPath string) *gin.Engine {
	router := gin.Default()
	router.Use(utils.CORSMiddleware())
	apiv1 := router.Group("/api/v1")
	adminapiv1 := router.Group("/admin/api/v1")
	apiv1.Use(utils.SetConfigMiddleWare(SECRET_KEY, CurrentPath))
	adminapiv1.Use(utils.SetConfigMiddleWare(SECRET_KEY, CurrentPath))
	{
		router.GET("/", controller.Index)
		adminapiv1.POST("/loginadmin", Admin.Sgin)
		adminapiv1.GET("/aaa", utils.AdminVerifyMiddleware(), controller.Index)
	}

	return router
}
