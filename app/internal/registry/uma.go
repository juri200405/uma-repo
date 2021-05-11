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
	us := &services.UmaService{Repo: ur}

	return &usecases.UmaServices{Uma: us}
}
