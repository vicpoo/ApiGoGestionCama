// GetMembresiaByIdUseCase.go
package application

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain/entities"
)

type GetMembresiaByIdUseCase struct {
	repo domain.IMembresia
}

func NewGetMembresiaByIdUseCase(repo domain.IMembresia) *GetMembresiaByIdUseCase {
	return &GetMembresiaByIdUseCase{repo: repo}
}

func (uc *GetMembresiaByIdUseCase) Run(id int32) (*entities.Membresia, error) {
	return uc.repo.GetById(id)
}
