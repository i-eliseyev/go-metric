package main

import (
	"github.com/i-eliseyev/go-metric/internal/server"
	"log"
)

func main() {
	log.Fatal(server.StartServer())
}
