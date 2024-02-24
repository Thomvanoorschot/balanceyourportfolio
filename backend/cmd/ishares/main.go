package main

import (
	"balanceyourportfolio/services/ishares"
	"context"
	"fmt"

	"balanceyourportfolio/clients"
	"balanceyourportfolio/config"
	"balanceyourportfolio/repositories/pgrepo"
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
