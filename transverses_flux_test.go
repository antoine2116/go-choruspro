package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestTransversesService_RecupererFormatFlux(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/formatflux", func(w http.ResponseWriter, r *http.Request) {
		v := new(struct{})
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeFormatFlux": [
				{
					"formatFlux": "f1"
				},
				{
					"formatFlux": "f2"
				}
			]
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.RecupererFormatFlux(ctx)

	if err != nil {
		t.Errorf("Transverses.RecupererFormatFlux returned error : %v", err)
	}

	want := &ListeFormatsFluxResponse{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Formats: []FormatFlux{{
			FormatFlux: "f1",
		}, {
			FormatFlux: "f2",
		}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererFormatFlux returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererFormatFlux", client, func() error {
		_, err := client.Transverses.RecupererFormatFlux(ctx)
		return err
	})
}

func TestTransversesService_ConsulterCompteRendu(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/consulterCR", func(w http.ResponseWriter, r *http.Request) {
		v := new(ConsulterCompteRenduOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeAppliPartenaire": "cap",
			"codeInterfaceFlux": "cif",
			"codeRetour": 0,
			"dateDepotFlux": ` + defaultISODateTimeStr + `,
			"dateHeureEtatCourantFlux": ` + defaultISODateTimeStr + `,
			"designationPartenaire": "dp",
			"etatCourantFlux": "ecl",
			"fichierCR": "f",
			"libelle": "l",
			"nomFichierFlux": "nff",
			"numeroFluxDepot": "nfd"
		}`))
	})

	ctx := context.Background()
	opt := ConsulterCompteRenduOptions{NumeroFluxDepot: "1"}
	got, err := client.Transverses.ConsulterCompteRendu(ctx, opt)

	if err != nil {
		t.Errorf("Transverses.ConsulterCompteRendu returned error : %v", err)
	}

	want := &ConsulterCompteRenduResponse{
		CodeAppliPartenaire:      "cap",
		CodeInterfaceFlux:        "cif",
		CodeRetour:               0,
		DateDepotFlux:            &defaultDate,
		DateHeureEtatCourantFlux: &defaultDate,
		DesignationPartenaire:    "dp",
		EtatCourantFlux:          "ecl",
		FichierCR:                "f",
		Libelle:                  "l",
		NomFichierFlux:           "nff",
		NumeroFluxDepot:          "nfd",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.ConsulterCompteRendu returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "ConsulterCompteRendu", client, func() error {
		_, err := client.Transverses.ConsulterCompteRendu(ctx, opt)
		return err
	})
}

func TestTransversesService_ConsulterCompteRenduDetaille(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/consulterCRDetaille", func(w http.ResponseWriter, r *http.Request) {
		v := new(ConsulterCompteRenduDetailleOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeInterfaceDepotFlux": "cidf",
			"codeRetour": 0,
			"dateDepotFlux": ` + defaultISODateTimeStr + `,
			"dateHeureEtatCourantFlux": ` + defaultISODateTimeStr + `,
			"etatCourantDepotFlux": "ecdf",
			"libelle": "l",
			"nomFichier": "nf",
			"listeErreurDP": [
				{
					"identifiantDestinataire": "id",
					"identifiantFournisseur": "if",
					"libelleErreurDP": "l",
					"numeroDP": "n"
				}
			],
			"listeErreurTechnique": [
				{
					"codeErreur": "c",
					"libelleErreur": "l",
					"natureErreur": "n"
				}
			]
		}
		`))
	})

	ctx := context.Background()
	opt := ConsulterCompteRenduDetailleOptions{NumeroFluxDepot: "1"}
	got, err := client.Transverses.ConsulterCompteRenduDetaille(ctx, opt)

	if err != nil {
		t.Errorf("Transverses.ConsulterCompteRenduDetaille returned error : %v", err)
	}

	want := &CompteRenduDetailleResponse{
		CodeInterfaceDepotFlux:   "cidf",
		CodeRetour:               0,
		DateDepotFlux:            &defaultDate,
		DateHeureEtatCourantFlux: &defaultDate,
		EtatCourantDepotFlux:     "ecdf",
		Libelle:                  "l",
		NomFichier:               "nf",
		ErreursDemandePaiement: []ErreurDemandePaiement{
			{
				IdentifiantDestinataire: "id",
				IdentifiantFournisseur:  "if",
				LibelleErreurDP:         "l",
				NumeroDP:                "n",
			},
		},
		ErreursTechniques: []ErreurTechnique{
			{
				CodeErreur:    "c",
				LibelleErreur: "l",
				NatureErreur:  "n",
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.ConsulterCompteRenduDetaille returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "ConsulterCompteRenduDetaille", client, func() error {
		_, err := client.Transverses.ConsulterCompteRenduDetaille(ctx, opt)
		return err
	})
}
