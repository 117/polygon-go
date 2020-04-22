# polygon

[![godoc](https://godoc.org/github.com/117/polygon?status.svg)](https://godoc.org/github.com/117/polygon)
[![goreportcard](https://goreportcard.com/badge/github.com/117/polygon)](https://goreportcard.com/report/github.com/117/polygon)

## Installation

`go get -u github.com/117/polygon@latest`

## Client

There are two ways to interact with the client. The first is using an
environment variable with the `DefaultClient`.

```console
$ export POLYGON_KEY yourKeyGoesHere
```

The second way is by creating your own client, this has the benefit of using
multiple keys.

```go
func main() {
    client, err := polygon.NewClient(&polygon.Credentials{Key: "yourKeyHere"})

    // an error occurs if they key is unauthorized
    if err != nil {
        panic(err)
    }
}
```

## Streaming

Your API key allows 1 simultaneous connection to each cluster.

| Cluster | URL                              | Enum             |
| :------ | :------------------------------- | :--------------- |
| Stocks  | `wss://socket.polygon.io/stocks` | `polygon.Stocks` |
| Forex   | `wss://socket.polygon.io/forex`  | `polygon.Forex`  |
| Crypto  | `wss://socket.polygon.io/crypto` | `polygon.Crypto` |

Connecting to these servers is easy.

```go
// this is a blocking method, wrap a for{} loop to reconnect on error
err := polygon.Stream(polygon.Stocks, []string{"Q.SPY"}, func(event polygon.StreamEvent) {
    switch event.String("ev") {
    case "Q":
        fmt.Println(fmt.Sprintf("received quote for %q symbol", event.String("sym")))
    }
})
```

## Methods

These are all the Polygon.io REST API methods supported by the wrapper.

| Method                                       | Returns                           | Example                               |
| :------------------------------------------- | :-------------------------------- | :------------------------------------ |
| `polygon.Tickers(*polygon.Parameters)`       | `ResponseTickers`                 | [See Example](#Tickers)               |
| `polygon.TickerTypes()`                      | `ResponseTickerTypes`             | [See Example](#TickerTypes)           |
| `polygon.TickerDetails(*polygon.Parameters)` | `ResponseTickerDetails`           | [See Example](#TickerDetails)         |
| `polygon.TickerNews()`                       | `ResponseTickerNews`              | [See Example](#TickerNews)            |
| `polygon.Markets()`                          | `ResponseMarkets`                 | [See Example](#Markets)               |
| `polygon.Locales()`                          | `ResponseLocales`                 | [See Example](#Locales)               |
| `polygon.StockSplits()`                      | `ResponseStockSplits`             | [See Example](#StockSplits)           |
| `polygon.StockDividends()`                   | `ResponseStockDividends`          | [See Example](#StockDividends)        |
| `polygon.StockFinancials()`                  | `ResponseStockFinancials`         | [See Example](#StockFinancials)       |
| `polygon.MarketStatus()`                     | `ResponseMarketStatus`            | [See Example](#MarketStatus)          |
| `polygon.MarketHolidays()`                   | `ResponseMarketHolidays`          | [See Example](#MarketHolidays)        |
| `polygon.Exchanges()`                        | `ResponseExchanges`               | [See Example](#Exchanges)             |
| `polygon.HistoricTrades()`                   | `ResponseHistoricTrades`          | [See Example](#HistoricTrades)        |
| `polygon.HistoricQuotes()`                   | `ResponseHistoricQuotes`          | [See Example](#HistoricQuotes)        |
| `polygon.LastTradeForATicker()`              | `ResponseLastTradeForATicker`     | [See Example](#LastTradeForATicker)   |
| `polygon.DailyOpenClose()`                   | `ResponseDailyOpenClose`          | [See Example](#DailyOpenClose)        |
| `polygon.ConditionMappings()`                | `map[string]string`               | [See Example](#ConditionMappings)     |
| `polygon.SnapshotAllTickers()`               | `ResponseSnapshotMultipleTickers` | [See Example](#SnapshotAllTickers)    |
| `polygon.SnapshotSingleTicker()`             | `ResponseSnapshotSingleTicker`    | [See Example](#SnapshotSingleTicker)  |
| `polygon.SnapshotGainersLosers()`            | `ResponseSnapshotMultipleTickers` | [See Example](#SnapshotGainersLosers) |
| `polygon.PreviousClose()`                    | `ResponsePreviousClose`           | [See Example](#PreviousClose)         |
| `polygon.Aggregates()`                       | `ResponseAggregates`              | [See Example](#Aggregates)            |
| `polygon.GroupedDaily()`                     | `ResponseAggregates`              | [See Example](#GroupedDaily)          |
| more coming soon...                          |                                   |                                       |

### Tickers

Query all ticker symbols which are supported by Polygon.io. This API includes
Indices, Crypto, FX, and Stocks/Equities.

```go
tickers, err := polygon.Tickers(&polygon.Parameters{
    Market: "stocks",
    // possibly more, check polygon docs or hover in VSC
})
```

### TickerTypes

Get the mapping of ticker types to descriptions / long names.

```go
types, err := polygon.TickerTypes()
```

### TickerDetails

Get the details of the symbol company/entity. These are important details which
offer an overview of the entity. Things like name, sector, description, logo and
similar companies.

```go
details, err := polygon.TickerDetails(&polygon.Parameters{
    Ticker: "AAPL",
    // possibly more, check polygon docs or hover in VSC
})
```

### TickerNews

Get news articles for this ticker.

```go
news, err := polygon.TickerNews(&polygon.Parameters{
    Ticker: "AAPL",
    // possibly more, check polygon docs or hover in VSC
})
```

### Markets

Get the list of currently supported markets.

```go
markets, err := polygon.Markets()
```

### Locales

Get the list of currently supported locales.

```go
locales, err := polygon.Locales()
```

### StockSplits

Get the historical splits for this symbol.

```go
splits, err := polygon.StockSplits(&polygon.Parameters{
    Ticker: "AAPL",
    // possibly more, check polygon docs or hover in VSC
})
```

### StockDividends

Get the historical divdends for this ticker.

```go
dividends, err := polygon.StockDividends(&polygon.Parameters{
    Ticker: "AAPL",
    // possibly more, check polygon docs or hover in VSC
})
```

### StockFinancials

Get the historical financials for this ticker.

```go
financials, err := polygon.StockFinancials(&polygon.Parameters{
    Ticker: "AAPL",
    // possibly more, check polygon docs or hover in VSC
})
```

### MarketStatus

Current status of each market.

```go
status, err := polygon.MarketStatus()
```

### MarketHolidays

Get upcoming market holidays and their open/close times.

```go
holidays, err := polygon.MarketHolidays()
```

### Exchanges

List of stock exchanges which are supported by Polygon.io.

```go
exchanges, err := polygon.Exchanges()
```

### HistoricTrades

Get historic trades for a ticker.

```go
trades, err := polygon.HistoricTrades(&polygon.Parameters{
    Ticker: "AAPL",
    Date:   "2020-01-01",
    Limit:  100,
    // possibly more, check polygon docs or hover in VSC
})
```

### HistoricQuotes

Get historic NBBO quotes for a ticker.

```go
quotes, err := polygon.HistoricQuotes(&polygon.Parameters{
    Ticker: "AAPL",
    Date:   "2020-01-01",
    Limit:  100,
    // possibly more, check polygon docs or hover in VSC
})
```

### LastTradeForATicker

Get the last trade for a given stock.

```go
trade, err := polygon.LastTradeForATicker(&polygon.Parameters{
    Ticker: "AAPL",
    // possibly more, check polygon docs or hover in VSC
})
```

### LastQuoteForATicker

Get the last quote tick for a given stock.

```go
quote, err := polygon.LastQuoteForATicker(&polygon.Parameters{
    Ticker: "AAPL",
    // possibly more, check polygon docs or hover in VSC
})
```

### DailyOpenClose

Get the open, close and afterhours prices of a symbol on a certain date.

```go
quotes, err := polygon.DailyOpenClose(&polygon.Parameters{
    Ticker: "AAPL",
    Date:   "2020-01-01",
    // possibly more, check polygon docs or hover in VSC
})
```

### ConditionMappings

The mappings for conditions on trades and quotes.

```go
mappings, err := polygon.ConditionMappings(&polygon.Parameters{
    TickType: "trades",
    // possibly more, check polygon docs or hover in VSC
})
```

### SnapshotAllTickers

Snapshot allows you to see all tickers current minute aggregate, daily aggregate
and last trade. As well as previous days aggregate and calculated change for
today.

```go
snapshots, err := polygon.SnapshotAllTickers()
```

### SnapshotSingleTicker

See the current snapshot of a single ticker.

```go
snapshot, err := polygon.SnapshotSingleTicker(&polygon.Parameters{
    Ticker: "AAPL",
    // possibly more, check polygon docs or hover in VSC
})
```

### SnapshotGainersLosers

See the current snapshot of the top 20 gainers or losers of the day at the
moment.

```go
snapshots, err := polygon.SnapshotGainersLosers(&polygon.Parameters{
    Direction: "gainers",
    // possibly more, check polygon docs or hover in VSC
})
```

### PreviousClose

Get the previous day close for the specified ticker.

```go
aggregate, err := polygon.PreviousClose(&polygon.Parameters{
    Ticker: "AAPL",
    // possibly more, check polygon docs or hover in VSC
})
```

### Aggregates

Get the previous day close for the specified ticker.

```go
aggregate, err := polygon.Aggregates(&polygon.Parameters{
    Ticker:     "AAPL",
    Timespan:   "day",
    From:       "2019-01-01",
    To:         "2019-02-01",
    Multiplier: 1,
    // possibly more, check polygon docs or hover in VSC
})
```

### GroupedDaily

Get the daily OHLC for entire markets.

```go
aggregates, err := polygon.GroupedDaily(&polygon.Parameters{
    Locale: "US",
    Market: "stocks",
    Date:   "2019-02-01",
    // possibly more, check polygon docs or hover in VSC
})
```
