package main

import (
	"fmt"
	"log"
	"time"

	"github.com/marko03kostic/betfair-stream-client/cache"
	"github.com/marko03kostic/betfair-stream-client/client"
)

func main() {
	config, err := client.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	mc := cache.NewMarketCache()

	sc := cache.NewStatusCache()

	c := client.NewExchangeStreamClient(config.AppKey, config.Session, sc, mc)

	err1 := c.Connect()
	if err1 != nil {
		log.Fatalf("failed to connect: %v", err1)
	}
	defer c.Close()

	err3 := c.SendAuthenticationMessage()
	if err3 != nil {
		log.Fatalf("failed to connect: %v", err3)
	}

	mc.AddMarket("1.230452417")

	marketIds := [1]string{"1.230452417"}

	err4 := c.SendMarketSubscriptionMessage(marketIds[:])
	if err4 != nil {
		log.Fatalf("failed to connect: %v", err4)
	}
	time.Sleep(200 * time.Second)
	fmt.Println(mc.Markets)
}
