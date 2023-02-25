package controller

import (
	"AccountManage/database"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UpStatusProjects(c *gin.Context) {
	var form ProjectsID
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	projects, err := database.ProjectsCheckID(form.ID)
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
	if projects.NewStatus == 1 {
		NewStatus = 0
		FuckStr = "解锁"
	}
	projects.UpStatusProjects(NewStatus)

	c.JSON(http.StatusOK, gin.H{
		"status":   0,
		"message":  strings.Join([]string{"成功", FuckStr, "后台"}, ""),
		"projects": projects,
	})
}
