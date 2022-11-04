package api

import (
	"github.com/Lemm8/Docentes-CollegeAPI.git/helpers"
	"github.com/aws/aws-lambda-go/events"
)

func GetClientes(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return helpers.GetClientesResponse(), nil
}

func PostCliente(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return helpers.PostClienteResponse(), nil
}

func PutCliente(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return helpers.PutClienteResponse(), nil
}

func DeleteCliente(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return helpers.DeleteClienteResponse(), nil
}
