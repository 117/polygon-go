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
import polygon

polygon.ApiKey("YOUR-API-KEY-HERE")

tickers, err := polygon.Tickers(&polygon.Parameters{
    Market
})

if err != nil {
    panic(err)
}

for _, ticker := range tickers {
    // access each ticker here
}
```

### TickerTypes

Get the mapping of ticker types to descriptions / long names.

```go
import polygon

polygon.ApiKey("YOUR-API-KEY-HERE")

types, err := polygon.TickerTypes()

if err != nil {
    panic(err)
}

// access types here
```

### TickerDetails

Get the details of the symbol company/entity. These are important details which
offer an overview of the entity. Things like name, sector, description, logo and
similar companies.

```go
import polygon

polygon.ApiKey("YOUR-API-KEY-HERE")

details, err := polygon.TickerDetails(&polygon.Parameters{
    Market
})

if err != nil {
    panic(err)
}

// access details here
```
