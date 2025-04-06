// CreateAlmacenUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/domain/entities"
)

type CreateAlmacenUseCase struct {
	repo repositories.IAlmacen
}

func NewCreateAlmacenUseCase(repo repositories.IAlmacen) *CreateAlmacenUseCase {
	return &CreateAlmacenUseCase{repo: repo}
}

func (uc *CreateAlmacenUseCase) Run(almacen *entities.Almacen) (*entities.Almacen, error) {
	err := uc.repo.Save(almacen)
	if err != nil {
		return nil, err
	}
	return almacen, nil
}
