// MySQLMembresiaRepository.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/core"
	repositories "github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain"
	"github.com/vicpoo/ApiGoGestionCama/nuevo_proyecto/src/membresia/domain/entities"
)

type MySQLMembresiaRepository struct {
	conn *sql.DB
}

func NewMySQLMembresiaRepository() repositories.IMembresia {
	conn := core.GetDB()
	return &MySQLMembresiaRepository{conn: conn}
}

// Función auxiliar para parsear fechas de manera flexible
func parseDatabaseDate(dateStr string) (time.Time, error) {
	// Intentar con formato ISO 8601 completo (con tiempo)
	if t, err := time.Parse(time.RFC3339, dateStr); err == nil {
		return t, nil
	}
	// Intentar con formato ISO 8601 sin zona horaria
	if t, err := time.Parse("2006-01-02T15:04:05", dateStr); err == nil {
		return t, nil
	}
	// Intentar con solo fecha
	return time.Parse("2006-01-02", dateStr)
}

func (mysql *MySQLMembresiaRepository) Save(membresia *entities.Membresia) error {
	query := `
		INSERT INTO Membresia (usuario_id, fecha_inicio, fecha_fin, estado)
		VALUES (?, ?, ?, ?)
	`
	result, err := mysql.conn.Exec(
		query,
		membresia.UsuarioID,
		membresia.FechaInicio.Format("2006-01-02"),
		membresia.FechaFin.Format("2006-01-02"),
		membresia.Estado,
	)
	if err != nil {
		log.Println("Error al guardar la membresía:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID:", err)
		return err
	}
	membresia.ID = int32(id)

	return nil
}

func (mysql *MySQLMembresiaRepository) Update(membresia *entities.Membresia) error {
	query := `
		UPDATE Membresia
		SET usuario_id = ?, fecha_inicio = ?, fecha_fin = ?, estado = ?
		WHERE id = ?
	`
	result, err := mysql.conn.Exec(
		query,
		membresia.UsuarioID,
		membresia.FechaInicio.Format("2006-01-02"),
		membresia.FechaFin.Format("2006-01-02"),
		membresia.Estado,
		membresia.ID,
	)
	if err != nil {
		log.Println("Error al actualizar la membresía:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("membresía con ID %d no encontrada", membresia.ID)
	}

	return nil
}

func (mysql *MySQLMembresiaRepository) Delete(id int32) error {
	query := "DELETE FROM Membresia WHERE id = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar la membresía:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("membresía con ID %d no encontrada", id)
	}

	return nil
}

func (mysql *MySQLMembresiaRepository) GetById(id int32) (*entities.Membresia, error) {
	query := `
		SELECT 
			id, 
			usuario_id, 
			DATE_FORMAT(fecha_inicio, '%Y-%m-%d') as fecha_inicio,
			DATE_FORMAT(fecha_fin, '%Y-%m-%d') as fecha_fin,
			estado
		FROM Membresia
		WHERE id = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var membresia entities.Membresia
	var fechaInicioStr, fechaFinStr string

	err := row.Scan(
		&membresia.ID,
		&membresia.UsuarioID,
		&fechaInicioStr,
		&fechaFinStr,
		&membresia.Estado,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("membresía con ID %d no encontrada", id)
		}
		log.Println("Error al buscar la membresía por ID:", err)
		return nil, err
	}

	// Parsear fechas ya formateadas desde la consulta
	membresia.FechaInicio, err = time.Parse("2006-01-02", fechaInicioStr)
	if err != nil {
		log.Println("Error al parsear fecha_inicio:", err)
		return nil, err
	}

	membresia.FechaFin, err = time.Parse("2006-01-02", fechaFinStr)
	if err != nil {
		log.Println("Error al parsear fecha_fin:", err)
		return nil, err
	}

	return &membresia, nil
}

func (mysql *MySQLMembresiaRepository) GetAll() ([]entities.Membresia, error) {
	query := `
		SELECT 
			id, 
			usuario_id, 
			DATE_FORMAT(fecha_inicio, '%Y-%m-%d') as fecha_inicio,
			DATE_FORMAT(fecha_fin, '%Y-%m-%d') as fecha_fin,
			estado
		FROM Membresia
		ORDER BY fecha_inicio DESC
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todas las membresías:", err)
		return nil, err
	}
	defer rows.Close()

	var membresias []entities.Membresia
	for rows.Next() {
		var membresia entities.Membresia
		var fechaInicioStr, fechaFinStr string

		err := rows.Scan(
			&membresia.ID,
			&membresia.UsuarioID,
			&fechaInicioStr,
			&fechaFinStr,
			&membresia.Estado,
		)
		if err != nil {
			log.Println("Error al escanear la membresía:", err)
			return nil, err
		}

		// Parsear fechas ya formateadas
		membresia.FechaInicio, err = time.Parse("2006-01-02", fechaInicioStr)
		if err != nil {
			log.Println("Error al parsear fecha_inicio:", err)
			return nil, err
		}

		membresia.FechaFin, err = time.Parse("2006-01-02", fechaFinStr)
		if err != nil {
			log.Println("Error al parsear fecha_fin:", err)
			return nil, err
		}

		membresias = append(membresias, membresia)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return membresias, nil
}

