package router

import (
	"AccountManage/controller"
	Admin "AccountManage/controller/Admin"
	Projects "AccountManage/controller/Projects"
	User "AccountManage/controller/User"
	utils "AccountManage/utils"

	"github.com/gin-gonic/gin"
)

// InitRouter make router
func InitRouter(SECRET_KEY, CurrentPath string) *gin.Engine {
	router := gin.Default()
	router.Use(utils.CORSMiddleware())
	adminapiv1 := router.Group("/admin/api/v1")
	adminapiv1.Use(utils.SetConfigMiddleWare(SECRET_KEY, CurrentPath))
	{
		router.GET("/", controller.Index)
		adminapiv1.POST("/AddAdmin", utils.AdminVerifyMiddleware(), Admin.AddAdmin)
		adminapiv1.PUT("/RePassword", utils.AdminVerifyMiddleware(), Admin.ResetPassword)
		adminapiv1.DELETE("/DelAdmin", utils.AdminVerifyMiddleware(), Admin.DeleteAdmin)
		adminapiv1.GET("/CheckLogin", utils.AdminVerifyMiddleware(), Admin.CheckLogin)
		adminapiv1.GET("/AdminList", utils.AdminVerifyMiddleware(), Admin.AdminList)
		adminapiv1.PUT("/UpStatus", utils.AdminVerifyMiddleware(), Admin.UpStatusAdmin)
		adminapiv1.POST("/AdminLogin", Admin.Sgin)
		adminapiv1.POST("/AddUser", utils.AdminVerifyMiddleware(), User.AddUser)
		adminapiv1.PUT("/RePasswordUser", utils.AdminVerifyMiddleware(), User.UserResetPassword)
		adminapiv1.DELETE("/DelUser", utils.AdminVerifyMiddleware(), User.DeleteUser)
		adminapiv1.GET("/UserList", utils.AdminVerifyMiddleware(), User.UsersList)
		adminapiv1.PUT("/UpStatusUser", utils.AdminVerifyMiddleware(), User.UpStatusUser)
		adminapiv1.GET("/aaa", utils.AdminVerifyMiddleware(), controller.Index)
		adminapiv1.POST("/AddProjects", utils.AdminVerifyMiddleware(), Projects.AddProjects)
		adminapiv1.DELETE("/DelProjects", utils.AdminVerifyMiddleware(), Projects.DeleteProjects)
		adminapiv1.GET("/ProjectsList", utils.AdminVerifyMiddleware(), Projects.ProjectsList)
		adminapiv1.PUT("/UpStatusProjects", utils.AdminVerifyMiddleware(), Projects.UpStatusProjects)
	}

	return router
}
