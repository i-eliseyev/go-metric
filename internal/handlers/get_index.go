package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/i-eliseyev/go-metric/internal/storage"
	"strconv"
)

func HandleIndex(ctx *fiber.Ctx) error {
	metrics := make(map[string]map[string]string, len(storage.MetricStorage.Metrics))

	for name, metric := range storage.MetricStorage.Metrics {
		metrics[name] = map[string]string{
			"Type":  metric.Type,
			"Value": strconv.FormatFloat(metric.Val, 'f', -1, 64),
		}
	}

	return ctx.Render("index", metrics)
}
