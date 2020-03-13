package stocks

import (
	"fmt"
	"polygon/client"
)

// Parameters ...
type Parameters struct {
	Ticker         string `json:",omitempty" url:",omitempty"`
	Date           string `json:",omitempty" url:",omitempty"`
	Timestamp      int    `json:",omitempty" url:",omitempty"`
	TimestampLimit int    `json:",omitempty" url:",omitempty"`
	Reverse        bool   `json:",omitempty" url:",omitempty"`
	Limit          int    `json:",omitempty" url:",omitempty"`
	TickType       string `json:",omitempty" url:",omitempty"`
	Direction      string `json:",omitempty" url:",omitempty"`
	Unadjusted     bool   `json:",omitempty" url:",omitempty"`
	Multiplier     int    `json:",omitempty" url:",omitempty"`
	Timespan       string `json:",omitempty" url:",omitempty"`
	From           string `json:",omitempty" url:",omitempty"`
	To             string `json:",omitempty" url:",omitempty"`
	Locale         string `json:",omitempty" url:",omitempty"`
	Market         string `json:",omitempty" url:",omitempty"`
}

// Exchange ...
type Exchange struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	Market string `json:"market"`
	Mic    string `json:"mic"`
	Name   string `json:"name"`
	Tape   string `json:"tape"`
}

// Exchanges ...
func Exchanges() ([]Exchange, error) {
	var exchanges []Exchange

	err := client.Request(client.Endpoint("/v1/meta/exchanges", nil), &exchanges)

	return exchanges, err
}

// Trade ...
type Trade struct {
	Ticker            string  `json:"T"`
	Timestamp         int64   `json:"t"`
	TimestampExchange int64   `json:"y"`
	TimestampTRF      int64   `json:"f"`
	SequenceNumber    int     `json:"q"`
	TradeID           string  `json:"i"`
	ExchangeID        int     `json:"x"`
	Size              int     `json:"s"`
	Conditions        []int   `json:"c"`
	Price             float64 `json:"p"`
	Tape              int     `json:"z"`
}

// HistoricTrades ...
func HistoricTrades(parameters *Parameters) ([]Trade, error) {
	var response struct {
		Results []Trade `json:"results"`
	}

	err := client.Request(client.Endpoint(fmt.Sprintf("/v2/ticks/stocks/trades/%s/%s", parameters.Ticker, parameters.Date), &parameters), &response)

	return response.Results, err
}

// Quote ...
type Quote struct {
	Ticker            string  `json:"T"`
	Timestamp         int64   `json:"t"`
	TimestampExchange int64   `json:"y"`
	TimestampTRF      int64   `json:"f"`
	SequenceNumber    int     `json:"q"`
	Conditions        []int   `json:"c"`
	Indicators        []int   `json:"i"`
	Bid               float64 `json:"p"`
	BidExchangeID     int     `json:"x"`
	BidSize           int     `json:"s"`
	Ask               float64 `json:"P"`
	AskExchangeID     int     `json:"X"`
	AskSize           int     `json:"S"`
	Tape              int     `json:"z"`
}

// HistoricQuotes ...
func HistoricQuotes(parameters *Parameters) ([]Quote, error) {
	var response struct {
		Results []Quote `json:"results"`
	}

	err := client.Request(client.Endpoint(fmt.Sprintf("/v2/ticks/stocks/nbbo/%s/%s", parameters.Ticker, parameters.Date), &parameters), &response)

	return response.Results, err
}

// LastTradeForASymbol ...
func LastTradeForASymbol(parameters *Parameters) (Trade, error) {
	var response struct {
		Last struct {
			Price      float64 `json:"price"`
			Size       int     `json:"size"`
			Exchange   int     `json:"exchange"`
			Condition1 int     `json:"cond1"`
			Condition2 int     `json:"cond2"`
			Condition3 int     `json:"cond3"`
			Condition4 int     `json:"cond4"`
			Timestamp  int64   `json:"timestamp"`
		} `json:"last"`
	}

	err := client.Request(client.Endpoint(fmt.Sprintf("/v1/last/stocks/%s", parameters.Ticker), &parameters), &response)

	return Trade{
		Price:      response.Last.Price,
		Size:       response.Last.Size,
		ExchangeID: response.Last.Exchange,
		Conditions: []int{
			response.Last.Condition1,
			response.Last.Condition2,
			response.Last.Condition3,
			response.Last.Condition4,
		},
		Timestamp: response.Last.Timestamp,
	}, err
}

// LastQuoteForASymbol ...
func LastQuoteForASymbol(parameters *Parameters) (Quote, error) {
	var response struct {
		Last Quote `json:"last"`
	}

	err := client.Request(client.Endpoint(fmt.Sprintf("/v1/last_quote/stocks/%s", parameters.Ticker), &parameters), &response)

	return response.Last, err
}

// Aggregate ...
type Aggregate struct {
	Ticker     string  `json:"T"`
	Volume     int     `json:"v"`
	Open       float64 `json:"o"`
	Close      float64 `json:"c"`
	High       float64 `json:"h"`
	Low        float64 `json:"l"`
	Timestamp  int64   `json:"t"`
	Items      int     `json:"n"`
	Condition1 int     `json:"c1"`
	Condition2 int     `json:"c2"`
	Condition3 int     `json:"c3"`
	Condition4 int     `json:"c4"`
}

