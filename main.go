package main

import (
	"github.com/InformaticsResearchCenter/gowes/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	routes.MainRoute(router)
	router.Run(":5000")
}
