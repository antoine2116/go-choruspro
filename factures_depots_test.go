package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestFacturesService_SoumettreFacture(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/factures/v1/soumettre", func(w http.ResponseWriter, r *http.Request) {
		v := new(SoumettreFactureOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"dateDepot": ` + `"2023-01-01"` + `,
			"empreinteCertificatDepot": "e",
			"identifiantFactureCPP": 1,
			"identifiantStructure": "i",
			"libelle": "l",
			"numeroFacture": "F001",
			"statutFacture": "s"
		}`))
	})

	ctx := context.Background()
	opt := SoumettreFactureOptions{
		References: SoumettreFactureReferences{
			NumeroFactureOrigine: "F001",
		},
	}
	got, err := client.Factures.SoumettreFacture(ctx, opt)

	if err != nil {
		t.Errorf("Factures.SoumettreFacture returned error : %v", err)
	}

	want := &SoumettreFactureResponse{
		CodeRetour:               1,
		DateDepot:                &Date{defaultDate},
		EmpreinteCertificatDepot: "e",
		IdentifiantFactureCPP:    1,
		IdentifiantStructure:     "i",
		Libelle:                  "l",
		NumeroFacture:            "F001",
		StatutFacture:            "s",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Factures.SoumettreFacture returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "SoumettreFacture", client, func() error {
		_, err := client.Factures.SoumettreFacture(ctx, opt)
		return err
	})
}

func TestFacturesService_DeposerFlux(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/factures/v1/deposer/flux", func(w http.ResponseWriter, r *http.Request) {
		v := new(DeposerFluxFactureOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"dateDepot": ` + defaultDateStr + `,
			"libelle": "l",
			"numeroFluxDepot": "n",
			"syntaxeFlux": "s"
		}`))
	})

	ctx := context.Background()
	opt := DeposerFluxFactureOptions{}
	got, err := client.Factures.DeposerFlux(ctx, opt)

	if err != nil {
		t.Errorf("Factures.DeposerFlux returned error : %v", err)
	}

	want := &DeposerFluxFactureResponse{
		CodeRetour:      1,
		DateDepot:       &Date{defaultDate},
		Libelle:         "l",
		NumeroFluxDepot: "n",
		SyntaxeFlux:     "s",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Factures.DeposerFlux returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "DeposerFlux", client, func() error {
		_, err := client.Factures.DeposerFlux(ctx, opt)
		return err
	})
}

func TestFacturesService_CompleterFacture(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/factures/v1/completer", func(w http.ResponseWriter, r *http.Request) {
		v := new(CompleterFactureOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"dateTraitement": ` + defaultDateStr + `,
			"identifiantFactureCPP": 1,
			"libelle": "l",
			"numeroFacture": "n"
		}`))
	})

	ctx := context.Background()
	opt := CompleterFactureOptions{}
	got, err := client.Factures.CompleterFacture(ctx, opt)

	if err != nil {
		t.Errorf("Factures.CompleterFacture returned error : %v", err)
	}

	want := &CompleterFactureResponse{
		CodeRetour:            1,
		DateTraitement:        &Date{defaultDate},
		IdentifiantFactureCPP: 1,
		Libelle:               "l",
		NumeroFacture:         "n",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Factures.CompleterFacture returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "CompleterFacture", client, func() error {
		_, err := client.Factures.CompleterFacture(ctx, opt)
		return err
	})
}
