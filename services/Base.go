package services

import (
	"github.com/gin-gonic/gin"
)
type User struct{
	Id string
	Username string
}

/**
 * User: lizhengxinag
 * createtime: 17-05-14 14:04
 * functionRole:验证登录
 */
func isLogin(c *gin.Context) User {
	Token :=  c.Param("Token")
	user := User{}
	if Token != "" {
		user.Id = "2011312050"
		user.Username = "wj"
	} else {
		user.Id = "2011312050"
		user.Username = "wj"
	}
	return user
}
