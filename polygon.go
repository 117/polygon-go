package polygon

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	rest      = "https://api.polygon.io"
	wss       = "wss://socket.polygon.io"
	wssAlpaca = "wss://alpaca.socket.polygon.io"
)

var apiKey string

// ApiKey ...
func ApiKey(key string) {
	apiKey = key
}

func request(url string, pointer interface{}) error {
	response, err := http.Get(url)

	if err != nil {
		return err
	}

	switch response.StatusCode {
	case 401:
		return errors.New("unauthorized, check your API key")
	case 404:
		return errors.New("the specified resource was not found")
	case 409:
		return errors.New("parameter is invalid or incorrect")
	}

	return json.NewDecoder(response.Body).Decode(&pointer)
}

func endpoint(path string, options interface{}) string {
	values, _ := query.Values(options)
	base := fmt.Sprintf("%s%s?apiKey=%s", rest, path, apiKey)

	if encoded := values.Encode(); len(encoded) > 0 {
		base = fmt.Sprintf("%s&%s", base, strings.ToLower(encoded))
	}

	return base
}

// Parameters ...
type Parameters struct {
	Ticker         string `json:",omitempty" url:",omitempty"`
	Date           string `json:",omitempty" url:",omitempty"`
	Timestamp      int    `json:",omitempty" url:",omitempty"`
	TimestampLimit int    `json:",omitempty" url:",omitempty"`
	Reverse        bool   `json:",omitempty" url:",omitempty"`
	Limit          int    `json:",omitempty" url:",omitempty"`
	TickType       string `json:",omitempty" url:",omitempty"`
	Direction      string `json:",omitempty" url:",omitempty"`
	Unadjusted     bool   `json:",omitempty" url:",omitempty"`
	Multiplier     int    `json:",omitempty" url:",omitempty"`
	Timespan       string `json:",omitempty" url:",omitempty"`
	From           string `json:",omitempty" url:",omitempty"`
	To             string `json:",omitempty" url:",omitempty"`
	Symbol         string `json:",omitempty" url:",omitempty"`
	Sort           string `json:",omitempty" url:",omitempty"`
	Type           string `json:",omitempty" url:",omitempty"`
	Market         string `json:",omitempty" url:",omitempty"`
	Locale         string `json:",omitempty" url:",omitempty"`
	Search         string `json:",omitempty" url:",omitempty"`
	PerPage        int    `json:",omitempty" url:",omitempty"`
	Page           int    `json:",omitempty" url:",omitempty"`
	Active         bool   `json:",omitempty" url:",omitempty"`
}

// Locale ...
type Locale struct {
	Locale string `json:"locale"`
	Name   string `json:"name"`
}

// Locales ...
func Locales() ([]Locale, error) {
	var locales []Locale

	err := request(endpoint("/v2/reference/locales", nil), &locales)

	return locales, err
}

// Market ...
type Market struct {
	Market      string `json:"market"`
	Description string `json:"desc"`
}

// Markets ...
func Markets() ([]Market, error) {
	var markets []Market

	err := request(endpoint("/v2/reference/markets", nil), &markets)

	return markets, err
}

