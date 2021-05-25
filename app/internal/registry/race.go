package registry

import (
	"github.com/juri200405/uma-repo/app/internal/infrastructures/dao/mysql"
	"github.com/juri200405/uma-repo/app/internal/infrastructures/datastore"
	"github.com/juri200405/uma-repo/app/internal/domain/services"
	"github.com/juri200405/uma-repo/app/internal/application/usecases"
)

type RaceRegistry struct {
}

func NewRaceRegistry() *RaceRegistry {
	return &RaceRegistry{}
}

func (r RaceRegistry) GetRaceUsecase() usecases.RaceUsecase {
	db := mysql.GetConnection()

	rr := &datastore.RaceRepository{Db: db}
	rs := &services.RaceService{Repo: rr}

	return &usecases.RaceServices{Race: rs}
}
