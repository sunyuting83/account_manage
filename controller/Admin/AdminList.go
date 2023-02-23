package controller

import (
	"AccountManage/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AdminList(c *gin.Context) {
	var page string = c.DefaultQuery("page", "0")
	pageInt, _ := strconv.Atoi(page)
	var admin database.Manager
	count, err := admin.GetCount()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "失败",
		})
		return
	}
	dataList, err := database.GetAdminList(pageInt)
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
