package router

import (
	"AccountManage/controller"
	Admin "AccountManage/controller/Admin"
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
		adminapiv1.POST("/addadmin", utils.AdminVerifyMiddleware(), Admin.AddAdmin)
		adminapiv1.PUT("/repassword", utils.AdminVerifyMiddleware(), Admin.ResetPassword)
		adminapiv1.DELETE("/deladmin", utils.AdminVerifyMiddleware(), Admin.DeleteAdmin)
		adminapiv1.GET("/checklogin", utils.AdminVerifyMiddleware(), Admin.CheckLogin)
		adminapiv1.GET("/adminlist", utils.AdminVerifyMiddleware(), Admin.AdminList)
		adminapiv1.PUT("/upstatus", utils.AdminVerifyMiddleware(), Admin.UpStatusAdmin)
		adminapiv1.POST("/loginadmin", Admin.Sgin)
		adminapiv1.POST("/adduser", utils.AdminVerifyMiddleware(), User.AddUser)
		adminapiv1.PUT("/userrepassword", utils.AdminVerifyMiddleware(), User.UserResetPassword)
		adminapiv1.DELETE("/deluser", utils.AdminVerifyMiddleware(), User.DeleteUser)
		adminapiv1.GET("/userlist", utils.AdminVerifyMiddleware(), User.UsersList)
		adminapiv1.PUT("/userupstatus", utils.AdminVerifyMiddleware(), User.UpStatusUser)
		adminapiv1.GET("/aaa", utils.AdminVerifyMiddleware(), controller.Index)
	}

	return router
}
