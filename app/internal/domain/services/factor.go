package services

import (
	"github.com/juri200405/uma-repo/app/internal/domain/models"
	"github.com/juri200405/uma-repo/app/internal/domain/repositories"
)

type FactorService struct {
	Repo repositories.FactorRepository
}

func(s *FactorService) Register(factor *models.Factor) error {
	return s.Repo.Register(factor)
}

func(s *FactorService) GetAll() ([]models.Factor, error) {
	return s.Repo.GetAll()
}

func(s *FactorService) GetAllSorted() ([]models.Factor, error) {
	return s.Repo.GetAllSorted()
}

func(s *FactorService) GetSlice(ids []uint) ([]models.Factor, error) {
	return s.Repo.GetSlice(ids)
}
