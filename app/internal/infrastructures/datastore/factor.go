package datastore

import (
	"gorm.io/gorm"

	"github.com/juri200405/uma-repo/app/internal/domain/models"
)

type FactorRepository struct {
	Db *gorm.DB
}

func (r *FactorRepository) Register(factor *models.Factor) error {
	return r.Db.Create(factor).Error
}

func (r *FactorRepository) GetAll() (factors []models.Factor, err error) {
	err = r.Db.Find(&factors).Error
	return
}

func (r *FactorRepository) GetAllSorted(t []string) (factors []models.Factor, err error) {
	err = r.Db.Where("type in (?)", t).Order("name, star").Find(&factors).Error
	return
}

func (r *FactorRepository) GetSlice(ids []uint) (factors []models.Factor, err error) {
	if len(ids) == 0 {
		return
	}
	err = r.Db.Where(ids).Find(&factors).Error
	return
}

func (r *FactorRepository) GetByID(id uint) (factor models.Factor, err error) {
	err = r.Db.First(&factor, id).Error
	return
}

func (r *FactorRepository) Delete(id uint) error {
	return r.Db.Delete(&models.Factor{}, id).Error
}

func (r *FactorRepository) Update(factor *models.Factor) error {
	return r.Db.Save(factor).Error
}