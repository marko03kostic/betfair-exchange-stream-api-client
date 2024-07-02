package main

import (
	"log"
	"time"
	"github.com/marko03kostic/betfair-stream-client/client"
)

func main() {
	config, err := client.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	c := client.NewExchangeStreamClient(config.AppKey, config.Session)

	err1 := c.Connect()
	if err1 != nil {
		log.Fatalf("failed to connect: %v", err1)
	}
	defer c.Close()

	c.SendAuthenticationMessage()

	marketIds := [1]string{"1.230181554"}

	c.SendMarketSubscriptionMessage(marketIds[:])

	time.Sleep(20 * time.Second)
}
