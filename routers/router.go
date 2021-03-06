package routers

import (
	"eloizhang/danmu/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/writedanmu", controllers.WriterDanmu)
	router.POST("/regist", controllers.Regist)
	router.POST("/login", controllers.Login)

	return router
}
