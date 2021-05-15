package datastore

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/juri200405/uma-repo/app/internal/domain/models"
)

type UmaRepository struct {
	Db *gorm.DB
}

func (r *UmaRepository) Register(uma *models.Uma) error {
	return r.Db.Create(uma).Error
}

func (r *UmaRepository) GetAll() (umas []models.Uma, err error) {
	err = r.Db.Preload(clause.Associations).Find(&umas).Error
	return
}

func (r *UmaRepository) GetNames() (umas []models.Uma, err error) {
	err = r.Db.Distinct("name").Order("name").Find(&umas).Error
	return
}
