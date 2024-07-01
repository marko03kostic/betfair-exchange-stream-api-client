package main

import (
	"log"
	"time"
)

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	client := NewExchangeStreamClient(config.AppKey, config.Session)

	err1 := client.Connect()
	if err1 != nil {
		log.Fatalf("failed to connect: %v", err1)
	}
	defer client.Close()

	client.SendAuthenticationMessage()

	marketIds := [1]string{"1.230179589"}

	client.SendMarketSubscriptionMessage(marketIds[:])

	time.Sleep(20 * time.Second)
}
