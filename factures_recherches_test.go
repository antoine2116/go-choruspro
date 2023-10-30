package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestFacturesService_RechercherDemandePaiement(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/factures/v1/rechercher/demandePaiement", func(w http.ResponseWriter, r *http.Request) {
		v := new(RechercherDemandePaiementOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"libelle": "l",
			"listeDemandePaiement": [
				{
					"cadreFacturation": "c",
					"codeJuridiction": "c",
					"codeServiceDestinataire": "c",
					"codeServiceSollicitation": "c",
					"commentaireEtatCourant": "c",
					"dateDepot": ` + defaultDateStr + `,
					"dateFournisseur": ` + defaultDateStr + `,
					"dateHeureEtatCourant": ` + defaultDateTimeStr + `,
					"dpOuDpArchive": "d",
					"etatCourant": "e",
					"idDemandePaiement": 1,
					"idDossierFacturation": 1,
					"idServiceDestinataire": 1,
					"idServiceFournisseur": 1,
					"idServiceMoa": 1,
					"idServiceMoe": 1,
					"idStructureDestinataire": 1,
					"idStructureFournisseur": 1,
					"idStructureMoa": 1,
					"idStructureMoe": 1,
					"idStructureValideur1": 1,
					"idStructureValideur2": 1,
					"identifiantStructureDestinataire": "i",
					"identifiantStructureFournisseur": "i",
					"libelleCodePostaleDeptDR": "l",
					"montantAPayer": 1,
					"montantHT": 1,
					"montantTTC": 1,
					"nomServiceDestinataire": "n",
					"nomServiceSollicitation": "n",
					"nombreResultatTotal": "n",
					"numeroBonCommande": "n",
					"numeroDemandePaiement": "n",
					"numeroDossierFacturation": "n",
					"numeroDpMandat": "n",
					"numeroFactureOrigine": "n",
					"numeroMarche": "n",
					"raisonSocialeDeptDR": "r",
					"raisonSocialeJuridiction": "r",
					"raisonSocialeTGI": "r",
					"siretDestinataireQualif": "s",
					"siretFournisseurQualif": "s",
					"typeDemandePaiement": "t",
					"typeFacture": "t",
					"typeFactureTravaux": "t"
				}
			]
		}`))
	})

	ctx := context.Background()
	opt := RechercherDemandePaiementOptions{}
	got, err := client.Factures.RechercherDemandePaiement(ctx, opt)

	if err != nil {
		t.Errorf("Factures.RechercherDemandePaiement returned error : %v", err)
	}

	want := &RechercherDemandePaiementResponse{
		CodeRetour: 1,
		Libelle:    "l",
		DemandesPaiement: &[]DemandePaiement{
			{
				CadreFacturation:                 "c",
				CodeJuridiction:                  "c",
				CodeServiceDestinataire:          "c",
				CodeServiceSollicitation:         "c",
				CommentaireEtatCourant:           "c",
				DateDepot:                        &Date{defaultDate},
				DateFournisseur:                  &Date{defaultDate},
				DateHeureEtatCourant:             &defaultDate,
				DpOuDpArchive:                    "d",
				EtatCourant:                      "e",
				IdDemandePaiement:                1,
				IdDossierFacturation:             1,
				IdServiceDestinataire:            1,
				IdServiceFournisseur:             1,
				IdServiceMoa:                     1,
				IdServiceMoe:                     1,
				IdStructureDestinataire:          1,
				IdStructureFournisseur:           1,
				IdStructureMoa:                   1,
				IdStructureMoe:                   1,
				IdStructureValideur1:             1,
				IdStructureValideur2:             1,
				IdentifiantStructureDestinataire: "i",
				IdentifiantStructureFournisseur:  "i",
				LibelleCodePostaleDeptDR:         "l",
				MontantAPayer:                    1,
				MontantHT:                        1,
				MontantTTC:                       1,
				NomServiceDestinataire:           "n",
				NomServiceSollicitation:          "n",
				NombreResultatTotal:              "n",
				NumeroBonCommande:                "n",
				NumeroDemandePaiement:            "n",
				NumeroDossierFacturation:         "n",
				NumeroDpMandat:                   "n",
				NumeroFactureOrigine:             "n",
				NumeroMarche:                     "n",
				RaisonSocialeDeptDR:              "r",
				RaisonSocialeJuridiction:         "r",
				RaisonSocialeTGI:                 "r",
				SiretDestinataireQualif:          "s",
				SiretFournisseurQualif:           "s",
				TypeDemandePaiement:              "t",
				TypeFacture:                      "t",
				TypeFactureTravaux:               "t",
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Factures.RechercherDemandePaiement returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RechercherDemandePaiement", client, func() error {
		_, err := client.Factures.RechercherDemandePaiement(ctx, opt)
		return err
	})
}

func TestFacturesService_RechercherFactureParRecipiendaire(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/factures/v1/rechercher/recipiendaire", func(w http.ResponseWriter, r *http.Request) {
		v := new(RechercherFactureRecipiendaireOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"libelle": "l",
			"listeFactures": [
				{
					"codeDestinataire": "c",
					"codeFournisseur": "c",
					"codeMOA": "c",
					"codeMOE": "c",
					"codeServiceExecutant": "c",
					"codeServiceFournisseur": "c",
					"commentaireEtatCourant": "c",
					"dateDepot": ` + defaultDateStr + `,
					"dateFacture": ` + defaultDateStr + `,
					"dateHeureEtatCourant": ` + defaultDateTimeStr + `,
					"designationDestinataire": "d",
					"designationFournisseur": "d",
					"designationMOA": "d",
					"designationMOE": "d",
					"devise": "d",
					"factureTelechargeeParDestinataire": true,
					"idDestinataire": 1,
					"idFacture": 1,
					"idServiceExecutant": 1,
					"montantAPayer": 1,
					"montantHT": 1,
					"montantTTC": 1,
					"nomServiceExecutant": "n",
					"nomServiceFournisseur": "n",
					"numeroBonCommande": "n",
					"numeroFacture": "n",
					"numeroMarche": "n",
					"statut": "s",
					"taille": 1,
					"typeDemandePaiement": "t",
					"typeFacture": "t",
					"typeFactureTravaux": "t",
					"typeIdentifiantFournisseur": "t",
					"typeIdentifiantMOA": "t",
					"typeIdentifiantMOE": "t"
				}
			],
			"nbResultatsParPage": 1,
			"pageCourante": 1,
			"pages": 1,
			"total": 1
		}`))
	})

	ctx := context.Background()
	opt := RechercherFactureRecipiendaireOptions{}
	got, err := client.Factures.RechercherFactureParRecipiendaire(ctx, opt)

	if err != nil {
		t.Errorf("Factures.RechercherFactureParRecipiendaire returned error : %v", err)
	}

	want := &RechercherFactureRecipiendaireResponse{
		CodeRetour: 1,
		Libelle:    "l",
		Factures: &[]FactureRecherche{
			{
				CodeDestinataire:                  "c",
				CodeFournisseur:                   "c",
				CodeMOA:                           "c",
				CodeMOE:                           "c",
				CodeServiceExecutant:              "c",
				CodeServiceFournisseur:            "c",
				CommentaireEtatCourant:            "c",
				DateDepot:                         &Date{defaultDate},
				DateFacture:                       &Date{defaultDate},
				DateHeureEtatCourant:              &defaultDate,
				DesignationDestinataire:           "d",
				DesignationFournisseur:            "d",
				DesignationMOA:                    "d",
				DesignationMOE:                    "d",
				Devise:                            "d",
				FactureTelechargeeParDestinataire: true,
				IdDestinataire:                    1,
				IdFacture:                         1,
				IdServiceExecutant:                1,
				MontantAPayer:                     1,
				MontantHT:                         1,
				MontantTTC:                        1,
				NomServiceExecutant:               "n",
				NomServiceFournisseur:             "n",
				NumeroBonCommande:                 "n",
				NumeroFacture:                     "n",
				NumeroMarche:                      "n",
				Statut:                            "s",
				Taille:                            1,
				TypeDemandePaiement:               "t",
				TypeFacture:                       "t",
				TypeFactureTravaux:                "t",
				TypeIdentifiantFournisseur:        "t",
				TypeIdentifiantMOA:                "t",
				TypeIdentifiantMOE:                "t",
			},
		},
		NbResultatsParPage: 1,
		PageCourante:       1,
		Pages:              1,
		Total:              1,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Factures.RechercherFactureParRecipiendaire returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RechercherFactureParRecipiendaire", client, func() error {
		_, err := client.Factures.RechercherFactureParRecipiendaire(ctx, opt)
		return err
	})
}
