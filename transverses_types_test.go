package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestTransversesService_RecupererTypesDemandePaiement(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/typedp", func(w http.ResponseWriter, r *http.Request) {
		v := new(struct{})
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "GCU_MSG_01_000",
			"listeTypeDemandePaiement": [
				{
					"typeDemandePaiement": "t1"
				},
				{
					"typeDemandePaiement": "t2"
				}
			]
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.RecupererTypesDemandePaiement(ctx)

	if err != nil {
		t.Errorf("Transverses.RecupererTypesDemandePaiement returned error : %v", err)
	}

	want := &ListeTypesDemandePaiement{
		CodeRetour: 0,
		Libelle:    "GCU_MSG_01_000",
		Types: []TypeDemandePaiement{
			{
				Type: "t1",
			},
			{
				Type: "t2",
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererTypesDemandePaiement returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererTypesDemandePaiement", client, func() error {
		_, err := client.Transverses.RecupererTypesDemandePaiement(ctx)
		return err
	})
}

func TestTransversesService_RecupererTypesFactureTravaux(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/typefacturetravaux", func(w http.ResponseWriter, r *http.Request) {
		v := new(struct{})
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeTypeFactureTravaux": [
				{
					"typeFactureTravaux": "t1"
				},
				{
					"typeFactureTravaux": "t2"
				}
			]
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.RecupererTypesFactureTravaux(ctx)

	if err != nil {
		t.Errorf("Transverses.RecupererTypesFactureTravaux returned error : %v", err)
	}

	want := &ListeTypesFactureTravaux{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Types: []TypeFactureTravaux{
			{
				Type: "t1",
			},
			{
				Type: "t2",
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererTypesFactureTravaux returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererTypesFactureTravaux", client, func() error {
		_, err := client.Transverses.RecupererTypesFactureTravaux(ctx)
		return err
	})
}

func TestTransversesService_RecupererTypesIdentifiantsStructure(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/typeidentifiant/structure", func(w http.ResponseWriter, r *http.Request) {
		v := new(struct{})
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeTypeIdentifiant": [
				{
					"typeIdentifiantStructure": "t1"
				},
				{
					"typeIdentifiantStructure": "t2"
				}
			]
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.RecupererTypesIdentifiantsStructure(ctx)

	if err != nil {
		t.Errorf("Transverses.RecupererTypesIdentifiantsStructure returned error : %v", err)
	}

	want := &ListeTypesIdentifiantsStructure{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Types: []TypeIdentifiantStructure{
			{
				Type: "t1",
			},
			{
				Type: "t2",
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererTypesIdentifiantsStructure returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererTypesIdentifiantsStructure", client, func() error {
		_, err := client.Transverses.RecupererTypesIdentifiantsStructure(ctx)
		return err
	})
}
