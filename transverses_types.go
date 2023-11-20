package choruspro

import (
	"context"
	"net/http"
)

// ListeTypesDemandePaiementResponse est la structure de données représentant
// la réponse de la méthode RecupererTypesDemandePaiement.
type ListeTypesDemandePaiementResponse struct {
	CodeRetour int32                 `json:"codeRetour"`
	Libelle    string                `json:"libelle"`
	Types      []TypeDemandePaiement `json:"listeTypeDemandePaiement"`
}

// TypeDemandePaiement est la structure de données représentant
// un type de demande de paiement.
type TypeDemandePaiement struct {
	Type string `json:"typeDemandePaiement"`
}

// La méthode RecupererTypesDemandePaiement permet de récupérer les types de
// demandes de paiement pouvant être renseignées dans Chorus Pro.
func (s *TransversesService) RecupererTypesDemandePaiement(ctx context.Context) (*ListeTypesDemandePaiementResponse, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/recuperer/typedp", opts)
	if err != nil {
		return nil, err
	}

	res := new(ListeTypesDemandePaiementResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ListeTypesFactureTravauxResponse est la structure de données représentant
// la réponse de la méthode RecupererTypesFactureTravaux.
type ListeTypesFactureTravauxResponse struct {
	CodeRetour int32                `json:"codeRetour"`
	Libelle    string               `json:"libelle"`
	Types      []TypeFactureTravaux `json:"listeTypeFactureTravaux"`
}

// TypeFactureTravaux est la structure de données représentant
// un type de facture de travaux.
type TypeFactureTravaux struct {
	Type string `json:"typeFactureTravaux"`
}

// La méthode RecupererTypesFactureTravaux permet de renseigner les éléments
// de facturations pouvant être renseignés lors du dépôt d'un dossier
// de facturation de travaux.
func (s *TransversesService) RecupererTypesFactureTravaux(ctx context.Context) (*ListeTypesFactureTravauxResponse, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/recuperer/typefacturetravaux", opts)
	if err != nil {
		return nil, err
	}

	res := new(ListeTypesFactureTravauxResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ListeTypesIdentifiantsStructureResponse est la structure de données représentant
// la réponse de la méthode RecupererTypesIdentifiantsStructure.
type ListeTypesIdentifiantsStructureResponse struct {
	CodeRetour int32                      `json:"codeRetour"`
	Libelle    string                     `json:"libelle"`
	Types      []TypeIdentifiantStructure `json:"listeTypeIdentifiant"`
}

// TypeIdentifiantStructure est la structure de données représentant
// un type d'identifiant de structure.
type TypeIdentifiantStructure struct {
	Type string `json:"typeIdentifiantStructure"`
}

// La méthode recupererTypeIdentifiantStructure permet de récupérer les types
// d'identifiants pouvant être renseignés par un fournisseur ou un valideur de facture.
func (s *TransversesService) RecupererTypesIdentifiantsStructure(ctx context.Context) (*ListeTypesIdentifiantsStructureResponse, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/recuperer/typeidentifiant/structure", opts)
	if err != nil {
		return nil, err
	}

	res := new(ListeTypesIdentifiantsStructureResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
