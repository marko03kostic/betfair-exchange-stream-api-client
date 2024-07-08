package cache

import (
	"github.com/marko03kostic/betfair-stream-client/model"
)

type Market struct {
	ID               string
	Runners          map[int]*Runner
	TotalVolume      int
	MarketDefinition model.BetfairMarketDefinition
}

func NewMarket(id string) *Market {
	return &Market{
		ID:      id,
		Runners: make(map[int]*Runner),
	}
}

func (m *Market) Update(betfairMarketChange model.BetfairMarketChange) {
	if betfairMarketChange.Rc != nil {
		for _, runnerChange := range betfairMarketChange.Rc {
			id := runnerChange.ID
			_, ok := m.Runners[id]
			if !ok {
				m.AddRunner(id)
			}
			m.Runners[id].Update(runnerChange)

		}
	}
}

func (m *Market) AddRunner(id int) {
	runner := NewRunner(id)
	m.Runners[id] = runner
}
