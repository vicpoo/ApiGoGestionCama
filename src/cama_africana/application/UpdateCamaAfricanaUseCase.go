package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain/entities"
)

type UpdateCamaAfricanaUseCase struct {
	repo repositories.ICamaAfricana
}

func NewUpdateCamaAfricanaUseCase(repo repositories.ICamaAfricana) *UpdateCamaAfricanaUseCase {
	return &UpdateCamaAfricanaUseCase{repo: repo}
}

func (uc *UpdateCamaAfricanaUseCase) Run(camaAfricana *entities.CamaAfricana) (*entities.CamaAfricana, error) {
	err := uc.repo.Update(camaAfricana)
	if err != nil {
		return nil, err
	}
	return camaAfricana, nil
}
