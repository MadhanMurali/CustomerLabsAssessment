package route

import (
	"net/http"

	"github.com/MadhanMurali/CustomerLabsAssessment/api/handler"
)

func GetEventRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /event/", handler.PostEvent)

	return router
}
