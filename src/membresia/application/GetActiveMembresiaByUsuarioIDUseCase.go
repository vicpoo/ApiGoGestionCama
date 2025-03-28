// GetActiveMembresiaByUsuarioIDUseCase.go
package application

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain/entities"
)

type GetActiveMembresiaByUsuarioIDUseCase struct {
	repo domain.IMembresia
}

func NewGetActiveMembresiaByUsuarioIDUseCase(repo domain.IMembresia) *GetActiveMembresiaByUsuarioIDUseCase {
	return &GetActiveMembresiaByUsuarioIDUseCase{repo: repo}
}

func (uc *GetActiveMembresiaByUsuarioIDUseCase) Run(usuarioID int32) (*entities.Membresia, error) {
	return uc.repo.GetActiveByUsuarioID(usuarioID)
}
