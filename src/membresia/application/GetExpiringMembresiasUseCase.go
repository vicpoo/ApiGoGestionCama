// GetExpiringMembresiasUseCase.go
package application

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain/entities"
)

type GetExpiringMembresiasUseCase struct {
	repo domain.IMembresia
}

func NewGetExpiringMembresiasUseCase(repo domain.IMembresia) *GetExpiringMembresiasUseCase {
	return &GetExpiringMembresiasUseCase{repo: repo}
}

func (uc *GetExpiringMembresiasUseCase) Run(daysBeforeExpiration int) ([]entities.Membresia, error) {
	return uc.repo.GetExpiringMemberships(daysBeforeExpiration)
}
