package choruspro

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type ListeTauxTva struct {
	CodeRetour int32     `json:"codeRetour"`
	Libelle    string    `json:"libelle"`
	TauxTva    []TauxTva `json:"listeTauxTva"`
}

type TauxTva struct {
	Code    string  `json:"codeTauxTva"`
	Libelle string  `json:"libelleTauxTva"`
	Valeur  float32 `json:"valeurTauxTva"`
}

func (s *TransversesService) RecupererTauxTva(ctx context.Context, lang CodeLangue) (*ListeTauxTva, error) {
	opts := &struct {
		CodeLangue CodeLangue `json:"codeLangue,omitempty"`
	}{CodeLangue: lang}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/tauxtva", opts)
	if err != nil {
		return nil, err
	}

	taux := new(ListeTauxTva)

	err = s.client.doRequest(ctx, req, taux)
	if err != nil {
		return nil, err
	}

	return taux, nil
}

type ListeMotifsExonerationTva struct {
	CodeRetour int32                 `json:"codeRetour"`
	Libelle    string                `json:"libelle"`
	Motifs     []MotifExonerationTva `json:"listeMotifsExonerationTva"`
}

type MotifExonerationTva struct {
	Code    string `json:"codeMotifExonerationTva"`
	Libelle string `json:"libelleMotifExonerationTva"`
}

func (s *TransversesService) RecupererMotifsExonerationTva(ctx context.Context, lang CodeLangue) (*ListeMotifsExonerationTva, error) {
	opts := &struct {
		CodeLangue CodeLangue `json:"codeLangue,omitempty"`
	}{CodeLangue: lang}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/motifs/exonerationtva", opts)
	if err != nil {
		return nil, err
	}

	motifs := new(ListeMotifsExonerationTva)

	err = s.client.doRequest(ctx, req, motifs)
	if err != nil {
		return nil, err
	}

	return motifs, nil
}

type CompteRenduDetaille struct {
	CodeInterfaceDepotFlux   string                  `json:"codeInterfaceDepotFlux,omitempty"`
	CodeRetour               int32                   `json:"codeRetour"`
	DateDepotFlux            *time.Time              `json:"dateDepotFlux,omitempty"`
	DateHeureEtatCourantFlux *time.Time              `json:"dateHeureEtatCourantFlux,omitempty"`
	EtatCourantDepotFlux     string                  `json:"etatCourantDepotFlux,omitempty"`
	Libelle                  string                  `json:"libelle"`
	NomFichier               string                  `json:"nomFichier,omitempty"`
	ErreursDemandePaiement   []ErreurDemandePaiement `json:"listeErreurDP,omitempty"`
	ErreursTechniques        []ErreurTechnique       `json:"listeErreurTechnique,omitempty"`
}

type ErreurDemandePaiement struct {
	IdentifiantDestinataire string `json:"identifiantDestinataire,omitempty"`
	IdentifiantFournisseur  string `json:"identifiantFournisseur,omitempty"`
	LibelleErreurDP         string `json:"libelleErreurDP,omitempty"`
	NumeroDP                string `json:"numeroDP,omitempty"`
}

type ErreurTechnique struct {
	CodeErreur    string `json:"codeErreurn,omitempty"`
	LibelleErreur string `json:"libelleErreur,omitempty"`
	NatureErreur  string `json:"natureErreur,omitempty"`
}

type CompteRenduDetailleOptions struct {
	NumeroFluxDepot string      `json:"numeroFluxDepot"`
	SyntaxeFlux     SyntaxeFlux `json:"syntaxeFlux,omitempty"`
}

func (o CompteRenduDetailleOptions) Validate() error {
	if o.NumeroFluxDepot == "" {
		return fmt.Errorf("choruspro: NumeroFluxDepot is required")
	}

	return nil
}

func (s *TransversesService) ConsulterCompteRenduDetaille(ctx context.Context, opts CompteRenduDetailleOptions) (*CompteRenduDetaille, error) {
	err := opts.Validate()
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/consulterCRDetaille", opts)
	if err != nil {
		return nil, err
	}

	formats := new(CompteRenduDetaille)

	err = s.client.doRequest(ctx, req, formats)
	if err != nil {
		return nil, err
	}

	return formats, nil
}
