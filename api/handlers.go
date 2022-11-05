package api

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Lemm8/Docentes-CollegeAPI.git/database"
	"github.com/Lemm8/Docentes-CollegeAPI.git/helpers"
	"github.com/aws/aws-lambda-go/events"
)

func GetDocentes(ctx context.Context, db *sql.DB, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	docentes, err := database.FetchDocentes(ctx, db)

	if err != nil {
		return helpers.ServerError(string(err.Error())), nil
	}

	return helpers.GetDocentesResponse(docentes), nil
}

func PostDocente(ctx context.Context, db *sql.DB, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	docente := &helpers.Docente{}
	err := json.Unmarshal([]byte(req.Body), &docente)
	if err != nil {
		return helpers.ServerError(string(err.Error())), nil
	}

	docente, err = database.CreateDocente(ctx, db, docente.Nombre, docente.Apellido, docente.Matricula,
		docente.Fecha_Nacimiento, docente.Titulo, docente.Correo, docente.Telefono)

	if err != nil {
		return helpers.ServerError(string(err.Error())), nil
	}

	return helpers.PostDocenteResponse(docente), nil
}

func PutDocente(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return helpers.PutDocenteResponse(), nil
}

func DeleteDocente(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return helpers.DeleteDocenteResponse(), nil
}
