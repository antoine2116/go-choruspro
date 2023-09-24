package choruspro

import (
	"context"
	"errors"
	"net/http"
)

type ListeDevises struct {
	CodeRetour int32    `json:"codeRetour"`
	Libelle    string   `json:"libelle"`
	Devises    []Devise `json:"listeDevises"`
}

type Devise struct {
	Code    string `json:"codeDevise"`
	Libelle string `json:"libelleDevise"`
}

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

type AnnuaireDestinataire struct {
	CodeRetour int32  `json:"codeRetour"`
	Libelle    string `json:"libelle"`
	Fichier    string `json:"fichierResultat,omitempty"` // encoded in base64
}

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

type ListeModesDepot struct {
	CodeRetour int32       `json:"codeRetour"`
	Libelle    string      `json:"libelle"`
	ModesDepot []ModeDepot `json:"listeModeDepot"`
}

type ModeDepot struct {
	ModeDepot string `json:"modeDepot"`
}

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

type ListePays struct {
	CodeRetour int32  `json:"codeRetour"`
	Libelle    string `json:"libelle"`
	Pays       []Pays `json:"listePays"`
}

type Pays struct {
	Code    string `json:"codePays"`
	Libelle string `json:"libellePays"`
}

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

type ListeMotifsRefusFactureAValider struct {
	CodeRetour int32        `json:"codeRetour"`
	Libelle    string       `json:"libelle"`
	Motifs     []MotifRefus `json:"listeMotifRefusFactureAValider"`
}

type MotifRefus struct {
	Id      int64  `json:"idMotif"`
	Code    string `json:"codeMotif"`
	Libelle string `json:"libelleMotif"`
}

type ListeMotifsRefusFactureAValiderOptions struct {
	CodeCadreFacturation string `json:"codeCadreFacturationAValider,omitempty"`
}

func (o *ListeMotifsRefusFactureAValiderOptions) Validate() error {
	if o.CodeCadreFacturation == "" {
		return errors.New("CodeCadreFacturation is required")
	}

	return nil
}

func (s *TransversesService) RecupererMotifsRefusFactureAValider(ctx context.Context, opts ListeMotifsRefusFactureAValiderOptions) (*ListeMotifsRefusFactureAValider, error) {
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

type ListeModesReglement struct {
	CodeRetour int32           `json:"codeRetour"`
	Libelle    string          `json:"libelle"`
	Modes      []ModeReglement `json:"listeModePaiement"`
}

type ModeReglement struct {
	ModeReglement string `json:"modePaiement"`
}

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

type ListeCadresFacturation struct {
	CodeRetour int32              `json:"codeRetour"`
	Libelle    string             `json:"libelle"`
	Cadres     []CadreFacturation `json:"listeCadreFacturation"`
}

type CadreFacturation struct {
	Code string `json:"codeCadreFacturation"`
}

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

type ListeCoordonneesBancaires struct {
	CodeRetour  int32                  `json:"codeRetour"`
	Libelle     string                 `json:"libelle"`
	Coordonnees []CoordonneesBancaires `json:"listeCoordonneeBancaire"`
}

type CoordonneesBancaires struct {
	Id  int64  `json:"idCoordonneeBancaire"`
	Nom string `json:"nomCoordonneeBancaire"`
}

type ListeCoordonneesBancairesOptions struct {
	IdStructure int64 `json:"idStructure,omitempty"`
}

func (o ListeCoordonneesBancairesOptions) Validate() error {
	if o.IdStructure == 0 {
		return errors.New("IdStructure is required")
	}

	return nil
}

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
