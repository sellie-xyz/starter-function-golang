package helpers

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type Request events.APIGatewayProxyRequest

type Response events.APIGatewayProxyResponse

func Fail(err error, status int) (Response, error) {
	e := make(map[string]string, 0)
	e["message"] = err.Error()

	body, _ := json.Marshal(e)

	return Response{
		Body:       string(body),
		StatusCode: status,
		Headers: map[string] string {
			"Content-Type":   "application/json",
		},
	}, nil
}

func Success(data interface{}, status int) (Response, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return Fail(err, http.StatusInternalServerError)
	}

	return Response{
		Body:       string(body),
		StatusCode: status,
		Headers: map[string] string {
			"Content-Type":   "application/json",
		},
	}, nil
}
