package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestFacturesService_TraiterFactureAValider(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/factures/v1/traiter/factureAValider", func(w http.ResponseWriter, r *http.Request) {
		v := new(TraiterFactureAValiderOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"dateTraitement": ` + defaultDateStr + `,
			"idFacture": 1,
			"libelle": "l",
			"numeroFacture": "n",
			"statutFacture": "s"
		}`))
	})

	ctx := context.Background()
	opt := TraiterFactureAValiderOptions{}
	got, err := client.Factures.TraiterFactureAValider(ctx, opt)

	if err != nil {
		t.Errorf("Factures.TraiterFactureAValider returned error : %v", err)
	}

	want := &TraiterFactureAValiderResponse{
		CodeRetour:     1,
		DateTraitement: &Date{defaultDate},
		IdFacture:      1,
		Libelle:        "l",
		NumeroFacture:  "n",
		StatutFacture:  "s",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Factures.TraiterFactureAValider returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "TraiterFactureAValider", client, func() error {
		_, err := client.Factures.TraiterFactureAValider(ctx, opt)
		return err
	})
}

func TestFacturesService_TraiterFactureRecue(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/factures/v1/traiter/recue", func(w http.ResponseWriter, r *http.Request) {
		v := new(TraiterFactureRecueOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"dateTraitement": ` + defaultDateStr + `,
			"idFacture": 1,
			"libelle": "l",
			"nouveauStatut": "n",
			"numeroFacture": "n"
		}`))
	})

	ctx := context.Background()
	opt := TraiterFactureRecueOptions{}
	got, err := client.Factures.TraiterFactureRecue(ctx, opt)

	if err != nil {
		t.Errorf("Factures.TraiterFactureRecue returned error : %v", err)
	}

	want := &TraiterFactureRecueResponse{
		CodeRetour:     1,
		DateTraitement: &Date{defaultDate},
		IdFacture:      1,
		Libelle:        "l",
		NouveauStatut:  "n",
		NumeroFacture:  "n",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Factures.TraiterFactureRecue returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "TraiterFactureRecue", client, func() error {
		_, err := client.Factures.TraiterFactureRecue(ctx, opt)
		return err
	})
}

func TestFacturesService_TraiterRejet(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/factures/v1/traiterRejet", func(w http.ResponseWriter, r *http.Request) {
		v := new(TraiterRejetOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"libelle": "l"
		}`))
	})

	ctx := context.Background()
	opt := TraiterRejetOptions{}
	got, err := client.Factures.TraiterFacturesRejetees(ctx, opt)

	if err != nil {
		t.Errorf("Factures.TraiterRejet returned error : %v", err)
	}

	want := &TraiterRejetResponse{
		CodeRetour: 1,
		Libelle:    "l",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Factures.TraiterRejet returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "TraiterRejet", client, func() error {
		_, err := client.Factures.TraiterFacturesRejetees(ctx, opt)
		return err
	})
}
