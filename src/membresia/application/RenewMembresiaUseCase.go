// RenewMembresiaUseCase.go
package application

import (
	"time"

	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain"
)

type RenewMembresiaUseCase struct {
	repo domain.IMembresia
}

func NewRenewMembresiaUseCase(repo domain.IMembresia) *RenewMembresiaUseCase {
	return &RenewMembresiaUseCase{repo: repo}
}

func (uc *RenewMembresiaUseCase) Run(membresiaID int32, newEndDate time.Time) error {
	return uc.repo.RenewMembership(membresiaID, newEndDate)
}
