package choruspro

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type ListeFormatsFlux struct {
	CodeRetour int32        `json:"codeRetour"`
	Libelle    string       `json:"libelle"`
	Formats    []FormatFlux `json:"listeFormatFlux"`
}

type FormatFlux struct {
	FormatFlux string `json:"formatFlux,omitempty"`
}

func (s *TransversesService) RecupererFormatFlux(ctx context.Context) (*ListeFormatsFlux, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/recuperer/formatflux", opts)
	if err != nil {
		return nil, err
	}

	formats := new(ListeFormatsFlux)

	err = s.client.doRequest(ctx, req, formats)
	if err != nil {
		return nil, err
	}

	return formats, nil
}

type CompteRendu struct {
	CodeAppliPartenaire      string     `json:"codeAppliPartenaire,omitempty"`
	CodeInterfaceFlux        string     `json:"codeInterfaceFlux,omitempty"`
	CodeRetour               int        `json:"codeRetour"`
	DateDepotFlux            *time.Time `json:"dateDepotFlux,omitempty"`
	DateHeureEtatCourantFlux *time.Time `json:"dateHeureEtatCourantFlux,omitempty"`
	DesignationPartenaire    string     `json:"designationPartenaire,omitempty"`
	EtatCourantFlux          string     `json:"etatCourantFlux,omitempty"`
	FichierCR                string     `json:"fichierCR,omitempty"` // encoded in base64
	Libelle                  string     `json:"libelle"`
	NomFichierFlux           string     `json:"nomFichierFlux,omitempty"`
	NumeroFluxDepot          string     `json:"numeroFluxDepot,omitempty"`
}

type CompteRenduOptions struct {
	DateDepot       *time.Time  `json:"dateDepot,omitempty"`
	NumeroFluxDepot string      `json:"numeroFluxDepot"`
	SyntaxeFlux     SyntaxeFlux `json:"syntaxeFlux,omitempty"`
}

func (o CompteRenduOptions) Validate() error {
	if o.NumeroFluxDepot == "" {
		return errors.New("choruspro: NumeroFluxDepot is required")
	}

	return nil
}

func (s *TransversesService) ConsulterCompteRendu(ctx context.Context, opts CompteRenduOptions) (*CompteRendu, error) {
	err := opts.Validate()
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/consulterCR", opts)
	if err != nil {
		return nil, err
	}

	formats := new(CompteRendu)

	err = s.client.doRequest(ctx, req, formats)
	if err != nil {
		return nil, err
	}

	return formats, nil
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

	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/consulterCRDetaille", opts)
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
