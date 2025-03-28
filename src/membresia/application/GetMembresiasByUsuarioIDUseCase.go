// GetMembresiasByUsuarioIDUseCase.go
package application

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain/entities"
)

type GetMembresiasByUsuarioIDUseCase struct {
	repo domain.IMembresia
}

func NewGetMembresiasByUsuarioIDUseCase(repo domain.IMembresia) *GetMembresiasByUsuarioIDUseCase {
	return &GetMembresiasByUsuarioIDUseCase{repo: repo}
}

func (uc *GetMembresiasByUsuarioIDUseCase) Run(usuarioID int32) ([]entities.Membresia, error) {
	return uc.repo.GetByUsuarioID(usuarioID)
}
