package services

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
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
func isLogin(c *gin.Context,authority string) User {
	//需要忽略验证的模块
	ignoreValidation := map[string] int{
		"HomesLoginjudge" : 1,
		"dynamicGetdynamic" : 1,
	}
	fmt.Println("***********")
	fmt.Println(c.ContentType())
	fmt.Println("***********")
	Token :=  c.Param("Token")
	user := User{}
	if Token != "" {
		user.Id = "2011312050"
		user.Username = "wj"
	} else if ignoreValidation[authority] == 1 {
		user.Id = ""
		user.Username = ""
	} else {
		//强制登录
		data :=result("",0,0)
		c.JSON(http.StatusOK, data)
		c.Abort()
	}

	return user
}

/*
 * 封装返回的结果
 * $status＝１表示成功，$status＝０表示未登录，$status＝１0非法操作
 * @todo $status＝１0　要不要将该用户进行退出操作，并锁定该用户30min，若非登陆用户则进行Ip锁定
 */
func result(data interface{},status int,code int) Result {
	result := Result{}
	result.Code = code
	result.Status = status
	result.Data = data
	return result
}
