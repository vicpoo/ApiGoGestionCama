package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain/entities"
)

type GetAllCamasAfricanasUseCase struct {
	repo repositories.ICamaAfricana
}

func NewGetAllCamasAfricanasUseCase(repo repositories.ICamaAfricana) *GetAllCamasAfricanasUseCase {
	return &GetAllCamasAfricanasUseCase{repo: repo}
}

func (uc *GetAllCamasAfricanasUseCase) Run() ([]entities.CamaAfricana, error) {
	camas, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return camas, nil
}
