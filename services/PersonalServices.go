package services

import (
	"github.com/gin-gonic/gin"
	"g7love/model"
	"fmt"
)


/*
 * User: lizhengxiang
 * Date: 17-5-24
 * Time: 上午6:25
 * 获取自己自己或别人的基本信息
 * return 基本信息＋该用户是否是自己（否则返回有没有关注）等基本信息
 */
func Getuserinformation(c *gin.Context) Result {
	user :=  c.MustGet("user").(User)
	userid := c.PostForm("userid")
	if userid == "" {
		userid = user.Id
	}
	fmt.Println(userid)
	if userid != "" {
		resultData := model.Getuserinformation(userid)
		fmt.Println(resultData)
		if resultData.Userid != "" {
			if userid == user.Id {
				resultData.Self = 1
			}else {
				//@todo 获取关注与否
				resultData.Self = 2
			}
			return result(resultData,1,0)
		}
		//如果用户id为空表示非法操作
		result("",10,0)
	}
	return result("",0,0)
}