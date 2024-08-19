package handlers

import (
	"github.com/i-eliseyev/go-metric/internal/storage"
	"log"
	"net/http"
)

func HandleGauge(writer http.ResponseWriter, request *http.Request) {
	log.Println("Serving: ", request.URL.Path)
	metric, err := getMetricFromRequest(request, urlPrefixGauge)

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = storage.MetricStorage.UpdateGauge(metric)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
