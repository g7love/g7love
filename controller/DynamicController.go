package controller

/**
 * User: lizhengxiang
 * Date: 17-5-17
 * Time: 上午07:31
 */

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"g7love/services"
)
var Dynamic dynamic = dynamic{}
type dynamic struct{}

func (u *dynamic) Getdynamic(c *gin.Context) {
	result := services.Getdynamic(c)
	c.JSON(http.StatusOK, result)
}

func (u *dynamic) Posting(c *gin.Context) {
	result := services.Posting(c)
	c.JSON(http.StatusOK, result)
}