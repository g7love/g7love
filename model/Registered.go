package model
import (
	"g7love/database"
	"github.com/jinzhu/gorm"
	"strconv"
)
type Registered struct {
	gorm.Model
	Email		string	`gorm:"size:255"`
	School		int `sql:"type:int(4);NOT NULL;COMMENT:'学校'"`
	Name         string  `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
}

type School struct {
	Id			int `json:"id"`
	Name		string `json:"name"`
	parentid	int `json:"parentid"`
	deleted 	int `json:"deleted"`
}

func Creat(parentid string) []School {
	parent := 0
	if parentid != ""{
		parent,_ = strconv.Atoi(parentid)
	}
	var result []School
	db := database.GetDB()
	db.Where("parentid = ?",parent).Find(&result)
	return result
}