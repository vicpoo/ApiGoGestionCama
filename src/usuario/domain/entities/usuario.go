// usuario.go
package entities

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	ID        int32  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"not null"`
	Lastname  string `json:"lastname" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"-" gorm:"not null"`
	RoleID    *int32 `json:"roleId" gorm:"column:role_id"`
	IsPremium bool   `json:"isPremium" gorm:"default:false"`
}

// HashPassword hashes the user password
func (u *Usuario) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword compares input password with stored hash
func (u *Usuario) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// Setters
func (u *Usuario) SetID(id int32) {
	u.ID = id
}

func (u *Usuario) SetName(name string) {
	u.Name = name
}

func (u *Usuario) SetLastname(lastname string) {
	u.Lastname = lastname
}

func (u *Usuario) SetEmail(email string) {
	u.Email = email
}

func (u *Usuario) SetPassword(password string) {
	u.Password = password
}

func (u *Usuario) SetRoleID(roleID *int32) {
	u.RoleID = roleID
}

func (u *Usuario) SetIsPremium(isPremium bool) {
	u.IsPremium = isPremium
}

// Getters
func (u *Usuario) GetID() int32 {
	return u.ID
}

func (u *Usuario) GetName() string {
	return u.Name
}

func (u *Usuario) GetLastname() string {
	return u.Lastname
}

func (u *Usuario) GetEmail() string {
	return u.Email
}

func (u *Usuario) GetPassword() string {
	return u.Password
}

func (u *Usuario) GetRoleID() *int32 {
	return u.RoleID
}

func (u *Usuario) GetIsPremium() bool {
	return u.IsPremium
}

// ToJSON returns JSON representation without password
func (u *Usuario) ToJSON() ([]byte, error) {
	type Alias Usuario
	return json.Marshal(&struct {
		*Alias
		Password string `json:"password,omitempty"`
	}{
		Alias:    (*Alias)(u),
		Password: "",
	})
}

func NewUsuario(name, lastname, email, password string, roleID *int32, isPremium bool) *Usuario {
	return &Usuario{
		Name:      name,
		Lastname:  lastname,
		Email:     email,
		Password:  password,
		RoleID:    roleID,
		IsPremium: isPremium,
	}
}
