package model

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"g7love/database"
)

//各省份的学校
type School struct {
	gorm.Model
	Name		string	`gorm:"size:255"`	//省份学校专业
	Parentid	int		`gorm:"size:2"`  		//父id
}

type ResultSchool struct {
	Name		string	`json:"name"`	//省份学校专业
	Parentid	int		`json:"parentid"`  		//父id
}

func GetProvinces(parentid string) []ResultSchool {
	var school  School
	parent := 0
	if parentid != ""{
		parent,_ = strconv.Atoi(parentid)
	}
	var result []ResultSchool
	db := database.GetDB()
	db.Select("name,parentid").Where("parentid = ?",parent).Find(&school).Scan(&result)
	return result
}