package usecases

import (
	"github.com/juri200405/uma-repo/app/internal/domain/models"
	"github.com/juri200405/uma-repo/app/internal/domain/services"
)

type FactorUsecase interface {
	Register(*models.Factor) error
	GetAll() ([]models.Factor, error)
	Delete(uint) error
}

type FactorServices struct {
	Factor *services.FactorService
}

func (s *FactorServices) Register(factor *models.Factor) error {
	return s.Factor.Register(factor)
}

func (s *FactorServices) GetAll() ([]models.Factor, error) {
	return s.Factor.GetAll()
}

func (s *FactorServices) Delete(id uint) error {
	return s.Factor.Delete(id)
}