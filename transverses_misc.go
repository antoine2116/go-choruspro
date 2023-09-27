package choruspro

import (
	"context"
	"errors"
	"net/http"
)

type HealthCheck struct {
	Body            string `json:"body"`
	StatusCodeValue int    `json:"statusCodeValue"`
	StatusCode      string `json:"statusCode"`
}

func (s *TransversesService) HealthCheck(ctx context.Context) (*HealthCheck, error) {
	req, err := s.client.newRequest(ctx, http.MethodGet, "/cpro/transverses/v1/health-check", nil)
	if err != nil {
		return nil, err
	}

	hc := new(HealthCheck)

	err = s.client.doRequest(ctx, req, hc)
	if err != nil {
		return nil, err
	}

	return hc, nil
}

type HealthCheck struct {
	Body            string `json:"body"`
	StatusCodeValue int    `json:"statusCodeValue"`
	StatusCode      string `json:"statusCode"`
}

func (s *TransversesService) HealthCheck(ctx context.Context) (*HealthCheck, error) {
	req, err := s.client.newRequest(ctx, http.MethodGet, "/cpro/transverses/v1/health-check", nil)
	if err != nil {
		return nil, err
	}

	hc := new(HealthCheck)

	err = s.client.doRequest(ctx, req, hc)
	if err != nil {
		return nil, err
	}

	return hc, nil
}


// ListeDevises réprésente la réponse retournée par 
// TransversesService.RecupererDevises
type ListeDevises struct {
	CodeRetour int32    `json:"codeRetour"`
	Libelle    string   `json:"libelle"`
	Devises    []Devise `json:"listeDevises"`
}

// Devise représente un code devise
type Devise struct {
	Code    string `json:"codeDevise"`
	Libelle string `json:"libelleDevise"`
}

// RecupererDevises permet récupérer la liste des codes devises pouvant 
// être renseignés lors de la saisie dans une facture. lang est la langue dans 
// laquelle les devises seront retournées. Si lang n'est pas spécifié, la langue 
/// par défaut est le français (CodeLangueFr)
func (s *TransversesService) RecupererDevises(ctx context.Context, lang CodeLangue) (*ListeDevises, error) {
	opts := &struct {
		CodeLangue CodeLangue `json:"codeLangue,omitempty"`
	}{CodeLangue: lang}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/devises", opts)
	if err != nil {
		return nil, err
	}

	devise := new(ListeDevises)

	err = s.client.doRequest(ctx, req, devise)
	if err != nil {
		return nil, err
	}

	return devise, nil
}

// Annuaire réprésente la réponse retournée par 
// TransversesService.TelechargerAnnuaireDestinataire
// Le contenu de l'annuaire est encodé en base64
type AnnuaireDestinataire struct {
	CodeRetour int32  `json:"codeRetour"`
	Libelle    string `json:"libelle"`
	Fichier    string `json:"fichierResultat,omitempty"` // encoded in base64
}

