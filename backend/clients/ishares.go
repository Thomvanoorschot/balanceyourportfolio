package clients

import (
	"bufio"
	"bytes"
	"encoding/json"
	"etfinsight/services/ishares"
	"fmt"
	"github.com/gocolly/colly/v2"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"etfinsight/config"
)

type IShares struct {
	url    string
	client *http.Client
}

func NewIShares(cfg *config.Config) *IShares {
	return &IShares{url: cfg.ISharesUrl, client: &http.Client{}}
}

var tickerISINs = map[string]string{}

func (v *IShares) GetFunds(limit, offset int) (resp []ishares.FundResponse, err error) {
	c := colly.NewCollector(
		colly.AllowedDomains(strings.ReplaceAll(v.url, "https://", "")),
		colly.Async(true),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("html", func(request *colly.HTMLElement) {
		fundUrlRegex, err := regexp.Compile("<td class=\"links\"><a href=\"(.*)\">")
		if err != nil {
			return
		}
		fundUrlMatches := fundUrlRegex.FindAllStringSubmatch(request.Text, -1)
		for i, m := range fundUrlMatches {
			if i < offset {
				continue
			}
			if i > offset+limit {
				return
			}
			fundUrl := fmt.Sprintf("%s%s", v.url, m[1])
			hasVisitedFund, err := c.HasVisited(fundUrl)
			if err != nil {
				fmt.Println(err)
				return
			}
			if !hasVisitedFund {
				err := c.Visit(fundUrl)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	})

	var fundIdentifierMap = map[string]*ishares.FundResponse{}
	c.OnHTML("h1.product-title", func(e *colly.HTMLElement) {
		fundResponse, ok := fundIdentifierMap[e.Request.URL.Path]
		if !ok {
			fundResponse = &ishares.FundResponse{
				HoldingsTableIndex: map[string]int{},
			}
			fundIdentifierMap[e.Request.URL.Path] = fundResponse
		}

		fundUrlRegex, err := regexp.Compile("/uk/professional/en/products/([0-9]*)")
		if err != nil {
			return
		}
		fundName := strings.ReplaceAll(e.Text, "\n", "")
		fundResponse.FundName = fundName
		fundIdentifier := fundUrlRegex.FindStringSubmatch(e.Request.URL.Path)
		fundResponse.ExternalIdentifier = fundIdentifier[1]
	})
	c.OnHTML("select.date-dropdown", func(e *colly.HTMLElement) {
		fundResponse, ok := fundIdentifierMap[e.Request.URL.Path]
		if !ok {
			fundResponse = &ishares.FundResponse{
				HoldingsTableIndex: map[string]int{},
			}
			fundIdentifierMap[e.Request.URL.Path] = fundResponse
		}
		dateStr := e.ChildAttr("option", "value")
		date, err := time.Parse("20060102", dateStr)
		if err != nil {
			fmt.Println(err)
			return
		}
		holdingUrl := fmt.Sprintf("%s%s/1506575576011.ajax?tab=all&fileType=json&asOfDate=%s", v.url, e.Request.URL.Path, dateStr)
		holdings, err := v.GetHoldings(holdingUrl)
		if err != nil {
			fmt.Println(err)
			return
		}
		fundResponse.Holdings = holdings
		fundResponse.EffectiveDate = date
	})
	c.OnHTML("span.header-nav-data", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "USD") || strings.Contains(e.Text, "EUR") || strings.Contains(e.Text, "GBP") {
			currencyPrice := strings.ReplaceAll(e.Text, "\n", "")
			splitCurrencyPrice := strings.Split(currencyPrice, " ")
			currency := splitCurrencyPrice[0]
			price := splitCurrencyPrice[1]
			fundResponse, ok := fundIdentifierMap[e.Request.URL.Path]
			if !ok {
				fundResponse = &ishares.FundResponse{
					HoldingsTableIndex: map[string]int{},
				}
				fundIdentifierMap[e.Request.URL.Path] = fundResponse
			}
			fundResponse.Currency = currency
			priceFloat, err := strconv.ParseFloat(price, 64)
			if err != nil {
				fmt.Println(err)
				return
			}
			fundResponse.Price = priceFloat
		}
	})
	c.OnHTML("p.identifier", func(e *colly.HTMLElement) {
		fundResponse, ok := fundIdentifierMap[e.Request.URL.Path]
		if !ok {
			fundResponse = &ishares.FundResponse{
				HoldingsTableIndex: map[string]int{},
			}
			fundIdentifierMap[e.Request.URL.Path] = fundResponse
		}
		ticker := strings.ReplaceAll(e.Text, "\n", "")
		fundResponse.Ticker = ticker
	})
	c.OnHTML("td.colTicker", func(e *colly.HTMLElement) {
		fundResponse, ok := fundIdentifierMap[e.Request.URL.Path]
		if !ok {
			fundResponse = &ishares.FundResponse{
				HoldingsTableIndex: map[string]int{},
			}
			fundIdentifierMap[e.Request.URL.Path] = fundResponse
		}
		ticker := strings.ReplaceAll(e.Text, "\n", "")
		if ticker != "Ticker" {
			for _, mappedTicker := range fundResponse.Tickers {
				if mappedTicker == ticker {
					return
				}
			}
			fundResponse.Tickers = append(fundResponse.Tickers, ticker)
		}
	})
	c.OnHTML("div.col-totalNetAssetsFundLevel", func(e *colly.HTMLElement) {
		currencyNetAssets := e.ChildText("span.data")
		netAssetsFloat, err := strconv.ParseFloat(strings.ReplaceAll(currencyNetAssets[4:], ",", ""), 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		fundResponse, ok := fundIdentifierMap[e.Request.URL.Path]
		if !ok {
			fundResponse = &ishares.FundResponse{
				HoldingsTableIndex: map[string]int{},
			}
			fundIdentifierMap[e.Request.URL.Path] = fundResponse
		}
		fundResponse.NetAssets = netAssetsFloat
	})

	//var columns := []string{"colTicker", "colIssueName", "colSectorName", "colAssetClass", "colMarketValue", "colHoldingPercent", "colUnitsHeld", "colIsin"}

	c.OnHTML("th.colTicker", listenToCol("colTicker", fundIdentifierMap))
	c.OnHTML("th.colIssueName", listenToCol("colIssueName", fundIdentifierMap))
	c.OnHTML("th.colSectorName", listenToCol("colSectorName", fundIdentifierMap))
	c.OnHTML("th.colAssetClass", listenToCol("colAssetClass", fundIdentifierMap))
	c.OnHTML("th.colMarketValue", listenToCol("colMarketValue", fundIdentifierMap))
	c.OnHTML("th.colHoldingPercent", listenToCol("colHoldingPercent", fundIdentifierMap))
	c.OnHTML("th.colUnitsHeld", listenToCol("colUnitsHeld", fundIdentifierMap))
	c.OnHTML("th.colIsin", listenToCol("colIsin", fundIdentifierMap))
	c.OnHTML("th.colParValue", listenToCol("colParValue", fundIdentifierMap))

	err = c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 1 * time.Second,
		Delay:       1 * time.Second,
		Parallelism: 1,
	})
	if err != nil {
		return resp, err
	}
	err = c.Visit(fmt.Sprintf("%s/uk/professional/en/products/etf-investments?switchLocale=y&siteEntryPassthrough=true#/?productView=etf&", v.url))
	if err != nil {
		return resp, err
	}
	c.Wait()
	var identifiers []int
	for _, val := range fundIdentifierMap {
		intIdentifier, err := strconv.Atoi(val.ExternalIdentifier)
		if err != nil {
			fmt.Println(err)
			return resp, err
		}
		identifiers = append(identifiers, intIdentifier)
	}
	err = scrapeISINs(identifiers)
	if err != nil {
		return resp, err
	}
	for _, val := range fundIdentifierMap {
		isin, ok := tickerISINs[val.Ticker]
		if ok {
			val.ISIN = isin
		}
		resp = append(resp, *val)
	}

	return resp, err
}
func listenToCol(col string, fundIdentifierMap map[string]*ishares.FundResponse) func(element *colly.HTMLElement) {
	return func(e *colly.HTMLElement) {
		tableName, _ := e.DOM.Parent().Parent().Parent().Attr("id")
		if tableName != "allHoldingsTable" {
			return
		}
		fundResponse, ok := fundIdentifierMap[e.Request.URL.Path]
		if !ok {
			fundResponse = &ishares.FundResponse{
				HoldingsTableIndex: map[string]int{},
			}
			fundIdentifierMap[e.Request.URL.Path] = fundResponse
		}
		columnIndex, err := regexp.Compile("col(\\d+)")
		if err != nil {
			return
		}
		fundUrlMatches := columnIndex.FindStringSubmatch(e.Attr("class"))
		if len(fundUrlMatches) < 2 {
			return
		}
		idx, err := strconv.Atoi(fundUrlMatches[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		fundResponse.HoldingsTableIndex[col] = idx - 1
	}
}
func (v *IShares) GetHoldings(url string) (resp ishares.HoldingsResponse, err error) {
	r, err := v.client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return resp, err
	}
	b = bytes.TrimPrefix(b, []byte("\xef\xbb\xbf"))

	holdingsResponse := ishares.HoldingsResponse{}
	err = json.Unmarshal(b, &holdingsResponse)
	if err != nil {
		return resp, err
	}
	return holdingsResponse, nil
}

type OverviewQuery struct {
	ProductView              string   `json:"productView"`
	Portfolios               []int    `json:"portfolios"`
	DataPoints               []string `json:"dataPoints"`
	DcrPath                  string   `json:"dcrPath"`
	DisclosureContentDcrPath string   `json:"disclosureContentDcrPath"`
}

func scrapeISINs(identifiers []int) error {
	url := "https://www.ishares.com/uk/professional/en/product-screener/product-screener-v3.1.jsn?type=customized-excel"

	// Replace the following with your actual request payload
	body := OverviewQuery{
		ProductView:              "etf",
		Portfolios:               identifiers,
		DataPoints:               []string{"localExchangeTicker", "isin"},
		DcrPath:                  "/templatedata/config/product-screener-v3/data/en/uk/product-screener/ishares-product-screener-excel-config",
		DisclosureContentDcrPath: "/templatedata/content/article/data/en/one/DEFAULT/product-screener-all-disclaimer",
	}
	b, err := json.Marshal(body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	// Add other headers as needed

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(resp.Body)

	lineCounter := 0
	fundLineCounter := 0
	latestTicker := ""
	for scanner.Scan() {
		line := scanner.Text()
		if lineCounter < 44 {
			lineCounter++
			continue
		}
		if fundLineCounter == 0 {
			latestTicker = strings.Split(strings.Split(line, ">")[1], "<")[0]
			fundLineCounter++
			continue
		}
		if fundLineCounter == 3 {
			isin := strings.Split(strings.Split(line, ">")[1], "<")[0]
			tickerISINs[latestTicker] = isin
			fundLineCounter = -4
			continue
		}
		fundLineCounter++
	}
	return nil
}
