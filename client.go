package polygon

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

// Client null
type Client struct {
	Credentials *Credentials
}

// Locales - Get the list of currently supported locales
func (c *Client) Locales() (response ResponseLocales, _ error) {
	return response, request(endpoint(c.Credentials, "/v2/reference/locales", nil), &response)
}

// Markets - Get the list of currently supported markets
func (c *Client) Markets() (response ResponseMarkets, _ error) {
	return response, request(endpoint(c.Credentials, "/v2/reference/markets", nil), &response)
}

// MarketStatus - Current status of each market.
func (c *Client) MarketStatus() (response ResponseMarketStatus, _ error) {
	return response, request(endpoint(c.Credentials, "/v1/marketstatus/now", nil), &response)
}

// MarketHolidays - Get upcoming market holidays and their open/close times
func (c *Client) MarketHolidays() (response ResponseMarketHolidays, _ error) {
	return response, request(endpoint(c.Credentials, "/v1/marketstatus/upcoming", nil), &response)
}

// Tickers - Query all ticker symbols which are supported by Polygon.io.
// This API includes Indices, Crypto, FX, and Stocks/Equities.
func (c *Client) Tickers(parameters *Parameters) (response ResponseTickers, _ error) {
	return response, request(endpoint(c.Credentials, "/v2/reference/tickers", &parameters), &response)
}

// TickerTypes - Get the mapping of ticker types to descriptions / long names.
func (c *Client) TickerTypes() (response ResponseTypes, _ error) {
	return response, request(endpoint(c.Credentials, "/v2/reference/types", nil), &response)
}

// TickerDetails - Get the details of the symbol company/entity.
// These are important details which offer an overview of the entity.
// Things like name, sector, description, logo and similar companies.
func (c *Client) TickerDetails(parameters *Parameters) (response ResponseTickerDetails, _ error) {
	return response, request(endpoint(c.Credentials, fmt.Sprintf("/v1/meta/symbols/%s/company", parameters.Ticker), nil), &response)
}

// TickerNews - Get news articles for this ticker.
func (c *Client) TickerNews(parameters *Parameters) (response ResponseTickerNews, _ error) {
	return response, request(endpoint(c.Credentials, fmt.Sprintf("/v1/meta/symbols/%s/news", parameters.Ticker), &parameters), &response)
}

// StockSplits - Get the historical splits for this symbol.
func (c *Client) StockSplits(parameters *Parameters) (response ResponseStockSplits, _ error) {
	return response, request(endpoint(c.Credentials, fmt.Sprintf("/v2/reference/splits/%s", parameters.Ticker), &parameters), &response)
}

// StockDividends - Get the historical divdends for this ticker.
func (c *Client) StockDividends(parameters *Parameters) (response ResponseStockDividends, _ error) {
	return response, request(endpoint(c.Credentials, fmt.Sprintf("/v2/reference/dividends/%s", parameters.Ticker), &parameters), &response)
}

// StockFinancials - Get the historical financials for this ticker.
func (c *Client) StockFinancials(parameters *Parameters) (response ResponseStockFinancials, _ error) {
	return response, request(endpoint(c.Credentials, fmt.Sprintf("/v2/reference/financials/%s", parameters.Ticker), &parameters), &response)
}

// Exchanges - List of stock exchanges which are supported by Polygon.io.
func (c *Client) Exchanges() (response ResponseExchanges, _ error) {
	return response, request(endpoint(c.Credentials, "/v1/meta/exchanges", nil), &response)
}

// HistoricTrades - Get historic trades for a ticker.
func (c *Client) HistoricTrades(parameters *Parameters) (response ResponseHistoricTrades, _ error) {
	return response, request(endpoint(c.Credentials, fmt.Sprintf("/v2/ticks/stocks/trades/%s/%s", parameters.Ticker, parameters.Date), &parameters), &response)
}

// HistoricQuotes - Get historic NBBO quotes for a ticker.
func (c *Client) HistoricQuotes(parameters *Parameters) (response ResponseHistoricQuotes, _ error) {
	return response, request(endpoint(c.Credentials, fmt.Sprintf("/v2/ticks/stocks/nbbo/%s/%s", parameters.Ticker, parameters.Date), &parameters), &response)
}

// LastTradeForASymbol - Get the last trade for a given stock.
func (c *Client) LastTradeForASymbol(parameters *Parameters) (response ResponseLastTradeForASymbol, _ error) {
	return response, request(endpoint(c.Credentials, fmt.Sprintf("/v1/last/stocks/%s", parameters.Ticker), &parameters), &response)
}

