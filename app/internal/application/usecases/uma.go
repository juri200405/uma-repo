package usecases

import (
	"github.com/juri200405/uma-repo/app/internal/domain/models"
	"github.com/juri200405/uma-repo/app/internal/domain/services"
)

type UmaUsecase interface {
	Register(*models.Uma) error
}

type UmaServices struct {
	Uma *services.UmaService
}

func (s *UmaServices) Register(uma *models.Uma) error {
	return s.Uma.Register(uma)
}
