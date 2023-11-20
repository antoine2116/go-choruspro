package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestTransversesService_Healthcheck(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/health-check", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.Write([]byte(`{"statusCode":"200"}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.HealthCheck(ctx)
	if err != nil {
		t.Errorf("Transverses.HealthCheck returned error : %v", err)
	}

	want := &HealthCheck{StatusCode: "200"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.HealthCheck returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "HealthCheck", client, func() error {
		_, err := client.Transverses.HealthCheck(ctx)
		return err
	})
}

func TestTransversesService_RecupererDevises_FR(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/devises", func(w http.ResponseWriter, r *http.Request) {
		v := new(codeLangueOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeDevises": [
				{
					"codeDevise": "EUR",
					"libelleDevise": "Euro européen"
				}
			]
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.RecupererDevises(ctx, CodeLangueFr)
	if err != nil {
		t.Errorf("Transverses.RecupererDevises returned error : %v", err)
	}

	want := &ListeDevises{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Devises: []Devise{{
			Code:    "EUR",
			Libelle: "Euro européen",
		}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererDevises returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererDevises", client, func() error {
		_, err := client.Transverses.RecupererDevises(ctx, "")
		return err
	})
}

func TestTransversesService_RecupererDevises_EN(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/devises", func(w http.ResponseWriter, r *http.Request) {
		v := new(codeLangueOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeDevises": [
				{
					"codeDevise": "EUR",
					"libelleDevise": "European Euro"
				}
			]
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.RecupererDevises(ctx, CodeLangueEn)
	if err != nil {
		t.Errorf("Transverses.RecupererDevises returned error : %v", err)
	}

	want := &ListeDevises{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Devises: []Devise{{
			Code:    "EUR",
			Libelle: "European Euro",
		}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererDevises returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererDevises", client, func() error {
		_, err := client.Transverses.RecupererDevises(ctx, "")
		return err
	})
}

func TestTransversesService_TelechargerAnnuaireDestinataire(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/telecharger/annuaire/destinataire", func(w http.ResponseWriter, r *http.Request) {
		v := new(struct{})
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "GCU_MSG_01_000",
			"fichierResultat": "UEsDBBQACAgIAFQYQ1"
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.TelechargerAnnuaireDestinataire(ctx)
	if err != nil {
		t.Errorf("Transverses.TelechargerAnnuaireDestinataire returned error : %v", err)
	}

	want := &AnnuaireDestinataire{
		CodeRetour: 0,
		Libelle:    "GCU_MSG_01_000",
		Fichier:    "UEsDBBQACAgIAFQYQ1",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.TelechargerAnnuaireDestinataire returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "TelechargerAnnuaireDestinataire", client, func() error {
		_, err := client.Transverses.TelechargerAnnuaireDestinataire(ctx)
		return err
	})
}

func TestTransversesService_RecupererModesDepot(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/modedepot", func(w http.ResponseWriter, r *http.Request) {
		v := new(struct{})
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeModeDepot": [
				{
					"modeDepot": "SAISIE_WEB"
				},
				{
					"modeDepot": "PORTAIL_PDF"
				}
			]
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.RecupererModesDepot(ctx)
	if err != nil {
		t.Errorf("Transverses.RecupererModesDepot returned error : %v", err)
	}

	want := &ListeModesDepot{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		ModesDepot: []ModeDepot{{
			ModeDepot: "SAISIE_WEB",
		}, {
			ModeDepot: "PORTAIL_PDF",
		}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererModesDepot returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererModesDepot", client, func() error {
		_, err := client.Transverses.RecupererModesDepot(ctx)
		return err
	})
}

func TestTransversesService_RecupererPays_FR(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/pays", func(w http.ResponseWriter, r *http.Request) {
		v := new(codeLangueOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listePays": [
				{
					"codePays": "AD",
					"libellePays": "Andorre"
				},
				{
					"codePays": "FR",
					"libellePays": "France"
				}
			]
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.RecupererPays(ctx, CodeLangueFr)
	if err != nil {
		t.Errorf("Transverses.RecupererPays returned error : %v", err)
	}

	want := &ListePays{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Pays: []Pays{{
			Code:    "AD",
			Libelle: "Andorre",
		}, {
			Code:    "FR",
			Libelle: "France",
		}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererPays returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererPays", client, func() error {
		_, err := client.Transverses.RecupererPays(ctx, "")
		return err
	})
}

