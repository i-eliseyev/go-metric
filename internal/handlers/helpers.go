package handlers

import (
	"errors"
	"github.com/i-eliseyev/go-metric/internal/common"
	"log"
	"net/http"
	"strings"
)

var prefixTypeMap = map[string]string{
	urlPrefixGauge:   "gauge",
	urlPrefixCounter: "counter",
}

func getMetricFromRequest(request *http.Request, prefix string) (common.Metric, error) {
	parts := strings.Split(strings.TrimPrefix(request.URL.Path, prefix), "/")
	if len(parts) < 2 {
		log.Println("Bad url path: ", request.URL.Path)
		return common.Metric{}, errors.New("bad url path")
	}
	return common.Metric{
		Type: prefixTypeMap[prefix],
		Name: parts[0],
		Val:  parts[1],
	}, nil
}
