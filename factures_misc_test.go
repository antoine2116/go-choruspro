package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestFacturesService_TelechargerGroupe(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/factures/v1/telecharger/groupe", func(w http.ResponseWriter, r *http.Request) {
		v := new(TraiterFactureAValiderOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"fichierResultat": "f",
			"libelle": "l"
		}`))
	})

	ctx := context.Background()
	opt := TraiterFactureAValiderOptions{}
	got, err := client.Factures.TelechargerGroupe(ctx, opt)

	if err != nil {
		t.Errorf("Factures.TelechargerGroupe returned error : %v", err)
	}

	want := &TelechargerGroupeFactureResponse{
		CodeRetour:      1,
		FichierResultat: "f",
		Libelle:         "l",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Factures.TelechargerGroupe returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "TelechargerGroupe", client, func() error {
		_, err := client.Factures.TelechargerGroupe(ctx, opt)
		return err
	})
}
