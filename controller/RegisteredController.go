package controller
/**
 * User: lizhengxinag
 * createtime: 17-05-16 05:31
 * functionRole:注册模块
 */
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"g7love/services"
)
var Registered registered = registered{}
type registered struct{}

func (u *registered) Provinces(c *gin.Context) {
	result := services.Provinces(c)
	c.JSON(http.StatusOK, result)
}


/**
 *  用户注册 2017-5-21 14:19
 */
func (u *registered) Registered(c *gin.Context) {
	result := services.Registered(c)
	c.JSON(http.StatusOK, result)
}
