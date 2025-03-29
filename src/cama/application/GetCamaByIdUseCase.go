package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain/entities"
)

type GetCamaByIdUseCase struct {
	repo repositories.ICamaRepository
}

func NewGetCamaByIdUseCase(repo repositories.ICamaRepository) *GetCamaByIdUseCase {
	return &GetCamaByIdUseCase{repo: repo}
}

func (uc *GetCamaByIdUseCase) Run(id int32) (*entities.Cama, error) {
	cama, err := uc.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return cama, nil
}
