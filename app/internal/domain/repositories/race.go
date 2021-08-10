package repositories

import (
	"github.com/juri200405/uma-repo/app/internal/domain/models"
)

type RaceResultRepository interface {
	Register(*models.RaceResult) error
	Update(*models.RaceResult) error
	FindById(uint) (models.RaceResult, error)
	Delete(*models.RaceResult) error
}

type RaceRepository interface {
	Register(*models.Race) error
	GetAll(string) ([]models.Race, error)
	Delete(uint) error
	GetByID(uint) (models.Race, error)
	Update(*models.Race) error
}
