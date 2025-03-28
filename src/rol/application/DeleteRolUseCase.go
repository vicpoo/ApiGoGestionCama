// DeleteRolUseCase.go
package application

import repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/rol/domain"

type DeleteRolUseCase struct {
	repo repositories.IRol
}

func NewDeleteRolUseCase(repo repositories.IRol) *DeleteRolUseCase {
	return &DeleteRolUseCase{repo: repo}
}

func (uc *DeleteRolUseCase) Run(id int32) error {
	err := uc.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
