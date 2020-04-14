package polygon

import (
	"reflect"
	"time"
)

// Parameters - see method(s)
type Parameters struct {
	Ticker         string `json:",omitempty" url:",omitempty"`
	Date           string `json:",omitempty" url:",omitempty"`
	Timestamp      int64  `json:",omitempty" url:",omitempty"`
	TimestampLimit int64  `json:",omitempty" url:",omitempty"`
	Reverse        bool   `json:",omitempty" url:",omitempty"`
	Limit          int64  `json:",omitempty" url:",omitempty"`
	TickType       string `json:",omitempty" url:",omitempty"`
	Direction      string `json:",omitempty" url:",omitempty"`
	Unadjusted     bool   `json:",omitempty" url:",omitempty"`
	Multiplier     int64  `json:",omitempty" url:",omitempty"`
	Timespan       string `json:",omitempty" url:",omitempty"`
	From           string `json:",omitempty" url:",omitempty"`
	To             string `json:",omitempty" url:",omitempty"`
	Sort           string `json:",omitempty" url:",omitempty"`
	Type           string `json:",omitempty" url:",omitempty"`
	Market         string `json:",omitempty" url:",omitempty"`
	Locale         string `json:",omitempty" url:",omitempty"`
	Search         string `json:",omitempty" url:",omitempty"`
	PerPage        int64  `json:",omitempty" url:",omitempty"`
	Page           int64  `json:",omitempty" url:",omitempty"`
	Active         bool   `json:",omitempty" url:",omitempty"`
}

// ResponseLocales - see method(s)
type ResponseLocales struct {
	Status  string `json:"status"`
	Results []struct {
		Locale string `json:"locale"`
		Name   string `json:"name"`
	} `json:"results"`
}

// ResponseMarkets - see method(s)
type ResponseMarkets struct {
	Status  string `json:"status"`
	Results []struct {
		Market      string `json:"market"`
		Description string `json:"desc"`
	} `json:"results"`
}

// ResponseMarketStatus - see method(s)
type ResponseMarketStatus struct {
	Market     string `json:"market"`
	ServerTime string `json:"serverTime"`
	Exchanges  struct {
		Nyse   string `json:"nyse"`
		Nasdaq string `json:"nasdaq"`
		Otc    string `json:"otc"`
	} `json:"exchanges"`
	Currencies struct {
		Fx     string `json:"fx"`
		Crypto string `json:"crypto"`
	} `json:"currencies"`
}

// ResponseMarketHolidays - see method(s)
type ResponseMarketHolidays []struct {
	Exchange string    `json:"exchange"`
	Name     string    `json:"name"`
	Status   string    `json:"status"`
	Date     time.Time `json:"date"`
	Open     time.Time `json:"open"`
	Close    time.Time `json:"close"`
}

// ResponseTickers - see method(s)
type ResponseTickers []struct {
	Ticker      string `json:"ticker"`
	Name        string `json:"name"`
	Market      string `json:"market"`
	Locale      string `json:"locale"`
	Currency    string `json:"currency"`
	Active      bool   `json:"active"`
	PrimaryExch string `json:"primaryExch"`
	Type        string `json:"type"`
	Codes       struct {
		Cik     string `json:"cik"`
		Figiuid string `json:"figiuid"`
		Scfigi  string `json:"scfigi"`
		Cfigi   string `json:"cfigi"`
		Figi    string `json:"figi"`
	} `json:"codes"`
	Updated time.Time `json:"updated"`
	URL     string    `json:"url"`
}

// ResponseTypes - see method(s)
type ResponseTypes struct {
	Status  string `json:"status"`
	Results struct {
		Types      map[string]string `json:"types"`
		IndexTypes map[string]string `json:"indexTypes"`
	} `json:"results"`
}

// ResponseTickerDetails - see method(s)
type ResponseTickerDetails struct {
	Logo        string   `json:"logo"`
	Exchange    string   `json:"exchange"`
	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
	Cik         string   `json:"cik"`
	Bloomberg   string   `json:"bloomberg"`
	Lei         string   `json:"lei"`
	Sic         int64    `json:"sic"`
	Country     string   `json:"country"`
	Industry    string   `json:"industry"`
	Sector      string   `json:"sector"`
	Marketcap   int64    `json:"marketcap"`
	Employees   int64    `json:"employees"`
	Phone       string   `json:"phone"`
	Ceo         string   `json:"ceo"`
	URL         string   `json:"url"`
	Description string   `json:"description"`
	Similar     []string `json:"similar"`
	Tags        []string `json:"tags"`
}

