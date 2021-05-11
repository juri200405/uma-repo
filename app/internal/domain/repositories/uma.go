package repositories

import (
	"github.com/juri200405/uma-repo/app/internal/domain/models"
)

type UmaRepository interface {
	Register(*models.Uma) error
}
