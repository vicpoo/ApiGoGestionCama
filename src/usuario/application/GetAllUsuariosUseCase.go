// GetAllUsuariosUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain/entities"
)

type GetAllUsuariosUseCase struct {
	repo repositories.IUsuario
}

func NewGetAllUsuariosUseCase(repo repositories.IUsuario) *GetAllUsuariosUseCase {
	return &GetAllUsuariosUseCase{repo: repo}
}

func (uc *GetAllUsuariosUseCase) Run() ([]entities.Usuario, error) {
	usuarios, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}
