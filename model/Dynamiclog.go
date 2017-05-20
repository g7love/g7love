package model

import (
	"github.com/jinzhu/gorm"
	"fmt"
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
	fmt.Println(argsDynamiclog)
	return  argsDynamiclog
}