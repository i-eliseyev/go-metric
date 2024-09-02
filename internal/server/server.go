package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/i-eliseyev/go-metric/internal/routers"
	"log"
)

func StartServer() error {
	app := fiber.New(
		fiber.Config{
			IdleTimeout:  IdleTimeout,
			ReadTimeout:  ReadTimeout,
			WriteTimeout: WriteTimeout,
			Views:        html.New("./internal/templates", ".html"),
		},
	)
	routers.SetupRouters(app)
	err := app.Listen(Port)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}
