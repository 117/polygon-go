# polygon

## Methods

| Method                                        | Returns    | Example                       |
| :-------------------------------------------- | :--------- | :---------------------------- |
| `polygon.Tickers(\*polygon.Parameters)`       | `[]Ticker` | [See Example](#Tickers)       |
| `polygon.TickerTypes()`                       | `Types`    | [See Example](#TickerTypes)   |
| `polygon.TickerDetails(\*polygon.Parameters)` | `Details`  | [See Example](#TickerDetails) |

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
