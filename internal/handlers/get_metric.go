package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/i-eliseyev/go-metric/internal/common"
	"github.com/i-eliseyev/go-metric/internal/storage"
	"net/http"
)

func HandleGetMetric(ctx *fiber.Ctx) error {
	log.Info("Serving: ", ctx.OriginalURL())

	requestedMetric := new(common.Metric)
	err := ctx.BodyParser(requestedMetric)
	if err != nil {
		log.Warn(err)
		ctx.Status(http.StatusBadRequest)
		return err
	}

	metric, err := storage.MetricStorage.GetMetric(requestedMetric.ID)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return err
	}

	ctx.Status(http.StatusOK)
	return ctx.JSON(metric)
}
