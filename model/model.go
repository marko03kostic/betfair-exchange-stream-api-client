package model

type BetfairStatusMessage struct {
	Op                   string  `json:"op"`
	ID                   int     `json:"id"`
	StatusCode           string  `json:"statusCode"`
	ConnectionClosed     bool    `json:"connectionClosed"`
	ErrorCode            *string `json:"errorCode,omitempty"`
	ErrorMessage         *string `json:"errorMessage,omitempty"`
	ConnectionsAvailable *int    `json:"connectionsAvailable,omitempty"`
}

type BetfairAuthenticationMessage struct {
	Op      string `json:"op"`
	ID      int    `json:"id"`
	AppKey  string `json:"appKey"`
	Session string `json:"session"`
}

type BetfairConnectionMessage struct {
	Op           string `json:"op"`
	ID           int    `json:"id"`
	ConnectionId string `json:"connectionId"`
}

type BetfairSubscriptionMessage struct {
	Op                  string  `json:"op"`
	ID                  int     `json:"id"`
	SegmentationEnabled *bool   `json:"segmentationEnabled,omitempty"`
	ConflateMs          *int    `json:"conflateMs,omitempty"`
	HeartbeatMs         *int    `json:"heartbeatMs,omitempty"`
	InitialClk          *string `json:"initialClk,omitempty"`
	Clk                 *string `json:"clk,omitempty"`
}

type BetfairMarketFilter struct {
	CountryCodes      []string `json:"countryCodes,omitempty"`
	BettingTypes      []string `json:"bettingTypes,omitempty"`
	TurnInPlayEnabled *bool    `json:"turnInPlayEnabled,omitempty"`
	MarketTypes       []string `json:"marketTypes,omitempty"`
	Venues            []string `json:"venues,omitempty"`
	MarketIds         []string `json:"marketIds,omitempty"`
	EventTypeIds      []string `json:"eventTypeIds,omitempty"`
	EventIds          []string `json:"eventIds,omitempty"`
	BspMarket         *bool    `json:"bspMarket,omitempty"`
	RaceTypes         []string `json:"raceTypes,omitempty"`
}

type BetfairMarketDataFilter struct {
	LadderLevels *int     `json:"ladderLevels,omitempty"`
	Fields       []string `json:"fields,omitempty"`
}

type BetfairOrderFilter struct {
	IncludeOverallPosition        *bool    `json:"includeOverallPosition,omitempty"`
	AccountIds                    []int    `json:"accountIds,omitempty"`
	CustomerStrategyRefs          []string `json:"customerStrategyRefs,omitempty"`
	PartitionMatchedByStrategyRef *bool    `json:"partitionMatchedByStrategyRef,omitempty"`
}

type BetfairMarketSubscriptionMessage struct {
	Op                  string                  `json:"op"`
	ID                  int                     `json:"id"`
	SegmentationEnabled *bool                   `json:"segmentationEnabled,omitempty"`
	ConflateMs          *int                    `json:"conflateMs,omitempty"`
	HeartbeatMs         *int                    `json:"heartbeatMs,omitempty"`
	InitialClk          *string                 `json:"initialClk,omitempty"`
	Clk                 *string                 `json:"clk,omitempty"`
	MarketFilter        BetfairMarketFilter     `json:"marketFilter"`
	MarketDataFilter    BetfairMarketDataFilter `json:"marketDataFilter"`
}

type BetfairOrderSubscriptionMessage struct {
	Op                  string             `json:"op"`
	ID                  int                `json:"id"`
	SegmentationEnabled *bool              `json:"segmentationEnabled,omitempty"`
	ConflateMs          *int               `json:"conflateMs,omitempty"`
	HeartbeatMs         *int               `json:"heartbeatMs,omitempty"`
	InitialClk          *string            `json:"initialClk,omitempty"`
	Clk                 *string            `json:"clk,omitempty"`
	OrderFilter         BetfairOrderFilter `json:"orderFilter"`
}

type BetfairChangeMessage struct {
	Op          string `json:"op"`
	ID          int    `json:"id"`
	Ct          string `json:"ct"`
	SegmentType string `json:"segmentType"`
	ConflateMs  *int   `json:"conflateMs,omitempty"`
	Status      string `json:"status"`
	HeartbeatMs int    `json:"heartbeatMs"`
	Pt          int    `json:"pt"`
	InitialClk  string `json:"initialClk"`
	Clk         string `json:"clk"`
}

