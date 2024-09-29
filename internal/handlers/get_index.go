package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/i-eliseyev/go-metric/internal/storage"
	"strconv"
)

func HandleIndex(ctx *fiber.Ctx) error {
	metrics := make(map[string]map[string]string, len(storage.MetricStorage.Metrics))

	for name, metric := range storage.MetricStorage.Metrics {
		var delta, value string
		if metric.MType == "counter" {
			delta = strconv.FormatInt(*metric.Delta, 10)
		} else {
			value = strconv.FormatFloat(*metric.Value, 'f', -1, 64)
		}

		metrics[name] = map[string]string{
			"Type":  metric.MType,
			"Value": value,
			"Delta": delta,
		}
	}

	return ctx.Render("index", metrics)
}
