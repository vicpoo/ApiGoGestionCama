package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/almacen/domain/entities"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/core"
)

type MySQLAlmacenRepository struct {
	conn *sql.DB
}

func NewMySQLAlmacenRepository() repositories.IAlmacen {
	conn := core.GetDB()
	return &MySQLAlmacenRepository{conn: conn}
}

func (mysql *MySQLAlmacenRepository) Save(almacen *entities.Almacen) error {
	query := `
		INSERT INTO Almacen (tipo_cama_id, cantidad)
		VALUES (?, ?)
	`
	result, err := mysql.conn.Exec(
		query,
		almacen.TipoCamaID,
		almacen.Cantidad,
	)
	if err != nil {
		log.Println("Error al guardar registro en almacén:", err)
		return err
	}

	// Obtener el ID generado
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	almacen.ID = int32(id)

	return nil
}

func (mysql *MySQLAlmacenRepository) Update(almacen *entities.Almacen) error {
	query := `
		UPDATE Almacen
		SET tipo_cama_id = ?, cantidad = ?
		WHERE id = ?
	`
	result, err := mysql.conn.Exec(
		query,
		almacen.TipoCamaID,
		almacen.Cantidad,
		almacen.ID,
	)
	if err != nil {
		log.Println("Error al actualizar registro en almacén:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("registro en almacén con ID %d no encontrado", almacen.ID)
	}

	return nil
}

func (mysql *MySQLAlmacenRepository) Delete(id int32) error {
	query := "DELETE FROM Almacen WHERE id = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar registro de almacén:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("registro en almacén con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MySQLAlmacenRepository) GetById(id int32) (*entities.Almacen, error) {
	query := `
		SELECT id, tipo_cama_id, cantidad
		FROM Almacen
		WHERE id = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var almacen entities.Almacen
	err := row.Scan(
		&almacen.ID,
		&almacen.TipoCamaID,
		&almacen.Cantidad,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("registro en almacén con ID %d no encontrado", id)
		}
		log.Println("Error al buscar registro en almacén por ID:", err)
		return nil, err
	}

	return &almacen, nil
}

func (mysql *MySQLAlmacenRepository) GetAll() ([]entities.Almacen, error) {
	query := `
		SELECT id, tipo_cama_id, cantidad
		FROM Almacen
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los registros de almacén:", err)
		return nil, err
	}
	defer rows.Close()

	var almacenes []entities.Almacen
	for rows.Next() {
		var almacen entities.Almacen
		err := rows.Scan(
			&almacen.ID,
			&almacen.TipoCamaID,
			&almacen.Cantidad,
		)
		if err != nil {
			log.Println("Error al escanear registro de almacén:", err)
			return nil, err
		}
		almacenes = append(almacenes, almacen)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return almacenes, nil
}

func (mysql *MySQLAlmacenRepository) IncrementarCantidad(id int32, cantidad int32) error {
	query := `
		UPDATE Almacen
		SET cantidad = cantidad + ?
		WHERE id = ?
	`
	result, err := mysql.conn.Exec(
		query,
		cantidad,
		id,
	)
	if err != nil {
		log.Println("Error al incrementar cantidad en almacén:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("registro en almacén con ID %d no encontrado", id)
	}

	return nil
}
