package cache

import (
	"encoding/json"
	"github.com/marko03kostic/betfair-stream-client/model"
)

type StatusCache struct {
	ConnectionClosed     bool
	ConnectionsAvailable int
}

func NewStatusCache() *StatusCache {
	return &StatusCache{}
}

func (s *StatusCache) Parse(message string) error {

	var betfairStatusMessage model.BetfairStatusMessage

	err := json.Unmarshal([]byte(message), &betfairStatusMessage)
	if err != nil {
		return err
	}

	s.ConnectionClosed = betfairStatusMessage.ConnectionClosed

	if betfairStatusMessage.ConnectionsAvailable != nil {
		s.ConnectionsAvailable = *betfairStatusMessage.ConnectionsAvailable
	}

	return nil

}
