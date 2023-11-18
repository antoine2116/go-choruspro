package choruspro

import (
	"context"
	"net/http"
)

// TelechargerGroupeFactureOptions est la structure de données utlisée
// pour la méthode TelechargerGroupeFacture.
type TelechargerGroupeFactureOptions struct {
	// 2 values allowed: "OUI" or "NON"
	AvecPiecesJointesComplementaires string        `json:"avecPiecesJointesComplementaires,omitempty"`
	Format                           FormatFichier `json:"format,omitempty"`
	IdEspace                         int64         `json:"idEspace,omitempty"`
	IdUtilisateurCourant             int64         `json:"idUtilisateurCourant,omitempty"`
	ListeFactures                    []*IdFacture  `json:"listeFacture,omitempty"`
}

// TelechargerGroupeFactureResponse est la structure de données réprésentant
// la réponse de la méthode TelechargerGroupeFacture.
type TelechargerGroupeFactureResponse struct {
	CodeRetour      int    `json:"codeRetour,omitempty"`
	FichierResultat string `json:"fichierResultat,omitempty"` // base64 encoded
	Libelle         string `json:"libelle,omitempty"`
}

// La méthode TelechargerGroupeFacture permet de télécharger une ou plusieurs
// factures émises, reçues ou à valider en précisant le format de réception
// (XML, PDF ou ZIP) et les pièces jointes associées à ces factures. Le dossier
// téléchargé ne doit pas dépasser la taille maximum d'un dossier de facturation
// (120 Mo), quelque soit le nombre de factures téléchargées. Si le dossier
// récupéré dépasse cette taille, une erreur 413 sera remontée.
func (s *FacturesService) TelechargerGroupeFacture(ctx context.Context, opts TraiterFactureAValiderOptions) (*TelechargerGroupeFactureResponse, error) {
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

// RecupererStatutsVisiblesParValideurOptions est la structure de données utlisée
// pour la méthode RecupererStatutsVisiblesParValideur.
type RecupererStatutsVisiblesParValideurOptions struct {
	IdUtilisateurCourant int64 `json:"idUtilisateurCourant"`
}

// StatutVisible est la structure de données réprésentant un statut visible
// par un valideur. Elle est utilisée dans la réponse de la méthode
// RecupererStatutsVisiblesParValideur.
type StatutVisible struct {
	CodeStatut string `json:"codeStatutFacture"`
}

// RecupererStatutsVisiblesParValideurResponse est la structure de données réprésentant
// la réponse de la méthode RecupererStatutsVisiblesParValideur.
type RecupererStatutsVisiblesParValideurResponse struct {
	CodeRetour         int64           `json:"codeRetour"`
	Libelle            string          `json:"libelle"`
	ListeStatutFacture []StatutVisible `json:"listeStatutFacture"`
}

// La méthode RecupererStatutsFactureVisiblesParValideur permet à un valideur de
// récupérer la liste des statuts auxquels une facture peut être récupérée pour traitement.
func (s *FacturesService) RecupererStatutsVisiblesParValideur(ctx context.Context, opts RecupererStatutsVisiblesParValideurOptions) (*RecupererStatutsVisiblesParValideurResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/factures/v1/recuperer/statutsFactureVisibles/valideur", opts)
	if err != nil {
		return nil, err
	}

	res := new(RecupererStatutsVisiblesParValideurResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
