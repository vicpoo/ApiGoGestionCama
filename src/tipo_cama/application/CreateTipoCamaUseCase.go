package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/domain/entities"
)

type CreateTipoCamaUseCase struct {
	repo repositories.ITipoCama
}

func NewCreateTipoCamaUseCase(repo repositories.ITipoCama) *CreateTipoCamaUseCase {
	return &CreateTipoCamaUseCase{repo: repo}
}

func (uc *CreateTipoCamaUseCase) Run(tipoCama *entities.TipoCama) (*entities.TipoCama, error) {
	err := uc.repo.Save(tipoCama)
	if err != nil {
		return nil, err
	}
	return tipoCama, nil
}
