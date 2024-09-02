package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/i-eliseyev/go-metric/internal/storage"
	"net/http"
)

func HandleGetMetric(ctx *fiber.Ctx) error {
	log.Info("Serving: ", ctx.OriginalURL())

	metricType := utils.CopyString(ctx.Params("type"))
	if metricType != "counter" && metricType != "gauge" {
		log.Warnw("Wrong metric type", "type", metricType)
		ctx.Status(http.StatusBadRequest)
		return errors.New("invalid type")
	}

	metric, err := storage.MetricStorage.GetMetric(ctx.Params("name"))
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return err
	}

	ctx.Status(http.StatusOK)
	return ctx.JSON(metric)
}
