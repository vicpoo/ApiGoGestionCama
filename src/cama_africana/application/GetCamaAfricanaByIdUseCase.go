package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain/entities"
)

type GetCamaAfricanaByIdUseCase struct {
	repo repositories.ICamaAfricana
}

func NewGetCamaAfricanaByIdUseCase(repo repositories.ICamaAfricana) *GetCamaAfricanaByIdUseCase {
	return &GetCamaAfricanaByIdUseCase{repo: repo}
}

func (uc *GetCamaAfricanaByIdUseCase) Run(id int32) (*entities.CamaAfricana, error) {
	cama, err := uc.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return cama, nil
}
