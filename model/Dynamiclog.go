package model

import (
	"github.com/jinzhu/gorm"
	"g7love/database"
)

//点赞，转发，举报日志记录表
type Dynamiclog struct {
	gorm.Model
	DynamicId int `gorm:"size:7;not null;index:userid_dynamicId"` //动态id
	Userid string `gorm:"size:10;not null;index:userid_dynamicId"` //用户id
	ReportNum int `gorm:"size:2"`  //举报次数
	Praise	int `gorm:"size:2"`  //点赞次数
	ForwardingNum int `gorm:"size:2"` //转发次数
}


/**
 * 获取点赞等
 */

func ResultDynamicLog(userid string, dynamicId int) Dynamiclog {
	var argsDynamiclog Dynamiclog
	argsDynamiclog.Userid = userid
	argsDynamiclog.DynamicId = dynamicId
	db := database.GetDB()
	db.Where(argsDynamiclog).First(&argsDynamiclog)
	return  argsDynamiclog
}

/**
 * 点赞等处理
 */

func SaveDoEevaluation(userId string,dynamicId,typeData int)  {
	var argsDynamiclog Dynamiclog
	if typeData == 1 {
		argsDynamiclog.Userid = userId
		argsDynamiclog.DynamicId = dynamicId
		argsDynamiclog.Praise = 1
	} else if typeData == 2 {
		argsDynamiclog.Userid = userId
		argsDynamiclog.DynamicId = dynamicId
		argsDynamiclog.Praise = 1
	}
}