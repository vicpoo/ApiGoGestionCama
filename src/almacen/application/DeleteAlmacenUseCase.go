// DeleteAlmacenUseCase.go
package application

import repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/domain"

type DeleteAlmacenUseCase struct {
	repo repositories.IAlmacen
}

func NewDeleteAlmacenUseCase(repo repositories.IAlmacen) *DeleteAlmacenUseCase {
	return &DeleteAlmacenUseCase{repo: repo}
}

func (uc *DeleteAlmacenUseCase) Run(id int32) error {
	err := uc.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
