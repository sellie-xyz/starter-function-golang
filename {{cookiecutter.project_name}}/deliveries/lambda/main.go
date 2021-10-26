package main

import (
	"context"
	"{{cookiecutter.module}}"
	"{{cookiecutter.module}}/helpers"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type handler struct {
	usecase __cookiecutter_project_name__.ObjectService
}

func (h *handler) Create(ctx context.Context, body []byte) (helpers.Response, error) {

	log.Info("request received")

	createPaymentMethod := &__cookiecutter_project_name__.CreateObject{}
	if err := json.Unmarshal(body, &createPaymentMethod); err != nil {
		log.Info("invalid body received")
		return helpers.Fail(err, http.StatusBadRequest)
	}

	response, error := h.usecase.SaveObject(ctx, createPaymentMethod)
	if error != nil {
		return helpers.Fail(error, http.StatusInternalServerError)
	}

	return helpers.Success(response, http.StatusCreated)
}

func main() {
	usecase, err := __cookiecutter_project_name__.Init(false)
	if err != nil {
		log.Error(err)
	}

	h := &handler{usecase}
	lambda.Start(helpers.Router(h))
}
