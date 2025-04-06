// GetAlmacenByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/domain/entities"
)

type GetAlmacenByIdUseCase struct {
	repo repositories.IAlmacen
}

func NewGetAlmacenByIdUseCase(repo repositories.IAlmacen) *GetAlmacenByIdUseCase {
	return &GetAlmacenByIdUseCase{repo: repo}
}

func (uc *GetAlmacenByIdUseCase) Run(id int32) (*entities.Almacen, error) {
	almacen, err := uc.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return almacen, nil
}
