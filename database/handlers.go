package database

import (
	"context"
	"database/sql"

	"github.com/Lemm8/Docentes-CollegeAPI.git/helpers"
)

const getDocentesQuery = `SELECT * FROM Docente;`

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
	if rows.Next() {
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
	titulo string, correo string, telefono string) error {

	const sql = `INSERT INTO Docentes (Nombre, Apellido, Matricula, Fecha_Nacimiento, Titulo, Correo, Telefono) 
				 VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.ExecContext(ctx, sql, nombre, apellido, matricula, fecha_nacimiento, titulo, correo, telefono)
	return err
}
