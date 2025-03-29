package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain/entities"
)

type CreateCamaUseCase struct {
	repo repositories.ICamaRepository
}

func NewCreateCamaUseCase(repo repositories.ICamaRepository) *CreateCamaUseCase {
	return &CreateCamaUseCase{repo: repo}
}

func (uc *CreateCamaUseCase) Run(cama *entities.Cama) (*entities.Cama, error) {
	err := uc.repo.Save(cama)
	if err != nil {
		return nil, err
	}
	return cama, nil
}
