package model

import (
	"g7love/database"
	"github.com/jinzhu/gorm"
	"time"
)

/**
 * 获取帖子条件结构体
 */
type Dynamic struct {
	gorm.Model
	Userid string  `gorm:"size:10;not null;index"` //用户id
	Content string `gorm:"size:256"` //动态内容
	Pic1	string `gorm:"size:256"` //图片1
	Pic2	string `gorm:"size:256"` //图片2
	Pic3	string `gorm:"size:256"` //图片3
	Pic4	string `gorm:"size:256"` //图片4
	ReportNum int  `gorm:"size:3"`   //举报次数
	Praise int  `gorm:"size:3"`   //点赞次数
	Createtime int64 `gorm:"size:10"` //创建时间
	ForwardingNum int  `gorm:"size:3"`   //转发次数
}

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
 * 获取帖子
 */
func Getdynamic(userId string) []ResultsData {
	var result []ResultsData
	db := database.GetDB()
	var dynamicArg Dynamic
	dynamicArg.Userid = userId
	db.Table("dynamic").Select("dynamic.id,dynamic.userid,content,pic1,pic2,pic3,pic4,dynamic.createtime," +
			"praise,forwarding_num,report_num,registered.head_portrait,registered.background_image,registered.motto," +
			"registered.nickname,registered.birthday,registered.gender,school.name AS school").
			Joins("LEFT JOIN registered ON registered.userid = dynamic.userid").
			Joins("LEFT JOIN `school` ON registered.school = school.id").
			Where(&dynamicArg).Scan(&result)
	return result
}

func SavePosting( content string,userId string) int {
	var SavePosting  Dynamic
	SavePosting.Content  = content
	SavePosting.Userid = userId
	SavePosting.Createtime = time.Now().Unix()
	db := database.GetDB()
	db.Save(&SavePosting)
	var result int
	if err := db.Save(&SavePosting).Error; err != nil {
		result = 1
	} else {
		result = 0
	}
	return result
}