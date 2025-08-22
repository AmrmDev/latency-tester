package usecase

import (
	"github.com/AmrmDev/latency-tester/internal/domain"
)

type LatencyTesterUseCase struct {
	Tester domain.LatencyTester
}

func NewLatencyTesterUseCase(tester domain.LatencyTester) *LatencyTesterUseCase {
	return &LatencyTesterUseCase{Tester: tester}
}

func (uc *LatencyTesterUseCase) Execute (host string) (domain.LatencyTest, error) {
	return uc.Tester.Run(host)
}