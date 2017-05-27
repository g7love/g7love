package controller
/**
 * User: lizhengxiang
 * Date: 17-5-24
 * Time: 上午5:57
 */

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"g7love/services"
)
var Personal personal = personal{}
type personal struct{}


func (u *personal) Getuserinformation(c *gin.Context) {
	result := services.Getuserinformation(c)
	c.JSON(http.StatusOK, result)
}