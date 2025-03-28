// GetAllMembresiasUseCase.go
package application

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain/entities"
)

type GetAllMembresiasUseCase struct {
	repo domain.IMembresia
}

func NewGetAllMembresiasUseCase(repo domain.IMembresia) *GetAllMembresiasUseCase {
	return &GetAllMembresiasUseCase{repo: repo}
}

func (uc *GetAllMembresiasUseCase) Run() ([]entities.Membresia, error) {
	return uc.repo.GetAll()
}
