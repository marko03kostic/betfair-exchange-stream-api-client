package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type ExchangeStreamClient struct {
	appKey string
	session string
	address string
	conn    net.Conn
	closeCh chan struct{}
}

func NewExchangeStreamClient(appKey string, session string) *ExchangeStreamClient {
	return &ExchangeStreamClient{
		appKey: appKey,
		session: session,
		address: "stream-api.betfair.com:443",
		closeCh: make(chan struct{}),
	}
}

func (client *ExchangeStreamClient) Connect() error {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", client.address, conf)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	client.conn = conn
	go client.receiveLoop()
	return nil
}

func (client *ExchangeStreamClient) SendAuthenticationMessage() error {
	authMessage := BetfairAuthenticationMessage{
		Op:      "authentication",
		ID:      1,
		AppKey:  client.appKey,
		Session: client.session,
	}

	err := client.send(authMessage)
	if err != nil {
		return fmt.Errorf("failed to send auth message: %w", err)
	}

	return nil
}

func (client *ExchangeStreamClient) SendMarketSubscriptionMessage(marketIds []string) error {
	
	betfairMarketFilter := BetfairMarketFilter{
		MarketIds: marketIds,
	}
	
	
	marketSubscriptionMessage := BetfairMarketSubscriptionMessage{
		Op: "marketSubscription",
		ID: 2,
		MarketFilter: betfairMarketFilter,
	}

	err := client.send(marketSubscriptionMessage)
	if err != nil {
		return fmt.Errorf("failed to send auth message: %w", err)
	}

	return nil
}

func (client *ExchangeStreamClient) send(data any) error {
	if client.conn == nil {
		return fmt.Errorf("not connected")
	}

	b, err2 := json.Marshal(data)
	if err2 != nil {
		log.Fatalf("failed to connect: %v", err2)
	}

	b = append(b, "\r\n"...)

	_, err := client.conn.Write(b)
	if err != nil {
		return fmt.Errorf("failed to send data: %w", err)
	}

	return nil
}

func (client *ExchangeStreamClient) receiveLoop() {
	reader := bufio.NewReader(client.conn)
	for {
		select {
		case <-client.closeCh:
			return
		default:
			response, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("Error receiving data: %v", err)
				return
			}
			fmt.Println(response)
		}
	}
}

func (client *ExchangeStreamClient) Close() {
	if client.conn != nil {
		close(client.closeCh)
		client.conn.Close()
	}
}
