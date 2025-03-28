// GetUsuarioByEmailUseCase.go
package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain/entities"
)

type GetUsuarioByEmailUseCase struct {
	repo repositories.IUsuario
}

func NewGetUsuarioByEmailUseCase(repo repositories.IUsuario) *GetUsuarioByEmailUseCase {
	return &GetUsuarioByEmailUseCase{repo: repo}
}

func (uc *GetUsuarioByEmailUseCase) Run(email string) (*entities.Usuario, error) {
	usuario, err := uc.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return usuario, nil
}
