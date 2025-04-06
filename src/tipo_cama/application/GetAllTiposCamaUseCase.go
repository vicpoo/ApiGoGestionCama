// GetAllTiposCamaUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/domain/entities"
)

type GetAllTiposCamaUseCase struct {
	repo repositories.ITipoCama
}

func NewGetAllTiposCamaUseCase(repo repositories.ITipoCama) *GetAllTiposCamaUseCase {
	return &GetAllTiposCamaUseCase{repo: repo}
}

func (uc *GetAllTiposCamaUseCase) Run() ([]entities.TipoCama, error) {
	tiposCama, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return tiposCama, nil
}
