package usecases

import (
	"github.com/juri200405/uma-repo/app/internal/domain/models"
	"github.com/juri200405/uma-repo/app/internal/domain/services"
)

type RaceUsecase interface {
	Register(*models.Race) error
	GetRaces(string) ([]models.Race, error)
	Delete(uint) error
}

type RaceServices struct {
	Race *services.RaceService
}

func (s *RaceServices) Register(r *models.Race) error {
	return s.Race.Register(r)
}

func (s *RaceServices) GetRaces(sortKey string) ([]models.Race, error) {
	return s.Race.GetAll(sortKey)
}

func (s *RaceServices) Delete(id uint) error {
	return s.Race.Delete(id)
}