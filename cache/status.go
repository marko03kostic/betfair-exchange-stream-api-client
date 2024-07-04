package cache

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/marko03kostic/betfair-stream-client/model"
)

type StatusCache struct {
	ConnectionClosed     bool
	ConnectionsAvailable int
	ResponseChans        map[int]chan bool
	Mu                   sync.Mutex
}

func NewStatusCache() *StatusCache {
	return &StatusCache{
		ResponseChans: make(map[int]chan bool),
	}
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

	s.Mu.Lock()
	defer s.Mu.Unlock()

	if ch, ok := s.ResponseChans[betfairStatusMessage.ID]; ok {
		var Err error
		switch betfairStatusMessage.StatusCode {
		case "SUCCESS":
			ch <- true
		case "FAILURE":
			ch <- false
			if betfairStatusMessage.ErrorMessage != nil {
				Err = errors.New(*betfairStatusMessage.ErrorMessage)
			} else {
				Err = errors.New("betfair status message indicates failure")
			}
		default:
			ch <- false
			Err = errors.New("betfair status message indicates an unknown status code")
		}
		close(ch)
		delete(s.ResponseChans, betfairStatusMessage.ID)
		return Err
	}

	return nil
}
