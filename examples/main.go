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

	c, err := choruspro.NewClient().WithConfig(cfg)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	ctx := context.Background()

	res, err := c.Transverses.RecupererDevises(ctx, choruspro.CodeLangueFr)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	result, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Final Response: %v", string(result))
}
