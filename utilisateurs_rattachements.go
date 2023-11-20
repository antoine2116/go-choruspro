package choruspro

import (
	"context"
	"net/http"
)

// RattachementService est la structure de données représentant
// un rattachement à un service. Elle est utilisée dans la réponse de la méthode
// RecupererRattachementsUtilisateur.
type RattachementService struct {
	CodeService               string    `json:"codeService"`
	DateDbtService            *Datetime `json:"dateDbtService"`
	DateFinService            *Datetime `json:"dateFinService"`
	EstActif                  bool      `json:"estActif"`
	IdRattachementService     int       `json:"idRattachementService"`
	IdService                 int       `json:"idService"`
	LibelleService            string    `json:"libelleService"`
	StatutRattachementService string    `json:"statutRattachementService"`
	StatutService             string    `json:"statutService"`
}

// RattachementStructure est la structure de données représentant
// un rattachement à une structure. Elle est utilisée dans la réponse de la méthode
// RecupererRattachementsUtilisateur.
type RattachementStructure struct {
	DesignationStructure        string                `json:"designationStructure"`
	IdRattachementStructure     int64                 `json:"idRattachementStructure"`
	IdStructure                 int64                 `json:"idStructure"`
	IdentifiantStructure        string                `json:"identifiantStructure"`
	ServicesRattaches           []RattachementService `json:"listeServicesRattache"`
	RoleUtilisateur             string                `json:"roleUtilisateur"`
	StatutRattachementStructure string                `json:"statutRattachementStructure"`
	StatutUtilisateur           string                `json:"statutUtilisateur"`
	TypeIdentifiantStructure    string                `json:"typeIdentifiantStructure"`
}

// ListeRattachementsUtilisateurResponse est la structure de données réprésentant
// la réponse de la méthode RecupererRattachementsUtilisateur.
type ListeRattachementsUtilisateurResponse struct {
	CodeRetour    int32                   `json:"codeRetour"`
	Libelle       string                  `json:"libelle"`
	Rattachements []RattachementStructure `json:"listeStructureRattachement"`
	Pagination    *PaginationResponse     `json:"parametresRetour"`
}

// ListeRattachementsUtilisateurOptions est la structure de données utlisée
// pour appeler la méthode RecupererRattachementsUtilisateur.
type ListeRattachementsUtilisateurOptions struct {
	Pagination *PaginationOptions `json:"parametresRecherche"`
}

// La méthode recupererRattachementsMonCompteUtilisateur permet de rechercher les
// structures pour lesquelles l'utilisateur possède au moins un rattachement actif.
func (s *UtilisateursService) RecupererRattachementsUtilisateur(ctx context.Context, opts ListeRattachementsUtilisateurOptions) (*ListeRattachementsUtilisateurResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/utilisateurs/v1/monCompte/recuperer/rattachements", opts)
	if err != nil {
		return nil, err
	}

	res := new(ListeRattachementsUtilisateurResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
