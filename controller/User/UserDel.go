package controller

import (
	"AccountManage/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserID struct {
	ID int64 `form:"id" json:"id" xml:"id"  binding:"required"`
}

func DeleteUser(c *gin.Context) {
	var form UserID
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	user, err := database.UserCheckID(form.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}
	user.DeleteOne(form.ID)
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "成功删除用户",
		"id":      user.ID,
	})
}
