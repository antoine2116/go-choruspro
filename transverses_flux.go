package choruspro

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type SyntaxeFlux string

const (
	SyntaxeFluxE1UblInvoice    SyntaxeFlux = "IN_DP_E1_UBL_INVOICE"
	SyntaxeFluxE1Cii16B        SyntaxeFlux = "IN_DP_E1_CII_16B"
	SyntaxeFluxE2UblInvoiceMin SyntaxeFlux = "IN_DP_E2_UBL_INVOICE_MIN"
	SyntaxeFluxE2CppFactureMin SyntaxeFlux = "IN_DP_E2_CPP_FACTURE_MIN"
	SyntaxeFluxE2CiiMin16B     SyntaxeFlux = "IN_DP_E2_CII_MIN_16B"
	SyntaxeFluxE2CiiFacturx    SyntaxeFlux = "IN_DP_E2_CII_FACTURX"
	SyntaxeFluxE3UblInvoice    SyntaxeFlux = "IN_DP_E3_UBL_INVOICE"
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

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/formatflux", opts)
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

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/consulterCR", opts)
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
