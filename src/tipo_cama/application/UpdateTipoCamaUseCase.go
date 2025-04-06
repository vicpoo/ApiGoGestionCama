// UpdateTipoCamaUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/domain/entities"
)

type UpdateTipoCamaUseCase struct {
	repo repositories.ITipoCama
}

func NewUpdateTipoCamaUseCase(repo repositories.ITipoCama) *UpdateTipoCamaUseCase {
	return &UpdateTipoCamaUseCase{repo: repo}
}

func (uc *UpdateTipoCamaUseCase) Run(tipoCama *entities.TipoCama) (*entities.TipoCama, error) {
	err := uc.repo.Update(tipoCama)
	if err != nil {
		return nil, err
	}
	return tipoCama, nil
}
