package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func UnhandledMethod() *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  http.StatusBadRequest,
		Message: "Unhandled method, try again",
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       string(body),
	}
}

func ServerError(errMessage string) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  http.StatusInternalServerError,
		Message: errMessage,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       string(body),
	}
}

func ServerSuccess(msg string) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  http.StatusOK,
		Message: msg,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}
}

func GetClientesResponse() *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  http.StatusAccepted,
		Message: "GET Clientes",
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
		Body:       string(body),
	}
}

func PostClienteResponse() *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  http.StatusCreated,
		Message: "POST Cliente",
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
		Body:       string(body),
	}
}

func PutClienteResponse() *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  http.StatusOK,
		Message: "PUT Cliente",
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
		Body:       string(body),
	}
}

func DeleteClienteResponse() *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DefaultResponse{
		Status:  http.StatusOK,
		Message: "DELETE Cliente",
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
		Body:       string(body),
	}
}
