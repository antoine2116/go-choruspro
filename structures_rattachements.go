package choruspro

import (
	"context"
	"net/http"
	"time"
)

// RattachementServiceUtilisateur est la structure de données représentant un
// rattachement d'un utilisateur à un service.
type RattachementServiceUtilisateur struct {
	CodeService        string     `json:"codeService"`
	Commentaire        string     `json:"commentaire"`
	DateDebut          *time.Time `json:"dateDebut"`
	DateFin            *time.Time `json:"dateFin"`
	IdRattachement     int64      `json:"idRattachement"`
	IdService          int64      `json:"idService"`
	NomService         string     `json:"nomService"`
	StatutRattachement string     `json:"statutRattachement"`
}

// ListeRattachementsServicesUtilisateurResponse est la structure de données
// représentant la réponse de la méthode ListerRattachementsServicesUtilisateur.
type ListeRattachementsServicesUtilisateurResponse struct {
	CodeRetour           int32                            `json:"codeRetour"`
	Libelle              string                           `json:"libelle"`
	DesignationStructure string                           `json:"designationStructure"`
	IdStructure          int64                            `json:"idStructure"`
	IdentifiantStructure string                           `json:"identifiantStructure"`
	Rattachements        []RattachementServiceUtilisateur `json:"listeServices"`
	Pagination           *PaginationResponse              `json:"parametresRetour"`
	PremierService       bool                             `json:"premierService"`
}

// ListeRattachementsServicesUtilisateurOptions est la structure de données
// utilisée pour appeler la méthode ListerRattachementsServicesUtilisateur.
type ListeRattachementsServicesUtilisateurOptions struct {
	AdresseEmailConnexionUtilisateur string             `json:"adresseEmailConnexionUtilisateur"`
	IdStructure                      int64              `json:"idStructure"`
	IdUtilisateur                    int64              `json:"idUtilisateur"`
	IdentifiantStructure             string             `json:"identifiantStructure"`
	Pagination                       *PaginationOptions `json:"parametresListerRattachementsStructureUtilisateur"`
	TypeIdentifiantStructure         string             `json:"typeIdentifiantStructure"`
}

// La méthode RecupererRattachementsServicesUtilisateur récupère les informations
// nécessaires au formulaire de traitement des demandes de rattachement pour
// un utilisateur et une structure.
func (s *StructuresService) RecupererRattachementsServicesUtilisateur(ctx context.Context, opts ListeRattachementsServicesUtilisateurOptions) (*ListeRattachementsServicesUtilisateurResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/structures/v1/lister/rattachementsutilisateur", opts)
	if err != nil {
		return nil, err
	}

	res := new(ListeRattachementsServicesUtilisateurResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
