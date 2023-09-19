package choruspro

import (
	"context"
	"net/http"
)

type RecupererTauxTvaParams struct {
	CodeLangue string `json:"codeLangue"`
}

type RecupererTauxTvaResponse struct {
	CodeRetour   int32     `json:"codeRetour"`
	Libelle      string    `json:"libelle"`
	ListeTauxTva []TauxTva `json:"listeTauxTva"`
}

type TauxTva struct {
	CodeTauxTva    string  `json:"codeTauxTva"`
	LibelleTauxTva string  `json:"libelleTauxTva"`
	ValeurTauxTva  float32 `json:"valeurTauxTva"`
}

func (c *Client) RecupererTauxTva(ctx context.Context, params RecupererTauxTvaParams) (*RecupererTauxTvaResponse, error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/tauxtva", params)
	if err != nil {
		return nil, err
	}

	taux := new(RecupererTauxTvaResponse)

	err = c.doRequest(ctx, req, taux)
	if err != nil {
		return nil, err
	}

	return taux, nil
}
