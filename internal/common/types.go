package common

type Metric struct {
	Name string
	Type string
	Val  float64
}

type Metrics map[string]Metric
