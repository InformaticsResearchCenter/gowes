package routes

import (
	"github.com/InformaticsResearchCenter/gowes/controllers"
	"github.com/gin-gonic/gin"
)

func MainRoute(router *gin.Engine) {
	router.POST("/api", controllers.GetApi)
	router.GET("/", controllers.GetHome)
}
