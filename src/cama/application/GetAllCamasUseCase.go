// GetAllCamasUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain/entities"
)

type GetAllCamasUseCase struct {
	repo repositories.ICamaRepository
}

func NewGetAllCamasUseCase(repo repositories.ICamaRepository) *GetAllCamasUseCase {
	return &GetAllCamasUseCase{repo: repo}
}

func (uc *GetAllCamasUseCase) Run() ([]entities.Cama, error) {
	camas, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return camas, nil
}
