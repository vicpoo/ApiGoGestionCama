// CreateUsuarioUseCase.go
package application

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain/entities"
)

type CreateUsuarioUseCase struct {
	repo domain.IUsuario
}

func NewCreateUsuarioUseCase(repo domain.IUsuario) *CreateUsuarioUseCase {
	return &CreateUsuarioUseCase{repo: repo}
}

func (uc *CreateUsuarioUseCase) Run(usuario *entities.Usuario) (*entities.Usuario, error) {
	err := uc.repo.Save(usuario)
	if err != nil {
		return nil, err
	}
	return usuario, nil
}
