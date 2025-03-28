// DeleteUsuarioUseCase.go
package application

import repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain"

type DeleteUsuarioUseCase struct {
	repo repositories.IUsuario
}

func NewDeleteUsuarioUseCase(repo repositories.IUsuario) *DeleteUsuarioUseCase {
	return &DeleteUsuarioUseCase{repo: repo}
}

func (uc *DeleteUsuarioUseCase) Run(id int32) error {
	err := uc.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