// ResponseTickerNews - see method(s)
type ResponseTickerNews []struct {
	Symbols   []string  `json:"symbols"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Source    string    `json:"source"`
	Summary   string    `json:"summary"`
	Image     string    `json:"image"`
	Timestamp time.Time `json:"timestamp"`
	Keywords  []string  `json:"keywords"`
}

// ResponseStockSplits - see method(s)
type ResponseStockSplits struct {
	Status  string `json:"status"`
	Count   int64  `json:"count"`
	Results []struct {
		Ticker       string  `json:"ticker"`
		ExDate       string  `json:"exDate"`
		PaymentDate  string  `json:"paymentDate"`
		RecordDate   string  `json:"recordDate"`
		DeclaredDate string  `json:"declaredDate"`
		Ratio        float64 `json:"ratio"`
		Tofactor     int64   `json:"tofactor"`
		Forfactor    int64   `json:"forfactor"`
	} `json:"results"`
}

// ResponseStockDividends - see method(s)
type ResponseStockDividends struct {
	Status  string `json:"status"`
	Count   int64  `json:"count"`
	Results []struct {
		Symbol       string    `json:"symbol"`
		Type         string    `json:"type"`
		ExDate       time.Time `json:"exDate"`
		PaymentDate  time.Time `json:"paymentDate"`
		RecordDate   time.Time `json:"recordDate"`
		DeclaredDate time.Time `json:"declaredDate"`
		Amount       float64   `json:"amount"`
		Qualified    string    `json:"qualified"`
		Flag         string    `json:"flag"`
	} `json:"results"`
}

// ResponseStockFinancials - see method(s)
type ResponseStockFinancials struct {
	Status  string `json:"status"`
	Count   int64  `json:"count"`
	Results []struct {
		Ticker                                                 string  `json:"ticker"`
		Period                                                 string  `json:"period"`
		CalendarDate                                           string  `json:"calendarDate"`
		ReportPeriod                                           string  `json:"reportPeriod"`
		Updated                                                string  `json:"updated"`
		AccumulatedOtherComprehensiveIncome                    float64 `json:"accumulatedOtherComprehensiveIncome"`
		Assets                                                 float64 `json:"assets"`
		AssetsAverage                                          float64 `json:"assetsAverage"`
		AssetsCurrent                                          float64 `json:"assetsCurrent"`
		AssetTurnover                                          float64 `json:"assetTurnover"`
		AssetsNonCurrent                                       float64 `json:"assetsNonCurrent"`
		BookValuePerShare                                      float64 `json:"bookValuePerShare"`
		CapitalExpenditure                                     float64 `json:"capitalExpenditure"`
		CashAndEquivalents                                     float64 `json:"cashAndEquivalents"`
		CashAndEquivalentsUSD                                  float64 `json:"cashAndEquivalentsUSD"`
		CostOfRevenue                                          float64 `json:"costOfRevenue"`
		ConsolidatedIncome                                     float64 `json:"consolidatedIncome"`
		CurrentRatio                                           float64 `json:"currentRatio"`
		DebtToEquityRatio                                      float64 `json:"debtToEquityRatio"`
		Debt                                                   float64 `json:"debt"`
		DebtCurrent                                            float64 `json:"debtCurrent"`
		DebtNonCurrent                                         float64 `json:"debtNonCurrent"`
		DebtUSD                                                float64 `json:"debtUSD"`
		DeferredRevenue                                        float64 `json:"deferredRevenue"`
		DepreciationAmortizationAndAccretion                   float64 `json:"depreciationAmortizationAndAccretion"`
		Deposits                                               float64 `json:"deposits"`
		DividendYield                                          float64 `json:"dividendYield"`
		DividendsPerBasicCommonShare                           float64 `json:"dividendsPerBasicCommonShare"`
		EarningBeforeInterestTaxes                             float64 `json:"earningBeforeInterestTaxes"`
		EarningsBeforeInterestTaxesDepreciationAmortization    float64 `json:"earningsBeforeInterestTaxesDepreciationAmortization"`
		EBITDAMargin                                           float64 `json:"EBITDAMargin"`
		EarningsBeforeInterestTaxesDepreciationAmortizationUSD float64 `json:"earningsBeforeInterestTaxesDepreciationAmortizationUSD"`
		EarningBeforeInterestTaxesUSD                          float64 `json:"earningBeforeInterestTaxesUSD"`
		EarningsBeforeTax                                      float64 `json:"earningsBeforeTax"`
		EarningsPerBasicShare                                  float64 `json:"earningsPerBasicShare"`
		EarningsPerDilutedShare                                float64 `json:"earningsPerDilutedShare"`
		EarningsPerBasicShareUSD                               float64 `json:"earningsPerBasicShareUSD"`
		ShareholdersEquity                                     float64 `json:"shareholdersEquity"`
		AverageEquity                                          float64 `json:"averageEquity"`
		ShareholdersEquityUSD                                  float64 `json:"shareholdersEquityUSD"`
		EnterpriseValue                                        float64 `json:"enterpriseValue"`
		EnterpriseValueOverEBIT                                float64 `json:"enterpriseValueOverEBIT"`
		EnterpriseValueOverEBITDA                              float64 `json:"enterpriseValueOverEBITDA"`
		FreeCashFlow                                           float64 `json:"freeCashFlow"`
		FreeCashFlowPerShare                                   float64 `json:"freeCashFlowPerShare"`
		ForeignCurrencyUSDExchangeRate                         float64 `json:"foreignCurrencyUSDExchangeRate"`
		GrossProfit                                            float64 `json:"grossProfit"`
		GrossMargin                                            float64 `json:"grossMargin"`
		GoodwillAndIntangibleAssets                            float64 `json:"goodwillAndIntangibleAssets"`
		InterestExpense                                        float64 `json:"interestExpense"`
		InvestedCapital                                        float64 `json:"investedCapital"`
		InvestedCapitalAverage                                 float64 `json:"investedCapitalAverage"`
		Inventory                                              float64 `json:"inventory"`
		Investments                                            float64 `json:"investments"`
		InvestmentsCurrent                                     float64 `json:"investmentsCurrent"`
		InvestmentsNonCurrent                                  float64 `json:"investmentsNonCurrent"`
		TotalLiabilities                                       float64 `json:"totalLiabilities"`
		CurrentLiabilities                                     float64 `json:"currentLiabilities"`
		LiabilitiesNonCurrent                                  float64 `json:"liabilitiesNonCurrent"`
		MarketCapitalization                                   float64 `json:"marketCapitalization"`
		NetCashFlow                                            float64 `json:"netCashFlow"`
		NetCashFlowBusinessAcquisitionsDisposals               float64 `json:"netCashFlowBusinessAcquisitionsDisposals"`
		IssuanceEquityShares                                   float64 `json:"issuanceEquityShares"`
		IssuanceDebtSecurities                                 float64 `json:"issuanceDebtSecurities"`
		PaymentDividendsOtherCashDistributions                 float64 `json:"paymentDividendsOtherCashDistributions"`
		NetCashFlowFromFinancing                               float64 `json:"netCashFlowFromFinancing"`
		NetCashFlowFromInvesting                               float64 `json:"netCashFlowFromInvesting"`
		NetCashFlowInvestmentAcquisitionsDisposals             float64 `json:"netCashFlowInvestmentAcquisitionsDisposals"`
		NetCashFlowFromOperations                              float64 `json:"netCashFlowFromOperations"`
		EffectOfExchangeRateChangesOnCash                      float64 `json:"effectOfExchangeRateChangesOnCash"`
		NetIncome                                              float64 `json:"netIncome"`
		NetIncomeCommonStock                                   float64 `json:"netIncomeCommonStock"`
		NetIncomeCommonStockUSD                                float64 `json:"netIncomeCommonStockUSD"`
		NetLossIncomeFromDiscontinuedOperations                float64 `json:"netLossIncomeFromDiscontinuedOperations"`
		NetIncomeToNonControllingInterests                     float64 `json:"netIncomeToNonControllingInterests"`
		ProfitMargin                                           float64 `json:"profitMargin"`
		OperatingExpenses                                      float64 `json:"operatingExpenses"`
		OperatingIncome                                        float64 `json:"operatingIncome"`
		TradeAndNonTradePayables                               float64 `json:"tradeAndNonTradePayables"`
		PayoutRatio                                            float64 `json:"payoutRatio"`
		PriceToBookValue                                       float64 `json:"priceToBookValue"`
		PriceEarnings                                          float64 `json:"priceEarnings"`
		PriceToEarningsRatio                                   float64 `json:"priceToEarningsRatio"`
		PropertyPlantEquipmentNet                              float64 `json:"propertyPlantEquipmentNet"`
		PreferredDividendsIncomeStatementImpact                float64 `json:"preferredDividendsIncomeStatementImpact"`
		SharePriceAdjustedClose                                float64 `json:"sharePriceAdjustedClose"`
		PriceSales                                             float64 `json:"priceSales"`
		PriceToSalesRatio                                      float64 `json:"priceToSalesRatio"`
		TradeAndNonTradeReceivables                            float64 `json:"tradeAndNonTradeReceivables"`
		AccumulatedRetainedEarningsDeficit                     float64 `json:"accumulatedRetainedEarningsDeficit"`
		Revenues                                               float64 `json:"revenues"`
		RevenuesUSD                                            float64 `json:"revenuesUSD"`
		ResearchAndDevelopmentExpense                          float64 `json:"researchAndDevelopmentExpense"`
		ReturnOnAverageAssets                                  float64 `json:"returnOnAverageAssets"`
		ReturnOnAverageEquity                                  float64 `json:"returnOnAverageEquity"`
		ReturnOnInvestedCapital                                float64 `json:"returnOnInvestedCapital"`
		ReturnOnSales                                          float64 `json:"returnOnSales"`
		ShareBasedCompensation                                 float64 `json:"shareBasedCompensation"`
		SellingGeneralAndAdministrativeExpense                 float64 `json:"sellingGeneralAndAdministrativeExpense"`
		ShareFactor                                            float64 `json:"shareFactor"`
		Shares                                                 float64 `json:"shares"`
		WeightedAverageShares                                  float64 `json:"weightedAverageShares"`
		WeightedAverageSharesDiluted                           float64 `json:"weightedAverageSharesDiluted"`
		SalesPerShare                                          float64 `json:"salesPerShare"`
		TangibleAssetValue                                     float64 `json:"tangibleAssetValue"`
		TaxAssets                                              float64 `json:"taxAssets"`
		IncomeTaxExpense                                       float64 `json:"incomeTaxExpense"`
		TaxLiabilities                                         float64 `json:"taxLiabilities"`
		TangibleAssetsBookValuePerShare                        float64 `json:"tangibleAssetsBookValuePerShare"`
		WorkingCapital                                         float64 `json:"workingCapital"`
	} `json:"results"`
}

// ResponseExchanges - see method(s)
type ResponseExchanges []struct {
	ID     int64  `json:"id"`
	Type   string `json:"type"`
	Market string `json:"market"`
	Mic    string `json:"mic"`
	Name   string `json:"name"`
	Tape   string `json:"tape"`
}

// ResponseHistoricTrades - see method(s)
type ResponseHistoricTrades struct {
	ResultsCount int64  `json:"results_count"`
	DbLatency    int64  `json:"db_latency"`
	Success      bool   `json:"success"`
	Ticker       string `json:"ticker"`
	Results      []struct {
		Ticker            string  `json:"T"`
		Timestamp         int64   `json:"t"`
		TimestampExchange int64   `json:"y"`
		TimestampTRF      int64   `json:"f"`
		SequenceNumber    int64   `json:"q"`
		TradeID           string  `json:"i"`
		ExchangeID        int64   `json:"x"`
		Size              int64   `json:"s"`
		Conditions        []int   `json:"c"`
		Price             float64 `json:"p"`
		Tape              int64   `json:"z"`
	} `json:"results"`
}

// ResponseHistoricQuotes - see method(s)
type ResponseHistoricQuotes struct {
	ResultsCount int64  `json:"results_count"`
	DbLatency    int64  `json:"db_latency"`
	Success      bool   `json:"success"`
	Ticker       string `json:"ticker"`
	Results      []struct {
		Ticker            string  `json:"T"`
		Timestamp         int64   `json:"t"`
		TimestampExchange int64   `json:"y"`
		TimestampTRF      int64   `json:"f"`
		SequenceNumber    int64   `json:"q"`
		Conditions        []int   `json:"c"`
		Indicators        []int   `json:"i"`
		Bid               float64 `json:"p"`
		BidExchangeID     int64   `json:"x"`
		BidSize           int64   `json:"s"`
		Ask               float64 `json:"P"`
		AskExchangeID     int64   `json:"X"`
		AskSize           int64   `json:"S"`
		Tape              int64   `json:"z"`
	} `json:"results"`
}

// ResponseLastTradeForASymbol - see method(s)
type ResponseLastTradeForASymbol struct {
	Status string `json:"status"`
	Symbol string `json:"symbol"`
	Last   struct {
		Price      float64 `json:"price"`
		Size       int64   `json:"size"`
		Exchange   int64   `json:"exchange"`
		Condition1 int64   `json:"cond1"`
		Condition2 int64   `json:"cond2"`
		Condition3 int64   `json:"cond3"`
		Condition4 int64   `json:"cond4"`
		Timestamp  int64   `json:"timestamp"`
	} `json:"last"`
}

// ResponseLastQuoteForASymbol - see method(s)
type ResponseLastQuoteForASymbol struct {
	Status string `json:"status"`
	Symbol string `json:"symbol"`
	Last   struct {
		AskPrice    float64 `json:"askprice"`
		AskSize     int64   `json:"asksize"`
		AskExchange int64   `json:"askexchange"`
		BidPrice    float64 `json:"bidprice"`
		BidSize     int64   `json:"bidsize"`
		BidExchange int64   `json:"bidexchange"`
		Timestamp   int64   `json:"timestamp"`
	} `json:"last"`
}

// ResponseDailyOpenClose - see method(s)
type ResponseDailyOpenClose struct {
	Status     string  `json:"status"`
	From       string  `json:"from"`
	Symbol     string  `json:"AAPL"`
	Open       float64 `json:"open"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Close      float64 `json:"close"`
	AfterHours float64 `json:"afterHours"`
	Volume     int64   `json:"volume"`
}

