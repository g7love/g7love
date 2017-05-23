package services

import (
	"github.com/gin-gonic/gin"
	"g7love/model"
	"strconv"
	"fmt"
)

/*
 * 处理点赞等操作 2017-05-21 17:32
 */
func doEevaluation(userId string,id int ,typeData string) int {
	resultDynamicLog := model.ResultDynamicLog(userId,id,"")
	firstDynamicLog := model.ResultDynamicLog(userId,id,typeData)
	getFirstDynamic := model.Getdynamic(userId,id)
	var  result  int
	fmt.Println(resultDynamicLog)
	fmt.Println(firstDynamicLog)
	if resultDynamicLog.ID != 0 {
		//更新 1
		if firstDynamicLog.ID == 0 {
			if model.UpdateDynamic(uint(id),typeData,getFirstDynamic[0].Praise + 1,getFirstDynamic[0].ReportNum+1) == 1 {
				result = model.SaveDoEevaluation(userId,id,typeData,1,
					resultDynamicLog.ID)
			}
		}
	} else {
		//插入 2
		//先更新主表再更新次表
		if firstDynamicLog.ID == 0 {
			if model.UpdateDynamic(uint(id),typeData,getFirstDynamic[0].Praise + 1,getFirstDynamic[0].ReportNum+1) == 1{
				result = model.SaveDoEevaluation(userId,id,typeData,2,0)
			}
		}
	}
	return result
}
/**
 * 点赞,举报,转发 2017-05-21 17:01
 */
func Evaluation(c *gin.Context) Result {
	user :=  c.MustGet("user").(User)
	userId := user.Id
	id,_ := strconv.Atoi(c.PostForm("id"))
	arg := c.PostForm("arg")
	resultData := doEevaluation(userId,id,arg)
	return result(resultData,1,0)
}
