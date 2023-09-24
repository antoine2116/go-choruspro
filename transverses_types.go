package choruspro

import (
	"context"
	"net/http"
)

type ListeTypesDemandePaiement struct {
	CodeRetour int32                 `json:"codeRetour"`
	Libelle    string                `json:"libelle"`
	Types      []TypeDemandePaiement `json:"listeTypeDemandePaiement"`
}

type TypeDemandePaiement struct {
	Type string `json:"typeDemandePaiement"`
}

func (s *TransversesService) RecupererTypesDemandePaiement(ctx context.Context) (*ListeTypesDemandePaiement, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/typedp", opts)
	if err != nil {
		return nil, err
	}

	types := new(ListeTypesDemandePaiement)

	err = s.client.doRequest(ctx, req, types)
	if err != nil {
		return nil, err
	}

	return types, nil
}

type ListeTypesFactureTravaux struct {
	CodeRetour          int32                `json:"codeRetour"`
	Libelle             string               `json:"libelle"`
	TypesFactureTravaux []TypeFactureTravaux `json:"listeTypeFactureTravaux"`
}

type TypeFactureTravaux struct {
	TypeFactureTravaux string `json:"typeFactureTravaux"`
}

func (s *TransversesService) RecupererTypesFactureTravaux(ctx context.Context) (*ListeTypesFactureTravaux, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/typefacturetravaux", opts)
	if err != nil {
		return nil, err
	}

	types := new(ListeTypesFactureTravaux)

	err = s.client.doRequest(ctx, req, types)
	if err != nil {
		return nil, err
	}

	return types, nil
}

type ListeTypesIdentifiantsStructure struct {
	CodeRetour int32                      `json:"codeRetour"`
	Libelle    string                     `json:"libelle"`
	Types      []TypeIdentifiantStructure `json:"listeTypeIdentifiant"`
}

type TypeIdentifiantStructure struct {
	Type string `json:"typeIdentifiantStructure"`
}

func (s *TransversesService) RecupererTypesIdentifiantsStructure(ctx context.Context) (*ListeTypesIdentifiantsStructure, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/typeidentifiant/structure", opts)
	if err != nil {
		return nil, err
	}

	types := new(ListeTypesIdentifiantsStructure)

	err = s.client.doRequest(ctx, req, types)
	if err != nil {
		return nil, err
	}

	return types, nil
}
