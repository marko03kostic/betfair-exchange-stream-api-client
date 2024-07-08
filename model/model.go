package model

type IBetfairMessage interface {
	SetID(int)
}

type BetfairStatusMessage struct {
	Op                   string  `json:"op"`
	ID                   int     `json:"id"`
	StatusCode           string  `json:"statusCode"`
	ConnectionClosed     bool    `json:"connectionClosed"`
	ErrorCode            *string `json:"errorCode,omitempty"`
	ErrorMessage         *string `json:"errorMessage,omitempty"`
	ConnectionsAvailable *int    `json:"connectionsAvailable,omitempty"`
}

func (msg *BetfairStatusMessage) SetID(id int) {
	msg.ID = id
}

type BetfairAuthenticationMessage struct {
	Op      string `json:"op"`
	ID      int    `json:"id"`
	AppKey  string `json:"appKey"`
	Session string `json:"session"`
}

func (msg *BetfairAuthenticationMessage) SetID(id int) {
	msg.ID = id
}

type BetfairConnectionMessage struct {
	Op           string `json:"op"`
	ID           int    `json:"id"`
	ConnectionId string `json:"connectionId"`
}

func (msg *BetfairConnectionMessage) SetID(id int) {
	msg.ID = id
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

func (msg *BetfairMarketSubscriptionMessage) SetID(id int) {
	msg.ID = id
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

func (msg *BetfairOrderSubscriptionMessage) SetID(id int) {
	msg.ID = id
}

type BetfairChangeMessage struct {
	Op          string `json:"op"`
	ID          int    `json:"id"`
	Ct          string `json:"ct"`
	SegmentType string `json:"segmentType"`
	ConflateMs  *int   `json:"conflateMs,omitempty"`
	Status      string `json:"status"`
	HeartbeatMs int    `json:"heartbeatMs"`
	Pt          int64  `json:"pt"`
	InitialClk  string `json:"initialClk"`
	Clk         string `json:"clk"`
}

type BetfairMarketDefinition struct {
	Venue                 string          `json:"venue,omitempty"`
	BspMarket             bool            `json:"bspMarket,omitempty"`
	TurnInPlayEnabled     bool            `json:"turnInPlayEnabled,omitempty"`
	PersistenceEnabled    bool            `json:"persistenceEnabled,omitempty"`
	MarketBaseRate        float64         `json:"marketBaseRate,omitempty"`
	EventId               string          `json:"eventId,omitempty"`
	EventTypeId           string          `json:"eventTypeId,omitempty"`
	NumberOfWinners       int             `json:"numberOfWinners,omitempty"`
	BettingType           string          `json:"bettingType,omitempty"`
	MarketType            string          `json:"marketType,omitempty"`
	MarketTime            string          `json:"marketTime,omitempty"`
	SuspendTime           string          `json:"suspendTime,omitempty"`
	BspReconciled         bool            `json:"bspReconciled,omitempty"`
	Complete              bool            `json:"complete,omitempty"`
	InPlay                bool            `json:"inPlay,omitempty"`
	CrossMatching         bool            `json:"crossMatching,omitempty"`
	RunnersVoidable       bool            `json:"runnersVoidable,omitempty"`
	NumberOfActiveRunners int             `json:"numberOfActiveRunners,omitempty"`
	BetDelay              int             `json:"betDelay,omitempty"`
	Status                string          `json:"status,omitempty"`
	Runners               []BetfairRunner `json:"runners,omitempty"`
	Regulators            []string        `json:"regulators,omitempty"`
	DiscountAllowed       bool            `json:"discountAllowed,omitempty"`
	Timezone              string          `json:"timezone,omitempty"`
	OpenDate              string          `json:"openDate,omitempty"`
	Version               int64           `json:"version,omitempty"`
}

type BetfairRunner struct {
	Status       string `json:"status,omitempty"`
	SortPriority int    `json:"sortPriority,omitempty"`
	ID           int    `json:"id,omitempty"`
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
	Pt          int64                 `json:"pt"`
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
	Pt          int64                       `json:"pt"`
	InitialClk  string                      `json:"initialClk"`
	Clk         string                      `json:"clk"`
	Oc          []BetfairOrderAccountChange `json:"oc"`
}
