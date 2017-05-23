package model
import (
	"github.com/jinzhu/gorm"
	"g7love/database"
	"time"
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

/**
 *获取用户基本信息
 */

type  Getuserinfo struct {
	Userid string `json:"userid"` //用户id
	Nickname string `json:"nickname"` //昵称
	Email  string `json:"email"` //邮箱
	Provinces int `json:"provinces"`  //省份
	School string `json:"school"`  //学校
	Birthday int64  `json:"birthday"`//生日
	Gender  int `json:"gender"` //性别: 0,男 1,女
	Motto string `json:"motto"` //个性签名
	HeadPortrait string  `json:"headPortrait"` //头像
	BackgroundImage string `json:"backgroundImage"` //背景条
	Thumb int `json:"thumb"` //主页点赞
	Self int `json:"self"`
	CreatedAt time.Time `json:"createtime"`
}

func Getuserinformation( userId string) Getuserinfo {
	var result Getuserinfo
	db := database.GetDB()
	var registeredArg Registered
	registeredArg.Userid = userId
	db.Table("registered").Select("registered.head_portrait,registered.background_image," +
			"registered.motto,registered.nickname,registered.userid,registered.created_at,registered.birthday,registered.gender," +
			"registered.thumb,school.name as school").
			Joins("LEFT JOIN `school` ON registered.school = school.id").
			Where(&registeredArg).Scan(&result)
	return result
}