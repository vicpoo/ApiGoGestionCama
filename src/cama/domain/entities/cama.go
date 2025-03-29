package entities

type Cama struct {
	ID        int32  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Modelo    string `json:"modelo" gorm:"column:modelo;not null"`
	TipoID    int32  `json:"tipo_id" gorm:"column:tipo_id"`
	UsuarioID *int32 `json:"usuario_id,omitempty" gorm:"column:usuario_id"` // Usamos puntero para permitir NULL
}

// Setters
func (c *Cama) SetID(id int32) {
	c.ID = id
}

func (c *Cama) SetModelo(modelo string) {
	c.Modelo = modelo
}

func (c *Cama) SetTipoID(tipoID int32) {
	c.TipoID = tipoID
}

func (c *Cama) SetUsuarioID(usuarioID *int32) {
	c.UsuarioID = usuarioID
}

// Getters
func (c *Cama) GetID() int32 {
	return c.ID
}

func (c *Cama) GetModelo() string {
	return c.Modelo
}

func (c *Cama) GetTipoID() int32 {
	return c.TipoID
}

func (c *Cama) GetUsuarioID() *int32 {
	return c.UsuarioID
}

func NewCama(modelo string, tipoID int32, usuarioID *int32) *Cama {
	return &Cama{
		Modelo:    modelo,
		TipoID:    tipoID,
		UsuarioID: usuarioID,
	}
}
