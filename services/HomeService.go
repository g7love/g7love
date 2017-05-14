package services
import (
	"github.com/gin-gonic/gin"
)

/*
 * 判断登陆　2016-10-12 23:44
 */
var user1 User
func LoginJudge(c *gin.Context) User {
	user1 =  isLogin(c)
	return user1
}