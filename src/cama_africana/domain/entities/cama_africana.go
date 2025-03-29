// cama_africana.go
package entities

type CamaAfricana struct {
	ID        int32  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	CamaID    int32  `json:"cama_id" gorm:"column:cama_id;unique;not null"`
	UsuarioID *int32 `json:"usuario_id" gorm:"column:usuario_id"`
}

// Setters
func (ca *CamaAfricana) SetID(id int32) {
	ca.ID = id
}

func (ca *CamaAfricana) SetCamaID(camaID int32) {
	ca.CamaID = camaID
}

func (ca *CamaAfricana) SetUsuarioID(usuarioID *int32) {
	ca.UsuarioID = usuarioID
}

// Getters
func (ca *CamaAfricana) GetID() int32 {
	return ca.ID
}

func (ca *CamaAfricana) GetCamaID() int32 {
	return ca.CamaID
}

func (ca *CamaAfricana) GetUsuarioID() *int32 {
	return ca.UsuarioID
}

func NewCamaAfricana(camaID int32, usuarioID *int32) *CamaAfricana {
	return &CamaAfricana{
		CamaID:    camaID,
		UsuarioID: usuarioID,
	}
}
