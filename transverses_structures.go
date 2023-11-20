package choruspro

import (
	"context"
	"net/http"
)

// ListeStructuresActivesResponse est la structure de données représentant
// la réponse de la méthode RecupererStructuresActives.
type ListeStructuresActivesResponse struct {
	CodeRetour int32             `json:"codeRetour"`
	Libelle    string            `json:"libelle"`
	Structures []StructureActive `json:"listeStructures"`
}

// StructureActive est la structure de données représentant
// une structure active.
type StructureActive struct {
	Designation string `json:"designationStructure"`
	IdCPP       int64  `json:"idStructureCPP"`
	Identifiant string `json:"identifiant"`
}

// La méthode RecupererStructuresActivesFactureTravaux permet de
// récupérer la liste des structures actives auxquelles un utilisateur
// est rattaché et habilité à l'espace "factures de travaux".
func (s *TransversesService) RecupererStructuresActivesFactureTravaux(ctx context.Context) (*ListeStructuresActivesResponse, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/recuperer/structures/actives/facturetravaux", opts)
	if err != nil {
		return nil, err
	}

	res := new(ListeStructuresActivesResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// La méthode RecupererStructuresActivesFournisseur permet de récupérer
// la liste des structures actives auxquelles un utilisateur est rattaché
// et habilité à l’espace "factures émises".
func (s *TransversesService) RecupererStructuresActivesFournisseur(ctx context.Context) (*ListeStructuresActivesResponse, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/recuperer/structures/actives/fournisseur", opts)
	if err != nil {
		return nil, err
	}

	res := new(ListeStructuresActivesResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// La méthode RecupererStructuresActivesDestinataire permet de récupérer
// la liste des structures actives auxquelles un utilisateur est rattaché et
// habilité à l'espace "factures reçues".
func (s *TransversesService) RecupererStructuresActivesDestinataire(ctx context.Context) (*ListeStructuresActivesResponse, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/recuperer/structures/actives/destinataire", opts)
	if err != nil {
		return nil, err
	}

	res := new(ListeStructuresActivesResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// InformationsSIRETResponse est la structure de données représentant
// la réponse de la méthode ConsulterInformationsSIRET.
type InformationsSIRETResponse struct {
	CodeRetour                    int32  `json:"codeRetour"`
	Libelle                       string `json:"libelle"`
	Adresse                       string `json:"adresse,omitempty"`
	CategorieEntreprise           string `json:"categorieEntreprise,omitempty"`
	CategorieJuridique            string `json:"categorieJuridique,omitempty"`
	Civilite                      string `json:"civilite,omitempty"`
	DateCreationEntreprise        string `json:"dateCreationEntreprise,omitempty"`
	DateCreationEtablissement     string `json:"dateCreationEtablissement,omitempty"`
	DateReactivationEtablissement string `json:"dateReactivationEtablissement,omitempty"`
	EstActif                      bool   `json:"estActif,omitempty"`
	LocalisationSiege             string `json:"localisationSiege,omitempty"`
	NonDiffusibleInsee            bool   `json:"nonDiffusibleInsee,omitempty"`
	NumeroInterneClassement       string `json:"numeroInterneClassement,omitempty"`
	RaisonSociale                 string `json:"raisonSociale,omitempty"`
	Siege                         string `json:"siege,omitempty"`
	Siren                         string `json:"siren,omitempty"`
	Siret                         string `json:"siret,omitempty"`
	SiretPredecesseurSuccesseur   string `json:"siretPredecesseurSuccesseur,omitempty"`
}

// ConsulterInformationsSIRETOptions est la structure de données utilisée
// pour appeler la méthode ConsulterInformationsSIRET.
type ConsulterInformationsSIRETOptions struct {
	Siret string `json:"siretStructure"`
}

// La méthode ConsulterInformationsSIRET permet de récupérer les données
// de la structure de type SIRET correspondant aux paramètres dans le
// référentiel de l'INSEE (service exposé ne fonctionnant pas sur
// l'espace de qualification).
func (s *TransversesService) ConsulterInformationsSIRET(ctx context.Context, opts ConsulterInformationsSIRETOptions) (*InformationsSIRETResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/consulter/information/siret", opts)
	if err != nil {
		return nil, err
	}

	res := new(InformationsSIRETResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ListeDestinatairesResponse est la structure de données représentant
// la réponse de la méthode RechercherDestinataires.
type ListeDestinatairesResponse struct {
	CodeRetour    int32                   `json:"codeRetour"`
	Libelle       string                  `json:"libelle"`
	Destinataires []DestinataireRecherche `json:"listeDestinataires"`
	Pagination    PaginationResponse      `json:"parametresRetour"`
}

// DestinataireRecherche est la structure de données représentant
// un destinataire de la liste des destinataires.
type DestinataireRecherche struct {
	IdStructureCPP    int64  `json:"idStructureCPP"`
	Siret             string `json:"siret"`
	Nom               string `json:"nomDestinataire,omitempty"`
	AdresseCodePostal string `json:"adresseCodePostal,omitempty"`
	AdresseVille      string `json:"adresseVille,omitempty"`
}

// ListeDestinatairesOptions est la structure de données utilisée
// pour appeler la méthode RechercherDestinataires.
type ListeDestinatairesOptions struct {
	Identifiant        string             `json:"identifiant,omitempty"`
	NomVilleCodePostal string             `json:"nomVilleCodePostal,omitempty"`
	TypeStructure      string             `json:"typeStructure,omitempty"`
	Pagination         *PaginationOptions `json:"parametresRecherche,omitempty"`
}

// La méthode RechercherDestinataires permet de rechercher les
// informations légales d'une structure publique.
func (s *TransversesService) RechercherDestinataires(ctx context.Context, opts ListeDestinatairesOptions) (*ListeDestinatairesResponse, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "cpro/transverses/v1/rechercher/destinataire", opts)
	if err != nil {
		return nil, err
	}

	res := new(ListeDestinatairesResponse)

	err = s.client.doRequest(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
