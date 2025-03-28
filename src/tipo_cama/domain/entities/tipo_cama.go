// tipo_cama.go
package entities

type TipoCama struct {
	ID     int32  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Nombre string `json:"nombre" gorm:"column:nombre;not null"`
	Clima  string `json:"clima" gorm:"column:clima;not null"`
}

// Setters
func (t *TipoCama) SetID(id int32) {
	t.ID = id
}

func (t *TipoCama) SetNombre(nombre string) {
	t.Nombre = nombre
}

func (t *TipoCama) SetClima(clima string) {
	t.Clima = clima
}

// Getters
func (t *TipoCama) GetID() int32 {
	return t.ID
}

func (t *TipoCama) GetNombre() string {
	return t.Nombre
}

func (t *TipoCama) GetClima() string {
	return t.Clima
}

func NewTipoCama(nombre, clima string) *TipoCama {
	return &TipoCama{
		Nombre: nombre,
		Clima:  clima,
	}
}
