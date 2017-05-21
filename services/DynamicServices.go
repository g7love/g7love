package services

import (
	"github.com/gin-gonic/gin"
	"g7love/model"
	"time"
)

/*
 * 获取动态　2017-05-20 12:57
 */
func Getdynamic(c *gin.Context) Result {
	user :=  c.MustGet("user").(User)
	id := user.Id
	resultData := model.Getdynamic(id)
	currentTime := time.Now().Unix()
	for i:=0; i < len(resultData);i++ {
		resultData[i].Time = TimeDifference(resultData[i].Createtime,currentTime)
		if user.Id != "" {
			resultDynamicLog := model.ResultDynamicLog(user.Id,resultData[i].Id)
			resultData[i].ReportNumTag = resultDynamicLog.ReportNum
			resultData[i].PraiseTag = resultDynamicLog.Praise
			resultData[i].ForwardingNumTag = resultDynamicLog.ForwardingNum
		}else {
			resultData[i].ReportNumTag = 0
			resultData[i].PraiseTag = 0
			resultData[i].ForwardingNumTag = 0
		}
	}
	return result(resultData,1,0)
}

/*
 * 获取动态　2017-05-20 12:57
 */
func Posting(c *gin.Context) Result {
	user :=  c.MustGet("user").(User)
	userId := user.Id
	content := c.PostForm("count")
	resultData := model.SavePosting(content,userId)
	return result(resultData,1,0)
}