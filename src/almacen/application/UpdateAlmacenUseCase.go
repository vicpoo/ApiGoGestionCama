// UpdateAlmacenUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/domain/entities"
)

type UpdateAlmacenUseCase struct {
	repo repositories.IAlmacen
}

func NewUpdateAlmacenUseCase(repo repositories.IAlmacen) *UpdateAlmacenUseCase {
	return &UpdateAlmacenUseCase{repo: repo}
}

func (uc *UpdateAlmacenUseCase) Run(almacen *entities.Almacen) (*entities.Almacen, error) {
	err := uc.repo.Update(almacen)
	if err != nil {
		return nil, err
	}
	return almacen, nil
}
