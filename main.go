package main

import (
	"github.com/marko03kostic/betfair-stream-client/client"
	"github.com/marko03kostic/betfair-stream-client/cache"
	"log"
	"time"
)

func main() {
	config, err := client.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	sc := cache.NewStatusCache()

	c := client.NewExchangeStreamClient(config.AppKey, config.Session, sc)

	err1 := c.Connect()
	if err1 != nil {
		log.Fatalf("failed to connect: %v", err1)
	}
	defer c.Close()

	c.SendAuthenticationMessage()

	marketIds := [1]string{"1.230332162"}

	c.SendMarketSubscriptionMessage(marketIds[:])

	time.Sleep(20 * time.Second)
}
