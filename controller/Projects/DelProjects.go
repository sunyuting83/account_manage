package controller

import (
	"AccountManage/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectsID struct {
	ID int64 `form:"id" json:"id" xml:"id"  binding:"required"`
}

func DeleteProjects(c *gin.Context) {
	var form ProjectsID
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	Projects, err := database.ProjectsCheckID(form.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	Projects.DeleteOne(form.ID)
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "成功删除后台",
		"id":      Projects.ID,
	})
}
