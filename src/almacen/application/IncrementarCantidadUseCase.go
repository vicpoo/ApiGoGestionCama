package application

import repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/domain"

type IncrementarCantidadUseCase struct {
	repo repositories.IAlmacen
}

func NewIncrementarCantidadUseCase(repo repositories.IAlmacen) *IncrementarCantidadUseCase {
	return &IncrementarCantidadUseCase{repo: repo}
}

func (uc *IncrementarCantidadUseCase) Run(id int32, cantidad int32) error {
	err := uc.repo.IncrementarCantidad(id, cantidad)
	if err != nil {
		return err
	}
	return nil
}
