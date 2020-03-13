package reference

import (
	"fmt"
	"polygon/client"
	"time"
)

// Parameters ...
type Parameters struct {
	Symbol  string `json:",omitempty" url:",omitempty"`
	Sort    string `json:",omitempty" url:",omitempty"`
	Type    string `json:",omitempty" url:",omitempty"`
	Market  string `json:",omitempty" url:",omitempty"`
	Locale  string `json:",omitempty" url:",omitempty"`
	Search  string `json:",omitempty" url:",omitempty"`
	PerPage int    `json:",omitempty" url:",omitempty"`
	Page    int    `json:",omitempty" url:",omitempty"`
	Active  bool   `json:",omitempty" url:",omitempty"`
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
func Tickers(parameters *Parameters) ([]Ticker, error) {
	var tickers []Ticker

	err := client.Request(client.Endpoint("/v2/reference/tickers", &parameters), &tickers)

	return tickers, err
}

// Types ...
type Types struct {
	Types      map[string]string `json:"types"`
	IndexTypes map[string]string `json:"indexTypes"`
}

// TickerTypes ...
func TickerTypes() (Types, error) {
	var types Types

	err := client.Request(client.Endpoint("/v2/reference/types", nil), &types)

	return types, err
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

	err := client.Request(client.Endpoint(fmt.Sprintf("/v1/meta/symbols/%s/company", parameters.Symbol), nil), &details)

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

	err := client.Request(client.Endpoint(fmt.Sprintf("/v1/meta/symbols/%s/news", parameters.Symbol), &parameters), &news)

	return news, err
}

// Market ...
type Market struct {
	Market      string `json:"market"`
	Description string `json:"desc"`
}

// Markets ...
func Markets() ([]Market, error) {
	var markets []Market

	err := client.Request(client.Endpoint("/v2/reference/markets", nil), &markets)

	return markets, err
}

// Locale ...
type Locale struct {
	Locale string `json:"locale"`
	Name   string `json:"name"`
}

// Locales ...
func Locales() ([]Locale, error) {
	var locales []Locale

	err := client.Request(client.Endpoint("/v2/reference/locales", nil), &locales)

	return locales, err
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

	err := client.Request(client.Endpoint(fmt.Sprintf("/v2/reference/splits/%s", parameters.Symbol), &parameters), &splits)

	return splits, err
}

// Dividends ...
type Dividends struct {
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
func StockDividends(parameters *Parameters) ([]Dividends, error) {
	var dividends []Dividends

	err := client.Request(client.Endpoint(fmt.Sprintf("/v2/reference/dividends/%s", parameters.Symbol), &parameters), &dividends)

	return dividends, err
}

// Financials ...
type Financials struct {
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
func StockFinancials(parameters *Parameters) ([]Financials, error) {
	var response struct {
		Results []Financials `json:"results"`
	}

	err := client.Request(client.Endpoint(fmt.Sprintf("/v2/reference/financials/%s", parameters.Symbol), &parameters), &response)

	return response.Results, err
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
func MarketStatus() (Status, error) {
	var status Status

	err := client.Request(client.Endpoint("/v1/marketstatus/now", nil), &status)

	return status, err
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
func MarketHolidays() ([]Holiday, error) {
	var holidays []Holiday

	err := client.Request(client.Endpoint("/v1/marketstatus/upcoming", nil), &holidays)

	return holidays, err
}
