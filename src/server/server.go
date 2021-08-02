package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetServer(port string, router *mux.Router) *http.Server {
	server := http.Server{
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	server.Handler = router
	return &server
}
