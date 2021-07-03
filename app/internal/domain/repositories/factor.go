package repositories

import (
	"github.com/juri200405/uma-repo/app/internal/domain/models"
)

type FactorRepository interface {
	Register(*models.Factor) error
	GetAll() ([]models.Factor, error)
	GetAllSorted([]string) ([]models.Factor, error)
	GetSlice([]uint) ([]models.Factor, error)
	GetByID(uint) (models.Factor, error)
	Delete(uint) error
	Update(*models.Factor) error
}