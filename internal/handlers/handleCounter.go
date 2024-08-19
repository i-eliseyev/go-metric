package handlers

import (
	"github.com/i-eliseyev/go-metric/internal/storage"
	"log"
	"net/http"
)

func HandleCounter(writer http.ResponseWriter, request *http.Request) {
	log.Println("Serving: ", request.URL.Path)
	metric, err := getMetricFromRequest(request, urlPrefixCounter)

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	err = storage.MetricStorage.UpdateCounter(metric)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
