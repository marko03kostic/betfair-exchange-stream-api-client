package cache

import (
	"github.com/marko03kostic/betfair-stream-client/model"
)

type Runner struct {
	ID              int
	FullPriceLadder map[string]map[float32]float32
	SingleValues map[string]float32
}

func NewRunner(id int) *Runner {
	return &Runner{
		ID:              id,
		FullPriceLadder: make(map[string]map[float32]float32),
		SingleValues: make(map[string]float32),
	}
}

func (r *Runner) Update(runnerChange model.BetfairRunnerChange) {

	r.UpdateFullPriceLadder(runnerChange.Atb, "atb")
	r.UpdateFullPriceLadder(runnerChange.Atl, "atl")
	r.UpdateFullPriceLadder(runnerChange.Trd, "trd")
	r.UpdateFullPriceLadder(runnerChange.Spb, "spb")
	r.UpdateFullPriceLadder(runnerChange.Spl, "spl")

	r.UpdateSingleValue(runnerChange.Tv, "tv")
	r.UpdateSingleValue(runnerChange.Ltp, "ltp")
	r.UpdateSingleValue(runnerChange.Spn, "spn")
	r.UpdateSingleValue(runnerChange.Spf, "spf")
	r.UpdateSingleValue(runnerChange.Hc, "hc")
}

func (r *Runner) UpdateFullPriceLadder(list [][2]float32, selection string) {
	if list == nil {
		return
	}

	if _, ok := r.FullPriceLadder[selection]; !ok {
		r.FullPriceLadder[selection] = make(map[float32]float32)
	}

	for _, i := range list {
		price := i[0]
		size := i[1]
		r.FullPriceLadder[selection][price] = size
	}
}

func (r *Runner) UpdateLevelBasedLadder(list [][3]float32, selection string) {
	if list == nil {
		return
	}
	
}

func (r *Runner) UpdateSingleValue(singleValue float32, selection string) {
	if singleValue != 0{
		r.SingleValues[selection] = singleValue
	}
}