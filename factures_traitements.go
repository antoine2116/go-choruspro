package choruspro

import (
	"context"
	"net/http"
)

// TraiterFactureAValiderOptions est la structure de données utilisée pour appeler
// la méthode TraiterFactureAValider.
type TraiterFactureAValiderOptions struct {
	Action         ActionFacture `json:"action,omitempty"`
	ComplementInfo string        `json:"complementInfo,omitempty"`
	IdFacture      int64         `json:"idFacture,omitempty"`
	IdPieceJointe  int64         `json:"idPieceJointe,omitempty"`
	IdValideur     int64         `json:"idValideur,omitempty"`
	MotifRefus     string        `json:"motifRefus,omitempty"`
	TypeValideur   string        `json:"typeValideur,omitempty"`
}

// TraiterFactureAValiderResponse est la réponse renvoyée par la méthode
// TraiterFactureAValider.
type TraiterFactureAValiderResponse struct {
	CodeRetour     int    `json:"codeRetour,omitempty"`
	DateTraitement *Date  `json:"dateTraitement,omitempty"`
	IdFacture      int64  `json:"idFacture,omitempty"`
	Libelle        string `json:"libelle,omitempty"`
	NumeroFacture  string `json:"numeroFacture,omitempty"`
	StatutFacture  string `json:"statutFacture,omitempty"`
}

// La méthode traiterFactureAValider permet à un valideur de modifier le
// statut d'une facture à valider en renseignant le cas échéant un motif de refus.
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

// TraiterFactureRecueOptions est la structure de données utilisée pour appeler
// la méthode TraiterFactureRecue.
type TraiterFactureRecueOptions struct {
	IdFacture            int64  `json:"idFacture"`
	IdUtilisateurCourant int64  `json:"idUtilisateurCourant"`
	Motif                string `json:"motif"`
	NouveauStatut        string `json:"nouveauStatut"`
	NumeroDPMandat       string `json:"numeroDPMandat"`
}

// TraiterFactureRecueResponse est la réponse renvoyée par la méthode
// TraiterFactureRecue.
type TraiterFactureRecueResponse struct {
	CodeRetour     int32  `json:"codeRetour"`
	DateTraitement *Date  `json:"dateTraitement"`
	IdFacture      int64  `json:"idFacture"`
	Libelle        string `json:"libelle"`
	NouveauStatut  string `json:"nouveauStatut"`
	NumeroFacture  string `json:"numeroFacture"`
}

// La méthode traiterFactureRecue permet de valider, rejeter ou suspendre
// une facture reçue. Lorsque le récipiendaire rejette ou suspend la facture,
// il est dans l'obligation de motiver son choix.
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

// DemandePaiementParam est la structure de données pour
// renseigner l'identifiant d'une demande de paiement.
// Cette structure est utilisée par la méthode TraiterRejet.
type DemandePaiementParam struct {
	IdDemandePaiement int64 `json:"idDemandePaiement"`
}

// TraiterRejetOptions est la structure de données utilisée pour appeler
// la méthode TraiterRejet.
type TraiterRejetOptions struct {
	CpltInfos            string                 `json:"cpltInfos"`
	ListeDemandePaiement []DemandePaiementParam `json:"listeDemandePaiement"`
	MotifRejet           string                 `json:"motifRejet"`
}

// TraiterRejetResponse est la réponse renvoyée par la méthode
// TraiterRejet.
type TraiterRejetResponse struct {
	CodeRetour int32  `json:"codeRetour"`
	Libelle    string `json:"libelle"`
}

// La méthode traiterFacturesRejetees a pour objectif d'indiquer, pour les factures
// ou factures de travaux données, que leur rejet a été traité par le fournisseur
func (s *FacturesService) TraiterFacturesRejetees(ctx context.Context, opts TraiterRejetOptions) (*TraiterRejetResponse, error) {
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
