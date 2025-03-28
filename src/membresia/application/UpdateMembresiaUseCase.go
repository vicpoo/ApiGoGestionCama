// UpdateMembresiaUseCase.go
package application

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain/entities"
)

type UpdateMembresiaUseCase struct {
	repo domain.IMembresia
}

func NewUpdateMembresiaUseCase(repo domain.IMembresia) *UpdateMembresiaUseCase {
	return &UpdateMembresiaUseCase{repo: repo}
}

func (uc *UpdateMembresiaUseCase) Run(membresia *entities.Membresia) (*entities.Membresia, error) {
	err := uc.repo.Update(membresia)
	if err != nil {
		return nil, err
	}
	return membresia, nil
}