// LastQuoteForASymbol - Get the last quote tick for a given stock.
func (c *Client) LastQuoteForASymbol(parameters *Parameters) (response ResponseLastQuoteForASymbol, _ error) {
	return response, request(endpoint(c.Credentials, fmt.Sprintf("/v1/last_quote/stocks/%s", parameters.Ticker), &parameters), &response)
}

// DailyOpenClose - Get the open, close and afterhours prices of a symbol on a certain date.
func (c *Client) DailyOpenClose(parameters *Parameters) (response ResponseDailyOpenClose, _ error) {
	return response, request(endpoint(c.Credentials, fmt.Sprintf("/v1/meta/conditions/%s", parameters.TickType), &parameters), &response)
}

// ConditionMappings - The mappings for conditions on trades and quotes.
func (c *Client) ConditionMappings(parameters *Parameters) (mappings map[string]string, _ error) {
	return mappings, request(endpoint(c.Credentials, fmt.Sprintf("/v1/meta/conditions/%s", parameters.TickType), &parameters), &mappings)
}

// SnapshotAllTickers - Snapshot allows you to see all tickers current minute aggregate, daily aggregate and last trade.
// As well as previous days aggregate and calculated change for today.
func (c *Client) SnapshotAllTickers() (response ResponseSnapshotMultipleTickers, _ error) {
	return response, request(endpoint(c.Credentials, "/v2/snapshot/locale/us/markets/stocks/tickers", nil), &response)
}

// SnapshotSingleTicker - See the current snapshot of a single ticker.
func (c *Client) SnapshotSingleTicker(parameters *Parameters) (response ResponseSnapshotSingleTicker, _ error) {
	return response, request(endpoint(c.Credentials, fmt.Sprintf("/v2/snapshot/locale/us/markets/stocks/tickers/%s", parameters.Ticker), nil), &response)
}

// SnapshotGainersLosers - See the current snapshot of the top 20 gainers or losers of the day at the moment.
func (c *Client) SnapshotGainersLosers(parameters *Parameters) (response ResponseSnapshotMultipleTickers, _ error) {
	return response, request(endpoint(c.Credentials, fmt.Sprintf("/v2/snapshot/locale/us/markets/stocks/%s", parameters.Direction), nil), &response)
}

// PreviousClose - Get the previous day close for the specified ticker.
func (c *Client) PreviousClose(parameters *Parameters) (response ResponsePreviousClose, _ error) {
	return response, request(endpoint(c.Credentials, fmt.Sprintf("/v2/aggs/ticker/%s/prev", parameters.Ticker), &parameters), &response)
}

// Aggregates - Get aggregates for a date range, in custom time window sizes.
func (c *Client) Aggregates(parameters *Parameters) (response ResponseAggregates, _ error) {
	return response, request(endpoint(c.Credentials, fmt.Sprintf("/v2/aggs/ticker/%s/range/%d/%s/%s/%s", parameters.Ticker, parameters.Multiplier, parameters.Timespan, parameters.From, parameters.To), nil), &response)
}

// GroupedDaily - Get the daily OHLC for entire markets.
func (c *Client) GroupedDaily(parameters *Parameters) (response ResponseAggregates, _ error) {
	return response, request(endpoint(c.Credentials, fmt.Sprintf("/v2/aggs/grouped/locale/%s/market/%s/%s", parameters.Locale, parameters.Market, parameters.Date), &parameters), &response)
}

// Stream null
func (c *Client) Stream(cluster StreamCluster, channels []string, handler func(StreamEvent)) error {
	connection, _, err := websocket.DefaultDialer.Dial(string(cluster), nil)

	if err != nil {
		return err
	}

	defer connection.Close()

	for _, action := range []string{
		fmt.Sprintf(`{"action":"auth","params":"%s"}`, c.Credentials.Key),
		fmt.Sprintf(`{"action":"subscribe","params":"%s"}`, strings.Join(channels, ",")),
	} {
		if err := connection.WriteMessage(websocket.TextMessage, []byte(action)); err != nil {
			return err
		}
	}

	for {
		_, message, err := connection.ReadMessage()

		if err != nil {
			return err
		}

		events := []StreamEvent{}

		// ignore any malformed JSON, it's not even our fault...
		json.Unmarshal(message, &events)

		for _, event := range events {
			handler(event)
		}
	}
}

// NewClient null
func NewClient(credentials *Credentials) (*Client, error) {
	client := &Client{
		Credentials: credentials,
	}

	// this will error if unauthorized
	if _, err := client.MarketStatus(); err != nil {
		return nil, err
	}

	return client, nil
}
