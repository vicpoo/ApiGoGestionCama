// almacen_repository.go
package domain

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/domain/entities"
)

type IAlmacen interface {
	Save(almacen *entities.Almacen) error
	Update(almacen *entities.Almacen) error
	Delete(id int32) error
	GetById(id int32) (*entities.Almacen, error)
	GetAll() ([]entities.Almacen, error)
	IncrementarCantidad(id int32, cantidad int32) error
}
