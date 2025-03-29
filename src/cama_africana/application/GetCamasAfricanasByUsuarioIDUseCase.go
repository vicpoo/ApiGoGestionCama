package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain/entities"
)

type GetCamasAfricanasByUsuarioIDUseCase struct {
	repo repositories.ICamaAfricana
}

func NewGetCamasAfricanasByUsuarioIDUseCase(repo repositories.ICamaAfricana) *GetCamasAfricanasByUsuarioIDUseCase {
	return &GetCamasAfricanasByUsuarioIDUseCase{repo: repo}
}

func (uc *GetCamasAfricanasByUsuarioIDUseCase) Run(usuarioID int32) ([]entities.CamaAfricana, error) {
	camas, err := uc.repo.GetByUsuarioID(usuarioID)
	if err != nil {
		return nil, err
	}
	return camas, nil
}
