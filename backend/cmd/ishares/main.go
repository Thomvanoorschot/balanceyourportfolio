package main

import (
	"context"
	"etfinsight/services/ishares"
	"fmt"

	"etfinsight/clients"
	"etfinsight/config"
	"etfinsight/repositories/pgrepo"
)

func main() {
	cfg := config.Load()

	repo := pgrepo.NewRepository(cfg)
	client := clients.NewIShares(cfg)
	figiClient := clients.NewFigi[ishares.FigiPayload, ishares.FigiResp]()
	svc := ishares.NewService(client, repo, figiClient)
	err := svc.UpsertFunds(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Done")
}
