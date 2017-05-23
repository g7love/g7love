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

func ResultDynamicLog(userid string, dynamicId int,typeData string) Dynamiclog {
	var argsDynamiclog Dynamiclog
	if typeData == "like" {
		argsDynamiclog.Praise = 1
	} else if typeData == "report" {
		argsDynamiclog.ReportNum = 1
	}
	argsDynamiclog.Userid = userid
	argsDynamiclog.DynamicId = dynamicId
	db := database.GetDB()
	db.Where(argsDynamiclog).First(&argsDynamiclog)
	return  argsDynamiclog
}

/**
 * 点赞等处理
 */

func SaveDoEevaluation(userId string,dynamicId int,typeData string,
		saveUpdate int,DynamicLogId uint)  int {
	var updtaDynamiclog Dynamiclog
	updtaDynamiclog.Userid = userId
	updtaDynamiclog.ID = DynamicLogId
	updtaDynamiclog.DynamicId = dynamicId
	if typeData == "like" {
		updtaDynamiclog.Praise = 1
	} else if typeData == "report" {
		updtaDynamiclog.ReportNum = 1
	}
	db := database.GetDB()
	var err error
	if saveUpdate == 1 {
		err = db.Model(&updtaDynamiclog).Update(&updtaDynamiclog).Error
	} else {
		err = db.Save(&updtaDynamiclog).Error
	}
	var result int
	if err != nil {
		result = 0
	} else {
		result = 1
	}
	return result
}