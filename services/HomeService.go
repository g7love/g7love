package services
import (
	"github.com/gin-gonic/gin"
)

/*
 * 判断登陆　2016-10-12 23:44
 */
func LoginJudge(c *gin.Context) Result {
	user :=  isLogin(c,"HomesLoginjudge")
	if user.Id != ""{
		return result(1,1,0)
	} else {
		return result(0,0,0)
	}
}