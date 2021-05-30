package datastore

import (
	"errors"

	"gorm.io/gorm"

	"github.com/juri200405/uma-repo/app/internal/domain/models"
	"github.com/juri200405/uma-repo/app/internal/domain/types"
)

type RaceRepository struct {
	Db *gorm.DB
}

func (r *RaceRepository) Register(race *models.Race) error {
	return r.Db.Create(race).Error
}

func (r *RaceRepository) GetAll() (races []models.Race, err error) {
	err = r.Db.Find(&races).Error
	return
}


type RaceResultRepository struct {
	Db *gorm.DB
}

func (r *RaceResultRepository) Register(race *models.RaceResult) (err error) {
	return r.Db.Create(race).Error
}

func (r *RaceResultRepository) Update(race *models.RaceResult) (err error) {
	return r.Db.Model(race).Updates(race).Error
}

func (r *RaceResultRepository) FindById(id uint) (race models.RaceResult, err error) {
	err = r.Db.Preload("Race").First(&race, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = types.ErrRecordNotFound
	}
	return
}

func (r *RaceResultRepository) Delete(race *models.RaceResult) (err error) {
	return r.Db.Delete(race).Error
}
