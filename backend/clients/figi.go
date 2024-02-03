package clients

import (
	"bytes"
	"encoding/json"
	"etfinsight/services/vanguard"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Figi struct {
}

func NewFigi() *Figi {
	return &Figi{}
}

func (f *Figi) GetFigi(payload []vanguard.FigiPayload) (r []vanguard.FigiResp, err error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling payload:", err)
		return r, err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "https://api.openfigi.com/v3/mapping", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return r, err
	}

	// Set the content type header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-OPENFIGI-APIKEY", "b3fd6244-4b64-4b3c-8a87-7a808f197a01")
	// Execute the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return r, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return r, err
	}
	if resp.StatusCode == 429 {
		<-time.NewTimer(7 * time.Second).C
		return f.GetFigi(payload)
	}
	var figiResp []vanguard.FigiResp
	err = json.Unmarshal(body, &figiResp)
	if err != nil {
		fmt.Println("Could not unmarshal body", err)
		return r, err
	}
	return figiResp, nil
}
