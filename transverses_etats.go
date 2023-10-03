package choruspro

import (
	"context"
	"errors"
	"net/http"
)

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

	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/recuperer/etat/typedp", opts)
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

type ListeEtatsTraitement struct {
	CodeRetour int32            `json:"codeRetour"`
	Libelle    string           `json:"libelle"`
	Etats      []EtatTraitement `json:"listeStatutsPossiblesPourTraitement"`
}

type EtatTraitement struct {
	Etat string `json:"statutPossiblePourTraitement"`
}

type ListeEtatsTraitementOptions struct {
	EtatCourant string `json:"statutCourant"`
}

func (o *ListeEtatsTraitementOptions) Validate() error {
	if o.EtatCourant == "" {
		return errors.New("EtatCourant is required")
	}

	return nil
}

func (s *TransversesService) RecupererEtatsTraitement(ctx context.Context, opts ListeEtatsTraitementOptions) (*ListeEtatsTraitement, error) {
	err := opts.Validate()
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/recuperer/etats/traitement", opts)
	if err != nil {
		return nil, err
	}

	etats := new(ListeEtatsTraitement)

	err = s.client.doRequest(ctx, req, etats)
	if err != nil {
		return nil, err
	}

	return etats, nil
}
