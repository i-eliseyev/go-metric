package main

import (
	"fmt"
	"github.com/i-eliseyev/go-metric/internal/common"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

var globalCounter int
var pollInterval = 2 * time.Second
var reportInterval = 10 * time.Second
var serverAddr = "127.0.0.1"
var serverPort = "8080"

func fillMetrics(metrics *common.Metrics) {
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
	globalCounter++
	(*metrics)["pollCount"] = common.Metric{
		Name: "PollCount",
		Type: "counter",
		Val:  globalCounter,
	}
	(*metrics)["randomValue"] = common.Metric{
		Name: "RandomValue",
		Type: "gauge",
		Val:  time.Now().UnixNano(),
	}
}

func reportMetrics(metrics *common.Metrics) {
	client := &http.Client{}
	for _, metric := range *metrics {
		url := fmt.Sprintf(
			"http://%s:%s/update/%s/%s/%v",
			serverAddr,
			serverPort,
			metric.Type,
			metric.Name,
			metric.Val,
		)
		request, err := http.NewRequest(http.MethodPost, url, nil)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		request.Header.Set("Content-Type", "text/plain")
		response, err := client.Do(request)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		log.Println("Status: ", response.Status)
		err = response.Body.Close()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
}

func main() {
	metrics := make(common.Metrics)
	fillMetricsTicker := time.NewTicker(pollInterval)
	reportMetricsTicker := time.NewTicker(reportInterval)
	defer fillMetricsTicker.Stop()
	defer reportMetricsTicker.Stop()

	for {
		select {
		case <-fillMetricsTicker.C:
			fillMetrics(&metrics)
		case <-reportMetricsTicker.C:
			reportMetrics(&metrics)
		}
	}
}
