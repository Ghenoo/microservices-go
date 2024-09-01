package repositories

import (
	"github.com/ghenoo/microservices-go/internal/entities"
)

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}