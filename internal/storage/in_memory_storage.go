package storage

import (
	"errors"
	"github.com/i-eliseyev/go-metric/internal/common"
	"log"
)

type MemStorage struct {
	Metrics common.Metrics
}

var MetricStorage *MemStorage

func init() {
	MetricStorage = &MemStorage{
		Metrics: make(map[string]common.Metric),
	}
}

func (ms *MemStorage) UpdateGauge(m *common.Metric) {
	ms.Metrics[m.Name] = *m
	log.Printf("Updated gauge %s value: %f", m.Name, m.Val)
}

func (ms *MemStorage) UpdateCounter(m *common.Metric) {
	if existingMetric, ok := ms.Metrics[m.Name]; ok {
		ms.Metrics[m.Name] = common.Metric{
			Name: m.Name,
			Type: m.Type,
			Val:  existingMetric.Val + m.Val,
		}
	} else {
		ms.Metrics[m.Name] = common.Metric{
			Name: m.Name,
			Type: m.Type,
			Val:  m.Val,
		}
	}
	log.Printf("Updated counter %s value: %f", m.Name, ms.Metrics[m.Name].Val)
}

func (ms *MemStorage) GetMetric(name string) (*common.Metric, error) {
	if m, ok := ms.Metrics[name]; ok {
		return &m, nil
	}

	log.Printf("Metric %s not found", name)
	return nil, errors.New("metric not found")
}
