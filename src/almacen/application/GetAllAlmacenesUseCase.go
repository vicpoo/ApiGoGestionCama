package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/domain/entities"
)

type GetAllAlmacenesUseCase struct {
	repo repositories.IAlmacen
}

func NewGetAllAlmacenesUseCase(repo repositories.IAlmacen) *GetAllAlmacenesUseCase {
	return &GetAllAlmacenesUseCase{repo: repo}
}

func (uc *GetAllAlmacenesUseCase) Run() ([]entities.Almacen, error) {
	almacenes, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return almacenes, nil
}
