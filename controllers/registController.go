package controllers

import (
	"github.com/gin-gonic/gin"
)

func Regist(c *gin.Context) {
	name := c.Request.FormValue("Name")
	passwd := c.Request.FormValue("Passwd")

}
