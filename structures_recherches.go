package choruspro

import (
	"context"
	"net/http"
)

// StructureRecherche est la structure de données représentant
// une structure de la liste des structures.
type StructureRecherche struct {
	Designation     string `json:"designationStructure"`
	IdStructureCPP  int64  `json:"idStructureCPP"`
	Identifiant     string `json:"identifiantStructure"`
	Statut          string `json:"statut"`
	TypeIdentifiant string `json:"typeIdentifiantStructure"`
}

// RechercherStructuresResponse est la structure de données représentant
// la réponse de la méthode RechercherStructures.
type RechercherStructuresResponse struct {
	CodeRetour int32                `json:"codeRetour"`
	Libelle    string               `json:"libelle"`
	Structures []StructureRecherche `json:"listeStructures"`
	Pagination *PaginationResponse  `json:"parametresRetour"`
}

// CriteresRechercheStructures est la structure de données représentant
// les critères de recherche d'une structure.
type CriteresRechercheStructures struct {
	AdresseCodePays   string `json:"adresseCodePays"`
	AdresseCodePostal string `json:"adresseCodePostal"`
	AdresseVille      string `json:"adresseVille"`
	EstMOA            bool   `json:"estMOA"`
	EstMOAUniquement  bool   `json:"estMOAUniquement"`
	Identifiant       string `json:"identifiantStructure"`
	Libelle           string `json:"libelleStructure"`
	Nom               string `json:"nomStructure"`
	Prenom            string `json:"prenomStructure"`
	RaisonSociale     string `json:"raisonSocialeStructure"`
	Statut            string `json:"statutStructure"`
	TypeIdentifiant   string `json:"typeIdentifiantStructure"`
	Type              string `json:"typeStructure"`
}

// RechercherStructuresOptions est la structure de données utilisée
// pour appeler la méthode RechercherStructures.
type RechercherStructuresOptions struct {
	Pagination                   *PaginationOptions           `json:"parametres"`
	RestreindreStructuresPrivees bool                         `json:"restreindreStructuresPrivees"`
	Criteres                     *CriteresRechercheStructures `json:"structure"`
}

// La méthode RechercherStructures permet à un gestionnaire de rechercher des structures.
func (s *StructuresService) RechercherStructures(ctx context.Context, opts RechercherStructuresOptions) (*RechercherStructuresResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/structures/v1/rechercher", opts)
	if err != nil {
		return nil, err
	}

	res := new(RechercherStructuresResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ServiceRecherche est la structure de données représentant
// un service de la liste des services.
type ServiceRecherche struct {
	Code     string `json:"codeService"`
	EstActif bool   `json:"estActif"`
	Id       int64  `json:"idService"`
	Libelle  string `json:"libelleService"`
	Statut   string `json:"statutService"`
}

// RechercherServicesResponse est la structure de données représentant
// la réponse de la méthode RechercherServices.
type RechercherServicesResponse struct {
	CodeRetour int32               `json:"codeRetour"`
	Libelle    string              `json:"libelle"`
	Services   []ServiceRecherche  `json:"listeServices"`
	Pagination *PaginationResponse `json:"parametresRetour"`
}

// RechercherServicesOptions est la structure de données utilisée
// pour appeler la méthode RechercherServices.
type RechercherServicesOptions struct {
	IdUtilisateurCourant int64              `json:"idUtilisateurCourant"`
	IdStructure          int64              `json:"idStructure"`
	Pagination           *PaginationOptions `json:"parametresRechercherServicesStructure"`
}

// La méthode RechercherServices permet de rechercher les services
// appartenant à une structure publique ou à une structure à laquelle
// l'utilisateur est rattaché.
func (s *StructuresService) RechercherServices(ctx context.Context, opts RechercherServicesOptions) (*RechercherServicesResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/structures/v1/rechercher/services", opts)
	if err != nil {
		return nil, err
	}

	res := new(RechercherServicesResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
