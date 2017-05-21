package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
type User struct{
	Id string
	Username string
}

type Result struct {
	Code int `json:"code"`
	Status int `json:"status"`
	Data interface{} `json:"data"`
}

/**
 * User: lizhengxinag
 * createtime: 17-05-14 14:04
 * functionRole:验证登录
 */
func isLogin(c *gin.Context,authority string) User {
	//需要忽略验证的模块
	ignoreValidation := map[string] int{
		"HomesLoginjudge" : 1,
		"dynamicGetdynamic" : 1,
	}
	Token :=  c.Param("Token")
	user := User{}
	if Token != "" {
		user.Id = "100177"
		user.Username = "wj"
		c.Abort()
	} else if ignoreValidation[authority] == 1 {
		user.Id = ""
		user.Username = ""
		c.Abort()
	} else {
		//强制登录
		data :=result("",0,0)
		c.JSON(http.StatusOK, data)
		c.Abort()
	}

	return user
}

/*
 * 封装返回的结果
 * $status＝１表示成功，$status＝０表示未登录，$status＝１0非法操作
 * @todo $status＝１0　要不要将该用户进行退出操作，并锁定该用户30min，若非登陆用户则进行Ip锁定
 */
func result(data interface{},status int,code int) Result {
	result := Result{}
	result.Code = code
	result.Status = status
	result.Data = data
	return result
}

/*
 * 计算时间差,返回时间差
 */
func TimeDifference(startdate int64, enddate int64 ) string {
	timediff := enddate - startdate
	timeDifference :=""
	days := int(timediff/86400)
	if days > 0 {
		timeDifference = strconv.Itoa(days)+"天前"
	} else {
		remain := timediff % 86400
		hours := int(remain/3600)
		if hours > 0 {
			timeDifference = strconv.Itoa(hours)+"小时前"
		} else {
			remain := remain % 3600
			mins := int(remain / 60)
			if mins > 0 {
				timeDifference = strconv.Itoa(mins)+"分钟前"
			} else {
				timeDifference = strconv.Itoa(mins)+"秒钟前"
			}
		}
	}
	return timeDifference
}

/**
 *  字符串截取 2017-5-21 14:19
 */
func SubstrString(str string,begin,length int) string {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}
