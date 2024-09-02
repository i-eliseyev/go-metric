package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/i-eliseyev/go-metric/internal/common"
	"github.com/i-eliseyev/go-metric/internal/storage"
	"net/http"
	"strconv"
)

func HandleUpdateMetric(ctx *fiber.Ctx) error {
	log.Info("Serving: ", ctx.OriginalURL())

	metricType := utils.CopyString(ctx.Params("type"))
	if metricType != "counter" && metricType != "gauge" {
		log.Warnw("Wrong metric type", "type", metricType)
		ctx.Status(http.StatusBadRequest)
		return errors.New("invalid type")
	}

	valueParam := utils.CopyString(ctx.Params("value"))
	metricValue, err := strconv.ParseFloat(valueParam, 64)
	if err != nil {
		log.Warn(err)
		ctx.Status(http.StatusBadRequest)
		return err
	}

	metric := common.Metric{
		Type: metricType,
		Name: utils.CopyString(ctx.Params("name")),
		Val:  metricValue,
	}

	if metricType == "counter" {
		storage.MetricStorage.UpdateCounter(&metric)
	} else {
		storage.MetricStorage.UpdateGauge(&metric)
	}

	ctx.Status(http.StatusOK)
	return nil
}
