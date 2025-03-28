package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/core"
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/tipo_cama/domain/entities"
)

type MySQLTipoCamaRepository struct {
	conn *sql.DB
}

func NewMySQLTipoCamaRepository() repositories.ITipoCama {
	conn := core.GetDB()
	return &MySQLTipoCamaRepository{conn: conn}
}

func (mysql *MySQLTipoCamaRepository) Save(tipoCama *entities.TipoCama) error {
	query := `
		INSERT INTO Tipo_Cama (nombre, clima)
		VALUES (?, ?)
	`
	result, err := mysql.conn.Exec(
		query,
		tipoCama.Nombre,
		tipoCama.Clima,
	)
	if err != nil {
		log.Println("Error al guardar el tipo de cama:", err)
		return err
	}

	// Obtener el ID generado
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	tipoCama.ID = int32(id)

	return nil
}

func (mysql *MySQLTipoCamaRepository) Update(tipoCama *entities.TipoCama) error {
	query := `
		UPDATE Tipo_Cama
		SET nombre = ?, clima = ?
		WHERE id = ?
	`
	result, err := mysql.conn.Exec(
		query,
		tipoCama.Nombre,
		tipoCama.Clima,
		tipoCama.ID,
	)
	if err != nil {
		log.Println("Error al actualizar el tipo de cama:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("tipo de cama con ID %d no encontrado", tipoCama.ID)
	}

	return nil
}

func (mysql *MySQLTipoCamaRepository) Delete(id int32) error {
	query := "DELETE FROM Tipo_Cama WHERE id = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el tipo de cama:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("tipo de cama con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MySQLTipoCamaRepository) GetById(id int32) (*entities.TipoCama, error) {
	query := `
		SELECT id, nombre, clima
		FROM Tipo_Cama
		WHERE id = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var tipoCama entities.TipoCama
	err := row.Scan(
		&tipoCama.ID,
		&tipoCama.Nombre,
		&tipoCama.Clima,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tipo de cama con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el tipo de cama por ID:", err)
		return nil, err
	}

	return &tipoCama, nil
}

func (mysql *MySQLTipoCamaRepository) GetAll() ([]entities.TipoCama, error) {
	query := `
		SELECT id, nombre, clima
		FROM Tipo_Cama
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los tipos de cama:", err)
		return nil, err
	}
	defer rows.Close()

	var tiposCama []entities.TipoCama
	for rows.Next() {
		var tipoCama entities.TipoCama
		err := rows.Scan(
			&tipoCama.ID,
			&tipoCama.Nombre,
			&tipoCama.Clima,
		)
		if err != nil {
			log.Println("Error al escanear el tipo de cama:", err)
			return nil, err
		}
		tiposCama = append(tiposCama, tipoCama)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return tiposCama, nil
}
