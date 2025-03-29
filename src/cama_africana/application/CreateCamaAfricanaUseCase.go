package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain/entities"
)

type CreateCamaAfricanaUseCase struct {
	repo repositories.ICamaAfricana
}

func NewCreateCamaAfricanaUseCase(repo repositories.ICamaAfricana) *CreateCamaAfricanaUseCase {
	return &CreateCamaAfricanaUseCase{repo: repo}
}

func (uc *CreateCamaAfricanaUseCase) Run(camaAfricana *entities.CamaAfricana) (*entities.CamaAfricana, error) {
	err := uc.repo.Save(camaAfricana)
	if err != nil {
		return nil, err
	}
	return camaAfricana, nil
}
