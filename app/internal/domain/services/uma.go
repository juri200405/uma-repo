package services

import (
	"github.com/juri200405/uma-repo/app/internal/domain/models"
	"github.com/juri200405/uma-repo/app/internal/domain/repositories"
)

type UmaService struct {
	Repo repositories.UmaRepository
}

func (s *UmaService) Register(uma *models.Uma) error {
	return s.Repo.Register(uma)
}

func (s *UmaService) GetAll() ([]models.Uma, error) {
	return s.Repo.GetAll()
}

func (s *UmaService) GetNames() ([]models.Uma, error) {
	return s.Repo.GetNames()
}

func (s *UmaService) FindById(id uint) (models.Uma, error) {
	return s.Repo.FindById(id)
}

func (s *UmaService) Update(uma *models.Uma) error {
	return s.Repo.Update(uma)
}

func (s *UmaService) RemoveWhiteFactor(uma *models.Uma, factors []models.Factor) error {
	for _, f := range uma.WhiteFactors {
		if !inWhiteFactor(factors, &f) {
			if err := s.Repo.RemoveWhiteFactor(uma, &f); err != nil {
				return err
			}
		}
	}
	return nil
}

func inWhiteFactor(fs []models.Factor, f *models.Factor) bool {
	for _, v := range fs {
		if v.ID == f.ID {
			return true
		}
	}
	return false
}
