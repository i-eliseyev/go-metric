package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/i-eliseyev/go-metric/internal"
	"github.com/i-eliseyev/go-metric/internal/common"
	"github.com/i-eliseyev/go-metric/internal/security"
	"github.com/i-eliseyev/go-metric/internal/storage"
	"net/http"
)

func HandleUpdateMetric(ctx *fiber.Ctx) error {
	log.Info("Serving: ", ctx.OriginalURL())

	metric := new(common.Metric)
	err := ctx.BodyParser(metric)
	if err != nil {
		log.Warn(err)
		ctx.Status(http.StatusBadRequest)
		return err
	}

	if *internal.SECRET != "" {
		isSignatureValid := security.ValidateSignature(metric, internal.SECRET)
		if !isSignatureValid {
			err := errors.New("bad signature")
			ctx.Status(http.StatusBadRequest)
			return err
		}
	}

	if metric.MType == "counter" {
		storage.MetricStorage.UpdateCounter(metric)
	} else {
		storage.MetricStorage.UpdateGauge(metric)
	}

	ctx.Status(http.StatusOK)
	return ctx.JSON(metric)
}
