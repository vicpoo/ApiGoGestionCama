package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain/entities"
)

type GetCamasByTipoIDUseCase struct {
	repo repositories.ICamaRepository
}

func NewGetCamasByTipoIDUseCase(repo repositories.ICamaRepository) *GetCamasByTipoIDUseCase {
	return &GetCamasByTipoIDUseCase{repo: repo}
}

func (uc *GetCamasByTipoIDUseCase) Run(tipoID int32) ([]entities.Cama, error) {
	camas, err := uc.repo.GetByTipoID(tipoID)
	if err != nil {
		return nil, err
	}
	return camas, nil
}
