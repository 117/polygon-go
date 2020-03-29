package polygon

import "fmt"

// Locales ...
func Locales() (locales []Locale, _ error) {
	return locales, req(end("/v2/reference/locales", nil), &locales)
}

// Markets ...
func Markets() (markets []Market, _ error) {
	return markets, req(end("/v2/reference/markets", nil), &markets)
}

// MarketStatus ...
func MarketStatus() (status Status, _ error) {
	return status, req(end("/v1/marketstatus/now", nil), &status)
}

// MarketHolidays ...
func MarketHolidays() (holidays []Holiday, _ error) {
	return holidays, req(end("/v1/marketstatus/upcoming", nil), &holidays)
}

// Tickers ...
func Tickers(parameters *Parameters) (tickers []Ticker, _ error) {
	return tickers, req(end("/v2/reference/tickers", &parameters), &tickers)
}

// TickerTypes ...
func TickerTypes() (types Types, _ error) {
	return types, req(end("/v2/reference/types", nil), &types)
}

// TickerDetails ...
func TickerDetails(parameters *Parameters) (details Details, _ error) {
	return details, req(end(fmt.Sprintf("/v1/meta/symbols/%s/company", parameters.Symbol), nil), &details)
}

// TickerNews ...
func TickerNews(parameters *Parameters) (news []News, _ error) {
	return news, req(end(fmt.Sprintf("/v1/meta/symbols/%s/news", parameters.Symbol), &parameters), &news)
}

// StockSplits ...
func StockSplits(parameters *Parameters) (splits []Split, _ error) {
	return splits, req(end(fmt.Sprintf("/v2/reference/splits/%s", parameters.Symbol), &parameters), &splits)
}

// StockDividends ...
func StockDividends(parameters *Parameters) (dividends []Dividend, _ error) {
	return dividends, req(end(fmt.Sprintf("/v2/reference/dividends/%s", parameters.Symbol), &parameters), &dividends)
}

// StockFinancials ...
func StockFinancials(parameters *Parameters) ([]Financial, error) {
	var response struct {
		Results []Financial `json:"results"`
	}

	return response.Results, req(end(fmt.Sprintf("/v2/reference/financials/%s", parameters.Symbol), &parameters), &response)
}

// Exchanges ...
func Exchanges() (exchanges []Exchange, _ error) {
	return exchanges, req(end("/v1/meta/exchanges", nil), &exchanges)
}

// HistoricTrades ...
func HistoricTrades(parameters *Parameters) (trades []Trade, _ error) {
	var response struct {
		Results []Trade `json:"results"`
	}

	return response.Results, req(end(fmt.Sprintf("/v2/ticks/stocks/trades/%s/%s", parameters.Ticker, parameters.Date), &parameters), &response)
}

// HistoricQuotes ...
func HistoricQuotes(parameters *Parameters) ([]Quote, error) {
	var response struct {
		Results []Quote `json:"results"`
	}

	return response.Results, req(end(fmt.Sprintf("/v2/ticks/stocks/nbbo/%s/%s", parameters.Ticker, parameters.Date), &parameters), &response)
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
	}, req(end(fmt.Sprintf("/v1/last/stocks/%s", parameters.Ticker), &parameters), &response)
}

// LastQuoteForASymbol ...
func LastQuoteForASymbol(parameters *Parameters) (Quote, error) {
	var response struct {
		Last Quote `json:"last"`
	}

	return response.Last, req(end(fmt.Sprintf("/v1/last_quote/stocks/%s", parameters.Ticker), &parameters), &response)
}

// DailyOpenCloseAfterHours ...
func DailyOpenCloseAfterHours(parameters *Parameters) (daily Daily, _ error) {
	return daily, req(end(fmt.Sprintf("/v1/meta/conditions/%s", parameters.TickType), &parameters), &daily)
}

// ConditionMappings ...
func ConditionMappings(parameters *Parameters) (mappings map[string]string, _ error) {
	return mappings, req(end(fmt.Sprintf("/v1/meta/conditions/%s", parameters.TickType), &parameters), &mappings)
}

// SnapshotAllTickers ...
func SnapshotAllTickers() ([]Snapshot, error) {
	var response struct {
		Tickers []Snapshot `json:"tickers"`
	}

	return response.Tickers, req(end("/v2/snapshot/locale/us/markets/stocks/tickers", nil), &response)
}

// SnapshotSingleTicker ...
func SnapshotSingleTicker(parameters *Parameters) (Snapshot, error) {
	var response struct {
		Ticker Snapshot `json:"ticker"`
	}

	return response.Ticker, req(end(fmt.Sprintf("/v2/snapshot/locale/us/markets/stocks/tickers/%s", parameters.Ticker), nil), &response)
}

// SnapshotGainersLosers ...
func SnapshotGainersLosers(parameters *Parameters) ([]Snapshot, error) {
	var response struct {
		Tickers []Snapshot `json:"tickers"`
	}

	return response.Tickers, req(end(fmt.Sprintf("/v2/snapshot/locale/us/markets/stocks/%s", parameters.Direction), nil), &response)
}

// PreviousClose ...
func PreviousClose(parameters *Parameters) (Aggregate, error) {
	var response struct {
		Results Aggregate `json:"results"`
	}

	return response.Results, req(end(fmt.Sprintf("/v2/aggs/ticker/%s/prev", parameters.Ticker), &parameters), &response)
}

// Aggregates ...
func Aggregates(parameters *Parameters) ([]Aggregate, error) {
	var response struct {
		Results []Aggregate `json:"results"`
	}

	return response.Results, req(end(fmt.Sprintf("/v2/aggs/ticker/%s/range/%d/%s/%s/%s", parameters.Ticker, parameters.Multiplier, parameters.Timespan, parameters.From, parameters.To), nil), &response)
}

// GroupedDaily ...
func GroupedDaily(parameters *Parameters) ([]Aggregate, error) {
	var response struct {
		Results []Aggregate `json:"results"`
	}

	return response.Results, req(end(fmt.Sprintf("/v2/aggs/grouped/locale/%s/market/%s/%s", parameters.Locale, parameters.Market, parameters.Date), &parameters), &response)
}
