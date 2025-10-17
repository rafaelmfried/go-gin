package main

import (
	"gin-quickstart/cmd/api/controllers"
	"gin-quickstart/internal/repositories"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	routes := router.Group("/categories")
	inMemoryCategoryRepository := repositories.NewInMemotyCategoryRepository()
	
	routes.POST("", func(ctx *gin.Context) {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)
	})

	routes.GET("", func(ctx *gin.Context) {
		controllers.ListCategory(ctx, inMemoryCategoryRepository)
	})
}