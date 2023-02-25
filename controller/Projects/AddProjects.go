package controller

import (
	"AccountManage/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Projects struct {
	UsersID      string
	ProjectsName string
	NewStatus    string
}

func AddProjects(c *gin.Context) {
	var form Projects
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}

	if len(form.UsersID) < 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "haven't userid",
		})
		return
	}
	if len(form.ProjectsName) < 6 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "haven't projects name",
		})
		return
	}
	if len(form.NewStatus) < 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "haven't status",
		})
		return
	}

	user, err := database.UserCheckUserName(form.ProjectsName)
	if err != nil && err.Error() != "record not found" {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	if len(user.UserName) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "用户名已存在",
		})
		return
	}
	UsersIDInt := StrToUInt(form.UsersID)
	NewStatusInt, _ := strconv.Atoi(form.NewStatus)
	var projects *database.Projects
	projects = &database.Projects{
		UsersID:      UsersIDInt,
		ProjectsName: form.ProjectsName,
		NewStatus:    NewStatusInt,
	}
	err = projects.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "添加成功",
		"data":    user,
	})
}

func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}
