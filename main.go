package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.GET("/ping", response)
  router.Run()
}

func response(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ping 2",
	})
}