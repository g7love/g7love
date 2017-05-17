package services

import (
	"github.com/gin-gonic/gin"
	"g7love/model"
)


/*
 * 获取省份，学校　2016-10-16 05:43
 */
func Getdynamic(c *gin.Context) Result {
	// user :=  isLogin(c,"dynamicGetdynamic")
	resultData := model.Getdynamic()
	return result(resultData,1,0)

}
