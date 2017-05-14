package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"g7love/services"
)
var Homes homes = homes{}
type homes struct{}

func (u *homes) Loginjudge(c *gin.Context) {
	abc := services.LoginJudge(c)
	c.JSON(http.StatusOK, gin.H{"todos": abc})
}