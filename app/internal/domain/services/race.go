package services

import (
	"github.com/juri200405/uma-repo/app/internal/domain/models"
	"github.com/juri200405/uma-repo/app/internal/domain/repositories"
)

type RaceService struct {
	Repo repositories.RaceRepository
}

func (s *RaceService) Register(r *models.Race) error {
	return s.Repo.Register(r)
}

func (s *RaceService) Find(r *models.Race) error {
	_, err := s.Repo.Find(r)
	return err
}

func (s *RaceService) GetAll() ([]models.Race, error) {
	return s.Repo.GetAll()
}


type RaceResultService struct {
	Repo repositories.RaceResultRepository
}

func (s *RaceResultService) Register(r *models.RaceResult) error {
	return s.Repo.Register(r)
}