type BetfairMarketDefinition struct {
	ID                    string  `json:"id"`
	Venue                 string  `json:"venue"`
	BspMarket             bool    `json:"bspMarket"`
	TurnInPlayEnabled     bool    `json:"turnInPlayEnabled"`
	PersistenceEnabled    bool    `json:"persistenceEnabled"`
	MarketBaseRate        float64 `json:"marketBaseRate"`
	EventId               string  `json:"eventId"`
	EventTypeId           string  `json:"eventTypeId"`
	NumberOfWinners       int     `json:"numberOfWinners"`
	BettingType           string  `json:"bettingType"`
	MarketType            string  `json:"marketType"`
	MarketTime            string  `json:"marketTime"`
	SuspendTime           string  `json:"suspendTime"`
	BspReconciled         bool    `json:"bspReconciled"`
	Complete              bool    `json:"complete"`
	InPlay                bool    `json:"inPLay"`
	CrossMatching         bool    `json:"crossMatching"`
	RunnersVoidable       bool    `json:"runnersVoidable"`
	NumberOfActiveRunners int     `json:"numberOfActiveRunners"`
	BetDelay              int     `json:"betDelay"`
	Status                string  `json:"status"`
	Regulators            string  `json:"regulators"`
	DiscountAllowed       bool    `json:"discountAllowed"`
	Timezone              string  `json:"timezone"`
}

type BetfairRunnerChange struct {
	ID    int         `json:"id"`
	Con   bool        `json:"con"`
	Tv    float64     `json:"tv"`
	Ltp   float64     `json:"ltp"`
	Spn   float64     `json:"spn"`
	Spf   float64     `json:"spf"`
	Batb  [][]float64 `json:"batb"`
	Batl  [][]float64 `json:"batl"`
	Bdatb [][]float64 `json:"bdatb"`
	Bdatl [][]float64 `json:"bdatl"`
	Atb   [][]float64 `json:"atb"`
	Atl   [][]float64 `json:"atl"`
	Spb   [][]float64 `json:"spb"`
	Spl   [][]float64 `json:"spl"`
	Trd   [][]float64 `json:"trd"`
	Hc    float64     `json:"hc"`
}

type BetfairMarketChange struct {
	Rc               []BetfairRunnerChange   `json:"rc"`
	Img              bool                    `json:"img"`
	Tv               float64                 `json:"tv"`
	MarketDefinition BetfairMarketDefinition `json:"marketDefinition"`
	ID               string                  `json:"id"`
}

type BetfairUnmatchedOrder struct {
	ID     string  `json:"id"`
	P      float64 `json:"p"`
	S      float64 `json:"s"`
	Bsp    float64 `json:"bsp"`
	Side   string  `json:"side"`
	Status string  `json:"status"`
	Pt     string  `json:"pt"`
	Ot     string  `json:"ot"`
	Pd     int     `json:"pd"`
	Md     int     `json:"md"`
	Cd     int     `json:"cd"`
	Ld     int     `json:"ld"`
	Lsrc   string  `json:"lsrc"`
	Avp    float64 `json:"avp"`
	Sm     float64 `json:"sm"`
	Sr     float64 `json:"sr"`
	Sl     float64 `json:"sl"`
	Sc     float64 `json:"sc"`
	Sv     float64 `json:"sv"`
	Rac    string  `json:"rac"`
	Rc     string  `json:"rc"`
	Rfo    string  `json:"rfo"`
	Rfs    string  `json:"rfs"`
}

type BetfairOrderChange struct {
	FullImage bool                    `json:"fullImage"`
	ID        string                  `json:"id"`
	Hc        *string                 `json:"hc,omitempty"`
	Uo        []BetfairUnmatchedOrder `json:"uo"`
	Mb        [][]float64             `json:"mb"`
	Ml        [][]float64             `json:"ml"`
}

type BetfairOrderAccountChange struct {
	Closed    bool                 `json:"closed"`
	ID        string               `json:"id"`
	FullImage bool                 `json:"fullImage"`
	Orc       []BetfairOrderChange `json:"orc"`
}

type BetfairMarketChangeMessage struct {
	Op          string                `json:"op"`
	ID          int                   `json:"id"`
	Ct          string                `json:"ct"`
	SegmentType string                `json:"segmentType"`
	ConflateMs  *int                  `json:"conflateMs,omitempty"`
	Status      string                `json:"status"`
	HeartbeatMs int                   `json:"heartbeatMs"`
	Pt          int                   `json:"pt"`
	InitialClk  string                `json:"initialClk"`
	Clk         string                `json:"clk"`
	Mc          []BetfairMarketChange `json:"mc"`
}

type BetfairOrderChangeMessage struct {
	Op          string                      `json:"op"`
	ID          int                         `json:"id"`
	Ct          string                      `json:"ct"`
	SegmentType string                      `json:"segmentType"`
	ConflateMs  *int                        `json:"conflateMs,omitempty"`
	Status      string                      `json:"status"`
	HeartbeatMs int                         `json:"heartbeatMs"`
	Pt          int                         `json:"pt"`
	InitialClk  string                      `json:"initialClk"`
	Clk         string                      `json:"clk"`
	Oc          []BetfairOrderAccountChange `json:"oc"`
}
