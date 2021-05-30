package repositories

import (
	"github.com/juri200405/uma-repo/app/internal/domain/models"
)

type UmaRepository interface {
	Register(*models.Uma) error
	GetAll() ([]models.Uma, error)
	GetNames() ([]models.Uma, error)
	FindById(uint) (models.Uma, error)
	Update(*models.Uma) error
	RemoveWhiteFactor(*models.Uma, *models.Factor) error
}
