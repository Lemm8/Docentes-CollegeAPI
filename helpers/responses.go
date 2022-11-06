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

func GetDocentesResponse(docentes []*Docente) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&ListDocentesResponseStruct{
		Status:   http.StatusOK,
		Docentes: docentes,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
		Body:       string(body),
	}
}

func GetDocenteResponse(docente *Docente) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DocenteResponseStruct{
		Status:  http.StatusAccepted,
		Message: "OK",
		Docente: docente,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
		Body:       string(body),
	}
}

func PostDocenteResponse(docente *Docente) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DocenteResponseStruct{
		Status:  http.StatusCreated,
		Message: "Docente creado",
		Docente: docente,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
		Body:       string(body),
	}
}

func PutDocenteResponse(docente *Docente) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DocenteResponseStruct{
		Status:  http.StatusAccepted,
		Message: "Docente actualizado",
		Docente: docente,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
		Body:       string(body),
	}
}

func DeleteDocenteResponse(docente *Docente) *events.APIGatewayProxyResponse {
	// FORMAT RESPONSE
	body, _ := json.Marshal(&DocenteResponseStruct{
		Status:  http.StatusAccepted,
		Message: "Docente eliminado",
		Docente: docente,
	})

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
		Body:       string(body),
	}
}
