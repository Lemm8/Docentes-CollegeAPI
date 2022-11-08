package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	// "reflect"
	"time"

	"github.com/Lemm8/Docentes-CollegeAPI.git/helpers"
	"github.com/Lemm8/Docentes-CollegeAPI.git/validators"
)

const getDocentesQuery = `SELECT * FROM Docente;`

const getDocenteQuery = `SELECT * FROM Docente WHERE ID = ?;`

const insertDocenteSQL = `INSERT INTO Docente (Nombre, Apellido, Matricula, Fecha_Nacimiento, Titulo, Correo, Telefono) 
VALUES (?, ?, ?, ?, ?, ?, ?);`

var updateDocenteSQL = `UPDATE Docente SET Nombre = ?, Apellido = ?, Matricula = ?, Fecha_Nacimiento = ?, Titulo = ?,
Correo = ?, Telefono = ? WHERE ID = ?;`

const deleteDocenteSQL = `DELETE FROM Docente WHERE ID = ?;`

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

// GET ALL DOCENTES FROM DB
func FetchDocente(ctx context.Context, db *sql.DB, id int) (*helpers.Docente, error) {
	// QUERY DOCENTE BY ID
	row := db.QueryRowContext(ctx, getDocenteQuery, id)

	docente := &helpers.Docente{}
	if err := row.Scan(&docente.ID, &docente.Nombre, &docente.Apellido, &docente.Matricula,
		&docente.Fecha_Nacimiento, &docente.Titulo, &docente.Correo, &docente.Telefono); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("No existe un docente con el ID %v", id))
		}
		return nil, err
	}
	return docente, nil
}

// CREATE A DOCENTE AND INSERT INTO TABLE
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

// UPDATE A DOCENTE
func UpdateDocente(ctx context.Context, db *sql.DB, id int, docente *helpers.Docente) (*helpers.Docente, error) {
	// CHECK IF DOCENTE EXISTS
	_, err := FetchDocente(ctx, db, id)
	if err != nil {
		return nil, err
	}

	if !validators.IsValidDate(docente.Fecha_Nacimiento) {
		return nil, errors.New("Invalid date format (must be YYYY-MM-DD)")
	}

	newFecha, err := time.Parse("2006-01-02", docente.Fecha_Nacimiento)

	if err != nil {
		return nil, err
	}

	_, err = db.ExecContext(ctx, updateDocenteSQL, docente.Nombre, docente.Apellido,
		docente.Matricula, newFecha, docente.Titulo, docente.Correo, docente.Telefono, id)
	if err != nil {
		return nil, err
	}

	updatedDocente := helpers.Docente{
		ID:               int(id),
		Nombre:           docente.Nombre,
		Apellido:         docente.Apellido,
		Matricula:        docente.Matricula,
		Fecha_Nacimiento: docente.Fecha_Nacimiento,
		Titulo:           docente.Titulo,
		Correo:           docente.Correo,
		Telefono:         docente.Telefono,
	}

	return &updatedDocente, nil

}

// DELETE A DOCENTE
func DeleteDocente(ctx context.Context, db *sql.DB, id int) (*helpers.Docente, error) {
	docente, err := FetchDocente(ctx, db, id)
	if err != nil {
		return nil, err
	}

	_, err = db.ExecContext(ctx, deleteDocenteSQL, id)
	if err != nil {
		return nil, err
	}

	return docente, nil
}
