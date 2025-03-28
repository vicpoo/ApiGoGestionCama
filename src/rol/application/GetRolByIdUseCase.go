// GetRolByIdUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/rol/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/rol/domain/entities"
)

type GetRolByIdUseCase struct {
	repo repositories.IRol
}

func NewGetRolByIdUseCase(repo repositories.IRol) *GetRolByIdUseCase {
	return &GetRolByIdUseCase{repo: repo}
}

func (uc *GetRolByIdUseCase) Run(id int32) (*entities.Rol, error) {
	rol, err := uc.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return rol, nil
}
