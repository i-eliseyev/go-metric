package main

import (
	"github.com/i-eliseyev/go-metric/internal/agent"
	"github.com/i-eliseyev/go-metric/internal/common"
	"time"
)

func main() {
	metrics := make(common.Metrics)
	fillMetricsTicker := time.NewTicker(agent.PollInterval)
	reportMetricsTicker := time.NewTicker(agent.ReportInterval)
	defer fillMetricsTicker.Stop()
	defer reportMetricsTicker.Stop()

	for {
		select {
		case <-fillMetricsTicker.C:
			agent.FillMetrics(&metrics)
		case <-reportMetricsTicker.C:
			agent.ReportMetrics(&metrics)
		}
	}
}
