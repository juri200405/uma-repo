package datastore

import (
	"gorm.io/gorm"

	"github.com/juri200405/uma-repo/app/internal/domain/models"
)

type RaceRepository struct {
	Db *gorm.DB
}

func (r *RaceRepository) Register(race *models.Race) error {
	return r.Db.Create(race).Error
}

func (r *RaceRepository) Find(race *models.Race) (fr models.Race, err error) {
	err = r.Db.Where(race).First(&fr).Error
	return
}

func (r *RaceRepository) GetAll() (races []models.Race, err error) {
	err = r.Db.Find(&races).Error
	return
}


type RaceResultRepository struct {
	Db *gorm.DB
}

func (r *RaceResultRepository) Register(race *models.RaceResult) error {
	return r.Db.Create(race).Error
}
