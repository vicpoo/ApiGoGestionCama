package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain/entities"
)

type UpdateCamaUseCase struct {
	repo repositories.ICamaRepository
}

func NewUpdateCamaUseCase(repo repositories.ICamaRepository) *UpdateCamaUseCase {
	return &UpdateCamaUseCase{repo: repo}
}

func (uc *UpdateCamaUseCase) Run(cama *entities.Cama) (*entities.Cama, error) {
	err := uc.repo.Update(cama)
	if err != nil {
		return nil, err
	}
	return cama, nil
}
