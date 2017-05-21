package services

import (
	"github.com/gin-gonic/gin"
	"g7love/model"
	"strconv"
	"time"
	"math/rand"
	"fmt"
)

/*
 * 获取省份，学校　2016-10-16 05:43
 */
func Provinces(c *gin.Context) Result {
	parentid := c.PostForm("parentid")
	resultData := model.GetProvinces(parentid)
	return result(resultData,1,0)
}

/**
 *  生成用户id 2017-5-21 14:19
 */
func getUserId(num_id string)  string{
	UserId,_ := strconv.Atoi(SubstrString(num_id,0,4))
	UserId +=1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(UserId)+strconv.Itoa(r.Intn(100))
}

/**
 *  用户注册 2017-5-21 14:19
 */
func Registered(c *gin.Context) Result {
	provinces,_ := strconv.Atoi(c.PostForm("provinces"))
	school,_ := strconv.Atoi(c.PostForm("school"))
	email := c.PostForm("email")
	nickname := c.PostForm("nickname")
	year := c.PostForm("year")
	mouth := c.PostForm("mouth")
	mouthStr := []rune(mouth)
	mouthLen := len(mouthStr)
	if mouthLen < 2 {
		mouth = "0" + mouth
	}
	day := c.PostForm("day")
	dayStr := []rune(day)
	dayLen := len(dayStr)
	if dayLen < 2 {
		day = "0" + day
	}
	date := mouth+"/"+day+"/"+year
	birthday, _ := time.Parse("01/02/2006", date)
	fmt.Println(birthday)
	userId := getUserId("100000")
	resultData := model.RegisteredSave(provinces,school,birthday.Unix(),email,nickname,userId)
	return result(resultData,1,0)
}
