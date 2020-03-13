# polygon

## Methods

| Method                                       | Returns       | Example                         |
| :------------------------------------------- | :------------ | :------------------------------ |
| `polygon.Tickers(*polygon.Parameters)`       | `[]Ticker`    | [See Example](#Tickers)         |
| `polygon.TickerTypes()`                      | `Types`       | [See Example](#TickerTypes)     |
| `polygon.TickerDetails(*polygon.Parameters)` | `Details`     | [See Example](#TickerDetails)   |
| `polygon.TickerNews()`                       | `[]News`      | [See Example](#TickerNews)      |
| `polygon.Markets()`                          | `[]Market`    | [See Example](#Markets)         |
| `polygon.Locales()`                          | `[]Locale`    | [See Example](#Locales)         |
| `polygon.StockSplits()`                      | `[]Split`     | [See Example](#StockSplits)     |
| `polygon.StockDividends()`                   | `[]Dividend`  | [See Example](#StockDividends)  |
| `polygon.StockFinancials()`                  | `[]Financial` | [See Example](#StockFinancials) |
| `polygon.MarketStatus()`                     | `Status`      | [See Example](#MarketStatus)    |
| `polygon.MarketHolidays()`                   | `[]Holiday`   | [See Example](#MarketHolidays)  |

### Tickers

Query all ticker symbols which are supported by Polygon.io. This API includes
Indices, Crypto, FX, and Stocks/Equities.

```go
tickers, err := polygon.Tickers(&polygon.Parameters{
    Market: "stocks",
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
})
```
