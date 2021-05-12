package repositories

import (
	"github.com/juri200405/uma-repo/app/internal/domain/models"
)

type FactorRepository interface {
	Register(*models.Factor) error
	GetAll() ([]models.Factor, error)
}