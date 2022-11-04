package main

import (
	"context"
	"database/sql"

	// "fmt"

	// "os"

	"github.com/Lemm8/Docentes-CollegeAPI.git/api"
	"github.com/Lemm8/Docentes-CollegeAPI.git/database"
	"github.com/Lemm8/Docentes-CollegeAPI.git/helpers"
	"github.com/Lemm8/Docentes-CollegeAPI.git/validators"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

var logger *zap.Logger
var db *sql.DB

func init() {

	l, _ := zap.NewProduction()
	logger = l

}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	logger.Info("Received evnet", zap.Any("method", event.HTTPMethod), zap.Any("body", event.Body))

	// VALIDATE PATH
	if !validators.IsValidPath(event.Path) {
		return helpers.ServerError("This path doesn't exist"), nil
	}

	// CONNECT TO DB
	dbConnection, err := database.GetConnection()
	if err != nil {
		return helpers.ServerError("Can't connect to database"), nil
	}
	db = dbConnection

	switch event.HTTPMethod {
	case "GET":
		apiResponse, err := api.GetClientes(event)
		if err != nil {
			return helpers.ServerError("Can't GET clientes"), nil
		}
		return apiResponse, nil

	case "POST":
		apiResponse, err := api.PostCliente(event)
		if err != nil {
			return helpers.ServerError("Can't POST clientes"), nil
		}
		return apiResponse, nil

	case "PUT":
		apiResponse, err := api.PutCliente(event)
		if err != nil {
			return helpers.ServerError("Can't POST clientes"), nil
		}
		return apiResponse, nil

	case "DELETE":
		apiResponse, err := api.DeleteCliente(event)
		if err != nil {
			return helpers.ServerError("Can't POST clientes"), nil
		}
		return apiResponse, nil

	default:
		return helpers.UnhandledMethod(), nil
	}
}

func main() {
	lambda.Start(handler)
}
