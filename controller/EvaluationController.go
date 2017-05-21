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
var Evaluation evaluation = evaluation{}
type evaluation struct{}

/**
 * 点赞,举报,转发 2017-05-21 17:01
 */
func (u *evaluation) Evaluation(c *gin.Context) {
	result := services.Evaluation(c)
	c.JSON(http.StatusOK, result)
}

