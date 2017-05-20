package model

import (
	"g7love/database"
)

type ResultsData struct {
	Id int `json:"id"`
	Userid int `json:"userid"`
	Content string `json:"content"`
	Pic1 string `json:"pic_1"`
	Pic2 string `json:"pic_2"`
	Pic3 string `json:"pic_3"`
	Pic4 string `json:"pic_4"`
	Createtime []uint8 `json:"createtime"`
	Praise int `json:"praise"`
	ForwardingNum int `json:"forwarding_num"`
	ReportNum int `json:"report_num"`
	HeadPortrait string `json:"head_portrait"`
	BackgroundImage string `json:"background_image"`
	Motto string `json:"motto"`
	Nickname string `json:"nickname"`
	Birthday string `json:"birthday"`
	Gender string `json:"gender"`
	School string `json:"school"`
	Time string `json:"time"`
	ReportNumTag int `json:"report_num_tag"`
	PraiseTag int `json:"praise_tag"`
	ForwardingNumTag int `json:"forwarding_num_tag"`
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
	for i:=0; i < len(result);i++ {
		result[i].Time = "2天前"
		result[i].ReportNum = 1
		result[i].PraiseTag =2
		result[i].ForwardingNumTag = 54
	}

	/*data, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)*/
	return result
}