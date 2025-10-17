package controllers

import (
	"gin-quickstart/internal/repositories"
	use_cases "gin-quickstart/internal/use-cases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(context *gin.Context, repository repositories.ICategoryRepository) {
	var body createCategoryInput

	if err := context.ShouldBindBodyWithJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}

	useCase := use_cases.NewCreateCategoryUseCase(repository)

	err := useCase.Execute(body.Name)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}