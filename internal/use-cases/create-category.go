package use_cases

import (
	"github.com/ghenoo/microservices-go/internal/repositories"

	"github.com/ghenoo/microservices-go/internal/entities"
)

type listCategoriesUseCase struct {
	repository repositories.ICategoryRepository
}

func NewListCategoriesUseCase(repository repositories.ICategoryRepository) *listCategoriesUseCase {
	return &listCategoriesUseCase{repository}
}

func (u *listCategoriesUseCase) Execute() ([]*entities.Category, error) {
	categories, err := u.repository.List()

	if err != nil {
		return nil, err
	}

	return categories, nil
}