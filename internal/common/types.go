package common

type Metric struct {
	Name string
	Type string
	Val  interface{}
}

type Metrics map[string]Metric
