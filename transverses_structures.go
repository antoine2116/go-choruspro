package choruspro

import (
	"context"
	"errors"
	"net/http"
)

type ListeStrcturesActives struct {
	CodeRetour int32             `json:"codeRetour"`
	Libelle    string            `json:"libelle"`
	Structures []StructureActive `json:"listeStructures"`
}

type StructureActive struct {
	Designation string `json:"designationStructure"`
	IdCPP       int64  `json:"idStrctureCPP"`
	Identifiant string `json:"identifiant"`
}

func (s *TransversesService) RecupererStructuresActivesFactureTravaux(ctx context.Context) (*ListeStrcturesActives, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/structures/actives/facturetravaux", opts)
	if err != nil {
		return nil, err
	}

	structures := new(ListeStrcturesActives)

	err = s.client.doRequest(ctx, req, structures)
	if err != nil {
		return nil, err
	}

	return structures, nil
}

func (s *TransversesService) RecupererStructuresActivesDestinataireFactureTravaux(ctx context.Context) (*ListeStrcturesActives, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/structures/actives/destinataire", opts)
	if err != nil {
		return nil, err
	}

	structures := new(ListeStrcturesActives)

	err = s.client.doRequest(ctx, req, structures)
	if err != nil {
		return nil, err
	}

	return structures, nil
}

func (s *TransversesService) RecupererStructuresActivesFournisseur(ctx context.Context) (*ListeStrcturesActives, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/structures/actives/fournisseur", opts)
	if err != nil {
		return nil, err
	}

	structures := new(ListeStrcturesActives)

	err = s.client.doRequest(ctx, req, structures)
	if err != nil {
		return nil, err
	}

	return structures, nil
}

func (s *TransversesService) RecupererStructuresActivesDestinataire(ctx context.Context) (*ListeStrcturesActives, error) {
	opts := struct{}{}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/recuperer/structures/actives/destinataire", opts)
	if err != nil {
		return nil, err
	}

	structures := new(ListeStrcturesActives)

	err = s.client.doRequest(ctx, req, structures)
	if err != nil {
		return nil, err
	}

	return structures, nil
}

type InformationsSIRET struct {
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

func (s *TransversesService) ConsulterInformationsSIRET(ctx context.Context, siret string) (*InformationsSIRET, error) {
	if siret == "" {
		return nil, errors.New("siret is required")
	}

	opts := &struct {
		SiretStrucutre string `json:"siretStructure"`
	}{SiretStrucutre: siret}

	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/consulter/information/siret", opts)
	if err != nil {
		return nil, err
	}

	infos := new(InformationsSIRET)

	err = s.client.doRequest(ctx, req, infos)
	if err != nil {
		return nil, err
	}

	return infos, nil
}

type ListeDestinataires struct {
	CodeRetour    int32              `json:"codeRetour"`
	Libelle       string             `json:"libelle"`
	Destinataires []Destinataire     `json:"listeDestinataires"`
	Pagination    PaginationResponse `json:"parametresRetour"`
}

type Destinataire struct {
	IdStructureCPP    int64  `json:"idStructureCPP"`
	Siret             string `json:"siret"`
	Nom               string `json:"nomDestinataire,omitempty"`
	AdresseCodePostal string `json:"adresseCodePostal,omitempty"`
	AdresseVille      string `json:"adresseVille,omitempty"`
}

type ListeDestinatairesOptions struct {
	Identifiant        string             `json:"identifiant,omitempty"`
	NomVilleCodePostal string             `json:"nomVilleCodePostal,omitempty"`
	TypeStructure      string             `json:"typeStructure,omitempty"`
	Pagination         *PaginationOptions `json:"parametresRecherche,omitempty"`
}

func (s *TransversesService) RechercherDestinataires(ctx context.Context, opts ListeDestinatairesOptions) (*ListeDestinataires, error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/cpro/transverses/v1/rechercher/destinataire", opts)
	if err != nil {
		return nil, err
	}

	destinataires := new(ListeDestinataires)

	err = s.client.doRequest(ctx, req, destinataires)
	if err != nil {
		return nil, err
	}

	return destinataires, nil
}
