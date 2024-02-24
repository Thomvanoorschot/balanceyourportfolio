package main

import (
	"context"
	"fmt"

	"balanceyourportfolio/clients"
	"balanceyourportfolio/config"
	"balanceyourportfolio/repositories/pgrepo"
	"balanceyourportfolio/services/vanguard"
)

func main() {
	cfg := config.Load()

	repo := pgrepo.NewRepository(cfg)
	client := clients.NewVanguard(cfg)
	figiClient := clients.NewFigi[vanguard.FigiPayload, vanguard.FigiResp]()
	svc := vanguard.NewService(client, repo, figiClient)
	err := svc.UpsertFunds(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Done")
}
