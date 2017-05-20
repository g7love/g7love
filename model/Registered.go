package model
import (
	"github.com/jinzhu/gorm"
)

type Registered struct {
	gorm.Model
	Userid string `gorm:"size:10;not null;index"` //用户id
	Nickname string `gorm:"size:36"` //昵称
	Email  string `gorm:"size:50;not null"` //邮箱
	Password string `gorm:"size:50;not null"` //密码
	Provinces int `gorm:"size:3;not null"`  //省份
	School int `gorm:"size:4;not null"`  //学校
	Birthday int64 `gorm:"size:10;not null"` //生日
	Gender  int `gorm:"size:1;not null"` //性别: 0,男 1,女
	Motto string `gorm:"size:50"` //个性签名
	HeadPortrait string  `gorm:"size:256"` //头像
	BackgroundImage string `gorm:"size:256"` //背景条
	Thumb int `gorm:"size:5"` //主页点赞
}