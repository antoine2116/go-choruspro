package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestTransversesService_RecupererTauxTva(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/tauxtva", func(w http.ResponseWriter, r *http.Request) {
		v := new(codeLangueOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeTauxTva": [
				{
					"codeTauxTva": "c1",
					"libelleTauxTva": "l1",
					"valeurTauxTva": 10
				},
				{
					"codeTauxTva": "c2",
					"libelleTauxTva": "l2",
					"valeurTauxTva": 20
				}
			]
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.RecupererTauxTva(ctx, CodeLangueFr)
	if err != nil {
		t.Errorf("Transverses.RecupererTauxTva returned error : %v", err)
	}

	want := &ListeTauxTva{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Taux: []TauxTva{
			{
				Code:    "c1",
				Libelle: "l1",
				Valeur:  10,
			},
			{
				Code:    "c2",
				Libelle: "l2",
				Valeur:  20,
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererTauxTva returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererTauxTva", client, func() error {
		_, err := client.Transverses.RecupererTauxTva(ctx, "")
		return err
	})
}

func TestTransversesService_RecupererMotifsExonerationTva(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/motifs/exonerationtva", func(w http.ResponseWriter, r *http.Request) {
		v := new(codeLangueOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeMotifsExonerationTva": [
				{
					"codeMotifExonerationTva": "c1",
					"libelleMotifExonerationTva": "l1"
				},
				{
					"codeMotifExonerationTva": "c2",
					"libelleMotifExonerationTva": "l2"
				}
			]
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.RecupererMotifsExonerationTva(ctx, CodeLangueFr)
	if err != nil {
		t.Errorf("Transverses.RecupererMotifsExonerationTva returned error : %v", err)
	}

	want := &ListeMotifsExonerationTva{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Motifs: []MotifExonerationTva{
			{
				Code:    "c1",
				Libelle: "l1",
			},
			{
				Code:    "c2",
				Libelle: "l2",
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererMotifsExonerationTva returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererMotifsExonerationTva", client, func() error {
		_, err := client.Transverses.RecupererMotifsExonerationTva(ctx, "")
		return err
	})
}
