package domain

import (
	"time"
)

type LatencyTest struct {
	Host string
	Duration time.Duration
	Timestamp time.Time
}



