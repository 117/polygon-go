# polygon

[![godoc](https://godoc.org/github.com/117/polygon?status.svg)](https://godoc.org/github.com/117/polygon)
[![goreportcard](https://goreportcard.com/badge/github.com/117/polygon)](https://goreportcard.com/report/github.com/117/polygon)

## Installation

`go get -u github.com/117/polygon`

## Authentication

Before calling any methods you must provide your
[Polygon.io](https://polygon.io/docs/#getting-started) API key.

```go
import polygon

func init() {
    polygon.APIKey("YOUR-API-KEY-GOES-HERE")
}
```

## Methods

| Method                                       | Returns             | Example                                  |
| :------------------------------------------- | :------------------ | :--------------------------------------- |
| `polygon.Tickers(*polygon.Parameters)`       | `[]Ticker`          | [See Example](#Tickers)                  |
| `polygon.TickerTypes()`                      | `Types`             | [See Example](#TickerTypes)              |
| `polygon.TickerDetails(*polygon.Parameters)` | `Details`           | [See Example](#TickerDetails)            |
| `polygon.TickerNews()`                       | `[]News`            | [See Example](#TickerNews)               |
| `polygon.Markets()`                          | `[]Market`          | [See Example](#Markets)                  |
| `polygon.Locales()`                          | `[]Locale`          | [See Example](#Locales)                  |
| `polygon.StockSplits()`                      | `[]Split`           | [See Example](#StockSplits)              |
| `polygon.StockDividends()`                   | `[]Dividend`        | [See Example](#StockDividends)           |
| `polygon.StockFinancials()`                  | `[]Financial`       | [See Example](#StockFinancials)          |
| `polygon.MarketStatus()`                     | `Status`            | [See Example](#MarketStatus)             |
| `polygon.MarketHolidays()`                   | `[]Holiday`         | [See Example](#MarketHolidays)           |
| `polygon.Exchanges()`                        | `[]Exchange`        | [See Example](#Exchanges)                |
| `polygon.HistoricTrades()`                   | `[]Trade`           | [See Example](#HistoricTrades)           |
| `polygon.HistoricQuotes()`                   | `[]Quote`           | [See Example](#HistoricQuotes)           |
| `polygon.LastTradeForASymbol()`              | `Trade`             | [See Example](#LastTradeForASymbol)      |
| `polygon.DailyOpenCloseAfterHours()`         | `Aggregate(...)`    | [See Example](#DailyOpenCloseAfterHours) |
| `polygon.ConditionMappings()`                | `map[string]string` | [See Example](#ConditionMappings)        |
| `polygon.SnapshotAllTickers()`               | `[]Snapshot`        | [See Example](#SnapshotAllTickers)       |
| `polygon.SnapshotSingleTicker()`             | `Snapshot`          | [See Example](#SnapshotSingleTicker)     |
| `polygon.SnapshotGainersLosers()`            | `[]Snapshot`        | [See Example](#SnapshotGainersLosers)    |
| `polygon.PreviousClose()`                    | `Aggregate`         | [See Example](#PreviousClose)            |
| `polygon.Aggregates()`                       | `[]Aggregate`       | [See Example](#Aggregates)               |
| `polygon.GroupedDaily()`                     | `[]Aggregate`       | [See Example](#GroupedDaily)             |
| more coming soon...                          |                     |                                          |

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
    Symbol: "AAPL",
    // possibly more, check polygon docs or hover in VSC
})
```

### TickerNews

Get news articles for this ticker.

```go
news, err := polygon.TickerNews(&polygon.Parameters{
    Symbol: "AAPL",
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
    Symbol: "AAPL",
    // possibly more, check polygon docs or hover in VSC
})
```

### StockDividends

Get the historical divdends for this ticker.

```go
dividends, err := polygon.StockDividends(&polygon.Parameters{
    Symbol: "AAPL",
    // possibly more, check polygon docs or hover in VSC
})
```

### StockFinancials

Get the historical financials for this ticker.

```go
financials, err := polygon.StockFinancials(&polygon.Parameters{
    Symbol: "AAPL",
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
    Symbol: "AAPL",
    Date:   "2020-01-01",
    Limit:  100,
    // possibly more, check polygon docs or hover in VSC
})
```

### HistoricQuotes

Get historic NBBO quotes for a ticker.

```go
quotes, err := polygon.HistoricQuotes(&polygon.Parameters{
    Symbol: "AAPL",
    Date:   "2020-01-01",
    Limit:  100,
    // possibly more, check polygon docs or hover in VSC
})
```

### LastTradeForASymbol

Get the last trade for a given stock.

```go
trade, err := polygon.LastTradeForASymbol(&polygon.Parameters{
    Symbol: "AAPL",
    // possibly more, check polygon docs or hover in VSC
})
```

### LastQuoteForASymbol

Get the last quote tick for a given stock.

```go
quote, err := polygon.LastQuoteForASymbol(&polygon.Parameters{
    Symbol: "AAPL",
    // possibly more, check polygon docs or hover in VSC
})
```

### DailyOpenCloseAfterHours

Get the open, close and afterhours prices of a symbol on a certain date.

```go
quotes, err := polygon.DailyOpenCloseAfterHours(&polygon.Parameters{
    Symbol: "AAPL",
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
