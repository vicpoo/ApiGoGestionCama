package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain/entities"
)

type GetCamaAfricanaByCamaIDUseCase struct {
	repo repositories.ICamaAfricana
}

func NewGetCamaAfricanaByCamaIDUseCase(repo repositories.ICamaAfricana) *GetCamaAfricanaByCamaIDUseCase {
	return &GetCamaAfricanaByCamaIDUseCase{repo: repo}
}

func (uc *GetCamaAfricanaByCamaIDUseCase) Run(camaID int32) (*entities.CamaAfricana, error) {
	cama, err := uc.repo.GetByCamaID(camaID)
	if err != nil {
		return nil, err
	}
	return cama, nil
}
