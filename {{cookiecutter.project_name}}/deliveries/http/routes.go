package http

import (
	"context"
	"{{cookiecutter.module}}"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const fiveSecondsTimeout = time.Second * 5

type delivery struct {
	usecase "{{cookiecutter.module}}".ObjectService
}

func writeErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func (d *delivery) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), fiveSecondsTimeout)
	defer cancel()

	decoder := json.NewDecoder(r.Body)
	createObject := &"{{cookiecutter.module}}".CreateObject{}
	if err := decoder.Decode(&createObject); err != nil {
		writeErr(w, err)
		return
	}

	response, error := d.usecase.SaveObject(ctx, createObject)
	if error != nil {
		writeErr(w, error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func Routes() (*mux.Router, error) {
	usecase, err := starter_function.Init(true)
	if err != nil {
		log.Panic(err)
	}

	delivery := &delivery{usecase}

	r := mux.NewRouter()
	r.HandleFunc("/\"{{cookiecutter.module}}\"", delivery.Create).Methods("POST")

	return r, nil
}
