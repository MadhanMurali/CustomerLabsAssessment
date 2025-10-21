package main

import (
	"log"
	"net/http"

	"github.com/MadhanMurali/CustomerLabsAssessment/api/route"
	"github.com/MadhanMurali/CustomerLabsAssessment/api/worker"
)

func main() {
	router := http.NewServeMux()

	router.Handle("/", route.GetEventRouter())

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go worker.SetupSendEventWorker()

	log.Println("starting server on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
