package services

import (
	"github.com/juri200405/uma-repo/app/internal/domain/models"
	"github.com/juri200405/uma-repo/app/internal/domain/repositories"
)

type UmaService struct {
	Repo repositories.UmaRepository
}

func(s *UmaService) Register(uma *models.Uma) error {
	return s.Repo.Register(uma)
}

func(s *UmaService) GetAll() ([]models.Uma, error) {
	return s.Repo.GetAll()
}

func(s *UmaService) GetNames() ([]models.Uma, error) {
	return s.Repo.GetNames()
}
