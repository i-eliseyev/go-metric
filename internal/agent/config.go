package agent

import "time"

var GlobalCounter float64

const (
	PollInterval   = 2 * time.Second
	ReportInterval = 10 * time.Second
	ServerAddr     = "127.0.0.1"
	ServerPort     = "8080"
	RetryCount     = 4
	RetryWaitTime  = 1 * time.Second
)
