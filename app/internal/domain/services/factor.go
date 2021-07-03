package services

import (
	"github.com/juri200405/uma-repo/app/internal/domain/models"
	"github.com/juri200405/uma-repo/app/internal/domain/repositories"
)

type FactorService struct {
	Repo repositories.FactorRepository
}

func (s *FactorService) Register(factor *models.Factor) error {
	return s.Repo.Register(factor)
}

func (s *FactorService) GetAll() ([]models.Factor, error) {
	return s.Repo.GetAll()
}

func (s *FactorService) GetAllSorted(t []string) ([]models.Factor, error) {
	return s.Repo.GetAllSorted(t)
}

func (s *FactorService) GetSlice(ids []uint) ([]models.Factor, error) {
	return s.Repo.GetSlice(ids)
}

func (s *FactorService) GetByID(id uint) (models.Factor, error) {
	return s.Repo.GetByID(id)
}

func (s *FactorService) Delete(id uint) error {
	return s.Repo.Delete(id)
}

func (s *FactorService) Update(factor *models.Factor) error {
	return s.Repo.Update(factor)
}