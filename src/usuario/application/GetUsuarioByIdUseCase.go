package application

import (
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain/entities"
)

type GetUsuarioByIdUseCase struct {
	repo repositories.IUsuario
}

func NewGetUsuarioByIdUseCase(repo repositories.IUsuario) *GetUsuarioByIdUseCase {
	return &GetUsuarioByIdUseCase{repo: repo}
}

func (uc *GetUsuarioByIdUseCase) Run(id int32) (*entities.Usuario, error) {
	usuario, err := uc.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return usuario, nil
}
