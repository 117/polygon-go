package stocks

import (
	"fmt"
	"polygon/client"
	"testing"
)

const apiKey = "YOUR_API_KEY_GOES_HERE"

func TestExchanges(t *testing.T) {
	client.UseAPIKey(apiKey)

	exchanges, err := Exchanges()

	if err != nil {
		t.Error(err)
	}

	for _, exchange := range exchanges {
		fmt.Println(exchange)
	}
}

func TestHistoricTrades(t *testing.T) {
	client.UseAPIKey(apiKey)

	exchanges, err := HistoricTrades(&Parameters{
		Ticker: "CLVS",
		Date:   "2020-03-10",
		Limit:  10,
	})

	if err != nil {
		t.Error(err)
	}

	for _, exchange := range exchanges {
		fmt.Println(exchange)
	}
}

func TestHistoricQuotes(t *testing.T) {
	client.UseAPIKey(apiKey)

	exchanges, err := HistoricQuotes(&Parameters{
		Ticker: "CLVS",
		Date:   "2020-03-10",
		Limit:  10,
	})

	if err != nil {
		t.Error(err)
	}

	for _, exchange := range exchanges {
		fmt.Println(exchange)
	}
}

func TestLastTradeForASymbol(t *testing.T) {
	client.UseAPIKey(apiKey)

	trade, err := LastTradeForASymbol(&Parameters{Ticker: "CLVS"})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(trade)
}

func TestLastQuoteForASymbol(t *testing.T) {
	client.UseAPIKey(apiKey)

	trade, err := LastQuoteForASymbol(&Parameters{Ticker: "CLVS"})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(trade)
}
