package storage

import (
	"errors"
	"github.com/i-eliseyev/go-metric/internal/common"
	"github.com/i-eliseyev/go-metric/internal/utils"
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
	ms.Metrics[m.ID] = *m
	log.Printf("Updated gauge %s value: %f", m.ID, *m.Value)
}

func (ms *MemStorage) UpdateCounter(m *common.Metric) {
	if existingMetric, ok := ms.Metrics[m.ID]; ok {
		ms.Metrics[m.ID] = common.Metric{
			ID:    m.ID,
			MType: m.MType,
			Delta: utils.AddFloat64Ptr(existingMetric.Delta, m.Delta),
		}
	} else {
		ms.Metrics[m.ID] = common.Metric{
			ID:    m.ID,
			MType: m.MType,
			Delta: m.Delta,
		}
	}
	log.Printf("Updated counter %s value: %d", m.ID, *ms.Metrics[m.ID].Delta)
}

func (ms *MemStorage) GetMetric(name string) (*common.Metric, error) {
	if m, ok := ms.Metrics[name]; ok {
		return &m, nil
	}

	log.Printf("Metric %s not found", name)
	return nil, errors.New("metric not found")
}
