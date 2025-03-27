package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/rol/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/rol/domain/entities"
)

type GetAllRolesUseCase struct {
	repo repositories.IRol
}

func NewGetAllRolesUseCase(repo repositories.IRol) *GetAllRolesUseCase {
	return &GetAllRolesUseCase{repo: repo}
}

func (uc *GetAllRolesUseCase) Run() ([]entities.Rol, error) {
	roles, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return roles, nil
}