// TelechargerAnnuaireDestinataire permet de télécharger l'annuaire des
// destinataires. Le fichier est retourné encodé en base64.
func (s *TransversesService) TelechargerAnnuaireDestinataire(ctx context.Context) (*AnnuaireDestinataire, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/telecharger/annuaire/destinataire", opts)
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

// ListeModesDepot réprésente la réponse retournée par
// TransversesService.RecupererModesDepot
type ListeModesDepot struct {
	CodeRetour int32       `json:"codeRetour"`
	Libelle    string      `json:"libelle"`
	ModesDepot []ModeDepot `json:"listeModeDepot"`
}

// ModeDepot représente un mode de dépôt ou de saisie d'une facture
type ModeDepot struct {
	ModeDepot string `json:"modeDepot"`
}

// RecupererModesDepot permet de récupérer la liste des modes de dépôt ou de
// saisie d'une facture
func (s *TransversesService) RecupererModesDepot(ctx context.Context) (*ListeModesDepot, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/modedepot", opts)
	if err != nil {
		return nil, err
	}

	modes := new(ListeModesDepot)

	err = s.client.doRequest(ctx, req, modes)
	if err != nil {
		return nil, err
	}

	return modes, nil
}

// ListePays réprésente la réponse retournée par TransversesService.RecupererPays
type ListePays struct {
	CodeRetour int32  `json:"codeRetour"`
	Libelle    string `json:"libelle"`
	Pays       []Pays `json:"listePays"`
}

// Pays représente un code pays (ISO 3166)
type Pays struct {
	Code    string `json:"codePays"`
	Libelle string `json:"libellePays"`
}

// RecupererPays permet de récupérer la liste des codes pays (liste ISO 3166)
// dans la langue spécifiée par lang. Si lang n'est pas spécifié, la langue par
// défaut est le français (CodeLangueFr)
func (s *TransversesService) RecupererPays(ctx context.Context, lang CodeLangue) (*ListePays, error) {
	opts := &struct {
		CodeLangue CodeLangue `json:"codeLangue,omitempty"`
	}{CodeLangue: lang}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/pays", opts)
	if err != nil {
		return nil, err
	}

	pays := new(ListePays)

	err = s.client.doRequest(ctx, req, pays)
	if err != nil {
		return nil, err
	}

	return pays, nil
}


// ListeMotifsRefusFactureAValider réprésente la réponse retournée par
// TransversesService.RecupererMotifsRefusFactureAValider
type ListeMotifsRefusFactureAValider struct {
	CodeRetour int32        `json:"codeRetour"`
	Libelle    string       `json:"libelle"`
	Motifs     []MotifRefus `json:"listeMotifRefusFactureAValider"`
}

// MotifRefus représente un motif de refus de facture
type MotifRefus struct {
	Id      int64  `json:"idMotif"`
	Code    string `json:"codeMotif"`
	Libelle string `json:"libelleMotif"`
}

// ListeMotifsRefusFactureAValiderOptions représente les paramètres de la requête
// TransversesService.RecupererMotifsRefusFactureAValider
type ListeMotifsRefusFactureAValiderOptions struct {
	CodeCadreFacturation string `json:"codeCadreFacturationAValider,omitempty"`
}

func (o *ListeMotifsRefusFactureAValiderOptions) Validate() error {
	if o.CodeCadreFacturation == "" {
		return errors.New("CodeCadreFacturation is required")
	}

	return nil
}

// RecupererMotifsRefusFactureAValider permet de récupérer la liste des motifs
func (s *TransversesService) RecupererMotifsRefusFactureAValider(ctx context.Context, opts ListeMotifsRefusFactureAValiderOptions) (*ListeMotifsRefusFactureAValider, error) {
	err := opts.Validate()
	if err != nil {
		return nil, err
	}
	
	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/motifs/refus/Facture/AValider", opts)
	if err != nil {
		return nil, err
	}

	motifs := new(ListeMotifsRefusFactureAValider)

	err = s.client.doRequest(ctx, req, motifs)
	if err != nil {
		return nil, err
	}

	return motifs, nil
}

// ListeModesReglement réprésente la réponse retournée par
// TransversesService.RecupererModesReglement
type ListeModesReglement struct {
	CodeRetour int32           `json:"codeRetour"`
	Libelle    string          `json:"libelle"`
	Modes      []ModeReglement `json:"listeModePaiement"`
}

// Mode Reglement représente un mode de réglement
type ModeReglement struct {
	ModeReglement string `json:"modePaiement"`
}

// RecupererModesReglement permet de récupérer la liste des modes de règlement
// pouvant être renseignées lors de la saisie d'une facture pour renseigner le
// mode de règlement de la demande de paiement déposée par le fournisseur.
func (s *TransversesService) RecupererModesReglement(ctx context.Context) (*ListeModesReglement, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/modereglements", opts)
	if err != nil {
		return nil, err
	}

	modes := new(ListeModesReglement)

	err = s.client.doRequest(ctx, req, modes)
	if err != nil {
		return nil, err
	}

	return modes, nil
}

// ListeCadresFacturation réprésente la réponse retournée par
// TransversesService.RecupererCadresFacturation
type ListeCadresFacturation struct {
	CodeRetour int32              `json:"codeRetour"`
	Libelle    string             `json:"libelle"`
	Cadres     []CadreFacturation `json:"listeCadreFacturation"`
}


// CadreFacturation représente un cadre de facturation
type CadreFacturation struct {
	Code string `json:"codeCadreFacturation"`
}

// RecupererCadresFacturation permet de récupérer la liste des cadres de
// facturation pouvant être renseignés lors de la saisie d'une facture.
// typeDP est le type de demande de paiement. La liste des types de demande peut
// être récupérée avec TransversesService.RecupererTypesDemandePaiement
func (s *TransversesService) RecupererCadresFacturation(ctx context.Context, typeDP string) (*ListeCadresFacturation, error) {
	opts := struct {
		TypeDemande string `json:"typeDemandePaiement,omitempty"`
	}{TypeDemande: typeDP}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/cadrefacturation", opts)
	if err != nil {
		return nil, err
	}

	cadres := new(ListeCadresFacturation)

	err = s.client.doRequest(ctx, req, cadres)
	if err != nil {
		return nil, err
	}

	return cadres, nil
}

// ListeCoordonnesBancaires réprésente la réponse retournée par
// TransversesService.RecupererCoordonneesBancairesValides
type ListeCoordonneesBancaires struct {
	CodeRetour  int32                  `json:"codeRetour"`
	Libelle     string                 `json:"libelle"`
	Coordonnees []CoordonneeBancaire `json:"listeCoordonneeBancaire"`
}

// CoordonneeBancaire représente une coordonnée bancaire valide
type CoordonneeBancaire struct {
	Id  int64  `json:"idCoordonneeBancaire"`
	Nom string `json:"nomCoordonneeBancaire"`
}

// ListeCoordonneesBancairesOptions représente les paramètres de la requête
// TransversesService.RecupererCoordonneesBancairesValides
type ListeCoordonneesBancairesOptions struct {
	IdStructure int64 `json:"idStructure,omitempty"`
}

func (o ListeCoordonneesBancairesOptions) Validate() error {
	if o.IdStructure == 0 {
		return errors.New("IdStructure is required")
	}

	return nil
}

// RecupererCoordonneesBancairesValides permet de récupérer la liste des
// coordonnées bancaires renseignées pour une structure à laquelle l'utilisateur
// courant est rattaché.
func (s *TransversesService) RecupererCoordonneesBancairesValides(ctx context.Context, opts ListeCoordonneesBancairesOptions) (*ListeCoordonneesBancaires, error) {
	err := opts.Validate()
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/coordbanc/valides", opts)
	if err != nil {
		return nil, err
	}

	coord := new(ListeCoordonneesBancaires)

	err = s.client.doRequest(ctx, req, coord)
	if err != nil {
		return nil, err
	}

	return coord, nil
}
