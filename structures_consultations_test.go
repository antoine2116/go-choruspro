package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestUtilisateursService_ConsulterStructure(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/structures/v1/consulter", func(w http.ResponseWriter, r *http.Request) {
		v := new(ConsulterStructureOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"libelle": "l",
			"adressePostaleDuSiege": {
				"adresse": "a",
				"codePostal": "c",
				"complementAdresse1": "c",
				"complementAdresse2": "c",
				"fax": "f",
				"indicatifFax": "i",
				"indicatifTelephone": "i",
				"pays": "p",
				"telephone": "t",
				"ville": "v"
			},
			"informationGenerales": {
				"dateExpirationMotPasseCompteTech": ` + defaultISODateTimeStr + `,
				"emailStructure": "e",
				"estValideurDeleguee": true,
				"idStructureCPP": 1,
				"idStructureOrigine": 1,
				"identifiantStructure": "i",
				"libelleStructure": "l",
				"listeStructureDupliquees": [
					{
						"structureDupliquee": 1
					},
					{
						"structureDupliquee": 1
					}
				],
				"nomStructure": "n",
				"numeroRcsStructure": "n",
				"numeroTVA": "n",
				"prenomStructure": "p",
				"raisonSocialeStructure": "r",
				"structurePriveePublique": "p",
				"typeIdentifiantStructure": "t"
			},
			"parametres": {
				"codeServiceDoitEtreRenseigne": true,
				"connexionEDI": true,
				"estMoaUniquement": true,
				"estValideurDelegue": true,
				"gestionNumeroEJOuCodeService": true,
				"getaMoa": true,
				"nonDiffusibleInsee": true,
				"numeroEJDoitEtreRenseigne": true,
				"recevoirDonneesViaEDI": true,
				"statutMiseEnPaiementNestPasRemonte": true
			}
		}`))
	})

	ctx := context.Background()
	opt := ConsulterStructureOptions{}
	got, err := client.Structures.ConsulterStructure(ctx, opt)

	if err != nil {
		t.Errorf("Structures.ConsulterStructure returned error : %v", err)
	}

	want := &ConsulterStructureResponse{
		CodeRetour: 1,
		Libelle:    "l",
		AdressePostale: &AdressePostale{
			Adresse:            "a",
			CodePostal:         "c",
			ComplementAdresse1: "c",
			ComplementAdresse2: "c",
			Fax:                "f",
			IndicatifFax:       "i",
			IndicatifTelephone: "i",
			Pays:               "p",
			Telephone:          "t",
			Ville:              "v",
		},
		InformationGenerales: &InformationGeneralesStructure{
			DateExpirationMotPasseCompteTech: &defaultDate,
			Email:                            "e",
			EstValideurDeleguee:              true,
			IdCPP:                            1,
			IdOrigine:                        1,
			Identifiant:                      "i",
			Libelle:                          "l",
			ListeStructureDupliquees: []StructureDupliquee{
				{
					StructureDupliquee: 1,
				},
				{
					StructureDupliquee: 1,
				},
			},
			Nom:             "n",
			NumeroRcs:       "n",
			NumeroTVA:       "n",
			Prenom:          "p",
			RaisonSociale:   "r",
			PriveePublique:  "p",
			TypeIdentifiant: "t",
		},
		Parametrage: &ParametrageStructure{
			CodeServiceDoitEtreRenseigne:       true,
			ConnexionEDI:                       true,
			EstMoaUniquement:                   true,
			EstValideurDelegue:                 true,
			GestionNumeroEJOuCodeService:       true,
			GetaMoa:                            true,
			NonDiffusibleInsee:                 true,
			NumeroEJDoitEtreRenseigne:          true,
			RecevoirDonneesViaEDI:              true,
			StatutMiseEnPaiementNestPasRemonte: true,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Structures.ConsulterStructure returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "ConsulterStructure", client, func() error {
		_, err := client.Structures.ConsulterStructure(ctx, opt)
		return err
	})
}

func TestUtilisateursService_ConsulterService(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/structures/v1/consulter/service", func(w http.ResponseWriter, r *http.Request) {
		v := new(ConsulterServiceOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"libelle": "l",
			"adressePostale": {
				"adresse": "a",
				"codePostal": "c",
				"complementAdresse1": "c",
				"complementAdresse2": "c",
				"fax": "f",
				"indicatifFax": "i",
				"indicatifTelephone": "i",
				"pays": "p",
				"telephone": "t",
				"ville": "v"
			},
			"informationsGenerales": {
				"codeService": "c",
				"descriptionService": "d",
				"nomService": "n"
			},
			"parametres": {
				"dateCreation": ` + defaultDateTimeStr + `,
				"dateDebutValidite": ` + defaultDateTimeStr + `,
				"dateFinValidite": ` + defaultDateTimeStr + `,
				"dateModification": ` + defaultDateTimeStr + `,
				"miseEnPaiement": true,
				"numeroEngagement": true
			}
		}`))
	})

	ctx := context.Background()
	opt := ConsulterServiceOptions{}
	got, err := client.Structures.ConsulterService(ctx, opt)

	if err != nil {
		t.Errorf("Structures.ConsulterService returned error : %v", err)
	}

	want := &ConsulterServiceResponse{
		CodeRetour: 1,
		Libelle:    "l",
		AdressePostale: &AdressePostale{
			Adresse:            "a",
			CodePostal:         "c",
			ComplementAdresse1: "c",
			ComplementAdresse2: "c",
			Fax:                "f",
			IndicatifFax:       "i",
			IndicatifTelephone: "i",
			Pays:               "p",
			Telephone:          "t",
			Ville:              "v",
		},
		InformationsGenerales: &InformationsGeneralesService{
			Code:        "c",
			Description: "d",
			Nom:         "n",
		},
		Parametrage: &ParametrageService{
			DateCreation:      &Datetime{defaultDate},
			DateDebutValidite: &Datetime{defaultDate},
			DateFinValidite:   &Datetime{defaultDate},
			DateModification:  &Datetime{defaultDate},
			MiseEnPaiement:    true,
			NumeroEngagement:  true,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Structures.ConsulterService returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "ConsulterService", client, func() error {
		_, err := client.Structures.ConsulterService(ctx, opt)
		return err
	})
}
