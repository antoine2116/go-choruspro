package choruspro

import (
	"context"
	"net/http"
)

type TraiterFactureAValiderOptions struct {
	Action         ActionFacture `json:"action,omitempty"`
	ComplementInfo string        `json:"complementInfo,omitempty"`
	IdFacture      int64         `json:"idFacture,omitempty"`
	IdPieceJointe  int64         `json:"idPieceJointe,omitempty"`
	IdValideur     int64         `json:"idValideur,omitempty"`
	MotifRefus     string        `json:"motifRefus,omitempty"`
	TypeValideur   string        `json:"typeValideur,omitempty"`
}

type TraiterFactureAValiderResponse struct {
	CodeRetour     int    `json:"codeRetour,omitempty"`
	DateTraitement *Date  `json:"dateTraitement,omitempty"`
	IdFacture      int64  `json:"idFacture,omitempty"`
	Libelle        string `json:"libelle,omitempty"`
	NumeroFacture  string `json:"numeroFacture,omitempty"`
	StatutFacture  string `json:"statutFacture,omitempty"`
}

func (s *FacturesService) TraiterFactureAValider(ctx context.Context, opts TraiterFactureAValiderOptions) (*TraiterFactureAValiderResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/traiter/factureAValider", opts)
	if err != nil {
		return nil, err
	}

	res := new(TraiterFactureAValiderResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
