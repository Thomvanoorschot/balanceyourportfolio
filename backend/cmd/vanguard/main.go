package main

import (
	"context"
	"fmt"

	"etfinsight/clients"
	"etfinsight/config"
	"etfinsight/repositories/pgrepo"
	"etfinsight/services/vanguard"
)

func main() {
	cfg := config.Load()

	repo := pgrepo.NewRepository(cfg)
	client := clients.NewVanguard(cfg)
	svc := vanguard.NewService(client, repo)
	err := svc.UpsertFunds(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Done")
}
