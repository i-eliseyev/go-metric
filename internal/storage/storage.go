package storage

import (
	"github.com/i-eliseyev/go-metric/internal/common"
	"log"
	"strconv"
)

type Storage interface {
	CreateStorage() *Storage
	UpdateGauge(m *common.Metric) error
	UpdateCounter(m *common.Metric) error
}

type MemStorage struct {
	Metrics common.Metrics
}

func CreateStorage() *MemStorage {
	obj := MemStorage{}
	obj.Metrics = make(map[string]common.Metric)
	return &obj
}

func (ms *MemStorage) UpdateGauge(m common.Metric) error {
	ms.Metrics[m.Name] = m
	log.Printf("Updated %s value: %s", m.Name, ms.Metrics[m.Name].Val)
	return nil
}

func (ms *MemStorage) UpdateCounter(m common.Metric) error {
	oldMetric, ok := ms.Metrics[m.Name]
	newValue, err := strconv.Atoi(m.Val.(string))
	if err != nil {
		log.Printf("Unexpected type of counter value: %s", m.Val)
		return err
	}
	if !ok {
		ms.Metrics[m.Name] = common.Metric{
			Name: m.Name,
			Type: m.Type,
			Val:  newValue,
		}
	} else {
		ms.Metrics[m.Name] = common.Metric{
			Name: m.Name,
			Type: m.Type,
			Val:  oldMetric.Val.(int) + newValue,
		}
	}
	log.Printf("Updated %s value: %d", m.Name, ms.Metrics[m.Name].Val)
	return nil
}

var MetricStorage = CreateStorage()
