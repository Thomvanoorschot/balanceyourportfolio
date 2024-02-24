package clients

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"balanceyourportfolio/config"
)

type Vanguard struct {
	url    string
	client *http.Client
}

func NewVanguard(cfg *config.Config) *Vanguard {
	return &Vanguard{url: cfg.VanguardUrl, client: &http.Client{}}
}

func (v *Vanguard) GetFunds(ei []string) ([]byte, error) {
	variables := map[string]interface{}{
		"portIds": ei,
	}
	return v.queryGraphql(variables, fundsQuery)
}

func (v *Vanguard) queryGraphql(variables map[string]interface{}, query string) ([]byte, error) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"query":     query,
		"variables": variables,
	})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", v.url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-consumer-id", "nl0")

	resp, err := v.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	return io.ReadAll(resp.Body)
}

const fundsQuery = `
query FundsHoldingsQuery($portIds: [String!]!) {
	polarisAnalyticsHistory(portIds: $portIds){
		monthly{
			valuation{
				fund{
					items{
						OSCLTSHQTY{
							effectiveDate
							outstandingShare
						}
					}
				}
			}
		}
	}
	funds(portIds: $portIds) {
		profile {
		  portId
		  fundFullName
		  fundCurrency
		  listings {
			portId
			identifiers(
			  altIds: ["RIC"]
			) {
			  altId
			  altIdValue
			}
		  }
		  identifiers(
			altIds: ["ISIN"]
		  ) {
			altId
			altIdValue
		  }
		}
		pricingDetails{
			navPrices{
				items{
					price
				}
		  	}
		}
		brokerBasketData {
			holdings(limit: -1){
				ticker,
				isin
			}
		}
		holdings {
			originalHoldings(
				limit: -1
			) {
				totalHoldings
				items {
					ticker
					issueTypeName
					numberOfShares
					marketValPercent
					marketValue
					name
					sectorName
					sedol
					CUSIP
					countryCode
				}
			}
		}
	}
}
`
