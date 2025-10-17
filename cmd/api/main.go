package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.GET("/helthz", helthz)
	CategoryRoutes(router)
  router.Run(":8000")
}

func helthz(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}