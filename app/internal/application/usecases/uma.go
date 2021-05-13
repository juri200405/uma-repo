package usecases

import (
	"github.com/juri200405/uma-repo/app/internal/domain/models"
	"github.com/juri200405/uma-repo/app/internal/domain/services"
)

type UmaUsecase interface {
	Register(*models.Uma) error
	GetAll() ([]models.Uma, error)
	GetNames() ([]models.Uma, error)
	GetFactorList() ([]models.Factor, error)
}

type UmaServices struct {
	Uma *services.UmaService
	Factor *services.FactorService
}

func (s *UmaServices) Register(uma *models.Uma) error {
	return s.Uma.Register(uma)
}

func (s *UmaServices) GetAll() ([]models.Uma, error) {
	return s.Uma.GetAll()
}

func (s *UmaServices) GetNames() ([]models.Uma, error) {
	return s.Uma.GetNames()
}

func (s *UmaServices) GetFactorList() ([]models.Factor, error) {
	return s.Factor.GetAllSorted()
}