// DailyOpenCloseAfterHours ...
func DailyOpenCloseAfterHours(parameters *Parameters) (Aggregate, Aggregate, Aggregate, error) {
	var aggregate struct {
		Symbol     string    `json:"symbol"`
		Open       Aggregate `json:"open"`
		Close      Aggregate `json:"close"`
		AfterHours Aggregate `json:"afterHours"`
	}

	err := client.Request(client.Endpoint(fmt.Sprintf("/v1/open-close/%s/%s", parameters.Ticker, parameters.Date), &parameters), &aggregate)

	return aggregate.Open, aggregate.Close, aggregate.AfterHours, err
}

// ConditionMappings ...
func ConditionMappings(parameters *Parameters) (map[string]string, error) {
	var mappings map[string]string

	err := client.Request(client.Endpoint(fmt.Sprintf("/v1/meta/conditions/%s", parameters.TickType), &parameters), &mappings)

	return mappings, err
}

// Snapshot ...
type Snapshot struct {
	Ticker string `json:"ticker"`
	Day    struct {
		Close  float64 `json:"c"`
		High   float64 `json:"h"`
		Low    float64 `json:"l"`
		Open   float64 `json:"o"`
		Volume float64 `json:"v"`
	} `json:"day"`
	LastTrade struct {
		Condition1 int     `json:"c1"`
		Condition2 int     `json:"c2"`
		Condition3 int     `json:"c3"`
		Condition4 int     `json:"c4"`
		Exchange   int     `json:"e"`
		Price      float64 `json:"p"`
		Size       int     `json:"s"`
		Timestamp  int64   `json:"t"`
	} `json:"lastTrade"`
	LastQuote struct {
		Ask       int   `json:"p"`
		AskSize   int   `json:"s"`
		Bid       int   `json:"P"`
		BidSize   int   `json:"S"`
		Timestamp int64 `json:"t"`
	} `json:"lastQuote"`
	Min struct {
		Close  float64 `json:"c"`
		High   float64 `json:"h"`
		Low    float64 `json:"l"`
		Open   float64 `json:"o"`
		Volume float64 `json:"v"`
	} `json:"min"`
	PrevDay struct {
		Close  float64 `json:"c"`
		High   float64 `json:"h"`
		Low    float64 `json:"l"`
		Open   float64 `json:"o"`
		Volume float64 `json:"v"`
	} `json:"prevDay"`
	TodaysChange     float64 `json:"todaysChange"`
	TodaysChangePerc float64 `json:"todaysChangePerc"`
	Updated          int64   `json:"updated"`
}

// SnapshotAllTickers ...
func SnapshotAllTickers() ([]Snapshot, error) {
	var response struct {
		Tickers []Snapshot `json:"tickers"`
	}

	err := client.Request(client.Endpoint("/v2/snapshot/locale/us/markets/stocks/tickers", nil), &response)

	return response.Tickers, err
}

// SnapshotSingleTicker ...
func SnapshotSingleTicker(parameters *Parameters) (Snapshot, error) {
	var response struct {
		Ticker Snapshot `json:"ticker"`
	}

	err := client.Request(client.Endpoint(fmt.Sprintf("/v2/snapshot/locale/us/markets/stocks/tickers/%s", parameters.Ticker), nil), &response)

	return response.Ticker, err
}

// SnapshotGainersLosers ...
func SnapshotGainersLosers(parameters *Parameters) ([]Snapshot, error) {
	var response struct {
		Tickers []Snapshot `json:"tickers"`
	}

	err := client.Request(client.Endpoint(fmt.Sprintf("/v2/snapshot/locale/us/markets/stocks/%s", parameters.Direction), nil), &response)

	return response.Tickers, err
}

// PreviousClose ...
func PreviousClose(parameters *Parameters) (Aggregate, error) {
	var response struct {
		Results Aggregate `json:"results"`
	}

	err := client.Request(client.Endpoint(fmt.Sprintf("/v2/aggs/ticker/%s/prev", parameters.Ticker), &parameters), &response)

	return response.Results, err
}

// Aggregates ...
func Aggregates(parameters *Parameters) ([]Aggregate, error) {
	var response struct {
		Results []Aggregate `json:"results"`
	}

	err := client.Request(client.Endpoint(fmt.Sprintf("/v2/aggs/ticker/%s/range/%d/%s/%s/%s", parameters.Ticker, parameters.Multiplier, parameters.Timespan, parameters.From, parameters.To), nil), &response)

	return response.Results, err
}

// GroupedDaily ...
func GroupedDaily(parameters *Parameters) ([]Aggregate, error) {
	var response struct {
		Results []Aggregate `json:"results"`
	}

	err := client.Request(client.Endpoint(fmt.Sprintf("/v2/aggs/grouped/locale/%s/market/%s/%s", parameters.Locale, parameters.Market, parameters.Date), &parameters), &response)

	return response.Results, err
}
