// auth_usecase.go
package application

import (
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/core/auth"
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain/entities"
)

type AuthUseCase struct {
	repo repositories.IUsuario
}

func NewAuthUseCase(repo repositories.IUsuario) *AuthUseCase {
	return &AuthUseCase{repo: repo}
}

func (uc *AuthUseCase) Login(email, password string) (string, *entities.Usuario, error) {
	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		return "", nil, err
	}

	if err := user.CheckPassword(password); err != nil {
		return "", nil, err
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
