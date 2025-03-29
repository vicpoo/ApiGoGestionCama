// cama_africana_repository.go
package domain

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain/entities"
)

type ICamaAfricana interface {
	Save(camaAfricana *entities.CamaAfricana) error
	Update(camaAfricana *entities.CamaAfricana) error
	Delete(id int32) error
	GetById(id int32) (*entities.CamaAfricana, error)
	GetAll() ([]entities.CamaAfricana, error)

	// Métodos adicionales específicos para CamaAfricana
	GetByCamaID(camaID int32) (*entities.CamaAfricana, error)
	GetByUsuarioID(usuarioID int32) ([]entities.CamaAfricana, error)

	// Nuevo método para asignar cama a usuario
	AssignCamaToUsuario(camaID int32, usuarioID *int32) (*entities.CamaAfricana, error)
}