// ResponseSnapshotMultipleTickers - see method(s)
type ResponseSnapshotMultipleTickers struct {
	Status  string `json:"status"`
	Tickers []struct {
		Ticker string `json:"ticker"`
		Day    struct {
			Open   float64 `json:"o"`
			High   float64 `json:"h"`
			Low    float64 `json:"l"`
			Close  float64 `json:"c"`
			Volume float64 `json:"v"`
		} `json:"day"`
		LastTrade struct {
			Price      float64 `json:"p"`
			Size       int64   `json:"s"`
			Exchange   int64   `json:"e"`
			Condition1 int64   `json:"c1"`
			Condition2 int64   `json:"c2"`
			Condition3 int64   `json:"c3"`
			Condition4 int64   `json:"c4"`
			Timestamp  int64   `json:"t"`
		} `json:"lastTrade"`
		LastQuote struct {
			BidPrice  float64 `json:"p"`
			BidSize   int64   `json:"s"`
			AskPrice  float64 `json:"P"`
			AskSize   int64   `json:"S"`
			Timestamp int64   `json:"t"`
		} `json:"lastQuote"`
		Minute struct {
			Open   float64 `json:"o"`
			High   float64 `json:"h"`
			Low    float64 `json:"l"`
			Close  float64 `json:"c"`
			Volume float64 `json:"v"`
		} `json:"minute"`
		PreviousDay struct {
			Open   float64 `json:"o"`
			High   float64 `json:"h"`
			Low    float64 `json:"l"`
			Close  float64 `json:"c"`
			Volume float64 `json:"v"`
		} `json:"prevDay"`
		TodaysChange     float64 `json:"todaysChange"`
		TodaysChangePerc float64 `json:"todaysChangePerc"`
		Updated          int64   `json:"updated"`
	} `json:"tickers"`
}

