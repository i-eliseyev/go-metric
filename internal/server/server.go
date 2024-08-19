package server

import (
	"github.com/i-eliseyev/go-metric/internal/handlers"
	"log"
	"net/http"
	"time"
)

const PORT = ":8080"

func StartServer() error {

	mux := http.NewServeMux()
	server := http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	mux.Handle("/update/gauge/", http.HandlerFunc(handlers.HandleGauge))
	mux.Handle("/update/counter/", http.HandlerFunc(handlers.HandleCounter))

	log.Println("Ready to work!")

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
