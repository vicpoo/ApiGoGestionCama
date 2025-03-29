package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/cama/domain/entities"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/core"
)

type MySQLCamaRepository struct {
	conn *sql.DB
}

func NewMySQLCamaRepository() repositories.ICamaRepository {
	conn := core.GetDB()
	return &MySQLCamaRepository{conn: conn}
}

func (mysql *MySQLCamaRepository) Save(cama *entities.Cama) error {
	query := `
		INSERT INTO Cama (modelo, tipo_id, usuario_id)
		VALUES (?, ?, ?)
	`
	result, err := mysql.conn.Exec(
		query,
		cama.Modelo,
		cama.TipoID,
		cama.UsuarioID,
	)
	if err != nil {
		log.Println("Error al guardar la cama:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	cama.ID = int32(id)

	return nil
}

func (mysql *MySQLCamaRepository) Update(cama *entities.Cama) error {
	query := `
		UPDATE Cama
		SET modelo = ?, tipo_id = ?, usuario_id = ?
		WHERE id = ?
	`
	result, err := mysql.conn.Exec(
		query,
		cama.Modelo,
		cama.TipoID,
		cama.UsuarioID,
		cama.ID,
	)
	if err != nil {
		log.Println("Error al actualizar la cama:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("cama con ID %d no encontrada", cama.ID)
	}

	return nil
}

func (mysql *MySQLCamaRepository) Delete(id int32) error {
	query := "DELETE FROM Cama WHERE id = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar la cama:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("cama con ID %d no encontrada", id)
	}

	return nil
}

func (mysql *MySQLCamaRepository) GetById(id int32) (*entities.Cama, error) {
	query := `
		SELECT id, modelo, tipo_id, usuario_id
		FROM Cama
		WHERE id = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var cama entities.Cama
	err := row.Scan(
		&cama.ID,
		&cama.Modelo,
		&cama.TipoID,
		&cama.UsuarioID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("cama con ID %d no encontrada", id)
		}
		log.Println("Error al buscar la cama por ID:", err)
		return nil, err
	}

	return &cama, nil
}

func (mysql *MySQLCamaRepository) GetAll() ([]entities.Cama, error) {
	query := `
		SELECT id, modelo, tipo_id, usuario_id
		FROM Cama
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todas las camas:", err)
		return nil, err
	}
	defer rows.Close()

	var camas []entities.Cama
	for rows.Next() {
		var cama entities.Cama
		err := rows.Scan(
			&cama.ID,
			&cama.Modelo,
			&cama.TipoID,
			&cama.UsuarioID,
		)
		if err != nil {
			log.Println("Error al escanear la cama:", err)
			return nil, err
		}
		camas = append(camas, cama)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return camas, nil
}

func (mysql *MySQLCamaRepository) GetByUsuarioID(usuarioID int32) ([]entities.Cama, error) {
	query := `
		SELECT id, modelo, tipo_id, usuario_id
		FROM Cama
		WHERE usuario_id = ?
	`
	rows, err := mysql.conn.Query(query, usuarioID)
	if err != nil {
		log.Println("Error al obtener camas por usuario ID:", err)
		return nil, err
	}
	defer rows.Close()

	var camas []entities.Cama
	for rows.Next() {
		var cama entities.Cama
		err := rows.Scan(
			&cama.ID,
			&cama.Modelo,
			&cama.TipoID,
			&cama.UsuarioID,
		)
		if err != nil {
			log.Println("Error al escanear la cama:", err)
			return nil, err
		}
		camas = append(camas, cama)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return camas, nil
}

func (mysql *MySQLCamaRepository) GetByTipoID(tipoID int32) ([]entities.Cama, error) {
	query := `
		SELECT id, modelo, tipo_id, usuario_id
		FROM Cama
		WHERE tipo_id = ?
	`
	rows, err := mysql.conn.Query(query, tipoID)
	if err != nil {
		log.Println("Error al obtener camas por tipo ID:", err)
		return nil, err
	}
	defer rows.Close()

	var camas []entities.Cama
	for rows.Next() {
		var cama entities.Cama
		err := rows.Scan(
			&cama.ID,
			&cama.Modelo,
			&cama.TipoID,
			&cama.UsuarioID,
		)
		if err != nil {
			log.Println("Error al escanear la cama:", err)
			return nil, err
		}
		camas = append(camas, cama)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return camas, nil
}
