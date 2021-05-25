package usecases

import (
	"github.com/juri200405/uma-repo/app/internal/domain/models"
	"github.com/juri200405/uma-repo/app/internal/domain/services"
)

type UmaUsecase interface {
	Register(*models.Uma, []uint, []models.RaceResult) error
	GetAll() ([]models.Uma, error)
	GetNames() ([]models.Uma, error)
	GetFactorList() ([]models.Factor, []models.Factor, []models.Factor, error)
	GetRaces() ([]models.Race, error)
}

type UmaServices struct {
	Uma *services.UmaService
	Factor *services.FactorService
	Race *services.RaceService
}

func (s *UmaServices) Register(uma *models.Uma, whiteFactorIds []uint, r []models.RaceResult) error {
	whiteFactors, err := s.Factor.GetSlice(whiteFactorIds)
	if err != nil {
		return err
	}
	uma.WhiteFactors = whiteFactors
	uma.RaceResults = r
	return s.Uma.Register(uma)
}

func (s *UmaServices) GetAll() ([]models.Uma, error) {
	return s.Uma.GetAll()
}

func (s *UmaServices) GetNames() ([]models.Uma, error) {
	return s.Uma.GetNames()
}

func (s *UmaServices) GetFactorList() ([]models.Factor, []models.Factor, []models.Factor, error) {
	blue, err := s.Factor.GetAllSorted("blue")
	if err != nil {
		return []models.Factor{}, []models.Factor{}, []models.Factor{}, err
	}
	red, err := s.Factor.GetAllSorted("red")
	if err != nil {
		return []models.Factor{}, []models.Factor{}, []models.Factor{}, err
	}
	white, err := s.Factor.GetAllSorted("white")
	if err != nil {
		return []models.Factor{}, []models.Factor{}, []models.Factor{}, err
	}
	return blue, red, white, nil
}

func (s *UmaServices) GetRaces() ([]models.Race, error) {
	return s.Race.GetAll()
}
