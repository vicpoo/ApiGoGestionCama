// mysql_usuario_repository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/core"
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/usuario/domain/entities"
)

type MySQLUsuarioRepository struct {
	conn *sql.DB
}

func NewMySQLUsuarioRepository() repositories.IUsuario {
	conn := core.GetDB()
	return &MySQLUsuarioRepository{conn: conn}
}

func (mysql *MySQLUsuarioRepository) Save(usuario *entities.Usuario) error {
	// Hash password before saving
	if err := usuario.HashPassword(); err != nil {
		log.Println("Error al hashear la contraseña:", err)
		return err
	}

	query := `
		INSERT INTO Usuario (name, lastname, email, password, role_id, is_premium)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := mysql.conn.Exec(
		query,
		usuario.Name,
		usuario.Lastname,
		usuario.Email,
		usuario.Password,
		usuario.RoleID,
		usuario.IsPremium,
	)
	if err != nil {
		log.Println("Error al guardar el usuario:", err)
		return err
	}

	// Get the inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID:", err)
		return err
	}
	usuario.ID = int32(id)

	return nil
}

func (mysql *MySQLUsuarioRepository) Update(usuario *entities.Usuario) error {
	// Only hash password if it's being updated
	if usuario.Password != "" {
		if err := usuario.HashPassword(); err != nil {
			log.Println("Error al hashear la contraseña:", err)
			return err
		}
	}

	query := `
		UPDATE Usuario
		SET name = ?, lastname = ?, email = ?, 
		    password = COALESCE(?, password), 
		    role_id = ?, is_premium = ?
		WHERE id = ?
	`
	result, err := mysql.conn.Exec(
		query,
		usuario.Name,
		usuario.Lastname,
		usuario.Email,
		usuario.Password,
		usuario.RoleID,
		usuario.IsPremium,
		usuario.ID,
	)
	if err != nil {
		log.Println("Error al actualizar el usuario:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario con ID %d no encontrado", usuario.ID)
	}

	return nil
}

func (mysql *MySQLUsuarioRepository) Delete(id int32) error {
	query := "DELETE FROM Usuario WHERE id = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el usuario:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MySQLUsuarioRepository) GetById(id int32) (*entities.Usuario, error) {
	query := `
		SELECT id, name, lastname, email, password, role_id, is_premium
		FROM Usuario
		WHERE id = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var usuario entities.Usuario
	err := row.Scan(
		&usuario.ID,
		&usuario.Name,
		&usuario.Lastname,
		&usuario.Email,
		&usuario.Password,
		&usuario.RoleID,
		&usuario.IsPremium,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el usuario por ID:", err)
		return nil, err
	}

	return &usuario, nil
}

func (mysql *MySQLUsuarioRepository) GetByEmail(email string) (*entities.Usuario, error) {
	query := `
		SELECT id, name, lastname, email, password, role_id, is_premium
		FROM Usuario
		WHERE email = ?
	`
	row := mysql.conn.QueryRow(query, email)

	var usuario entities.Usuario
	err := row.Scan(
		&usuario.ID,
		&usuario.Name,
		&usuario.Lastname,
		&usuario.Email,
		&usuario.Password,
		&usuario.RoleID,
		&usuario.IsPremium,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario con email %s no encontrado", email)
		}
		log.Println("Error al buscar el usuario por email:", err)
		return nil, err
	}

	return &usuario, nil
}

func (mysql *MySQLUsuarioRepository) GetAll() ([]entities.Usuario, error) {
	query := `
		SELECT id, name, lastname, email, password, role_id, is_premium
		FROM Usuario
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los usuarios:", err)
		return nil, err
	}
	defer rows.Close()

	var usuarios []entities.Usuario
	for rows.Next() {
		var usuario entities.Usuario
		err := rows.Scan(
			&usuario.ID,
			&usuario.Name,
			&usuario.Lastname,
			&usuario.Email,
			&usuario.Password,
			&usuario.RoleID,
			&usuario.IsPremium,
		)
		if err != nil {
			log.Println("Error al escanear el usuario:", err)
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return usuarios, nil
}
