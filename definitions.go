package main

type BetfairAuthenticationMessage struct {
	Op      string `json:"op"`
	ID      int    `json:"id"`
	AppKey  string `json:"appKey"`
	Session string `json:"session"`
}

type BetfairMarketSubscriptionMessage struct {
	Op string `json:"op"`
	ID int    `json:"id"`
	MarketFilter BetfairMarketFilter `json:"marketFilter"`
}

type BetfairMarketFilter struct {
    MarketIds []string `json:"marketIds"`
}