// membresia_repository.go
package domain

import (
	"time"

	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain/entities"
)

type IMembresia interface {
	Save(membresia *entities.Membresia) error
	Update(membresia *entities.Membresia) error
	Delete(id int32) error
	GetById(id int32) (*entities.Membresia, error)
	GetAll() ([]entities.Membresia, error)

	// Métodos adicionales específicos para membresía
	GetByUsuarioID(usuarioID int32) ([]entities.Membresia, error)
	GetActiveByUsuarioID(usuarioID int32) (*entities.Membresia, error)
	GetExpiringMemberships(daysBeforeExpiration int) ([]entities.Membresia, error)
	RenewMembership(membresiaID int32, newEndDate time.Time) error
}
