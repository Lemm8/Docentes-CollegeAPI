package main

import (
	"context"
	"database/sql"

	"github.com/Lemm8/Docentes-CollegeAPI.git/api"
	"github.com/Lemm8/Docentes-CollegeAPI.git/database"
	"github.com/Lemm8/Docentes-CollegeAPI.git/helpers"
	"github.com/Lemm8/Docentes-CollegeAPI.git/validators"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var db *sql.DB

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	// VALIDATE PATH
	if !validators.IsValidPath(event.Path) {
		return helpers.ServerError("This path doesn't exist"), nil
	}

	// CONNECT TO DB
	dbConnection, err := database.GetConnection()
	if err != nil {
		return helpers.ServerError(string(err.Error())), nil
	}
	db = dbConnection

	switch event.HTTPMethod {
	case "GET":
		apiResponse, err := api.GetDocentes(ctx, db, event)
		if err != nil {
			return helpers.ServerError(string(err.Error())), nil
		}
		return apiResponse, nil

	case "POST":
		apiResponse, err := api.PostDocente(ctx, db, event)
		if err != nil {
			return helpers.ServerError(string(err.Error())), nil
		}
		return apiResponse, nil

	case "PUT":
		apiResponse, err := api.PutDocente(ctx, db, event)
		if err != nil {
			return helpers.ServerError(string(err.Error())), nil
		}
		return apiResponse, nil

	case "DELETE":
		apiResponse, err := api.DeleteDocente(ctx, db, event)
		if err != nil {
			return helpers.ServerError(string(err.Error())), nil
		}
		return apiResponse, nil

	default:
		return helpers.UnhandledMethod(), nil
	}
}

func main() {
	lambda.Start(handler)
}
