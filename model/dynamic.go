package model

import (
	"g7love/database"
)

type ResultsData struct {
	Id int
	Userid int
	Content string
	Pic1 string
	Pic2 string
	Pic3 string
	Pic4 string
	Createtime []uint8
	Praise int
	ForwardingNum int
	ReportNum int
	HeadPortrait string
	BackgroundImage string
	Motto string
	Nickname string
	Birthday string
	Gender string
	School string
}

type Dynamic struct {
	Id int
	Userid int
}
func Getdynamic() []ResultsData {
	var result []ResultsData
	db := database.GetDB()
	var dynamicArg Dynamic
	dynamicArg.Userid = 2011312050
	db.Table("dynamic").Select("dynamic.id,dynamic.userid,content,pic1,pic2,pic3,pic4,dynamic.createtime," +
			"praise,forwardingNum,reportNum,registered.headPortrait,registered.backgroundImage,registered.motto," +
			"registered.nickname,registered.birthday,registered.gender,school.name AS school").
			Joins("LEFT JOIN registered ON registered.userid = dynamic.userid").
			Joins("LEFT JOIN `school` ON registered.school = school.id").
			Where("dynamic.deleted = ?",0).
			Where(&dynamicArg).Scan(&result)
	return result
}