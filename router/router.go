package router

import (
	"github.com/daltondiaz/automatic-train/controllers"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.POST("/v1/bitbucket", controllers.Push)
	router.Run("localhost:9191")
}
