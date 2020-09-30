package main

import (
	"fmt"

	"github.com/slimcdk/nin-graph/pkg/domains/dnk"
)

func main() {
	fmt.Printf("Populator!\n")

	dnk.StartPopulationWorkers()
}
