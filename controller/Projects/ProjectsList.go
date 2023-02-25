package controller

import (
	"AccountManage/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProjectsList(c *gin.Context) {
	var page string = c.DefaultQuery("page", "0")
	pageInt, _ := strconv.Atoi(page)
	var projects *database.Projects
	count, err := projects.GetCount()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "失败",
		})
		return
	}
	dataList, err := database.GetUsersList(pageInt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "失败",
		})
		return
	}
	Data := gin.H{
		"status": 0,
		"data":   dataList,
		"total":  count,
	}
	c.JSON(http.StatusOK, Data)
}
