package clients

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"etfinsight/config"
)

type IShares struct {
	url    string
	client *http.Client
}

func NewIShares(cfg *config.Config) *IShares {
	return &IShares{url: cfg.ISharesUrl, client: &http.Client{}}
}

func (v *IShares) GetFunds() ([]byte, error) {
	return v.query("253743", "ishares-sp-500-b-ucits-etf-acc-fund")
}

func (v *IShares) query(productId string, productName string) ([]byte, error) {
	fmt.Println(fmt.Sprintf("%s/uk/professional/en/products/%s/%s/1506575576011.ajax?tab=all&fileType=json&asOfDate=20240109", v.url, productId, productName))
	resp, err := v.client.Get(fmt.Sprintf("%s/uk/professional/en/products/%s/%s/1506575576011.ajax?tab=all&fileType=json&asOfDate=20240109", v.url, productId, productName))
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
