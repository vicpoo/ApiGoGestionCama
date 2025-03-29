package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama_africana/domain/entities"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/core"
)

type MySQLCamaAfricanaRepository struct {
	conn *sql.DB
}

func NewMySQLCamaAfricanaRepository() repositories.ICamaAfricana {
	conn := core.GetDB()
	return &MySQLCamaAfricanaRepository{conn: conn}
}

func (mysql *MySQLCamaAfricanaRepository) Save(cama *entities.CamaAfricana) error {
	query := `
		INSERT INTO Cama_Africana (cama_id, usuario_id)
		VALUES (?, ?)
	`
	result, err := mysql.conn.Exec(
		query,
		cama.CamaID,
		cama.UsuarioID,
	)
	if err != nil {
		log.Println("Error al guardar la cama africana:", err)
		return err
	}

	// Obtener el ID generado
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	cama.ID = int32(id)

	return nil
}

func (mysql *MySQLCamaAfricanaRepository) Update(cama *entities.CamaAfricana) error {
	query := `
		UPDATE Cama_Africana
		SET cama_id = ?, usuario_id = ?
		WHERE id = ?
	`
	result, err := mysql.conn.Exec(
		query,
		cama.CamaID,
		cama.UsuarioID,
		cama.ID,
	)
	if err != nil {
		log.Println("Error al actualizar la cama africana:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("cama africana con ID %d no encontrada", cama.ID)
	}

	return nil
}

func (mysql *MySQLCamaAfricanaRepository) Delete(id int32) error {
	query := "DELETE FROM Cama_Africana WHERE id = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar la cama africana:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("cama africana con ID %d no encontrada", id)
	}

	return nil
}

func (mysql *MySQLCamaAfricanaRepository) GetById(id int32) (*entities.CamaAfricana, error) {
	query := `
		SELECT id, cama_id, usuario_id
		FROM Cama_Africana
		WHERE id = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var cama entities.CamaAfricana
	var usuarioID sql.NullInt32 // Para manejar valores NULL

	err := row.Scan(
		&cama.ID,
		&cama.CamaID,
		&usuarioID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("cama africana con ID %d no encontrada", id)
		}
		log.Println("Error al buscar la cama africana por ID:", err)
		return nil, err
	}

	// Convertir sql.NullInt32 a *int32
	if usuarioID.Valid {
		cama.UsuarioID = &usuarioID.Int32
	} else {
		cama.UsuarioID = nil
	}

	return &cama, nil
}

func (mysql *MySQLCamaAfricanaRepository) GetAll() ([]entities.CamaAfricana, error) {
	query := `
		SELECT id, cama_id, usuario_id
		FROM Cama_Africana
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todas las camas africanas:", err)
		return nil, err
	}
	defer rows.Close()

	var camas []entities.CamaAfricana
	for rows.Next() {
		var cama entities.CamaAfricana
		var usuarioID sql.NullInt32

		err := rows.Scan(
			&cama.ID,
			&cama.CamaID,
			&usuarioID,
		)
		if err != nil {
			log.Println("Error al escanear la cama africana:", err)
			return nil, err
		}

		if usuarioID.Valid {
			cama.UsuarioID = &usuarioID.Int32
		} else {
			cama.UsuarioID = nil
		}

		camas = append(camas, cama)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return camas, nil
}

func (mysql *MySQLCamaAfricanaRepository) GetByCamaID(camaID int32) (*entities.CamaAfricana, error) {
	query := `
		SELECT id, cama_id, usuario_id
		FROM Cama_Africana
		WHERE cama_id = ?
	`
	row := mysql.conn.QueryRow(query, camaID)

	var cama entities.CamaAfricana
	var usuarioID sql.NullInt32

	err := row.Scan(
		&cama.ID,
		&cama.CamaID,
		&usuarioID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("cama africana con cama_id %d no encontrada", camaID)
		}
		log.Println("Error al buscar la cama africana por cama_id:", err)
		return nil, err
	}

	if usuarioID.Valid {
		cama.UsuarioID = &usuarioID.Int32
	} else {
		cama.UsuarioID = nil
	}

	return &cama, nil
}

func (mysql *MySQLCamaAfricanaRepository) GetByUsuarioID(usuarioID int32) ([]entities.CamaAfricana, error) {
	query := `
		SELECT id, cama_id, usuario_id
		FROM Cama_Africana
		WHERE usuario_id = ?
	`
	rows, err := mysql.conn.Query(query, usuarioID)
	if err != nil {
		log.Println("Error al obtener camas africanas por usuario_id:", err)
		return nil, err
	}
	defer rows.Close()

	var camas []entities.CamaAfricana
	for rows.Next() {
		var cama entities.CamaAfricana
		var usuarioID sql.NullInt32

		err := rows.Scan(
			&cama.ID,
			&cama.CamaID,
			&usuarioID,
		)
		if err != nil {
			log.Println("Error al escanear la cama africana:", err)
			return nil, err
		}

		if usuarioID.Valid {
			cama.UsuarioID = &usuarioID.Int32
		} else {
			cama.UsuarioID = nil
		}

		camas = append(camas, cama)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return camas, nil
}

func (mysql *MySQLCamaAfricanaRepository) AssignCamaToUsuario(camaID int32, usuarioID *int32) (*entities.CamaAfricana, error) {
	// Primero obtenemos la cama existente
	cama, err := mysql.GetByCamaID(camaID)
	if err != nil {
		return nil, err
	}

	// Actualizamos el usuario_id
	query := `
		UPDATE Cama_Africana
		SET usuario_id = ?
		WHERE id = ?
	`
	var usrID interface{}
	if usuarioID != nil {
		usrID = *usuarioID
	} else {
		usrID = nil
	}

	result, err := mysql.conn.Exec(
		query,
		usrID,
		cama.ID,
	)
	if err != nil {
		log.Println("Error al asignar cama a usuario:", err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("no se pudo asignar la cama al usuario")
	}

	// Obtenemos la cama actualizada para retornarla
	return mysql.GetByCamaID(camaID)
}
