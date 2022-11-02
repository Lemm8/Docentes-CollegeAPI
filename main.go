package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

var logger *zap.Logger

type DefaultResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func init() {

	l, _ := zap.NewProduction()
	logger = l
	defer logger.Sync()

}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	logger.Info("Received evnet", zap.Any("method", event.HTTPMethod), zap.Any("body", event.Body))

	if event.Path != "/Docentes" {
		// FORMAT RESPONSE
		body, _ := json.Marshal(&DefaultResponse{
			Status:  http.StatusNotFound,
			Message: "This path doesn't exist",
		})

		res = &events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       string(body),
		}
	} else {
		// FORMAT RESPONSE
		body, _ := json.Marshal(&DefaultResponse{
			Status:  http.StatusOK,
			Message: "Hola Mundo!",
		})

		res = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       string(body),
		}
	}

	return res, nil
}

func main() {
	lambda.Start(handler)
}
