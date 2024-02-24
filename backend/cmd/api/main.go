package main

import (
	"log"

	"balanceyourportfolio/api"
	"balanceyourportfolio/config"
)

func main() {
	cfg := config.Load()
	err := api.ListenAndServe(cfg.APIHost, cfg.APIPort)
	if err != nil {
		log.Fatal(err)
	}
}
