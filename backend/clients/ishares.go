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

const (
	productId   = "253743"
	productName = "ishares-sp-500-b-ucits-etf-acc-fund"
)

func (v *IShares) GetFunds() (resp []ishares.FundResponse, err error) {
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
		for _, m := range fundUrlMatches {
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
			fundResponse = &ishares.FundResponse{}
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
			fundResponse = &ishares.FundResponse{}
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
				fundResponse = &ishares.FundResponse{}
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
	c.OnHTML("div.col-totalNetAssetsFundLevel", func(e *colly.HTMLElement) {
		currencyNetAssets := e.ChildText("span.data")
		netAssetsFloat, err := strconv.ParseFloat(strings.ReplaceAll(currencyNetAssets[4:], ",", ""), 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		fundResponse, ok := fundIdentifierMap[e.Request.URL.Path]
		if !ok {
			fundResponse = &ishares.FundResponse{}
			fundIdentifierMap[e.Request.URL.Path] = fundResponse
		}
		fundResponse.NetAssets = netAssetsFloat
	})

	err = c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 3 * time.Second,
		Delay:       3 * time.Second,
		Parallelism: 1,
	})
	if err != nil {
		return resp, err
	}

	err = c.Visit(fmt.Sprintf("%s/uk/professional/en/products/etf-investments?switchLocale=y&siteEntryPassthrough=true", v.url))
	if err != nil {
		return resp, err
	}
	c.Wait()
	return resp, err
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

func scrapeISIN(ticker string) string {
	url := "https://www.ishares.com/uk/professional/en/product-screener/product-screener-v3.1.jsn?type=customized-excel"

	// Replace the following with your actual request payload
	payload := []byte(`{"productView":"all","portfolios":[239726,244049,239458,239763,239774,239619,253743,286083,311863,244050,251882,287737,309035,332655,239710,239708,239623,253742,239454,239566,239456,239451,264659,295689,239600,239500,239565,250989,228471,228472,228522,228523,228524,228634,228691,229050,243850,250990,250991,287647,287648,287649,298006,305353,306174,309085,315973,321412,327286,329827,251726,290619,295863,310017,331356,333156,239637,239572,251900,307527,307528,319355,335098,239665,258441,307241,307243,308634,239714,253741,304353,291299,239463,251795,291401,251715,287340,296771,296772,304347,228444,229526,230288,237602,269666,286433,290634,325695,251850,239627,290846,291392,316039],"dataPoints":["localExchangeTicker","fundName","productType","seriesBaseCurrencyCode","investorClassName","useOfProfits","ter_ocf","totalFundSizeInMillions","totalFundSizeInMillionsAsOf","domicile","isin"],"dcrPath":"/templatedata/config/product-screener-v3/data/en/uk/product-screener/ishares-product-screener-excel-config","disclosureContentDcrPath":"/templatedata/content/article/data/en/one/DEFAULT/product-screener-all-disclaimer"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
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
		fmt.Println("Error sending request:", err)
		return ""
	}

	scanner := bufio.NewScanner(resp.Body)

	startCounting := false
	lineCounter := 0
	for scanner.Scan() {
		line := scanner.Text()
		if startCounting {
			lineCounter++
		}
		if strings.Contains(line, "MCHI") {
			fmt.Println(line)
			startCounting = true
		}
		if lineCounter == 9 {
			fmt.Println(strings.Split(strings.Split(line, ">")[1], "<")[0])
		}
	}

}
