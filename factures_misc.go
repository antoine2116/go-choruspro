package choruspro

import (
	"context"
	"net/http"
)

type TelechargerGroupeFactureOptions struct {
	// 2 values allowed: "OUI" or "NON"
	AvecPiecesJointesComplementaires string        `json:"avecPiecesJointesComplementaires,omitempty"`
	Format                           FormatFichier `json:"format,omitempty"`
	IdEspace                         int64         `json:"idEspace,omitempty"`
	IdUtilisateurCourant             int64         `json:"idUtilisateurCourant,omitempty"`
	ListeFactures                    []*IdFacture  `json:"listeFacture,omitempty"`
}

type TelechargerGroupeFactureResponse struct {
	CodeRetour      int    `json:"codeRetour,omitempty"`
	FichierResultat string `json:"fichierResultat,omitempty"` // base64 encoded
	Libelle         string `json:"libelle,omitempty"`
}

func (s *FacturesService) TelechargerGroupe(ctx context.Context, opts TraiterFactureAValiderOptions) (*TelechargerGroupeFactureResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/telecharger/groupe", opts)
	if err != nil {
		return nil, err
	}

	res := new(TelechargerGroupeFactureResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