func TestTransversesService_RecupererPays_EN(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/pays", func(w http.ResponseWriter, r *http.Request) {
		v := new(codeLangueOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listePays": [
				{
					"codePays": "AD",
					"libellePays": "Andorra"
				},
				{
					"codePays": "FR",
					"libellePays": "France"
				}
			]
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.RecupererPays(ctx, CodeLangueEn)
	if err != nil {
		t.Errorf("Transverses.RecupererPays returned error : %v", err)
	}

	want := &ListePays{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Pays: []Pays{{
			Code:    "AD",
			Libelle: "Andorra",
		}, {
			Code:    "FR",
			Libelle: "France",
		}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererPays returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererPays", client, func() error {
		_, err := client.Transverses.RecupererPays(ctx, "")
		return err
	})
}

func TestTransversesService_RecupererMotifsRefusFactureAValider(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/motifs/refus/Facture/AValider", func(w http.ResponseWriter, r *http.Request) {
		v := new(ListeMotifsRefusFactureAValiderOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeMotifRefusFactureAValider": [
				{
					"idMotif": 1,
					"codeMotif": "c1",
					"libelleMotif": "l1"
				},
				{
					"idMotif": 2,
					"codeMotif": "c2",
					"libelleMotif": "l2"
				}
			]
		}`))
	})

	ctx := context.Background()
	opt := ListeMotifsRefusFactureAValiderOptions{"c"}
	got, err := client.Transverses.RecupererMotifsRefusFactureAValider(ctx, opt)

	if err != nil {
		t.Errorf("Transverses.RecupererMotifsRefusFactureAValider returned error : %v", err)
	}

	want := &ListeMotifsRefusFactureAValider{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Motifs: []MotifRefus{{
			Id:      1,
			Code:    "c1",
			Libelle: "l1",
		}, {
			Id:      2,
			Code:    "c2",
			Libelle: "l2",
		}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererMotifsRefusFactureAValider returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererMotifsRefusFactureAValider", client, func() error {
		_, err := client.Transverses.RecupererMotifsRefusFactureAValider(ctx, opt)
		return err
	})
}

func TestTransversesService_RecupererModesReglement(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/modereglements", func(w http.ResponseWriter, r *http.Request) {
		v := new(struct{})
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "GCU_MSG_01_000",
			"listeModePaiement": [
				{
					"modePaiement": "m1"
				},
				{
					"modePaiement": "m2"
				}
			]
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.RecupererModesReglement(ctx)

	if err != nil {
		t.Errorf("Transverses.RecupererModesReglement returned error : %v", err)
	}

	want := &ListeModesReglement{
		CodeRetour: 0,
		Libelle:    "GCU_MSG_01_000",
		Modes: []ModeReglement{{
			ModeReglement: "m1",
		}, {
			ModeReglement: "m2",
		}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererModesReglement returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererModesReglement", client, func() error {
		_, err := client.Transverses.RecupererModesReglement(ctx)
		return err
	})
}

func TestTransversesService_RecupererCadresFacturation(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/cadrefacturation", func(w http.ResponseWriter, r *http.Request) {
		v := new(struct{})
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeCadreFacturation": [
				{
					"codeCadreFacturation": "c1"
				},
				{
					"codeCadreFacturation": "c2"
				}
			]
		}`))
	})

	ctx := context.Background()
	got, err := client.Transverses.RecupererCadresFacturation(ctx, "")

	if err != nil {
		t.Errorf("Transverses.RecupererCadresFacturation returned error : %v", err)
	}

	want := &ListeCadresFacturation{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Cadres: []CadreFacturation{{
			Code: "c1",
		}, {
			Code: "c2",
		}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererCadresFacturation returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererCadresFacturation", client, func() error {
		_, err := client.Transverses.RecupererCadresFacturation(ctx, "")
		return err
	})
}

func TestTransversesService_RecupererCoordonneesBancairesValides(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/coordbanc/valides", func(w http.ResponseWriter, r *http.Request) {
		v := new(ListeCoordonneesBancairesOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "GCU_MSG_01_000",
			"listeCoordonneeBancaire": [
				{
					"idCoordonneeBancaire": 1,
					"nomCoordonneeBancaire": "n1"
				},
				{
					"idCoordonneeBancaire": 2,
					"nomCoordonneeBancaire": "n2"
				}
			]
		}`))
	})

	ctx := context.Background()
	opt := ListeCoordonneesBancairesOptions{1}
	got, err := client.Transverses.RecupererCoordonneesBancairesValides(ctx, opt)

	if err != nil {
		t.Errorf("Transverses.RecupererCoordonneesBancairesValides returned error : %v", err)
	}

	want := &ListeCoordonneesBancaires{
		CodeRetour: 0,
		Libelle:    "GCU_MSG_01_000",
		Coordonnees: []CoordonneeBancaire{{
			Id:  1,
			Nom: "n1",
		}, {
			Id:  2,
			Nom: "n2",
		}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererCoordonneesBancairesValides returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererCoordonneesBancairesValides", client, func() error {
		_, err := client.Transverses.RecupererCoordonneesBancairesValides(ctx, opt)
		return err
	})
}
