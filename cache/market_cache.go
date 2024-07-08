package cache

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/marko03kostic/betfair-stream-client/model"
)

type MarketCache struct {
	Markets            map[string]*Market
	HeartbeatThreshold time.Duration
	InitialClk         string
	Clk                string
	mu                 sync.Mutex
	timer              *time.Timer
	latencyThreshold   time.Duration
}

func NewMarketCache() *MarketCache {
	return &MarketCache{
		Markets:            make(map[string]*Market),
		HeartbeatThreshold: time.Duration(5000) * time.Millisecond,
		latencyThreshold:   time.Duration(5) * time.Millisecond,
	}
}

func (m *MarketCache) Parse(msg string) error {
	var betfairMarketChangeMessage model.BetfairMarketChangeMessage

	err := json.Unmarshal([]byte(msg), &betfairMarketChangeMessage)
	if err != nil {
		fmt.Println(err)
		return err
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	timeSent := time.Unix(0, betfairMarketChangeMessage.Pt*int64(time.Millisecond))
	timeSinceSent := time.Since(timeSent)

	if betfairMarketChangeMessage.Status != "" {
		return fmt.Errorf(betfairMarketChangeMessage.Status)
	}

	if betfairMarketChangeMessage.Clk != "" {
		m.Clk = betfairMarketChangeMessage.Clk
	}

	if betfairMarketChangeMessage.HeartbeatMs != 0 {
		m.HeartbeatThreshold = time.Duration(int64(float64(betfairMarketChangeMessage.HeartbeatMs) * 1.05)) * time.Millisecond
	}

	switch betfairMarketChangeMessage.Ct {
	case "SUB_IMAGE":
		fmt.Print("")
	case "RESUB_DELTA":
		fmt.Print("")
	case "HEARTBEAT":
		fmt.Print("")
	case "":
		fmt.Print("")
	}

	if betfairMarketChangeMessage.Mc != nil {
		for _, marketChange := range betfairMarketChangeMessage.Mc {
			id := marketChange.ID
			market, ok := m.Markets[id]
			if ok {
				market.Update(marketChange)
			} else {
				log.Fatalf("market %v not in markets", id)
			}
		}
	}

	if timeSinceSent > m.latencyThreshold {
		return fmt.Errorf("high latency")
	}

	m.resetTimer()

	return nil
}

func (m *MarketCache) resetTimer() {
	if m.timer != nil {
		m.timer.Stop()
	}
	m.timer = time.AfterFunc(m.HeartbeatThreshold, func() {
		m.mu.Lock()
		defer m.mu.Unlock()
		fmt.Println("Error: heartbeat missed")
	})
}

func (m *MarketCache) AddMarket(id string) {
	market := NewMarket(id)
	m.Markets[id] = market
}
