package domains

type Nation interface {
	SpawnPopulationWorkers()
	TotalTokensAvailable() uint
}
