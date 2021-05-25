package repositories

import (
	"github.com/juri200405/uma-repo/app/internal/domain/models"
)

type RaceResultRepository interface {
	Register(*models.RaceResult) error
}

type RaceRepository interface {
	Register(*models.Race) error
	Find(*models.Race) (models.Race, error)
	GetAll() ([]models.Race, error)
}
