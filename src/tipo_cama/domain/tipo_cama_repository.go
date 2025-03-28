// tipo_cama_repository.go
package domain

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/domain/entities"
)

type ITipoCama interface {
	Save(tipoCama *entities.TipoCama) error
	Update(tipoCama *entities.TipoCama) error
	Delete(id int32) error
	GetById(id int32) (*entities.TipoCama, error)
	GetAll() ([]entities.TipoCama, error)
}
