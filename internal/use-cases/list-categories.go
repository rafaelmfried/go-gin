package use_cases

import (
	"gin-quickstart/internal/entities"
	"gin-quickstart/internal/repositories"
)

type listCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewListCategoriesUseCase(repository repositories.ICategoryRepository) *listCategoryUseCase {
	return &listCategoryUseCase{
		repository,
	}
}

func (useCase *listCategoryUseCase) Execute() ([]*entities.Category, error) {
	categories, err := useCase.repository.List()

	if err != nil {
		return nil, err
	}

	return categories, nil
}