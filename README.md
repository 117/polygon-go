# polygon

## Methods

| Method                                | Example                    |
| ------------------------------------- | -------------------------- |
| `polygon.Tickers(*polygon.Parameters)` | [See Example](#Tickers) |

### Tickers

Query all ticker symbols which are supported by Polygon.io. This API includes
Indices, Crypto, FX, and Stocks/Equities.

```go
import polygon

polygon.Tickers(&polygon.Parameters{
    Market
})

```
