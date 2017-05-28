package controller

/**
 * User: lizhengxiang
 * Date: 17-5-28
 * Time: 上午8:52
 */

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"g7love/services"
)
var Upload upload = upload{}
type upload struct{}

func (u *upload) Upload(c *gin.Context) {
	result := services.Upload(c)
	c.JSON(http.StatusOK, result)
}