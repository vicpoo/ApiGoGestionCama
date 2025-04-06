// DeleteTipoCamaUseCase.go
package application

import repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/domain"

type DeleteTipoCamaUseCase struct {
	repo repositories.ITipoCama
}

func NewDeleteTipoCamaUseCase(repo repositories.ITipoCama) *DeleteTipoCamaUseCase {
	return &DeleteTipoCamaUseCase{repo: repo}
}

func (uc *DeleteTipoCamaUseCase) Run(id int32) error {
	err := uc.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
