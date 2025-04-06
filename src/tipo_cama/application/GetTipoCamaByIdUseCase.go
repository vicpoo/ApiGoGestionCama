// GetTipoCamaByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/domain/entities"
)

type GetTipoCamaByIdUseCase struct {
	repo repositories.ITipoCama
}

func NewGetTipoCamaByIdUseCase(repo repositories.ITipoCama) *GetTipoCamaByIdUseCase {
	return &GetTipoCamaByIdUseCase{repo: repo}
}

func (uc *GetTipoCamaByIdUseCase) Run(id int32) (*entities.TipoCama, error) {
	tipoCama, err := uc.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return tipoCama, nil
}
