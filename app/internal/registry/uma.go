package registry

import (
	"github.com/juri200405/uma-repo/app/internal/infrastructures/dao/mysql"
	"github.com/juri200405/uma-repo/app/internal/infrastructures/datastore"
	"github.com/juri200405/uma-repo/app/internal/domain/services"
	"github.com/juri200405/uma-repo/app/internal/application/usecases"
)

type UmaRegistry struct {
}

func NewUmaRegistry() *UmaRegistry {
	return &UmaRegistry{}
}

func (r UmaRegistry) GetUmaUsecase() usecases.UmaUsecase {
	db := mysql.GetConnection()

	ur := &datastore.UmaRepository{Db: db}
	fr := &datastore.FactorRepository{Db: db}
	us := &services.UmaService{Repo: ur}
	fs := &services.FactorService{Repo: fr}

	return &usecases.UmaServices{Uma: us, Factor: fs}
}
