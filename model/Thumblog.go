package model

import "github.com/jinzhu/gorm"

//给主页点赞
type Thumblog struct {
	gorm.Model
	Thumbupuserid string `gorm:"size:10;not null"`  //点赞用户id
	Userid string `gorm:"size:10;not null"`  //'被点赞用户id'
}
