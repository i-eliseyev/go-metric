package agent

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/i-eliseyev/go-metric/internal/common"
	"log"
	"os"
	"runtime"
	"time"
)

func FillMetrics(metrics *common.Metrics) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	(*metrics)["alloc"] = common.Metric{
		Name: "Alloc",
		Type: "gauge",
		Val:  float64(memStats.Alloc),
	}
	(*metrics)["buckHashSys"] = common.Metric{
		Name: "BuckHashSys",
		Type: "gauge",
		Val:  float64(memStats.BuckHashSys),
	}
	(*metrics)["frees"] = common.Metric{
		Name: "Frees",
		Type: "gauge",
		Val:  float64(memStats.Frees),
	}
	(*metrics)["gCCPUFraction"] = common.Metric{
		Name: "GCCPUFraction",
		Type: "gauge",
		Val:  memStats.GCCPUFraction,
	}
	(*metrics)["gSys"] = common.Metric{
		Name: "GSys",
		Type: "gauge",
		Val:  float64(memStats.GCSys),
	}
	(*metrics)["heapAlloc"] = common.Metric{
		Name: "HeapAlloc",
		Type: "gauge",
		Val:  float64(memStats.HeapAlloc),
	}
	(*metrics)["heapIdle"] = common.Metric{
		Name: "HeapIdle",
		Type: "gauge",
		Val:  float64(memStats.HeapIdle),
	}
	(*metrics)["heapInuse"] = common.Metric{
		Name: "HeapInuse",
		Type: "gauge",
		Val:  float64(memStats.HeapInuse),
	}
	(*metrics)["heapObjects"] = common.Metric{
		Name: "HeapObjects",
		Type: "gauge",
		Val:  float64(memStats.HeapObjects),
	}
	(*metrics)["heapReleased"] = common.Metric{
		Name: "HeapReleased",
		Type: "gauge",
		Val:  float64(memStats.HeapReleased),
	}
	(*metrics)["heapSys"] = common.Metric{
		Name: "HeapSys",
		Type: "gauge",
		Val:  float64(memStats.HeapSys),
	}
	(*metrics)["lastGC"] = common.Metric{
		Name: "LastGC",
		Type: "gauge",
		Val:  float64(memStats.LastGC),
	}
	(*metrics)["lookups"] = common.Metric{
		Name: "Lookups",
		Type: "gauge",
		Val:  float64(memStats.Lookups),
	}
	(*metrics)["mCacheInuse"] = common.Metric{
		Name: "MCacheInuse",
		Type: "gauge",
		Val:  float64(memStats.MCacheInuse),
	}
	(*metrics)["mCacheSys"] = common.Metric{
		Name: "MCacheSys",
		Type: "gauge",
		Val:  float64(memStats.MCacheSys),
	}
	(*metrics)["mSpanInuse"] = common.Metric{
		Name: "MSpanInuse",
		Type: "gauge",
		Val:  float64(memStats.MSpanInuse),
	}
	(*metrics)["mSpanSys"] = common.Metric{
		Name: "MSpanSys",
		Type: "gauge",
		Val:  float64(memStats.MSpanSys),
	}
	(*metrics)["mallocs"] = common.Metric{
		Name: "Mallocs",
		Type: "gauge",
		Val:  float64(memStats.Mallocs),
	}
	(*metrics)["nextGC"] = common.Metric{
		Name: "NextGC",
		Type: "gauge",
		Val:  float64(memStats.NextGC),
	}
	(*metrics)["numForcedGC"] = common.Metric{
		Name: "NumForcedGC",
		Type: "gauge",
		Val:  float64(memStats.NumForcedGC),
	}
	(*metrics)["numGC"] = common.Metric{
		Name: "NumGC",
		Type: "gauge",
		Val:  float64(memStats.NumGC),
	}
	(*metrics)["otherSys"] = common.Metric{
		Name: "OtherSys",
		Type: "gauge",
		Val:  float64(memStats.OtherSys),
	}
	(*metrics)["pauseTotalNs"] = common.Metric{
		Name: "PauseTotalNs",
		Type: "gauge",
		Val:  float64(memStats.PauseTotalNs),
	}
	(*metrics)["stackInuse"] = common.Metric{
		Name: "StackInuse",
		Type: "gauge",
		Val:  float64(memStats.StackInuse),
	}
	(*metrics)["stackSys"] = common.Metric{
		Name: "StackSys",
		Type: "gauge",
		Val:  float64(memStats.StackSys),
	}
	(*metrics)["sys"] = common.Metric{
		Name: "Sys",
		Type: "gauge",
		Val:  float64(memStats.Sys),
	}
	(*metrics)["totalAlloc"] = common.Metric{
		Name: "TotalAlloc",
		Type: "gauge",
		Val:  float64(memStats.TotalAlloc),
	}
	GlobalCounter++
	(*metrics)["pollCount"] = common.Metric{
		Name: "PollCount",
		Type: "counter",
		Val:  GlobalCounter,
	}
	(*metrics)["randomValue"] = common.Metric{
		Name: "RandomValue",
		Type: "gauge",
		Val:  float64(time.Now().UnixNano()),
	}
}

func ReportMetrics(metrics *common.Metrics) {
	client := resty.New()
	client.SetRetryCount(RetryCount).SetRetryWaitTime(RetryWaitTime)

	for _, metric := range *metrics {
		pathParams := map[string]string{
			"baseURL": ServerAddr,
			"port":    ServerPort,
			"type":    metric.Type,
			"name":    metric.Name,
			"value":   fmt.Sprintf("%f", metric.Val),
		}

		resp, err := client.
			R().
			SetHeader("Content-Type", "text/plain").
			SetPathParams(pathParams).
			Post("http://{baseURL}:{port}/update/{type}/{name}/{value}")

		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		log.Println("Status: ", resp.Status())
		log.Println(resp.Request.URL)
	}
}
