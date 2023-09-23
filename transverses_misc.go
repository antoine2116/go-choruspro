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

type ListeEtatsTypeDemandePaiement struct {
	CodeRetour int32                     `json:"codeRetour"`
	Libelle    string                    `json:"libelle"`
	Etats      []EtatTypeDemandePaiement `json:"listeEtatDemandePaiement"`
}

type EtatTypeDemandePaiement struct {
	Etat string `json:"etatDemandePaiement"`
}

type ListeEtatsTypeDemandePaiementOptions struct {
	TypeDemandePaiement string `json:"typeDemandePaiement,omitempty"`
}

func (o *ListeEtatsTypeDemandePaiementOptions) Validate() error {
	if o.TypeDemandePaiement == "" {
		return errors.New("TypeDemandePaiement is required")
	}

	return nil
}

func (s *TransversesService) RecupererEtatParTypeDemandePaiement(ctx context.Context, opts ListeEtatsTypeDemandePaiementOptions) (*ListeEtatsTypeDemandePaiement, error) {
	err := opts.Validate()
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/etat/typedp", opts)
	if err != nil {
		return nil, err
	}

	etats := new(ListeEtatsTypeDemandePaiement)

	err = s.client.doRequest(ctx, req, etats)
	if err != nil {
		return nil, err
	}

	return etats, nil
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
