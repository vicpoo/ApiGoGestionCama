package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain/entities"
)

type GetCamasByUsuarioIDUseCase struct {
	repo repositories.ICamaRepository
}

func NewGetCamasByUsuarioIDUseCase(repo repositories.ICamaRepository) *GetCamasByUsuarioIDUseCase {
	return &GetCamasByUsuarioIDUseCase{repo: repo}
}

func (uc *GetCamasByUsuarioIDUseCase) Run(usuarioID int32) ([]entities.Cama, error) {
	camas, err := uc.repo.GetByUsuarioID(usuarioID)
	if err != nil {
		return nil, err
	}
	return camas, nil
}