// ResponseSnapshotSingleTicker - see method(s)
type ResponseSnapshotSingleTicker struct {
	Status string `json:"status"`
	Ticker struct {
		Ticker string `json:"ticker"`
		Day    struct {
			Open   float64 `json:"o"`
			High   float64 `json:"h"`
			Low    float64 `json:"l"`
			Close  float64 `json:"c"`
			Volume float64 `json:"v"`
		} `json:"day"`
		LastTrade struct {
			Price      float64 `json:"p"`
			Size       int64   `json:"s"`
			Exchange   int64   `json:"e"`
			Condition1 int64   `json:"c1"`
			Condition2 int64   `json:"c2"`
			Condition3 int64   `json:"c3"`
			Condition4 int64   `json:"c4"`
			Timestamp  int64   `json:"t"`
		} `json:"lastTrade"`
		LastQuote struct {
			BidPrice  float64 `json:"p"`
			BidSize   int64   `json:"s"`
			AskPrice  float64 `json:"P"`
			AskSize   int64   `json:"S"`
			Timestamp int64   `json:"t"`
		} `json:"lastQuote"`
		Minute struct {
			Open   float64 `json:"o"`
			High   float64 `json:"h"`
			Low    float64 `json:"l"`
			Close  float64 `json:"c"`
			Volume float64 `json:"v"`
		} `json:"minute"`
		PreviousDay struct {
			Open   float64 `json:"o"`
			High   float64 `json:"h"`
			Low    float64 `json:"l"`
			Close  float64 `json:"c"`
			Volume float64 `json:"v"`
		} `json:"prevDay"`
		TodaysChange     float64 `json:"todaysChange"`
		TodaysChangePerc float64 `json:"todaysChangePerc"`
		Updated          int64   `json:"updated"`
	} `json:"ticker"`
}

