package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestTransversesService_RecupererStructuresActivesFactureTravaux(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/structures/actives/facturetravaux", func(w http.ResponseWriter, r *http.Request) {
		v := new(struct{})
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeStructures": [
				{
					"idStructureCPP": 1,
					"identifiant": "1",
					"designationStructure": "d1"
				},
				{
					"idStructureCPP": 2,
					"identifiant": "2",
					"designationStructure": "d2"
				}
			]
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.RecupererStructuresActivesFactureTravaux(ctx)
	if err != nil {
		t.Errorf("Transverses.RecupererStructuresActivesFactureTravaux returned error : %v", err)
	}

	want := &ListeStructuresActivesResponse{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Structures: []StructureActive{
			{
				IdCPP:       1,
				Identifiant: "1",
				Designation: "d1",
			},
			{
				IdCPP:       2,
				Identifiant: "2",
				Designation: "d2",
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererStructuresActivesFactureTravaux returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererStructuresActivesFactureTravaux", client, func() error {
		_, err := client.Transverses.RecupererStructuresActivesFactureTravaux(ctx)
		return err
	})
}

func TestTransversesService_RecupererStructuresActivesFournisseur(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/structures/actives/fournisseur", func(w http.ResponseWriter, r *http.Request) {
		v := new(struct{})
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeStructures": [
				{
					"idStructureCPP": 1,
					"identifiant": "1",
					"designationStructure": "d1"
				},
				{
					"idStructureCPP": 2,
					"identifiant": "2",
					"designationStructure": "d2"
				}
			]
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.RecupererStructuresActivesFournisseur(ctx)
	if err != nil {
		t.Errorf("Transverses.RecupererStructuresActivesFournisseur returned error : %v", err)
	}

	want := &ListeStructuresActivesResponse{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Structures: []StructureActive{
			{
				IdCPP:       1,
				Identifiant: "1",
				Designation: "d1",
			},
			{
				IdCPP:       2,
				Identifiant: "2",
				Designation: "d2",
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererStructuresActivesFournisseur returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererStructuresActivesFournisseur", client, func() error {
		_, err := client.Transverses.RecupererStructuresActivesFournisseur(ctx)
		return err
	})
}

func TestTransversesService_RecupererStructuresActivesDestinataire(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/structures/actives/destinataire", func(w http.ResponseWriter, r *http.Request) {
		v := new(struct{})
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeStructures": [
				{
					"idStructureCPP": 1,
					"identifiant": "1",
					"designationStructure": "d1"
				},
				{
					"idStructureCPP": 2,
					"identifiant": "2",
					"designationStructure": "d2"
				}
			]
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.RecupererStructuresActivesDestinataire(ctx)
	if err != nil {
		t.Errorf("Transverses.RecupererStructuresActivesDestinataire returned error : %v", err)
	}

	want := &ListeStructuresActivesResponse{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Structures: []StructureActive{
			{
				IdCPP:       1,
				Identifiant: "1",
				Designation: "d1",
			},
			{
				IdCPP:       2,
				Identifiant: "2",
				Designation: "d2",
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererStructuresActivesDestinataire returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererStructuresActivesDestinataire", client, func() error {
		_, err := client.Transverses.RecupererStructuresActivesDestinataire(ctx)
		return err
	})
}

func TestTransversesService_ConsulterInformationsSIRET(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/consulter/information/siret", func(w http.ResponseWriter, r *http.Request) {
		v := new(ConsulterInformationsSIRETOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "l",
			"adresse": "a",
			"categorieEntreprise": "ce",
			"categorieJuridique": "cj",
			"civilite": "c",
			"dateCreationEntreprise": "2000-01-01",
			"dateCreationEtablissement": "2000-01-01",
			"dateReactivationEtablissement": "2000-01-01",
			"estActif": true,
			"localisationSiege": "l",
			"nonDiffusibleInsee": false,
			"numeroInterneClassement": "n",
			"raisonSociale": "rs",
			"siege": "s",
			"siren": "s",
			"siret": "s",
			"siretPredecesseurSuccesseur": "s"
		}`))
	})

	ctx := context.Background()
	opt := ConsulterInformationsSIRETOptions{Siret: "s"}
	got, err := client.Transverses.ConsulterInformationsSIRET(ctx, opt)

	if err != nil {
		t.Errorf("Transverses.ConsulterInformationsSIRET returned error : %v", err)
	}

	want := &InformationsSIRETResponse{
		CodeRetour:                    0,
		Libelle:                       "l",
		Adresse:                       "a",
		CategorieEntreprise:           "ce",
		CategorieJuridique:            "cj",
		Civilite:                      "c",
		DateCreationEntreprise:        "2000-01-01",
		DateCreationEtablissement:     "2000-01-01",
		DateReactivationEtablissement: "2000-01-01",
		EstActif:                      true,
		LocalisationSiege:             "l",
		NonDiffusibleInsee:            false,
		NumeroInterneClassement:       "n",
		RaisonSociale:                 "rs",
		Siege:                         "s",
		Siren:                         "s",
		Siret:                         "s",
		SiretPredecesseurSuccesseur:   "s",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.ConsulterInformationsSIRET returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "ConsulterInformationsSIRET", client, func() error {
		_, err := client.Transverses.ConsulterInformationsSIRET(ctx, opt)
		return err
	})
}

func TestTransversesService_RechercherDestinataires(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/rechercher/destinataire", func(w http.ResponseWriter, r *http.Request) {
		v := new(ListeDestinatairesOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "l",
			"listeDestinataires": [
				{
					"idStructureCPP": 1,
					"siret": "s",
					"nomDestinataire": "n",
					"adresseCodePostal": "cp",
					"adresseVille": "v"
				}
			],
			"parametresRetour": {
				"nbResultatsParPage": 10,
				"pageCourante": 1,
				"pages": 1,
				"total": 1
			}
		}`))
	})

	ctx := context.Background()
	opt := ListeDestinatairesOptions{}
	got, err := client.Transverses.RechercherDestinataires(ctx, opt)

	if err != nil {
		t.Errorf("Transverses.RechercherDestinataires returned error : %v", err)
	}

	want := &ListeDestinatairesResponse{
		CodeRetour: 0,
		Libelle:    "l",
		Destinataires: []DestinataireRecherche{
			{
				IdStructureCPP:    1,
				Siret:             "s",
				Nom:               "n",
				AdresseCodePostal: "cp",
				AdresseVille:      "v",
			},
		},
		Pagination: PaginationResponse{
			NbResultatsParPage: 10,
			PageCourante:       1,
			Pages:              1,
			Total:              1,
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RechercherDestinataires returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RechercherDestinataires", client, func() error {
		_, err := client.Transverses.RechercherDestinataires(ctx, opt)
		return err
	})
}
