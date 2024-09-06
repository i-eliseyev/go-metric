package storage

import "github.com/i-eliseyev/go-metric/internal/common"

type Storage interface {
	UpdateGauge(m *common.Metric)
	UpdateCounter(m *common.Metric)
	GetMetric(name string) (*common.Metric, error)
}
