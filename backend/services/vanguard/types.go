package vanguard

type FundsResponse struct {
	Data FundsData `json:"data"`
}

type FundsData struct {
	Funds                     []Fund                    `json:"funds"`
	PolarisAnalyticsHistories []PolarisAnalyticsHistory `json:"polarisAnalyticsHistory"`
}
type NavItem struct {
	Price float64 `json:"price"`
}
type NavPrices struct {
	Items []NavItem `json:"items"`
}
type PricingDetails struct {
	NavPrices NavPrices `json:"navPrices"`
}
type Profile struct {
	FundFullName string                   `json:"fundFullName"`
	FundCurrency string                   `json:"fundCurrency"`
	PortId       string                   `json:"portId"`
	Listings     []FundInformationListing `json:"listings"`
	Identifiers  []Identifier             `json:"identifiers"`
}

type Identifier struct {
	AltId      string `json:"altId"`
	AltIdValue string `json:"altIdValue"`
}

type HoldingsItem struct {
	Ticker           string  `json:"ticker"`
	IssueTypename    string  `json:"issueTypeName"`
	NumberOfShares   float64 `json:"numberOfShares"`
	MarketValPercent float64 `json:"marketValPercent"`
	MarketValue      float64 `json:"marketValue"`
	Name             string  `json:"name"`
	SectorName       string  `json:"sectorName"`
	SEDOL            *string `json:"sedol"`
	CUSIP            *string `json:"CUSIP"`
	CountryCode      string  `json:"countryCode"`
}

type OriginalHoldings struct {
	TotalHoldings float64        `json:"totalHoldings"`
	Items         []HoldingsItem `json:"items"`
}
type Holdings struct {
	OriginalHoldings OriginalHoldings `json:"originalHoldings"`
}
type Fund struct {
	PricingDetails PricingDetails `json:"pricingDetails"`
	Profile        Profile        `json:"profile"`
	Holdings       Holdings       `json:"holdings"`
}

type FundInformation struct {
	Data FundInformationData `json:"data"`
}

type FundInformationData struct {
	Funds []FundInformationFund `json:"funds"`
}

type FundInformationFund struct {
	PortId  string                 `json:"portId"`
	Profile FundInformationProfile `json:"profile"`
}

type FundInformationProfile struct {
	Listings []FundInformationListing `json:"listings"`
}

type FundInformationListing struct {
	Identifiers []Identifier `json:"identifiers"`
}
type PolarisAnalyticsHistory struct {
	AnalyticsMonthly PolarisAnalyticsMonthly `json:"monthly"`
}
type PolarisAnalyticsMonthly struct {
	Valuation PolarisValuation `json:"valuation"`
}
type PolarisValuation struct {
	Fund PolarisFund `json:"fund"`
}
type PolarisFund struct {
	Items []PolarisItem `json:"items"`
}
type PolarisItem struct {
	OSCLTSHQTY []PolarisOSCLTSHQTY `json:"OSCLTSHQTY"`
}
type PolarisOSCLTSHQTY struct {
	EffectiveDate    string  `json:"effectiveDate"`
	OutstandingShare float64 `json:"outstandingShare"`
}
