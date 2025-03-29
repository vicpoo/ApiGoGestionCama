package domain

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain/entities"
)

type ICamaRepository interface {
	Save(cama *entities.Cama) error
	Update(cama *entities.Cama) error
	Delete(id int32) error
	GetById(id int32) (*entities.Cama, error)
	GetAll() ([]entities.Cama, error)
	GetByUsuarioID(usuarioID int32) ([]entities.Cama, error) // Nuevo método específico
	GetByTipoID(tipoID int32) ([]entities.Cama, error)       // Nuevo método específico
}
