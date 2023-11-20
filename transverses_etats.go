package choruspro

import (
	"context"
	"net/http"
)

// ListeEtatsTypeDemandePaiementResponse est la structure de données représentant
// la réponse de la méthode RecupererEtatParTypeDemandePaiement.
type ListeEtatsTypeDemandePaiementResponse struct {
	CodeRetour int32                     `json:"codeRetour"`
	Libelle    string                    `json:"libelle"`
	Etats      []EtatTypeDemandePaiement `json:"listeEtatDemandePaiement"`
}

// EtatTypeDemandePaiement est la structure de données représentant
// un état de demande de paiement.
type EtatTypeDemandePaiement struct {
	Etat string `json:"etatDemandePaiement"`
}

// ListeEtatsTypeDemandePaiementOptions est la structure de données utlisée
// pour appeler la méthode RecupererEtatParTypeDemandePaiement.
type ListeEtatsTypeDemandePaiementOptions struct {
	TypeDemandePaiement string `json:"typeDemandePaiement,omitempty"`
}

// La méthode recupererEtatParTypeDemandePaiement permet de récupérer
// l'ensemble des statuts pouvant s'appliquer à une demande de paiement
// en fonction du type de demande de paiement.
func (s *TransversesService) RecupererEtatParTypeDemandePaiement(ctx context.Context, opts ListeEtatsTypeDemandePaiementOptions) (*ListeEtatsTypeDemandePaiementResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/recuperer/etat/typedp", opts)
	if err != nil {
		return nil, err
	}

	etats := new(ListeEtatsTypeDemandePaiementResponse)

	err = s.client.doRequest(ctx, req, etats)
	if err != nil {
		return nil, err
	}

	return etats, nil
}

// ListeEtatsTraitementResponse est la structure de données représentant
// la réponse de la méthode RecupererEtatsPossiblesPourTraitement.
type ListeEtatsTraitementResponse struct {
	CodeRetour int32            `json:"codeRetour"`
	Libelle    string           `json:"libelle"`
	Etats      []EtatTraitement `json:"listeStatutsPossiblesPourTraitement"`
}

// EtatTraitement est la structure de données représentant
// le statut possible pour le traitement d'une facture.
type EtatTraitement struct {
	Etat string `json:"statutPossiblePourTraitement"`
}

// ListeEtatsTraitementOptions est la structure de données utlisée
// pour appeler la méthode RecupererEtatsPossiblesPourTraitement.
type ListeEtatsTraitementOptions struct {
	EtatCourant string `json:"statutCourant"`
}

// La méthode RecupererEtatsPossiblesPourTraitement permet de récupérer
// la liste des statuts pouvant être renseignés par un destinataire souhaitant
// traiter une facture en fonction du statut actuel de la demande de paiement.
func (s *TransversesService) RecupererEtatsPossiblesPourTraitement(ctx context.Context, opts ListeEtatsTraitementOptions) (*ListeEtatsTraitementResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/recuperer/etats/traitement", opts)
	if err != nil {
		return nil, err
	}

	etats := new(ListeEtatsTraitementResponse)

	err = s.client.doRequest(ctx, req, etats)
	if err != nil {
		return nil, err
	}

	return etats, nil
}
