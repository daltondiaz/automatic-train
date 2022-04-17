package controllers

import (
	"log"

	"github.com/daltondiaz/automatic-train/bitbucket"
	"github.com/gin-gonic/gin"
)

func Push(ctx *gin.Context) {
	var webhook bitbucket.Webhook

	if err := ctx.BindJSON(&webhook); err != nil {
		return
	}
	log.Printf("%v", webhook)
}
