package choruspro

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

var (
	baseUrl = "/cpro/transverses/v1"
)

type TransversesService service

type MotifsExonerationTva struct {
	CodeRetour int32                 `json:"codeRetour"`
	Libelle    string                `json:"libelle"`
	Motifs     []MotifExonerationTva `json:"listeMotifsExonerationTva"`
}

type MotifExonerationTva struct {
	Code    string `json:"codeMotifExonerationTva"`
	Libelle string `json:"libelleMotifExonerationTva"`
}

type MotifsExonerationTvaOptions struct {
	CodeLangue string `json:"codeLangue,omitempty"`
}

func (s *TransversesService) RecupererMotifsExonerationTva(ctx context.Context, opts MotifsExonerationTvaOptions) (*MotifsExonerationTva, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, baseUrl+"/recuperer/motifs/exonerationtva", opts)
	if err != nil {
		return nil, err
	}

	motifs := new(MotifsExonerationTva)

	err = s.client.doRequest(ctx, req, motifs)
	if err != nil {
		return nil, err
	}

	return motifs, nil
}

type FomatsFlux struct {
	CodeRetour int32        `json:"codeRetour"`
	Libelle    string       `json:"libelle"`
	Formats    []FormatFlux `json:"listeFormatFlux"`
}

type FormatFlux struct {
	FormatFlux string `json:"formatFlux,omitempty"`
}

type FormatsFluxOptions struct {
}

func (s *TransversesService) RecupererFormatFlux(ctx context.Context, opts FormatsFluxOptions) (*FormatsFluxOptions, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, baseUrl+"/recuperer/formatflux", opts)
	if err != nil {
		return nil, err
	}

	formats := new(FormatsFluxOptions)

	err = s.client.doRequest(ctx, req, formats)
	if err != nil {
		return nil, err
	}

	return formats, nil
}

type CompteRendu struct {
	CodeAppliPartenaire      string     `json:"codeAppliPartenaire,omitempty"`
	CodeInterfaceFlux        string     `json:"codeInterfaceFlux,omitempty"`
	CodeRetour               int        `json:"codeRetour,omitempty"`
	DateDepotFlux            *time.Time `json:"dateDepotFlux,omitempty"`
	DateHeureEtatCourantFlux *time.Time `json:"dateHeureEtatCourantFlux,omitempty"`
	DesignationPartenaire    string     `json:"designationPartenaire,omitempty"`
	EtatCourantFlux          string     `json:"etatCourantFlux,omitempty"`
	FichierCR                string     `json:"fichierCR,omitempty"` // encoded in base64
	Libelle                  string     `json:"libelle,omitempty"`
	NomFichierFlux           string     `json:"nomFichierFlux,omitempty"`
	NumeroFluxDepot          string     `json:"numeroFluxDepot,omitempty"`
}

type CompteRenduOptions struct {
	DateDepot       *time.Time `json:"dateDepot,omitempty"`
	NumeroFluxDepot string     `json:"numeroFluxDepot"`
	SyntaxeFlux     string     `json:"syntaxeFlux,omitempty"`
}

func (opts CompteRenduOptions) Validate() error {
	if opts.NumeroFluxDepot == "" {
		return fmt.Errorf("choruspro: NumeroFluxDepot is required")
	}

	return nil
}

func (s *TransversesService) ConsulterCompteRendu(ctx context.Context, opts CompteRenduOptions) (*CompteRendu, error) {
	err := opts.Validate()
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest(ctx, http.MethodPost, baseUrl+"/consulterCR", opts)
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

type Devises struct {
	CodeRetour int32    `json:"codeRetour"`
	Libelle    string   `json:"libelle"`
	Devises    []Devise `json:"listeDevises"`
}

type Devise struct {
	Code    string `json:"codeDevise"`
	Libelle string `json:"libelleDevise"`
}

type DevisesOptions struct {
	CodeLangue string `json:"codeLangue,omitempty"`
}

func (s *TransversesService) RecupererDevises(ctx context.Context, opts DevisesOptions) (*Devises, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, baseUrl+"/recuperer/devises", opts)
	if err != nil {
		return nil, err
	}

	devise := new(Devises)

	err = s.client.doRequest(ctx, req, devise)
	if err != nil {
		return nil, err
	}

	return devise, nil
}

type AnnuaireDestinataire struct {
	CodeRetour int32  `json:"codeRetour"`
	Libelle    string `json:"libelle"`
	Fichier    string `json:"fichierResultat,omitempty"` // encoded in base64
}

type AnnuaireDestinataireOptions struct {
}

func (s *TransversesService) TelechargerAnnuaireDestinataire(ctx context.Context, opts AnnuaireDestinataireOptions) (*AnnuaireDestinataire, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, baseUrl+"/telecharger/annuaire/destinataire", opts)
	if err != nil {
		return nil, err
	}

	annuaire := new(AnnuaireDestinataire)

	err = s.client.doRequest(ctx, req, annuaire)
	if err != nil {
		return nil, err
	}

	return annuaire, nil
}

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

type TauxTvaOptions struct {
	CodeLangue string `json:"codeLangue,omitempty"`
}

func (s *TransversesService) RecupererTauxTva(ctx context.Context, opts TauxTvaOptions) (*ListeTauxTva, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, baseUrl+"/recuperer/tauxtva", opts)
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
