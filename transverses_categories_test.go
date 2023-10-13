package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
	"time"
)

const (
	referenceTimeStr = `"2023-01-01T00:00:00Z"`
)

var (
	referenceTime = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
)

func TestTransversesService_RechercherCategoriesSollicitation(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/rechercher/categorieSollicitation", func(w http.ResponseWriter, r *http.Request) {
		v := new(ListeCategoriesSollicitationOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeCategories": [
				{
					"codeCategorie": "c1",
					"libelleCategorie": "c1",
					"idTechniqueCategorie": 1
				},
				{
					"codeCategorie": "c2",
					"libelleCategorie": "c2",
					"idTechniqueCategorie": 2
				}
			]
		}`))
	})

	ctx := context.Background()
	opt := ListeCategoriesSollicitationOptions{}
	got, err := client.Transverses.RechercherCategoriesSollicitation(ctx, opt)
	if err != nil {
		t.Errorf("Transverses.RechercherCategoriesSollicitation returned error : %v", err)
	}

	want := &ListeCategoriesSollicitation{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Categories: []CategorieSollicitation{{
			Code:        "c1",
			Libelle:     "c1",
			IdTechnique: 1,
		}, {
			Code:        "c2",
			Libelle:     "c2",
			IdTechnique: 2,
		}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RechercherCategoriesSollicitation returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RechercherCategoriesSollicitation", client, func() error {
		_, err := client.Transverses.RechercherCategoriesSollicitation(ctx, opt)
		return err
	})
}

func TestTransversesService_RechercherSousCategoriesSollicitation(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/rechercher/sousCategorieSollicitation", func(w http.ResponseWriter, r *http.Request) {
		v := new(ListeSousCategoriesSollicitationOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeSousCategories": [
				{
					"categorie": [
						{
							"codeCategorie": "c",
							"libelleCategorie": "l",
							"idTechniqueCategorie": 1
						}
					],
					"ssCategorie": [
						{
							"code": "c",
							"dateCreation": ` + referenceTimeStr + `,
							"dateDerniereModification": ` + referenceTimeStr + `,
							"estActif": true,
							"idTechniqueCategorie": 1,
							"libelle": "l"
						}
					]
				}
			]
		}`))
	})

	ctx := context.Background()
	opt := ListeSousCategoriesSollicitationOptions{}
	got, err := client.Transverses.RechercherSousCategoriesSollicitation(ctx, opt)
	if err != nil {
		t.Errorf("Transverses.RechercherSousCategoriesSollicitation returned error : %v", err)
	}

	want := &ListeSousCategoriesSollicitation{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		SousCategories: []SousCategoriesSolliciation{{
			Categories: []CategorieSollicitation{{
				Code:        "c",
				Libelle:     "l",
				IdTechnique: 1,
			}},
			SousCategories: []SousCategorieSollicitation{{
				Code:                     "c",
				DateCreation:             &referenceTime,
				DateDerniereModification: &referenceTime,
				EstActif:                 true,
				IdTechnique:              1,
				Libelle:                  "l",
			}},
		}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RechercherSousCategoriesSollicitation returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RechercherSousCategoriesSollicitation", client, func() error {
		_, err := client.Transverses.RechercherSousCategoriesSollicitation(ctx, opt)
		return err
	})
}
