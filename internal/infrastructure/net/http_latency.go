package netinfra

import (
	"net"
	"time"

	"github.com/AmrmDev/latency-tester/internal/domain"
)

type HttpLatencyTester struct{}

func NewHttpLatencyTester() *HttpLatencyTester {
	return &HttpLatencyTester{}
}

func (t *HttpLatencyTester) Run(host string) (domain.LatencyTest, error) {
	start := time.Now()

	conn, err := net.DialTimeout("tcp", host+":80", 3*time.Second)
	if err != nil {
		return domain.LatencyTest{}, err
	}
	defer conn.Close()

	duration := time.Since(start)

	return domain.LatencyTest{
		Host: host,
		Duration:  duration,
		Timestamp: time.Now(),
	}, nil

}