// ResponsePreviousClose - see method(s)
type ResponsePreviousClose struct {
	Ticker       string `json:"ticker"`
	Status       string `json:"status"`
	Adjusted     bool   `json:"adjusted"`
	QueryCount   int64  `json:"queryCount"`
	ResultsCount int64  `json:"resultsCount"`
	Results      []struct {
		Ticker    string  `json:"T"`
		Open      float64 `json:"o"`
		High      float64 `json:"h"`
		Low       float64 `json:"l"`
		Close     float64 `json:"c"`
		Volume    int64   `json:"v"`
		Timestamp int64   `json:"t"`
		Items     int64   `json:"n"`
	}
}

// ResponseAggregates - see method(s)
type ResponseAggregates struct {
	Ticker       string `json:"ticker"`
	Status       string `json:"status"`
	Adjusted     bool   `json:"adjusted"`
	QueryCount   int64  `json:"queryCount"`
	ResultsCount int64  `json:"resultsCount"`
	Results      []struct {
		Ticker     string  `json:"T"`
		Volume     int64   `json:"v"`
		Open       float64 `json:"o"`
		Close      float64 `json:"c"`
		High       float64 `json:"h"`
		Low        float64 `json:"l"`
		Timestamp  int64   `json:"t"`
		Items      int64   `json:"n"`
		Condition1 int64   `json:"c1"`
		Condition2 int64   `json:"c2"`
		Condition3 int64   `json:"c3"`
		Condition4 int64   `json:"c4"`
	}
}

