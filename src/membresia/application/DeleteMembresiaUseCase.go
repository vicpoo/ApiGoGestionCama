// DeleteMembresiaUseCase.go
package application

import "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain"

type DeleteMembresiaUseCase struct {
	repo domain.IMembresia
}

func NewDeleteMembresiaUseCase(repo domain.IMembresia) *DeleteMembresiaUseCase {
	return &DeleteMembresiaUseCase{repo: repo}
}

func (uc *DeleteMembresiaUseCase) Run(id int32) error {
	return uc.repo.Delete(id)
}
