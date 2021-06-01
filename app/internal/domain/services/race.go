package services

import (
	"errors"

	"github.com/juri200405/uma-repo/app/internal/domain/models"
	"github.com/juri200405/uma-repo/app/internal/domain/repositories"
	"github.com/juri200405/uma-repo/app/internal/domain/types"
)

type RaceService struct {
	Repo repositories.RaceRepository
}

func (s *RaceService) Register(r *models.Race) error {
	return s.Repo.Register(r)
}

func (s *RaceService) GetAll(sortKey string) ([]models.Race, error) {
	return s.Repo.GetAll(sortKey)
}

func (s *RaceService) Delete(id uint) error {
	return s.Repo.Delete(id)
}

type RaceResultService struct {
	Repo repositories.RaceResultRepository
}

func (s *RaceResultService) Update(races []models.RaceResult, uma *models.Uma) error {
	for _, r := range uma.RaceResults {
		if !inRaceResult(races, &r) {
			if err := s.Repo.Delete(&r); err != nil {
				return err
			}
		}
	}
	for _, r := range races {
		r.UmaID = uma.ID
		if _, err := s.Repo.FindById(r.ID); errors.Is(err, types.ErrRecordNotFound) {
			if e := s.Repo.Register(&r); e != nil {
				return e
			}
		} else if err != nil {
			return err
		} else {
			if e := s.Repo.Update(&r); e != nil {
				return e
			}
		}
	}
	return nil
}

func inRaceResult(rs []models.RaceResult, r *models.RaceResult) bool {
	for _, v := range rs {
		if v.ID == r.ID {
			return true
		}
	}
	return false
}
