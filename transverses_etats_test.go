package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestTransversesService_RecupererEtatParTypeDemandePaiement(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/etat/typedp", func(w http.ResponseWriter, r *http.Request) {
		v := new(ListeEtatsTypeDemandePaiementOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeEtatDemandePaiement": [
				{
					"etatDemandePaiement": "e1"
				},
				{
					"etatDemandePaiement": "e2"
				}
			]
		}`))
	})

	ctx := context.Background()
	opt := ListeEtatsTypeDemandePaiementOptions{"t"}
	got, err := client.Transverses.RecupererEtatParTypeDemandePaiement(ctx, opt)
	if err != nil {
		t.Errorf("Transverses.RecupererEtatParTypeDemandePaiement returned error : %v", err)
	}

	want := &ListeEtatsTypeDemandePaiementResponse{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Etats: []EtatTypeDemandePaiement{{
			Etat: "e1",
		}, {
			Etat: "e2",
		}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererEtatParTypeDemandePaiement returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererEtatParTypeDemandePaiement", client, func() error {
		_, err := client.Transverses.RecupererEtatParTypeDemandePaiement(ctx, opt)
		return err
	})
}

func TestTransversesService_RecupererEtatsPossiblesPourTraitement(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/etats/traitement", func(w http.ResponseWriter, r *http.Request) {
		v := new(ListeEtatsTraitementOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeStatutsPossiblesPourTraitement": [
				{
					"statutPossiblePourTraitement": "s1"
				},
				{
					"statutPossiblePourTraitement": "s2"
				}
			]
		}`))
	})

	ctx := context.Background()
	opt := ListeEtatsTraitementOptions{"t"}
	got, err := client.Transverses.RecupererEtatsPossiblesPourTraitement(ctx, opt)
	if err != nil {
		t.Errorf("Transverses.RecupererEtatsPossiblesPourTraitement returned error : %v", err)
	}

	want := &ListeEtatsTraitementResponse{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Etats: []EtatTraitement{{
			Etat: "s1",
		}, {
			Etat: "s2",
		}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererEtatsPossiblesPourTraitement returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererEtatsPossiblesPourTraitement", client, func() error {
		_, err := client.Transverses.RecupererEtatsPossiblesPourTraitement(ctx, opt)
		return err
	})
}
