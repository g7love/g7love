package model

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"g7love/database"
)

//各省份的学校
type School struct {
	gorm.Model
	Name		string	`gorm:"size:255"` //省份学校专业
	Parentid	int		`gorm:"size:2"`  //父id
}

func GetProvinces(parentid string) []School {
	parent := 0
	if parentid != ""{
		parent,_ = strconv.Atoi(parentid)
	}
	var result []School
	db := database.GetDB()
	db.Where("parentid = ?",parent).Find(&result)
	return result
}