package controller

import (
	"AccountManage/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AccountList(c *gin.Context) {
	var page string = c.DefaultQuery("page", "0")
	var projectsID string = c.DefaultQuery("projectsID", "0")
	pageInt, _ := strconv.Atoi(page)
	var account *database.Accounts
	count, err := account.GetCount(projectsID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "失败",
		})
		return
	}
	dataList, err := database.GetAccountList(pageInt)
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
