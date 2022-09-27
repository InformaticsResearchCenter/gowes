package controllers

import (
	"github.com/InformaticsResearchCenter/gowes/models"
	"github.com/gin-gonic/gin"
)

func GetApi(ctx *gin.Context) {
	var chat models.Chat
	ctx.BindJSON(&chat)
	chat.Phone_number = "false"
	ctx.JSON(200, &chat)
}

func GetHome(ctx *gin.Context) {
	ctx.JSON(200, "selamat datang")
}
