package services

import (
	"github.com/gin-gonic/gin"
	"g7love/model"
)

/*
 * 获取省份，学校　2016-10-16 05:43
 */
func Provinces(c *gin.Context) Result {
	parentid := c.Param("parentid")
	resultData := model.GetProvinces(parentid)
	return result(resultData,1,0)

}
