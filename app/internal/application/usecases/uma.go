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
	GetRaces(string) ([]models.Race, error)
	FindById(uint) (models.Uma, error)
	Update(*models.Uma, []uint, []models.RaceResult) error
}

type UmaServices struct {
	Uma        *services.UmaService
	Factor     *services.FactorService
	Race       *services.RaceService
	RaceResult *services.RaceResultService
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
	blue, err := s.Factor.GetAllSorted([]string{"blue"})
	if err != nil {
		return []models.Factor{}, []models.Factor{}, []models.Factor{}, err
	}
	red, err := s.Factor.GetAllSorted([]string{"red"})
	if err != nil {
		return []models.Factor{}, []models.Factor{}, []models.Factor{}, err
	}
	white, err := s.Factor.GetAllSorted([]string{"white", "green"})
	if err != nil {
		return []models.Factor{}, []models.Factor{}, []models.Factor{}, err
	}
	return blue, red, white, nil
}

func (s *UmaServices) GetRaces(sortKey string) ([]models.Race, error) {
	return s.Race.GetAll(sortKey)
}

func (s *UmaServices) FindById(id uint) (models.Uma, error) {
	return s.Uma.FindById(id)
}

func (s *UmaServices) Update(uma *models.Uma, whiteFactorIds []uint, r []models.RaceResult) error {
	whiteFactors, err := s.Factor.GetSlice(whiteFactorIds)
	if err != nil {
		return err
	}
	if err := s.Uma.RemoveWhiteFactor(uma, whiteFactors); err != nil {
		return err
	}
	uma.WhiteFactors = whiteFactors
	err = s.RaceResult.Update(r, uma)
	if err != nil {
		return err
	}
	uma.RaceResults = []models.RaceResult{}
	return s.Uma.Update(uma)
}
