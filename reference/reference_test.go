package reference

import (
	"fmt"
	"polygon/client"
	"testing"
)

const apiKey = "YOUR_API_KEY_GOES_HERE"

func TestStockFinancials(t *testing.T) {
	client.UseAPIKey(apiKey)

	financials, err := StockFinancials(&Parameters{
		Symbol: "CLVS",
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(financials[0].Debt)
}
