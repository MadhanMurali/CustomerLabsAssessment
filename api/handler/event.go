package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/MadhanMurali/CustomerLabsAssessment/api/worker"
)

func PostEvent(resWriter http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		log.Println(err)

		resWriter.Header().Set("Content-Type", "application/json")
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(map[string]string{"message": "failed to read body"})

		return
	}

	worker.EventChannel <- body

	resWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resWriter).Encode(map[string]string{"message": "ok"})
}