// Status ...
type Status struct {
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

// MarketStatus ...
func MarketStatus() (status Status, _ error) {
	return status, request(endpoint("/v1/marketstatus/now", nil), &status)
}

// Holiday ...
type Holiday struct {
	Exchange string    `json:"exchange"`
	Name     string    `json:"name"`
	Status   string    `json:"status"`
	Date     time.Time `json:"date"`
	Open     time.Time `json:"open"`
	Close    time.Time `json:"close"`
}

// MarketHolidays ...
func MarketHolidays() (holidays []Holiday, _ error) {
	return holidays, request(endpoint("/v1/marketstatus/upcoming", nil), &holidays)
}

// Ticker ...
type Ticker struct {
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

// Tickers ...
func Tickers(parameters *Parameters) (tickers []Ticker, _ error) {
	return tickers, request(endpoint("/v2/reference/tickers", &parameters), &tickers)
}

// Types ...
type Types struct {
	Types      map[string]string `json:"types"`
	IndexTypes map[string]string `json:"indexTypes"`
}

// TickerTypes ...
func TickerTypes() (types Types, _ error) {
	return types, request(endpoint("/v2/reference/types", nil), &types)
}

// Details ...
type Details struct {
	Logo        string   `json:"logo"`
	Exchange    string   `json:"exchange"`
	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
	Cik         string   `json:"cik"`
	Bloomberg   string   `json:"bloomberg"`
	Lei         string   `json:"lei"`
	Sic         int      `json:"sic"`
	Country     string   `json:"country"`
	Industry    string   `json:"industry"`
	Sector      string   `json:"sector"`
	Marketcap   int64    `json:"marketcap"`
	Employees   int      `json:"employees"`
	Phone       string   `json:"phone"`
	Ceo         string   `json:"ceo"`
	URL         string   `json:"url"`
	Description string   `json:"description"`
	Similar     []string `json:"similar"`
	Tags        []string `json:"tags"`
}

// TickerDetails ...
func TickerDetails(parameters *Parameters) (Details, error) {
	var details Details

	err := request(endpoint(fmt.Sprintf("/v1/meta/symbols/%s/company", parameters.Symbol), nil), &details)

	return details, err
}

// News ...
type News struct {
	Symbols   []string  `json:"symbols"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Source    string    `json:"source"`
	Summary   string    `json:"summary"`
	Image     string    `json:"image"`
	Timestamp time.Time `json:"timestamp"`
	Keywords  []string  `json:"keywords"`
}

// TickerNews ...
func TickerNews(parameters *Parameters) ([]News, error) {
	var news []News

	err := request(endpoint(fmt.Sprintf("/v1/meta/symbols/%s/news", parameters.Symbol), &parameters), &news)

	return news, err
}

// Split ...
type Split struct {
	Ticker       string  `json:"ticker"`
	ExDate       string  `json:"exDate"`
	PaymentDate  string  `json:"paymentDate"`
	RecordDate   string  `json:"recordDate"`
	DeclaredDate string  `json:"declaredDate"`
	Ratio        float64 `json:"ratio"`
	Tofactor     int     `json:"tofactor"`
	Forfactor    int     `json:"forfactor"`
}

// StockSplits ...
func StockSplits(parameters *Parameters) ([]Split, error) {
	var splits []Split

	err := request(endpoint(fmt.Sprintf("/v2/reference/splits/%s", parameters.Symbol), &parameters), &splits)

	return splits, err
}

// Dividend ...
type Dividend struct {
	Symbol       string    `json:"symbol"`
	Type         string    `json:"type"`
	ExDate       time.Time `json:"exDate"`
	PaymentDate  time.Time `json:"paymentDate"`
	RecordDate   time.Time `json:"recordDate"`
	DeclaredDate time.Time `json:"declaredDate"`
	Amount       float64   `json:"amount"`
	Qualified    string    `json:"qualified"`
	Flag         string    `json:"flag"`
}

// StockDividends ...
func StockDividends(parameters *Parameters) ([]Dividend, error) {
	var dividends []Dividend

	err := request(endpoint(fmt.Sprintf("/v2/reference/dividends/%s", parameters.Symbol), &parameters), &dividends)

	return dividends, err
}

// Financial ...
type Financial struct {
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
}

// StockFinancials ...
func StockFinancials(parameters *Parameters) ([]Financial, error) {
	var response struct {
		Results []Financial `json:"results"`
	}

	err := request(endpoint(fmt.Sprintf("/v2/reference/financials/%s", parameters.Symbol), &parameters), &response)

	return response.Results, err
}

// Exchange ...
type Exchange struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	Market string `json:"market"`
	Mic    string `json:"mic"`
	Name   string `json:"name"`
	Tape   string `json:"tape"`
}

// Exchanges ...
func Exchanges() ([]Exchange, error) {
	var exchanges []Exchange

	err := request(endpoint("/v1/meta/exchanges", nil), &exchanges)

	return exchanges, err
}

// Trade ...
type Trade struct {
	Ticker            string  `json:"T"`
	Timestamp         int64   `json:"t"`
	TimestampExchange int64   `json:"y"`
	TimestampTRF      int64   `json:"f"`
	SequenceNumber    int     `json:"q"`
	TradeID           string  `json:"i"`
	ExchangeID        int     `json:"x"`
	Size              int     `json:"s"`
	Conditions        []int   `json:"c"`
	Price             float64 `json:"p"`
	Tape              int     `json:"z"`
}

// HistoricTrades ...
func HistoricTrades(parameters *Parameters) ([]Trade, error) {
	var response struct {
		Results []Trade `json:"results"`
	}

	err := request(endpoint(fmt.Sprintf("/v2/ticks/stocks/trades/%s/%s", parameters.Ticker, parameters.Date), &parameters), &response)

	return response.Results, err
}

// Quote ...
type Quote struct {
	Ticker            string  `json:"T"`
	Timestamp         int64   `json:"t"`
	TimestampExchange int64   `json:"y"`
	TimestampTRF      int64   `json:"f"`
	SequenceNumber    int     `json:"q"`
	Conditions        []int   `json:"c"`
	Indicators        []int   `json:"i"`
	Bid               float64 `json:"p"`
	BidExchangeID     int     `json:"x"`
	BidSize           int     `json:"s"`
	Ask               float64 `json:"P"`
	AskExchangeID     int     `json:"X"`
	AskSize           int     `json:"S"`
	Tape              int     `json:"z"`
}

// HistoricQuotes ...
func HistoricQuotes(parameters *Parameters) ([]Quote, error) {
	var response struct {
		Results []Quote `json:"results"`
	}

	err := request(endpoint(fmt.Sprintf("/v2/ticks/stocks/nbbo/%s/%s", parameters.Ticker, parameters.Date), &parameters), &response)

	return response.Results, err
}

// LastTradeForASymbol ...
func LastTradeForASymbol(parameters *Parameters) (Trade, error) {
	var response struct {
		Last struct {
			Price      float64 `json:"price"`
			Size       int     `json:"size"`
			Exchange   int     `json:"exchange"`
			Condition1 int     `json:"cond1"`
			Condition2 int     `json:"cond2"`
			Condition3 int     `json:"cond3"`
			Condition4 int     `json:"cond4"`
			Timestamp  int64   `json:"timestamp"`
		} `json:"last"`
	}

	err := request(endpoint(fmt.Sprintf("/v1/last/stocks/%s", parameters.Ticker), &parameters), &response)

	return Trade{
		Price:      response.Last.Price,
		Size:       response.Last.Size,
		ExchangeID: response.Last.Exchange,
		Conditions: []int{
			response.Last.Condition1,
			response.Last.Condition2,
			response.Last.Condition3,
			response.Last.Condition4,
		},
		Timestamp: response.Last.Timestamp,
	}, err
}

// LastQuoteForASymbol ...
func LastQuoteForASymbol(parameters *Parameters) (Quote, error) {
	var response struct {
		Last Quote `json:"last"`
	}

	err := request(endpoint(fmt.Sprintf("/v1/last_quote/stocks/%s", parameters.Ticker), &parameters), &response)

	return response.Last, err
}

// Aggregate ...
type Aggregate struct {
	Ticker     string  `json:"T"`
	Volume     int     `json:"v"`
	Open       float64 `json:"o"`
	Close      float64 `json:"c"`
	High       float64 `json:"h"`
	Low        float64 `json:"l"`
	Timestamp  int64   `json:"t"`
	Items      int     `json:"n"`
	Condition1 int     `json:"c1"`
	Condition2 int     `json:"c2"`
	Condition3 int     `json:"c3"`
	Condition4 int     `json:"c4"`
}

// DailyOpenCloseAfterHours ...
func DailyOpenCloseAfterHours(parameters *Parameters) (Aggregate, Aggregate, Aggregate, error) {
	var aggregate struct {
		Symbol     string    `json:"symbol"`
		Open       Aggregate `json:"open"`
		Close      Aggregate `json:"close"`
		AfterHours Aggregate `json:"afterHours"`
	}

	err := request(endpoint(fmt.Sprintf("/v1/open-close/%s/%s", parameters.Ticker, parameters.Date), &parameters), &aggregate)

	return aggregate.Open, aggregate.Close, aggregate.AfterHours, err
}

// ConditionMappings ...
func ConditionMappings(parameters *Parameters) (map[string]string, error) {
	var mappings map[string]string

	err := request(endpoint(fmt.Sprintf("/v1/meta/conditions/%s", parameters.TickType), &parameters), &mappings)

	return mappings, err
}

// Snapshot ...
type Snapshot struct {
	Ticker string `json:"ticker"`
	Day    struct {
		Close  float64 `json:"c"`
		High   float64 `json:"h"`
		Low    float64 `json:"l"`
		Open   float64 `json:"o"`
		Volume float64 `json:"v"`
	} `json:"day"`
	LastTrade struct {
		Condition1 int     `json:"c1"`
		Condition2 int     `json:"c2"`
		Condition3 int     `json:"c3"`
		Condition4 int     `json:"c4"`
		Exchange   int     `json:"e"`
		Price      float64 `json:"p"`
		Size       int     `json:"s"`
		Timestamp  int64   `json:"t"`
	} `json:"lastTrade"`
	LastQuote struct {
		Ask       int   `json:"p"`
		AskSize   int   `json:"s"`
		Bid       int   `json:"P"`
		BidSize   int   `json:"S"`
		Timestamp int64 `json:"t"`
	} `json:"lastQuote"`
	Min struct {
		Close  float64 `json:"c"`
		High   float64 `json:"h"`
		Low    float64 `json:"l"`
		Open   float64 `json:"o"`
		Volume float64 `json:"v"`
	} `json:"min"`
	PrevDay struct {
		Close  float64 `json:"c"`
		High   float64 `json:"h"`
		Low    float64 `json:"l"`
		Open   float64 `json:"o"`
		Volume float64 `json:"v"`
	} `json:"prevDay"`
	TodaysChange     float64 `json:"todaysChange"`
	TodaysChangePerc float64 `json:"todaysChangePerc"`
	Updated          int64   `json:"updated"`
}

// SnapshotAllTickers ...
func SnapshotAllTickers() ([]Snapshot, error) {
	var response struct {
		Tickers []Snapshot `json:"tickers"`
	}

	err := request(endpoint("/v2/snapshot/locale/us/markets/stocks/tickers", nil), &response)

	return response.Tickers, err
}

// SnapshotSingleTicker ...
func SnapshotSingleTicker(parameters *Parameters) (Snapshot, error) {
	var response struct {
		Ticker Snapshot `json:"ticker"`
	}

	err := request(endpoint(fmt.Sprintf("/v2/snapshot/locale/us/markets/stocks/tickers/%s", parameters.Ticker), nil), &response)

	return response.Ticker, err
}

// SnapshotGainersLosers ...
func SnapshotGainersLosers(parameters *Parameters) ([]Snapshot, error) {
	var response struct {
		Tickers []Snapshot `json:"tickers"`
	}

	err := request(endpoint(fmt.Sprintf("/v2/snapshot/locale/us/markets/stocks/%s", parameters.Direction), nil), &response)

	return response.Tickers, err
}

// PreviousClose ...
func PreviousClose(parameters *Parameters) (Aggregate, error) {
	var response struct {
		Results Aggregate `json:"results"`
	}

	err := request(endpoint(fmt.Sprintf("/v2/aggs/ticker/%s/prev", parameters.Ticker), &parameters), &response)

	return response.Results, err
}

// Aggregates ...
func Aggregates(parameters *Parameters) ([]Aggregate, error) {
	var response struct {
		Results []Aggregate `json:"results"`
	}

	err := request(endpoint(fmt.Sprintf("/v2/aggs/ticker/%s/range/%d/%s/%s/%s", parameters.Ticker, parameters.Multiplier, parameters.Timespan, parameters.From, parameters.To), nil), &response)

	return response.Results, err
}

// GroupedDaily ...
func GroupedDaily(parameters *Parameters) ([]Aggregate, error) {
	var response struct {
		Results []Aggregate `json:"results"`
	}

	err := request(endpoint(fmt.Sprintf("/v2/aggs/grouped/locale/%s/market/%s/%s", parameters.Locale, parameters.Market, parameters.Date), &parameters), &response)

	return response.Results, err
}
