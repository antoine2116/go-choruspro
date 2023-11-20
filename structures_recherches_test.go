package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestUtilisateursService_RechercherStructures(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/structures/v1/rechercher", func(w http.ResponseWriter, r *http.Request) {
		v := new(RechercherStructuresOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"libelle": "l",
			"listeStructures": [
				{
					"designationStructure": "d",
					"idStructureCPP": 1,
					"identifiantStructure": "i",
					"statut": "s",
					"typeIdentifiantStructure": "t"
				}
			],
			"parametresRetour": {
				"pageCourante": 1,
				"pages": 1,
				"nbResultatsParPage": 1,
				"total": 1
			}
		}`))
	})

	ctx := context.Background()
	opt := RechercherStructuresOptions{}
	got, err := client.Structures.RechercherStructures(ctx, opt)

	if err != nil {
		t.Errorf("Structures.RechercherStructures returned error : %v", err)
	}

	want := &RechercherStructuresResponse{
		CodeRetour: 1,
		Libelle:    "l",
		Structures: []StructureRecherche{
			{
				Designation:     "d",
				IdStructureCPP:  1,
				Identifiant:     "i",
				Statut:          "s",
				TypeIdentifiant: "t",
			},
		},
		Pagination: &PaginationResponse{
			PageCourante:       1,
			Pages:              1,
			NbResultatsParPage: 1,
			Total:              1,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Structures.RechercherStructures returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RechercherStructures", client, func() error {
		_, err := client.Structures.RechercherStructures(ctx, opt)
		return err
	})
}

func TestUtilisateursService_RechercherServices(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/structures/v1/rechercher/services", func(w http.ResponseWriter, r *http.Request) {
		v := new(RechercherServicesOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"libelle": "l",
			"listeServices": [
				{
					"codeService": "c",
					"estActif": true,
					"idService": 1,
					"libelleService": "l",
					"statutService": "s"
				}
			],
			"parametresRetour": {
				"pageCourante": 1,
				"pages": 1,
				"nbResultatsParPage": 1,
				"total": 1
			}
		}`))
	})

	ctx := context.Background()
	opt := RechercherServicesOptions{}
	got, err := client.Structures.RechercherServices(ctx, opt)

	if err != nil {
		t.Errorf("Structures.RechercherServices returned error : %v", err)
	}

	want := &RechercherServicesResponse{
		CodeRetour: 1,
		Libelle:    "l",
		Services: []ServiceRecherche{
			{
				Code:     "c",
				EstActif: true,
				Id:       1,
				Libelle:  "l",
				Statut:   "s",
			},
		},
		Pagination: &PaginationResponse{
			PageCourante:       1,
			Pages:              1,
			NbResultatsParPage: 1,
			Total:              1,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Structures.RechercherServices returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RechercherServices", client, func() error {
		_, err := client.Structures.RechercherServices(ctx, opt)
		return err
	})
}
