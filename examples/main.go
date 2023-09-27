package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/antoine2116/go-choruspro"
)

func main() {
	cfg := &choruspro.ClientConfig{
		BaseUrl:      "<piste_url>",
		AuthUrl:      "<piste_oauth_url>",
		ClientId:     "<piste_client_id>",
		ClientSecret: "<piste_client_secret>",
		Login:        "<chorus_pro_technical_credentials>",
	}

	c := choruspro.NewClient().WithConfig(cfg)

	ctx := context.Background()

	res, err := c.Transverses.ConsulterCompteRendu(ctx, choruspro.CompteRenduOptions{
		NumeroFluxDepot: "",
	})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	result, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Final Response: %v", string(result))
}
