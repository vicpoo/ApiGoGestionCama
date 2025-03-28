// CreateMembresiaUseCase.go
package application

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain/entities"
)

type CreateMembresiaUseCase struct {
	repo domain.IMembresia
}

func NewCreateMembresiaUseCase(repo domain.IMembresia) *CreateMembresiaUseCase {
	return &CreateMembresiaUseCase{repo: repo}
}

func (uc *CreateMembresiaUseCase) Run(membresia *entities.Membresia) (*entities.Membresia, error) {
	err := uc.repo.Save(membresia)
	if err != nil {
		return nil, err
	}
	return membresia, nil
}
