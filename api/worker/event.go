package worker

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/MadhanMurali/CustomerLabsAssessment/pkg/event"
)

var EventChannel = make(chan []byte)

func SendEvent(body []byte) {
	event := event.Event{}
	event.LoadFromMinifiedEventBytes(body)

	postBody, err := json.Marshal(event)
	if err != nil {
		log.Println(err)
		return
	}
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("https://webhook.site/5386301d-a2e7-439d-a7a6-57cb3dab9894", "application/json", responseBody)
	if err != nil || resp.StatusCode != 200 {
		if err != nil {
			log.Println(err)
		} else {
			log.Println("response status code is not 200")
		}
		return
	}

	log.Println("event sent")
}

func SetupSendEventWorker() {
	for body := range EventChannel {
		SendEvent(body)
	}
}
