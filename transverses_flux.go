package choruspro

import (
	"context"
	"net/http"
	"time"
)

// ListeFormatsFluxResponse est la structure de données représentant
// la réponse de la méthode RecupererFormatFlux.
type ListeFormatsFluxResponse struct {
	CodeRetour int32        `json:"codeRetour"`
	Libelle    string       `json:"libelle"`
	Formats    []FormatFlux `json:"listeFormatFlux"`
}

// FormatFlux est la structure de données représentant
// un format de flux.
type FormatFlux struct {
	FormatFlux string `json:"formatFlux,omitempty"`
}

// La méthode RecupererFormatFlux permet de récupérer les valeurs
// pouvant être renseignées lors du dépôt d'un flux pour renseigner le format du flux.
func (s *TransversesService) RecupererFormatFlux(ctx context.Context) (*ListeFormatsFluxResponse, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/recuperer/formatflux", opts)
	if err != nil {
		return nil, err
	}

	formats := new(ListeFormatsFluxResponse)

	err = s.client.doRequest(ctx, req, formats)
	if err != nil {
		return nil, err
	}

	return formats, nil
}

// ConsulterCompteRenduResponse est la structure de données représentant
// la réponse de la méthode ConsulterCompteRendu.
type ConsulterCompteRenduResponse struct {
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

// ConsulterCompteRenduOptions est la structure de données représentant
// les options de la méthode ConsulterCompteRendu.
type ConsulterCompteRenduOptions struct {
	DateDepot       *time.Time  `json:"dateDepot,omitempty"`
	NumeroFluxDepot string      `json:"numeroFluxDepot"`
	SyntaxeFlux     SyntaxeFlux `json:"syntaxeFlux,omitempty"`
}

// Le service ConsulterCompteRendu permet de consulter les informations
// liées au dépôt d'un flux et de récupérer au format PDF le compte rendu
// de traitement du flux déposé via le portail ou le service exposé DeposerFluxFacture.
func (s *TransversesService) ConsulterCompteRendu(ctx context.Context, opts ConsulterCompteRenduOptions) (*ConsulterCompteRenduResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/consulterCR", opts)
	if err != nil {
		return nil, err
	}

	formats := new(ConsulterCompteRenduResponse)

	err = s.client.doRequest(ctx, req, formats)
	if err != nil {
		return nil, err
	}

	return formats, nil
}

// ErreurDemandePaiement est la structure de données représentant
// une erreur de demande de paiement. Elle est utilisée dans la
// structure CompteRenduDetaille.
type ErreurDemandePaiement struct {
	IdentifiantDestinataire string `json:"identifiantDestinataire,omitempty"`
	IdentifiantFournisseur  string `json:"identifiantFournisseur,omitempty"`
	LibelleErreurDP         string `json:"libelleErreurDP,omitempty"`
	NumeroDP                string `json:"numeroDP,omitempty"`
}

// ErreurTechnique est la structure de données représentant
// une erreur technique. Elle est utilisée dans la
// structure CompteRenduDetaille.
type ErreurTechnique struct {
	CodeErreur    string `json:"codeErreur,omitempty"`
	LibelleErreur string `json:"libelleErreur,omitempty"`
	NatureErreur  string `json:"natureErreur,omitempty"`
}

// CompteRenduDetailleResponse est la structure de données représentant
// le compte rendu détaillé d'un dépôt de flux.
type CompteRenduDetailleResponse struct {
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

// ConsulterCompteRenduDetailleOptions est la structure de données représentant
// les options de la méthode ConsulterCompteRenduDetaille.
type ConsulterCompteRenduDetailleOptions struct {
	NumeroFluxDepot string      `json:"numeroFluxDepot"`
	SyntaxeFlux     SyntaxeFlux `json:"syntaxeFlux,omitempty"`
}

// Le service ConsulterCompteRenduDetaille permet de consulter l'état d'intégration d'un
// flux émis en API, avec le cas échéant les erreurs identifiées par le système
// pour l'irrecevabilité du flux ou le rejet d'une ou plusieurs demandes de paiement.
func (s *TransversesService) ConsulterCompteRenduDetaille(ctx context.Context, opts ConsulterCompteRenduDetailleOptions) (*CompteRenduDetailleResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/consulterCRDetaille", opts)
	if err != nil {
		return nil, err
	}

	formats := new(CompteRenduDetailleResponse)

	err = s.client.doRequest(ctx, req, formats)
	if err != nil {
		return nil, err
	}

	return formats, nil
}
