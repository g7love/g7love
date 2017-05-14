package services

import (
	"github.com/gin-gonic/gin"
	"fmt"
)
type User struct{
	Id string
	Username string
}

type Result struct {
	Code int
	Status int
	Data interface{}
}



/**
 * User: lizhengxinag
 * createtime: 17-05-14 14:04
 * functionRole:验证登录
 */
func isLogin(c *gin.Context) User {
	fmt.Println(c.ContentType())
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

func result(data interface{},status int,code int) Result {
	result := Result{}
	result.Code = code
	result.Status = status
	result.Data = data
	return result
}
