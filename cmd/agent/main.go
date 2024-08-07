package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"
)

type Metric struct {
	Name string
	Type string
	Val  interface{}
}

type Metrics map[string]Metric

var globalCounter int
var pollInterval = 2 * time.Second
var reportInterval = 10 * time.Second
var serverAddr = "127.0.0.1"
var serverPort = "8080"

func fillMetrics(metrics *Metrics) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	(*metrics)["alloc"] = Metric{
		Name: "Alloc",
		Type: "gauge",
		Val:  float64(memStats.Alloc),
	}
	(*metrics)["buckHashSys"] = Metric{
		Name: "BuckHashSys",
		Type: "gauge",
		Val:  float64(memStats.BuckHashSys),
	}
	(*metrics)["frees"] = Metric{
		Name: "Frees",
		Type: "gauge",
		Val:  float64(memStats.Frees),
	}
	(*metrics)["gCCPUFraction"] = Metric{
		Name: "GCCPUFraction",
		Type: "gauge",
		Val:  memStats.GCCPUFraction,
	}
	(*metrics)["gSys"] = Metric{
		Name: "GSys",
		Type: "gauge",
		Val:  float64(memStats.GCSys),
	}
	(*metrics)["heapAlloc"] = Metric{
		Name: "HeapAlloc",
		Type: "gauge",
		Val:  float64(memStats.HeapAlloc),
	}
	(*metrics)["heapIdle"] = Metric{
		Name: "HeapIdle",
		Type: "gauge",
		Val:  float64(memStats.HeapIdle),
	}
	(*metrics)["heapInuse"] = Metric{
		Name: "HeapInuse",
		Type: "gauge",
		Val:  float64(memStats.HeapInuse),
	}
	(*metrics)["heapObjects"] = Metric{
		Name: "HeapObjects",
		Type: "gauge",
		Val:  float64(memStats.HeapObjects),
	}
	(*metrics)["heapReleased"] = Metric{
		Name: "HeapReleased",
		Type: "gauge",
		Val:  float64(memStats.HeapReleased),
	}
	(*metrics)["heapSys"] = Metric{
		Name: "HeapSys",
		Type: "gauge",
		Val:  float64(memStats.HeapSys),
	}
	(*metrics)["lastGC"] = Metric{
		Name: "LastGC",
		Type: "gauge",
		Val:  float64(memStats.LastGC),
	}
	(*metrics)["lookups"] = Metric{
		Name: "Lookups",
		Type: "gauge",
		Val:  float64(memStats.Lookups),
	}
	(*metrics)["mCacheInuse"] = Metric{
		Name: "MCacheInuse",
		Type: "gauge",
		Val:  float64(memStats.MCacheInuse),
	}
	(*metrics)["mCacheSys"] = Metric{
		Name: "MCacheSys",
		Type: "gauge",
		Val:  float64(memStats.MCacheSys),
	}
	(*metrics)["mSpanInuse"] = Metric{
		Name: "MSpanInuse",
		Type: "gauge",
		Val:  float64(memStats.MSpanInuse),
	}
	(*metrics)["mSpanSys"] = Metric{
		Name: "MSpanSys",
		Type: "gauge",
		Val:  float64(memStats.MSpanSys),
	}
	(*metrics)["mallocs"] = Metric{
		Name: "Mallocs",
		Type: "gauge",
		Val:  float64(memStats.Mallocs),
	}
	(*metrics)["nextGC"] = Metric{
		Name: "NextGC",
		Type: "gauge",
		Val:  float64(memStats.NextGC),
	}
	(*metrics)["numForcedGC"] = Metric{
		Name: "NumForcedGC",
		Type: "gauge",
		Val:  float64(memStats.NumForcedGC),
	}
	(*metrics)["numGC"] = Metric{
		Name: "NumGC",
		Type: "gauge",
		Val:  float64(memStats.NumGC),
	}
	(*metrics)["otherSys"] = Metric{
		Name: "OtherSys",
		Type: "gauge",
		Val:  float64(memStats.OtherSys),
	}
	(*metrics)["pauseTotalNs"] = Metric{
		Name: "PauseTotalNs",
		Type: "gauge",
		Val:  float64(memStats.PauseTotalNs),
	}
	(*metrics)["stackInuse"] = Metric{
		Name: "StackInuse",
		Type: "gauge",
		Val:  float64(memStats.StackInuse),
	}
	(*metrics)["stackSys"] = Metric{
		Name: "StackSys",
		Type: "gauge",
		Val:  float64(memStats.StackSys),
	}
	(*metrics)["sys"] = Metric{
		Name: "Sys",
		Type: "gauge",
		Val:  float64(memStats.Sys),
	}
	(*metrics)["totalAlloc"] = Metric{
		Name: "TotalAlloc",
		Type: "gauge",
		Val:  float64(memStats.TotalAlloc),
	}
	globalCounter++
	(*metrics)["pollCount"] = Metric{
		Name: "PollCount",
		Type: "counter",
		Val:  globalCounter,
	}
	(*metrics)["randomValue"] = Metric{
		Name: "RandomValue",
		Type: "gauge",
		Val:  time.Now().UnixNano(),
	}
}

func reportMetrics(metrics *Metrics) {
	client := &http.Client{}
	for _, metric := range *metrics {
		url := fmt.Sprintf("http://%s:%s/update/%s/%s/%v", serverAddr, serverPort, metric.Type, metric.Name, metric.Val)
		request, err := http.NewRequest(http.MethodPost, url, nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		request.Header.Set("Content-Type", "text/plain")
		response, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Статус-код ", response.Status)
		defer response.Body.Close()
	}
}

func main() {
	metrics := make(Metrics)
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
