package controller

import (
	"AccountManage/database"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UpStatusAdmin(c *gin.Context) {
	var form User
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	user, err := database.CheckID(form.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	var (
		NewStatus int    = 1
		FuckStr   string = "锁定"
	)
	if user.NewStatus == 1 {
		NewStatus = 0
		FuckStr = "解锁"
	}
	user.UpStatusAdmin(NewStatus)

	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": strings.Join([]string{"成功", FuckStr, "管理员"}, ""),
		"user":    NewStatus,
	})
}
