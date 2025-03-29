// almacen.go
package entities

type Almacen struct {
	ID         int32 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	TipoCamaID int32 `json:"tipo_cama_id" gorm:"column:tipo_cama_id"`
	Cantidad   int32 `json:"cantidad" gorm:"column:cantidad;not null;default:0"`
}

// Setters
func (a *Almacen) SetID(id int32) {
	a.ID = id
}

func (a *Almacen) SetTipoCamaID(tipoCamaID int32) {
	a.TipoCamaID = tipoCamaID
}

func (a *Almacen) SetCantidad(cantidad int32) {
	a.Cantidad = cantidad
}

// Getters
func (a *Almacen) GetID() int32 {
	return a.ID
}

func (a *Almacen) GetTipoCamaID() int32 {
	return a.TipoCamaID
}

func (a *Almacen) GetCantidad() int32 {
	return a.Cantidad
}

func NewAlmacen(tipoCamaID int32, cantidad int32) *Almacen {
	return &Almacen{
		TipoCamaID: tipoCamaID,
		Cantidad:   cantidad,
	}
}
