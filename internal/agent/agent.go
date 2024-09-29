package agent

import (
	"github.com/go-resty/resty/v2"
	"github.com/i-eliseyev/go-metric/internal/common"
	"github.com/i-eliseyev/go-metric/internal/utils"
	"log"
	"os"
	"runtime"
	"time"
)

func FillMetrics(metrics *common.Metrics) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	(*metrics)["alloc"] = common.Metric{
		ID:    "Alloc",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.Alloc),
	}
	(*metrics)["buckHashSys"] = common.Metric{
		ID:    "BuckHashSys",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.BuckHashSys),
	}
	(*metrics)["frees"] = common.Metric{
		ID:    "Frees",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.Frees),
	}
	(*metrics)["gCCPUFraction"] = common.Metric{
		ID:    "GCCPUFraction",
		MType: "gauge",
		Value: &memStats.GCCPUFraction,
	}
	(*metrics)["gSys"] = common.Metric{
		ID:    "GSys",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.GCSys),
	}
	(*metrics)["heapAlloc"] = common.Metric{
		ID:    "HeapAlloc",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.HeapAlloc),
	}
	(*metrics)["heapIdle"] = common.Metric{
		ID:    "HeapIdle",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.HeapIdle),
	}
	(*metrics)["heapInuse"] = common.Metric{
		ID:    "HeapInuse",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.HeapInuse),
	}
	(*metrics)["heapObjects"] = common.Metric{
		ID:    "HeapObjects",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.HeapObjects),
	}
	(*metrics)["heapReleased"] = common.Metric{
		ID:    "HeapReleased",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.HeapReleased),
	}
	(*metrics)["heapSys"] = common.Metric{
		ID:    "HeapSys",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.HeapSys),
	}
	(*metrics)["lastGC"] = common.Metric{
		ID:    "LastGC",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.LastGC),
	}
	(*metrics)["lookups"] = common.Metric{
		ID:    "Lookups",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.Lookups),
	}
	(*metrics)["mCacheInuse"] = common.Metric{
		ID:    "MCacheInuse",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.MCacheInuse),
	}
	(*metrics)["mCacheSys"] = common.Metric{
		ID:    "MCacheSys",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.MCacheSys),
	}
	(*metrics)["mSpanInuse"] = common.Metric{
		ID:    "MSpanInuse",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.MSpanInuse),
	}
	(*metrics)["mSpanSys"] = common.Metric{
		ID:    "MSpanSys",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.MSpanSys),
	}
	(*metrics)["mallocs"] = common.Metric{
		ID:    "Mallocs",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.Mallocs),
	}
	(*metrics)["nextGC"] = common.Metric{
		ID:    "NextGC",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.NextGC),
	}
	(*metrics)["numForcedGC"] = common.Metric{
		ID:    "NumForcedGC",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(uint64(memStats.NumForcedGC)),
	}
	(*metrics)["numGC"] = common.Metric{
		ID:    "NumGC",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(uint64(memStats.NumGC)),
	}
	(*metrics)["otherSys"] = common.Metric{
		ID:    "OtherSys",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.OtherSys),
	}
	(*metrics)["pauseTotalNs"] = common.Metric{
		ID:    "PauseTotalNs",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.PauseTotalNs),
	}
	(*metrics)["stackInuse"] = common.Metric{
		ID:    "StackInuse",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.StackInuse),
	}
	(*metrics)["stackSys"] = common.Metric{
		ID:    "StackSys",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.StackSys),
	}
	(*metrics)["sys"] = common.Metric{
		ID:    "Sys",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.Sys),
	}
	(*metrics)["totalAlloc"] = common.Metric{
		ID:    "TotalAlloc",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(memStats.TotalAlloc),
	}
	GlobalPollsCounter++
	(*metrics)["pollCount"] = common.Metric{
		ID:    "PollCount",
		MType: "counter",
		Delta: &GlobalPollsCounter,
	}
	(*metrics)["randomValue"] = common.Metric{
		ID:    "RandomValue",
		MType: "gauge",
		Value: utils.UInt64ToFloat64Ptr(uint64(time.Now().UnixNano())),
	}
}

func ReportMetrics(metrics *common.Metrics) {
	client := resty.New()
	client.SetRetryCount(RetryCount).SetRetryWaitTime(RetryWaitTime)

	for _, metric := range *metrics {

		pathParams := map[string]string{
			"baseURL": ServerAddr,
			"port":    ServerPort,
		}
		responseMetric := new(common.Metric)

		resp, err := client.
			R().
			SetBody(metric).
			SetPathParams(pathParams).
			SetResult(responseMetric).
			Post("http://{baseURL}:{port}/update/")

		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		log.Println("Status: ", resp.Status())
		log.Println("Response: ", responseMetric)
	}
}
