package repositories

import "gin-quickstart/internal/entities"

type inMemoryCategoryRepository struct {
	db []*entities.Category
}

func NewInMemotyCategoryRepository() *inMemoryCategoryRepository {
	return &inMemoryCategoryRepository{
		db: make([]*entities.Category, 0),
	}
}

func (repository *inMemoryCategoryRepository) Save(category *entities.Category) error {
	repository.db = append(repository.db, category)
	return nil
}

func (repository *inMemoryCategoryRepository) List() ([]*entities.Category, error) {
	return repository.db, nil
}