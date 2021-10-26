package main

import (
	delivery "{{cookiecutter.module}}/deliveries/http"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	router, err := delivery.Routes()
	if err != nil {
		log.Panic(err)
	}

	log.Println("Running on port: ", port)
	log.Panic(http.ListenAndServe(":"+port, router))
}
