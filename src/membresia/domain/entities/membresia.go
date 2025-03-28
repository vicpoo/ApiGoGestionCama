// membresia.go
package entities

import (
	"encoding/json"
	"time"
)

type Membresia struct {
	ID          int32     `json:"id" gorm:"primaryKey;autoIncrement"`
	UsuarioID   int32     `json:"usuarioId" gorm:"column:usuario_id;not null"`
	FechaInicio time.Time `json:"fechaInicio" gorm:"type:date;not null"`
	FechaFin    time.Time `json:"fechaFin" gorm:"type:date;not null"`
	Estado      int       `json:"estado" gorm:"type:tinyint(1);default:1"` // Cambiado a int para 1/0
}

// Setters
func (m *Membresia) SetID(id int32) {
	m.ID = id
}

func (m *Membresia) SetUsuarioID(usuarioID int32) {
	m.UsuarioID = usuarioID
}

func (m *Membresia) SetFechaInicio(fechaInicio time.Time) {
	m.FechaInicio = fechaInicio
}

func (m *Membresia) SetFechaFin(fechaFin time.Time) {
	m.FechaFin = fechaFin
}

func (m *Membresia) SetEstado(estado int) { // Cambiado a int
	m.Estado = estado
}

// Getters
func (m *Membresia) GetID() int32 {
	return m.ID
}

func (m *Membresia) GetUsuarioID() int32 {
	return m.UsuarioID
}

func (m *Membresia) GetFechaInicio() time.Time {
	return m.FechaInicio
}

func (m *Membresia) GetFechaFin() time.Time {
	return m.FechaFin
}

func (m *Membresia) GetEstado() int { // Cambiado a int
	return m.Estado
}

// ToJSON returns JSON representation
func (m *Membresia) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// NewMembresia creates a new Membresia instance
func NewMembresia(usuarioID int32, fechaInicio, fechaFin time.Time, estado int) *Membresia { // Cambiado a int
	return &Membresia{
		UsuarioID:   usuarioID,
		FechaInicio: fechaInicio,
		FechaFin:    fechaFin,
		Estado:      estado,
	}
}

// IsActive checks if the membership is currently active
func (m *Membresia) IsActive() bool {
	now := time.Now()
	return m.Estado == 1 && now.After(m.FechaInicio) && now.Before(m.FechaFin) // Modificado para comparar con 1
}