func (mysql *MySQLMembresiaRepository) GetByUsuarioID(usuarioID int32) ([]entities.Membresia, error) {
	query := `
		SELECT 
			id, 
			usuario_id, 
			DATE_FORMAT(fecha_inicio, '%Y-%m-%d') as fecha_inicio,
			DATE_FORMAT(fecha_fin, '%Y-%m-%d') as fecha_fin,
			estado
		FROM Membresia
		WHERE usuario_id = ?
		ORDER BY fecha_inicio DESC
	`
	rows, err := mysql.conn.Query(query, usuarioID)
	if err != nil {
		log.Println("Error al obtener membresías por usuario:", err)
		return nil, err
	}
	defer rows.Close()

	var membresias []entities.Membresia
	for rows.Next() {
		var membresia entities.Membresia
		var fechaInicioStr, fechaFinStr string

		err := rows.Scan(
			&membresia.ID,
			&membresia.UsuarioID,
			&fechaInicioStr,
			&fechaFinStr,
			&membresia.Estado,
		)
		if err != nil {
			log.Println("Error al escanear la membresía:", err)
			return nil, err
		}

		// Parsear fechas ya formateadas
		membresia.FechaInicio, err = time.Parse("2006-01-02", fechaInicioStr)
		if err != nil {
			log.Println("Error al parsear fecha_inicio:", err)
			return nil, err
		}

		membresia.FechaFin, err = time.Parse("2006-01-02", fechaFinStr)
		if err != nil {
			log.Println("Error al parsear fecha_fin:", err)
			return nil, err
		}

		membresias = append(membresias, membresia)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return membresias, nil
}

func (mysql *MySQLMembresiaRepository) GetActiveByUsuarioID(usuarioID int32) (*entities.Membresia, error) {
	query := `
		SELECT 
			id, 
			usuario_id, 
			DATE_FORMAT(fecha_inicio, '%Y-%m-%d') as fecha_inicio,
			DATE_FORMAT(fecha_fin, '%Y-%m-%d') as fecha_fin,
			estado
		FROM Membresia
		WHERE usuario_id = ? 
		AND estado = 1 
		AND fecha_inicio <= CURRENT_DATE() 
		AND fecha_fin >= CURRENT_DATE()
		LIMIT 1
	`
	row := mysql.conn.QueryRow(query, usuarioID)

	var membresia entities.Membresia
	var fechaInicioStr, fechaFinStr string

	err := row.Scan(
		&membresia.ID,
		&membresia.UsuarioID,
		&fechaInicioStr,
		&fechaFinStr,
		&membresia.Estado,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("membresía activa no encontrada para el usuario %d", usuarioID)
		}
		log.Println("Error al buscar membresía activa:", err)
		return nil, err
	}

	// Parsear fechas ya formateadas
	membresia.FechaInicio, err = time.Parse("2006-01-02", fechaInicioStr)
	if err != nil {
		log.Println("Error al parsear fecha_inicio:", err)
		return nil, err
	}

	membresia.FechaFin, err = time.Parse("2006-01-02", fechaFinStr)
	if err != nil {
		log.Println("Error al parsear fecha_fin:", err)
		return nil, err
	}

	return &membresia, nil
}

func (mysql *MySQLMembresiaRepository) GetExpiringMemberships(daysBeforeExpiration int) ([]entities.Membresia, error) {
	query := `
		SELECT 
			id, 
			usuario_id, 
			DATE_FORMAT(fecha_inicio, '%Y-%m-%d') as fecha_inicio,
			DATE_FORMAT(fecha_fin, '%Y-%m-%d') as fecha_fin,
			estado
		FROM Membresia
		WHERE estado = 1 
		AND fecha_fin BETWEEN CURRENT_DATE() AND DATE_ADD(CURRENT_DATE(), INTERVAL ? DAY)
	`
	rows, err := mysql.conn.Query(query, daysBeforeExpiration)
	if err != nil {
		log.Println("Error al obtener membresías por vencer:", err)
		return nil, err
	}
	defer rows.Close()

	var membresias []entities.Membresia
	for rows.Next() {
		var membresia entities.Membresia
		var fechaInicioStr, fechaFinStr string

		err := rows.Scan(
			&membresia.ID,
			&membresia.UsuarioID,
			&fechaInicioStr,
			&fechaFinStr,
			&membresia.Estado,
		)
		if err != nil {
			log.Println("Error al escanear la membresía:", err)
			return nil, err
		}

		// Parsear fechas ya formateadas
		membresia.FechaInicio, err = time.Parse("2006-01-02", fechaInicioStr)
		if err != nil {
			log.Println("Error al parsear fecha_inicio:", err)
			return nil, err
		}

		membresia.FechaFin, err = time.Parse("2006-01-02", fechaFinStr)
		if err != nil {
			log.Println("Error al parsear fecha_fin:", err)
			return nil, err
		}

		membresias = append(membresias, membresia)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return membresias, nil
}

func (mysql *MySQLMembresiaRepository) RenewMembership(membresiaID int32, newEndDate time.Time) error {
	query := `
		UPDATE Membresia
		SET fecha_fin = ?
		WHERE id = ?
	`
	result, err := mysql.conn.Exec(
		query,
		newEndDate.Format("2006-01-02"),
		membresiaID,
	)
	if err != nil {
		log.Println("Error al renovar la membresía:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("membresía con ID %d no encontrada", membresiaID)
	}

	return nil
}
