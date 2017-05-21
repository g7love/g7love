package services

import (
	"github.com/gin-gonic/gin"
	"g7love/model"
	"strconv"
)

/*
 * 处理点赞等操作 2017-05-21 17:32
 */
func doEevaluation(userId string,id,typeData int) int {
	resultDynamicLog := model.ResultDynamicLog(userId,id)
	if resultDynamicLog.ID != 0 {

	} else {
		//先更新住表再更新次表
		model.SaveDoEevaluation(userId,id,typeData)
	}
	return 1
}
/**
 * 点赞,举报,转发 2017-05-21 17:01
 */
func Evaluation(c *gin.Context) Result {
	user :=  c.MustGet("user").(User)
	userId := user.Id
	id,_ := strconv.Atoi(c.PostForm("id"))
	arg := c.PostForm("arg")
	var resultData int
	var  typeData  int
	if arg == "like" {
		typeData = 1
		resultData = doEevaluation(userId,id,typeData)
	} else if arg == "report" {
		typeData = 2
		resultData = doEevaluation(userId,id,typeData)
	}
	return result(resultData,1,0)
}
