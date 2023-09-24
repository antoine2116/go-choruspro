package choruspro

import (
	"context"
	"net/http"
)

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

func (s *TransversesService) RecupererTauxTva(ctx context.Context, lang CodeLangue) (*ListeTauxTva, error) {
	opts := &struct {
		CodeLangue CodeLangue `json:"codeLangue,omitempty"`
	}{CodeLangue: lang}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/tauxtva", opts)
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

type ListeMotifsExonerationTva struct {
	CodeRetour int32                 `json:"codeRetour"`
	Libelle    string                `json:"libelle"`
	Motifs     []MotifExonerationTva `json:"listeMotifsExonerationTva"`
}

type MotifExonerationTva struct {
	Code    string `json:"codeMotifExonerationTva"`
	Libelle string `json:"libelleMotifExonerationTva"`
}

func (s *TransversesService) RecupererMotifsExonerationTva(ctx context.Context, lang CodeLangue) (*ListeMotifsExonerationTva, error) {
	opts := &struct {
		CodeLangue CodeLangue `json:"codeLangue,omitempty"`
	}{CodeLangue: lang}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/motifs/exonerationtva", opts)
	if err != nil {
		return nil, err
	}

	motifs := new(ListeMotifsExonerationTva)

	err = s.client.doRequest(ctx, req, motifs)
	if err != nil {
		return nil, err
	}

	return motifs, nil
}
