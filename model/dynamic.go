package model

import (
	"g7love/database"
	"fmt"
)

/**
 * 返回帖子
 */
type ResultsData struct {
	Id int `json:"id"`
	Userid int `json:"userid"`
	Content string `json:"content"`
	Pic1 string `json:"pic1"`
	Pic2 string `json:"pic2"`
	Pic3 string `json:"pic3"`
	Pic4 string `json:"pic4"`
	Createtime int64 `json:"createtime"`
	Praise int `json:"praise"`
	ForwardingNum int `json:"forwardingNum"`
	ReportNum int `json:"reportNum"`
	HeadPortrait string `json:"headPortrait"`
	BackgroundImage string `json:"backgroundImage"`
	Motto string `json:"motto"`
	Nickname string `json:"nickname"`
	Birthday string `json:"birthday"`
	Gender string `json:"gender"`
	School string `json:"school"`
	Time string `json:"time"`
	ReportNumTag int `json:"reportNumTag"`
	PraiseTag int `json:"praiseTag"`
	ForwardingNumTag int `json:"forwardingNumTag"`
}

/**
 * 获取帖子条件结构体
 */
type Dynamic struct {
	Id int
	Userid string
}

/**
 * 获取帖子
 */
func Getdynamic(userId string) []ResultsData {
	var result []ResultsData
	db := database.GetDB()
	var dynamicArg Dynamic
	dynamicArg.Userid = userId
	db.Table("dynamic").Select("dynamic.id,dynamic.userid,content,pic1,pic2,pic3,pic4,dynamic.createtime," +
			"praise,forwardingNum,reportNum,registered.headPortrait,registered.backgroundImage,registered.motto," +
			"registered.nickname,registered.birthday,registered.gender,school.name AS school").
			Joins("LEFT JOIN registered ON registered.userid = dynamic.userid").
			Joins("LEFT JOIN `school` ON registered.school = school.id").
			Where("dynamic.deleted = ?",0).
			Where(&dynamicArg).Scan(&result)
	return result
}

/**
 * 获取帖子点赞等结构体
 */

type  Dynamiclog  struct {
	ReportNum int
	Praise int
	ForwardingNum int
	Userid string
	Dynamicid int
}

/**
 * 获取点赞等
 */

func ResultDynamicLog(userid string, dynamicId int) Dynamiclog {
	var argsDynamiclog Dynamiclog
	argsDynamiclog.Userid = userid
	argsDynamiclog.Dynamicid = dynamicId
	db := database.GetDB()
	db.Where(argsDynamiclog).First(&argsDynamiclog)
	fmt.Println(argsDynamiclog)
	return  argsDynamiclog
}