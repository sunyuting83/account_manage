package controller

import (
	BadgerDB "AccountManage/badger"
	"AccountManage/database"
	"AccountManage/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Projects struct {
	UsersID      string `form:"usersid" json:"usersid" xml:"usersid"  binding:"required"`
	ProjectsName string `form:"ProjectsName" json:"ProjectsName" xml:"ProjectsName"  binding:"required"`
}

type CacheValue struct {
	UsersID    string `json:"UsersID"`
	ProjectsID string `json:"ProjectsID"`
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

	if len(form.UsersID) != 0 {
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

	projectsIDInt := strconv.Itoa(int(projects.ID))
	projectsIDStr := string(projectsIDInt)
	d := time.Now()
	date := d.Format("2006-01-02_15:04:05")
	key := utils.MD5(strings.Join([]string{form.UsersID, date, projectsIDStr}, ""))
	key = key[:12]

	cache := &CacheValue{
		UsersID:    form.UsersID,
		ProjectsID: projectsIDStr,
	}
	CacheValues, _ := json.Marshal(&cache)

	BadgerDB.Set([]byte(key), CacheValues)

	projects.UpProjectsKey(key)

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
