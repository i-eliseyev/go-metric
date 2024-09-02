package server

import "time"

const (
	Port         = ":8080"
	IdleTimeout  = 10 * time.Second
	ReadTimeout  = time.Second
	WriteTimeout = time.Second
)
