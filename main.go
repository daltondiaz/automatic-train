package main

import (
	"log"

	"github.com/daltondiaz/automatic-train/bitbucket"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/v1/bitbucket", getBitbucketPush)
	router.Run("localhost:9191")
}

func getBitbucketPush(ctx *gin.Context) {
	var webhook bitbucket.Webhook

	if err := ctx.BindJSON(&webhook); err != nil {
		return
	}
	log.Printf("%v", webhook)

}
