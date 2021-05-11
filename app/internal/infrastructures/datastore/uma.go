package datastore

import (
	"gorm.io/gorm"

	"github.com/juri200405/uma-repo/app/internal/domain/models"
)

type UmaRepository struct {
	Db *gorm.DB
}

func (r *UmaRepository) Register(uma *models.Uma) error {
	return r.Db.Create(uma).Error
}
