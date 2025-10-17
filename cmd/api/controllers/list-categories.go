package controllers

import (
	"gin-quickstart/internal/repositories"
	use_cases "gin-quickstart/internal/use-cases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCategory(context *gin.Context, repository repositories.ICategoryRepository) {
	useCase := use_cases.NewListCategoriesUseCase(repository)

	categories, err := useCase.Execute()

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
		"categories": categories,
	})
}