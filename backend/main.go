package main

import (
	"log"

	"etfinsight/api"
	"etfinsight/config"
)

func main() {
	cfg := config.Load()
	err := api.ListenAndServe(cfg.APIHost, cfg.APIPort)
	if err != nil {
		log.Fatal(err)
	}

}
