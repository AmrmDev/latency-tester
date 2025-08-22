package domain

type LatencyRepository interface {
	Save(test LatencyTest) error
	findAll() ([]LatencyTest, error)
}