// WebSocketEvent null
type WebSocketEvent map[string]interface{}

// Interface null
func (w WebSocketEvent) Interface(key string) interface{} {
	return w[key]
}

// Float32 null
func (w WebSocketEvent) Float32(key string) float32 {
	if !w.ContainsKey(key) {
		return 0.0
	}

	if reflect.TypeOf(w[key]).Kind() != reflect.Float32 {
		return 0.0
	}

	return w[key].(float32)
}

// Float64 null
func (w WebSocketEvent) Float64(key string) float64 {
	if !w.ContainsKey(key) {
		return 0.0
	}

	if reflect.TypeOf(w[key]).Kind() != reflect.Float64 {
		return 0.0
	}

	return w[key].(float64)
}

// String null
func (w WebSocketEvent) String(key string) string {
	if !w.ContainsKey(key) {
		return ""
	}

	if reflect.TypeOf(w[key]).Kind() != reflect.String {
		return ""
	}

	return w[key].(string)
}

// Int null
func (w WebSocketEvent) Int(key string) int {
	if !w.ContainsKey(key) {
		return 0
	}

	if reflect.TypeOf(w[key]).Kind() != reflect.Int {
		return 0
	}

	return w[key].(int)
}

// Int64 null
func (w WebSocketEvent) Int64(key string) int64 {
	if !w.ContainsKey(key) {
		return 0
	}

	if reflect.TypeOf(w[key]).Kind() != reflect.Int64 {
		return 0
	}

	return w[key].(int64)
}

// ContainsKey null
func (w WebSocketEvent) ContainsKey(key string) bool {
	if _, has := w[key]; has {
		return true
	}

	return false
}
