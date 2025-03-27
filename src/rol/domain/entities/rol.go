// rol.go
package entities

type Rol struct {
	ID          int32  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Titulo      string `json:"titulo" gorm:"column:titulo;not null"`
	Descripcion string `json:"descripcion" gorm:"column:descripcion"`
}

// Setters
func (r *Rol) SetID(id int32) {
	r.ID = id
}

func (r *Rol) SetTitulo(titulo string) {
	r.Titulo = titulo
}

func (r *Rol) SetDescripcion(descripcion string) {
	r.Descripcion = descripcion
}

// Getters
func (r *Rol) GetID() int32 {
	return r.ID
}

func (r *Rol) GetTitulo() string {
	return r.Titulo
}

func (r *Rol) GetDescripcion() string {
	return r.Descripcion
}

func NewRol(titulo, descripcion string) *Rol {
	return &Rol{
		Titulo:      titulo,
		Descripcion: descripcion,
	}
}
