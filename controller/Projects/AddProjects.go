package controller

import (
	"AccountManage/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Projects struct {
	UsersID      string `form:"usersid" json:"usersid" xml:"usersid"  binding:"required"`
	ProjectsName string `form:"ProjectsName" json:"ProjectsName" xml:"ProjectsName"  binding:"required"`
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
	var projects *database.Projects
	UsersIDInt := StrToUInt(form.UsersID)
	projects = &database.Projects{
		UsersID:      UsersIDInt,
		ProjectsName: form.ProjectsName,
		NewStatus:    0,
	}
	err := projects.Insert()
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
		"data":    projects,
	})
}

func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}
