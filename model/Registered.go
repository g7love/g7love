package model
import (
	"github.com/jinzhu/gorm"
	"g7love/database"
)

type Registered struct {
	gorm.Model
	Userid string `gorm:"unique;size:10;not null;index"` //用户id
	Nickname string `gorm:"size:36"` //昵称
	Email  string `gorm:"unique;size:50;not null"` //邮箱
	Password string `gorm:"size:50;not null"` //密码
	Provinces int `gorm:"size:3;not null"`  //省份
	School int `gorm:"size:4;not null"`  //学校
	Birthday int64  //生日
	Gender  int `gorm:"size:1;not null"` //性别: 0,男 1,女
	Motto string `gorm:"size:50"` //个性签名
	HeadPortrait string  `gorm:"size:256"` //头像
	BackgroundImage string `gorm:"size:256"` //背景条
	Thumb int `gorm:"size:5"` //主页点赞
}

/**
 *  用户注册 2017-5-21 14:19
 */
func RegisteredSave(provinces,school int ,birthday int64,email,nickname,userId string) int {
	var RegisteredSave  Registered
	RegisteredSave.Provinces = provinces
	RegisteredSave.School = school
	RegisteredSave.Email = email
	RegisteredSave.Nickname = nickname
	RegisteredSave.Birthday = birthday
	RegisteredSave.Userid = userId
	RegisteredSave.Password = "2011"
	db := database.GetDB()
	var result int
	if err := db.Save(&RegisteredSave).Error; err != nil {
		result = 0
	} else {
		result = 1
	}
	return result
}