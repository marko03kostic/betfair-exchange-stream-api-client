package cache

import (
	"fmt"

	"github.com/marko03kostic/betfair-stream-client/model"
)

type Runner struct {
	ID              int
	FullPriceLadder map[string]map[float32]float32
}

func NewRunner(id int) *Runner {
	return &Runner{
		ID:              id,
		FullPriceLadder: make(map[string]map[float32]float32),
	}
}

func (r *Runner) Update(runnerChange model.BetfairRunnerChange) {
	fmt.Println(runnerChange)
}
