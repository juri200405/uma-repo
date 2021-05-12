package registry

import (
	"github.com/juri200405/uma-repo/app/internal/infrastructures/dao/mysql"
	"github.com/juri200405/uma-repo/app/internal/infrastructures/datastore"
	"github.com/juri200405/uma-repo/app/internal/domain/services"
	"github.com/juri200405/uma-repo/app/internal/application/usecases"
)

type FactorRegistry struct {
}

func NewFactorRegistry() *FactorRegistry {
	return &FactorRegistry{}
}

func (r FactorRegistry) GetFactorUsecase() usecases.FactorUsecase {
	db := mysql.GetConnection()

	fr := &datastore.FactorRepository{Db: db}
	fs := &services.FactorService{Repo: fr}

	return &usecases.FactorServices{Factor: fs}
}
