package database

import (
	"context"
	"database/sql"
)

func CreateDocente(ctx context.Context, db *sql.DB, nombre string, apellido string, matricula string, fecha_nacimiento string,
	titulo string, correo string, telefono string) error {

	const sql = `INSERT INTO Docentes (Nombre, Apellido, Matricula, Fecha_Nacimiento, Titulo, Correo, Telefono) 
				 VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.ExecContext(ctx, sql, nombre, apellido, matricula, fecha_nacimiento, titulo, correo, telefono)
	return err
}
