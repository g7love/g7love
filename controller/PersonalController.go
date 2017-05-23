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

/**
 * 点赞,举报,转发 2017-05-21 17:01
 */
func (u *personal) Getuserinformation(c *gin.Context) {
	result := services.Getuserinformation(c)
	c.JSON(http.StatusOK, result)
}