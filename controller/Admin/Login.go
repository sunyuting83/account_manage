package controller

import (
	BadgerDB "AccountManage/badger"
	"AccountManage/database"
	"AccountManage/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Node node
type Login struct {
	UserName string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

func Sgin(c *gin.Context) {
	var form Login
	if err := c.ShouldBind(&form); err != nil {
		// fmt.Println(form.UserName, form.Password)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  1,
			"message": err.Error(),
		})
		return
	}

	if len(form.UserName) <= 4 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "haven't username",
		})
		return
	}
	if len(form.Password) <= 5 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  1,
			"message": "haven't password",
		})
		return
	}
	secret_key, _ := c.Get("secret_key")
	SECRET_KEY := secret_key.(string)

	PASSWD := utils.MD5(strings.Join([]string{form.Password, SECRET_KEY}, ""))
	login, err := database.CheckAdminLogin(form.UserName, PASSWD)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "登陆失败",
		})
		return
	}
	user, err := BadgerDB.Get([]byte(form.UserName))
	if err != nil && err.Error() != "Key not found" {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "登陆失败",
		})
		return
	}
	if user != "" {
		TOKEN, err := utils.EncryptByAes([]byte(user), []byte(SECRET_KEY))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  1,
				"message": "登陆失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "登陆成功",
			"token":   TOKEN,
			"user":    login.UserName,
		})
		return
	}
	T := time.Now().Format("20060102150411")
	token := utils.MD5(strings.Join([]string{login.UserName, login.Password, T}, ""))
	// var ttl int64 = 60 * 5
	var ttl int64 = 60 * 60 * 24 * 90 // ttl以秒为单位
	// ASE加密token
	TOKEN, err := utils.EncryptByAes([]byte(token), []byte(SECRET_KEY))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  1,
			"message": "登陆失败",
		})
		return
	}
	BadgerDB.SetWithTTL([]byte(form.UserName), []byte(token), ttl)
	BadgerDB.SetWithTTL([]byte(token), []byte(token), ttl)
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "登陆成功",
		"token":   TOKEN,
		"user":    login.UserName,
	})
}
