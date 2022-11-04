package api

import (
	"context"
	"database/sql"

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

func PostDocente(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return helpers.PostDocenteResponse(), nil
}

func PutDocente(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return helpers.PutDocenteResponse(), nil
}

func DeleteDocente(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return helpers.DeleteDocenteResponse(), nil
}
