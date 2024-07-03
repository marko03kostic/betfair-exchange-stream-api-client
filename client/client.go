package client

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/marko03kostic/betfair-stream-client/cache"
	"github.com/marko03kostic/betfair-stream-client/model"
)

type ExchangeStreamClient struct {
	appKey      string
	session     string
	address     string
	conn        net.Conn
	closeCh     chan struct{}
	StatusCache *cache.StatusCache
}

func NewExchangeStreamClient(appKey string, session string, StatusCache *cache.StatusCache) *ExchangeStreamClient {
	return &ExchangeStreamClient{
		appKey:      appKey,
		session:     session,
		address:     "stream-api.betfair.com:443",
		closeCh:     make(chan struct{}),
		StatusCache: StatusCache,
	}
}

func (c *ExchangeStreamClient) Connect() error {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", c.address, conf)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	c.conn = conn
	go c.receiveLoop()
	return nil
}

func (c *ExchangeStreamClient) SendAuthenticationMessage() error {
	authMessage := model.BetfairAuthenticationMessage{
		Op:      "authentication",
		ID:      1,
		AppKey:  c.appKey,
		Session: c.session,
	}

	err := c.send(authMessage)
	if err != nil {
		return fmt.Errorf("failed to send auth message: %w", err)
	}

	return nil
}

func (c *ExchangeStreamClient) SendMarketSubscriptionMessage(marketIds []string) error {

	betfairMarketFilter := model.BetfairMarketFilter{
		MarketIds: marketIds,
	}

	marketSubscriptionMessage := model.BetfairMarketSubscriptionMessage{
		Op:           "marketSubscription",
		ID:           2,
		MarketFilter: betfairMarketFilter,
	}

	err := c.send(marketSubscriptionMessage)
	if err != nil {
		return fmt.Errorf("failed to send auth message: %w", err)
	}

	return nil
}

func (c *ExchangeStreamClient) send(msg model.IBetfairMessage) error {
	if c.conn == nil {
		return fmt.Errorf("not connected")
	}

	id := msg.GetID()

	b, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	b = append(b, "\r\n"...)

	responseChan := make(chan string, 1)
	c.StatusCache.Mu.Lock()
	c.StatusCache.ResponseChans[id] = responseChan
	c.StatusCache.Mu.Unlock()

	_, err = c.conn.Write(b)
	if err != nil {
		c.StatusCache.Mu.Lock()
		delete(c.StatusCache.ResponseChans, id)
		c.StatusCache.Mu.Unlock()
		return fmt.Errorf("failed to send data: %w", err)
	}

	select {
	case status := <-responseChan:
		if status == "SUCCESS" {
			fmt.Printf("Message %v success", id)
			return nil
		} else {
			return fmt.Errorf("operation failed with status: %s", status)
		}
	case <-time.After(5 * time.Second):
		c.StatusCache.Mu.Lock()
		delete(c.StatusCache.ResponseChans, id)
		c.StatusCache.Mu.Unlock()
		return fmt.Errorf("operation timed out")
	}
}

func (c *ExchangeStreamClient) Parse(message string) error {
	var msgMap map[string]interface{}

	err := json.Unmarshal([]byte(message), &msgMap)
	if err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	op, exists := msgMap["op"]
	if !exists {
		return errors.New("missing 'op' field in message")
	}

	opStr, ok := op.(string)
	if !ok {
		return errors.New("'op' field is not a string")
	}

	switch opStr {
	case "status":
		c.StatusCache.Parse(message)
	case "connection":
		fmt.Println("connection")
	case "mcm":
		fmt.Println("mcm")
	case "ocm":
		fmt.Println("ocm")
	default:
		return fmt.Errorf("unknown 'op' value: %s", opStr)
	}

	return nil
}

func (c *ExchangeStreamClient) receiveLoop() {
	reader := bufio.NewReader(c.conn)
	for {
		select {
		case <-c.closeCh:
			return
		default:
			response, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("Error receiving data: %v", err)
				return
			}
			c.Parse(response)
		}
	}
}

func (c *ExchangeStreamClient) Close() {
	if c.conn != nil {
		close(c.closeCh)
		c.conn.Close()
	}
}
