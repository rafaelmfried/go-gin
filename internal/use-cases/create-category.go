package use_cases

import (
	"gin-quickstart/internal/entities"
	"gin-quickstart/internal/repositories"
	"log"
)

type createCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{
		repository,
	}
}

func (useCase *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	// Todo persiste entity to db
	log.Println(category)

	err = useCase.repository.Save(category)

	if err != nil {
		return err
	}

	return nil
}