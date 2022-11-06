package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/Lemm8/Docentes-CollegeAPI.git/database"
	"github.com/Lemm8/Docentes-CollegeAPI.git/helpers"
	"github.com/aws/aws-lambda-go/events"

	"strconv"
)

func GetDocentes(ctx context.Context, db *sql.DB, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	if len(id) > 0 {
		docente := &helpers.Docente{}
		intID, err := strconv.Atoi(id)
		if err != nil {
			return nil, errors.New("el id no es válido")
		}
		docente, err = database.FetchDocente(ctx, db, intID)
		return helpers.GetDocenteResponse(docente), nil
	}

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

func PutDocente(ctx context.Context, db *sql.DB, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	docente := &helpers.Docente{}
	err := json.Unmarshal([]byte(req.Body), &docente)
	if err != nil {
		return helpers.ServerError(string(err.Error())), nil
	}

	id := req.QueryStringParameters["id"]
	if len(id) > 0 {
		intID, err := strconv.Atoi(id)
		if err != nil {
			return nil, errors.New("el id no es válido")
		}
		updatedDocente, err := database.UpdateDocente(ctx, db, intID, docente)
		if err != nil {
			return nil, err
		}
		return helpers.PutDocenteResponse(updatedDocente), nil
	}

	return nil, errors.New("se debe incluir el id del docente para actualizarlo")
}

func DeleteDocente(ctx context.Context, db *sql.DB, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	if len(id) > 0 {
		docente := &helpers.Docente{}
		intID, err := strconv.Atoi(id)
		if err != nil {
			return nil, errors.New("el id no es válido")
		}
		docente, err = database.DeleteDocente(ctx, db, intID)
		if err != nil {
			return helpers.ServerError(err.Error()), nil
		}
		return helpers.DeleteDocenteResponse(docente), nil
	}
	return helpers.ServerError("se debe incluir el id del docente para aliminarlo"), nil
}
