package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Lemm8/Docentes-CollegeAPI.git/helpers"
	"github.com/Lemm8/Docentes-CollegeAPI.git/validators"
)

const getDocentesQuery = `SELECT * FROM Docente;`

const insertDocenteSQL = `INSERT INTO Docente (Nombre, Apellido, Matricula, Fecha_Nacimiento, Titulo, Correo, Telefono) 
VALUES (?, ?, ?, ?, ?, ?, ?);`

// updateDocenteSQL := `UPDATE Docente SET (?) WHERE ID=(?);`

// GET ALL DOCENTES FROM DB
func FetchDocentes(ctx context.Context, db *sql.DB) ([]*helpers.Docente, error) {
	// QUERY ALL DOCENTES
	rows, err := db.QueryContext(ctx, getDocentesQuery)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// INIT SLICE
	docentes := make([]*helpers.Docente, 0)
	for rows.Next() {
		// APPEND DOCENTE TO SLICE
		docente := &helpers.Docente{}
		if err := rows.Scan(&docente.ID, &docente.Nombre, &docente.Apellido, &docente.Matricula,
			&docente.Fecha_Nacimiento, &docente.Titulo, &docente.Correo, &docente.Telefono); err != nil {
			return nil, err
		}

		docentes = append(docentes, docente)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return docentes, nil

}

func CreateDocente(ctx context.Context, db *sql.DB, nombre string, apellido string, matricula string, fecha_nacimiento string,
	titulo string, correo string, telefono string) (*helpers.Docente, error) {

	if !validators.IsValidDate(fecha_nacimiento) {
		return nil, errors.New("Invalid date format (must be YYYY-MM-DD)")
	}

	newFecha, err := time.Parse("2006-01-02", fecha_nacimiento)

	if err != nil {
		return nil, err
	}

	res, err := db.ExecContext(ctx, insertDocenteSQL, nombre, apellido, matricula, newFecha, titulo, correo, telefono)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	docente := helpers.Docente{
		ID:               int(id),
		Nombre:           nombre,
		Apellido:         apellido,
		Matricula:        matricula,
		Fecha_Nacimiento: fecha_nacimiento,
		Titulo:           titulo,
		Correo:           correo,
		Telefono:         telefono,
	}

	return &docente, nil
}
