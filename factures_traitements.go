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

type TraiterFactureRecueOptions struct {
	IdFacture            int64  `json:"idFacture"`
	IdUtilisateurCourant int64  `json:"idUtilisateurCourant"`
	Motif                string `json:"motif"`
	NouveauStatut        string `json:"nouveauStatut"`
	NumeroDPMandat       string `json:"numeroDPMandat"`
}

type TraiterFactureRecueResponse struct {
	CodeRetour     int32  `json:"codeRetour"`
	DateTraitement *Date  `json:"dateTraitement"`
	IdFacture      int64  `json:"idFacture"`
	Libelle        string `json:"libelle"`
	NouveauStatut  string `json:"nouveauStatut"`
	NumeroFacture  string `json:"numeroFacture"`
}

func (s *FacturesService) TraiterFactureRecue(ctx context.Context, opts TraiterFactureRecueOptions) (*TraiterFactureRecueResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/traiter/recue", opts)
	if err != nil {
		return nil, err
	}

	res := new(TraiterFactureRecueResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type DemandePaiementParam struct {
	IdDemandePaiement int64 `json:"idDemandePaiement"`
}

type TraiterRejetOptions struct {
	CpltInfos            string                 `json:"cpltInfos"`
	ListeDemandePaiement []DemandePaiementParam `json:"listeDemandePaiement"`
	MotifRejet           string                 `json:"motifRejet"`
}

type TraiterRejetResponse struct {
	CodeRetour int32  `json:"codeRetour"`
	Libelle    string `json:"libelle"`
}

func (s *FacturesService) TraiterRejet(ctx context.Context, opts TraiterRejetOptions) (*TraiterRejetResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/traiterRejet", opts)
	if err != nil {
		return nil, err
	}

	res := new(TraiterRejetResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
