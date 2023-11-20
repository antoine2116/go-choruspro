package choruspro

import (
	"context"
	"net/http"
)

// ListeTauxTvaResponse est la structure de données représentant
// la réponse de la méthode RecupererTauxTva.
type ListeTauxTvaResponse struct {
	CodeRetour int32     `json:"codeRetour"`
	Libelle    string    `json:"libelle"`
	Taux       []TauxTva `json:"listeTauxTva"`
}

// TauxTva est la structure de données représentant un taux de TVA.
type TauxTva struct {
	Code    string  `json:"codeTauxTva"`
	Libelle string  `json:"libelleTauxTva"`
	Valeur  float32 `json:"valeurTauxTva"`
}

// La méthode recupererTauxTva permet de récupérer la liste des taux
// de TVA applicables.
func (s *TransversesService) RecupererTauxTva(ctx context.Context, lang CodeLangue) (*ListeTauxTvaResponse, error) {
	opts := &codeLangueOptions{lang}

	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/recuperer/tauxtva", opts)
	if err != nil {
		return nil, err
	}

	res := new(ListeTauxTvaResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ListeMotifsExonerationTvaResponse est la structure de données représentant
// la réponse de la méthode RecupererMotifsExonerationTva.
type ListeMotifsExonerationTvaResponse struct {
	CodeRetour int32                 `json:"codeRetour"`
	Libelle    string                `json:"libelle"`
	Motifs     []MotifExonerationTva `json:"listeMotifsExonerationTva"`
}

// MotifExonerationTva est la structure de données représentant un motif
// d'exonération de TVA.
type MotifExonerationTva struct {
	Code    string `json:"codeMotifExonerationTva"`
	Libelle string `json:"libelleMotifExonerationTva"`
}

// La méthode RecupererMotifsExonerationTva permet de récupérer la liste des
// motifs d'exonération à la TVA pouvant être renseignés lors du dépôt d'une facture.
func (s *TransversesService) RecupererMotifsExonerationTva(ctx context.Context, lang CodeLangue) (*ListeMotifsExonerationTvaResponse, error) {
	opts := &codeLangueOptions{lang}

	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/recuperer/motifs/exonerationtva", opts)
	if err != nil {
		return nil, err
	}

	res := new(ListeMotifsExonerationTvaResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
