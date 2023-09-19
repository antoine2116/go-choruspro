package main

import (
	"context"
	"log"

	"github.com/antoine2116/go-choruspro"
)

func main() {
	c := choruspro.NewClient(&choruspro.ClientConfig{
		Url:          "<piste_url>",
		AuthUrl:      "<piste_oauth_url>",
		ClientId:     "<piste_client_id>",
		ClientSecret: "<piste_client_secret>",
		Login:        "<chorus_pro_technical_credentials>",
	})

	ctx := context.Background()

	taux, err := c.RecupererTauxTva(ctx, choruspro.RecupererTauxTvaParams{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Taux: %v", taux)
}
