package domain

type LatencyTester interface {
	Run(host string) (LatencyTest, error)
}
