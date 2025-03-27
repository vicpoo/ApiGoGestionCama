package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain/entities"
)

type UpdateUsuarioUseCase struct {
	repo repositories.IUsuario
}

func NewUpdateUsuarioUseCase(repo repositories.IUsuario) *UpdateUsuarioUseCase {
	return &UpdateUsuarioUseCase{repo: repo}
}

func (uc *UpdateUsuarioUseCase) Run(usuario *entities.Usuario) (*entities.Usuario, error) {
	err := uc.repo.Update(usuario)
	if err != nil {
		return nil, err
	}
	return usuario, nil
}
