package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestFacturesService_ConsulterFactureParRecipiendaire(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/factures/v1/consulter/recipiendaire", func(w http.ResponseWriter, r *http.Request) {
		v := new(ConsulterFactureOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"facture": {
				"cadreDeFacturation": {
					"codeCadreFacturation": "c",
					"codeServiceMoa": "c",
					"codeServiceMoe": "c",
					"codeValideur1": "c",
					"codeValideur2": "c",
					"dateValidation1": ` + defaultDateStr + `,
					"dateValidation2": ` + defaultDateStr + `,
					"designationMoa": "d",
					"designationMoe": "d",
					"designationValideur2": "d",
					"idMoa": 1,
					"idMoe": 1,
					"idServiceMoa": 1,
					"idServiceMoe": 1,
					"idValideur1": 1,
					"idValideur2": 1,
					"identifiantMoa": "i",
					"identifiantMoe": "i",
					"nomMoe": "n",
					"nomServiceMoa": "n",
					"nomServiceMoe": "n",
					"nomValideur1": "n",
					"nomValideur2": "n",
					"prenomMoe": "p",
					"prenomValideur1": "p",
					"prenomValideur2": "p",
					"raisonSocialeMoa": "r",
					"raisonSocialeMoe": "r",
					"raisonSocialeValideur1": "r",
					"raisonSocialeValideur2": "r",
					"typeIdentifiantMoa": "t",
					"typeIdentifiantMoe": "t",
					"typeIdentifiantValideur1": "t",
					"typeIdentifiantValideur2": "t"
				},
				"commentaire": "c",
				"destinataire": {
					"adresseDestinataireCodePays": "a",
					"adresseDestinataireCodePostal": "a",
					"adresseDestinataireComplement1": "a",
					"adresseDestinataireComplement2": "a",
					"adresseDestinataireDetail": "a",
					"adresseDestinataireId": 1,
					"adresseDestinataireLibellePays": "a",
					"adresseDestinataireVille": "a",
					"codeDestinataire": "c",
					"codeServiceExecutant": "c",
					"destinataireEtat": "d",
					"idDestinataire": 1,
					"idServiceExecutant": 1,
					"libelleDestinataire": "l",
					"libelleServiceExecutant": "l"
				},
				"fournisseur": {
					"adresseFournisseurCodePays": "a",
					"adresseFournisseurCodePostal": "a",
					"adresseFournisseurComplement1": "a",
					"adresseFournisseurComplement2": "a",
					"adresseFournisseurDetail": "a",
					"adresseFournisseurId": 1,
					"adresseFournisseurLibellePays": "a",
					"adresseFournisseurVille": "a",
					"affactureur": {
						"affactureurCode": "a",
						"affactureurCodePays": "a",
						"affactureurId": 1,
						"affactureurRaisonSociale": "a",
						"affactureurTypeIdentifiant": "a"
					},
					"codeFournisseur": "c",
					"codeServiceFournisseur": "c",
					"coordBancairesBicSwift": "c",
					"coordBancairesCodeGuichet": "c",
					"coordBancairesFournisseurCleIban": "c",
					"coordBancairesFournisseurCleRib": "c",
					"coordBancairesFournisseurCodeBanque": "c",
					"coordBancairesFournisseurCodePays": "c",
					"coordBancairesFournisseurCompteBancaire": "c",
					"coordBancairesFournisseurId": 1,
					"coordBancairesFournisseurLibelle": "c",
					"coordBancairesFournisseurType": "c",
					"designationFournisseur": "d",
					"idFournisseur": 1,
					"idServiceFournisseur": 1,
					"libelleServiceFournisseur": "l",
					"numeroRcsFournisseur": "n",
					"typeIdentifiantFournisseur": "t"
				},
				"identifiantFactureCPP": 1,
				"lignesDePoste": {
					"lignePoste": [
						{
							"lignePosteDenomination": "l",
							"lignePosteMontantHtApresRemise": 1,
							"lignePosteMontantRemiseHT": 1,
							"lignePosteMontantTtcApresRemise": 1,
							"lignePosteMontantTva": 1,
							"lignePosteMontantUnitaireHT": 1,
							"lignePosteNumero": 1,
							"lignePosteQuantite": 1,
							"lignePosteReference": "r",
							"lignePosteTauxTva": "t",
							"lignePosteTauxTvaManuel": 1,
							"lignePosteUniteCode": 1,
							"lignePosteUniteLibelle": "u"
						}
					],
					"nbResultatsParPageLignesPoste": 1,
					"pageCouranteLignesPoste": 1,
					"pagesLignesPoste": 1,
					"totalLignesPoste": 1
				},
				"listeDesPiecesJointes": {
					"pieceJointe": [
						{
							"pieceJointeDesignation": "p",
							"pieceJointeExtension": "e",
							"pieceJointeId": 1,
							"pieceJointeIdLiaison": 1,
							"pieceJointeNumeroLigneFacture": 1,
							"pieceJointeTypeCode": "t",
							"pieceJointeTypeLibelle": "t"
						}
					],
					"nbResultatsParPageListePiecesJointe": 1,
					"pageCouranteListePiecesJointe": 1,
					"pagesListePiecesJointe": 1,
					"totalListePiecesJointe": 1
				},
				"modeDepot": "m",
				"montantTotal": {
					"montantAPayer": 1,
					"montantHtTotal": 1,
					"montantRemiseGlobaleTTC": 1,
					"montantTVA": 1,
					"montantTtcAvantRemiseGlobalTTC": 1,
					"montantTtcTotal": 1,
					"motifRemiseGlobaleTTC": "m"
				},
				"numeroFacture": "n",
				"pieceJointePrincipale": {
					"idLiaisonPieceJointePrincipale": 1,
					"idPieceJointePrincipale": 1
				},
				"piecePrecedente": {
					"cadreFacturationPiecePrecedente": "c",
					"codeServiceExecutantPiecePrecedente": "c",
					"idDestinatairePiecePrecedente": 1,
					"idPiecePrecedente": 1,
					"idServiceExecutantPiecePrecedente": 1,
					"identifiantDestinatairePiecePrecedente": "i",
					"nomServiceExecutantPiecePrecedente": "n",
					"numeroPiecePrecedente": "n",
					"raisonSocialeDestinatairePiecePrecedente": "r"
				},
				"pieceSuivante": {
					"cadreFacturationPieceSuivante": "c",
					"idPieceSuivante": 1,
					"numeroPieceSuivante": "n"
				},
				"recapitulatifDeTva": {
					"ligneTva": [
						{
							"ligneTvaMontantBaseHtParTaux": 1,
							"ligneTvaMontantTvaParTaux": 1,
							"ligneTvaTauxManuel": 1,
							"ligneTvaTauxRefCode": "t",
							"ligneTvaTauxRefId": 1,
							"ligneTvaTauxRefLibelle": "t",
							"ligneTvaTauxRefValeur": 1
						}
					],
					"nbResultatsParPageLignesRecapitulativesTVA": 1,
					"pageCouranteLignesRecapitulativesTVA": 1,
					"pagesLignesRecapitulativesTVA": 1,
					"totalLignesRecapitulativesTVA": 1
				},
				"references": {
					"codeDeviseFacture": "c",
					"dateCreationFacture": ` + defaultDateTimeStr + `,
					"dateDepot": ` + defaultDateStr + `,
					"dateEcheancePaiement": ` + defaultDateStr + `,
					"dateFacture": ` + defaultDateStr + `,
					"libelleDeviseFacture": "l",
					"libelleMotifExonerationTva": "l",
					"modePaiement": "m",
					"motifExonerationTva": "m",
					"numeroBonCommande": "n",
					"numeroDpMandat": "n",
					"numeroFactureOrigine": "n",
					"numeroMarche": "n",
					"tailleTotalePiecesJointes": 1,
					"typeFacture": "t",
					"typeFactureTravaux": "t",
					"typeTva": "t"
				},
				"statutFacture": "s"
			},
			"libelle": "l"
		}
		`))
	})

	ctx := context.Background()
	opt := ConsulterFactureOptions{}
	got, err := client.Factures.ConsulterFactureParRecipiendaire(ctx, opt)

	if err != nil {
		t.Errorf("Factures.ConsulterFactureParRecipiendaire returned error : %v", err)
	}

	want := &ConsulterFactureResponse{
		CodeRetour: 1,
		Facture: Facture{
			CadreDeFacturation: CadreDeFacturation{
				Code:                     "c",
				CodeServiceMoa:           "c",
				CodeServiceMoe:           "c",
				CodeValideur1:            "c",
				CodeValideur2:            "c",
				DateValidation1:          &Date{defaultDate},
				DateValidation2:          &Date{defaultDate},
				DesignationMoa:           "d",
				DesignationMoe:           "d",
				DesignationValideur2:     "d",
				IdMoa:                    1,
				IdMoe:                    1,
				IdServiceMoa:             1,
				IdServiceMoe:             1,
				IdValideur1:              1,
				IdValideur2:              1,
				IdentifiantMoa:           "i",
				IdentifiantMoe:           "i",
				NomMoe:                   "n",
				NomServiceMoa:            "n",
				NomServiceMoe:            "n",
				NomValideur1:             "n",
				NomValideur2:             "n",
				PrenomMoe:                "p",
				PrenomValideur1:          "p",
				PrenomValideur2:          "p",
				RaisonSocialeMoa:         "r",
				RaisonSocialeMoe:         "r",
				RaisonSocialeValideur1:   "r",
				RaisonSocialeValideur2:   "r",
				TypeIdentifiantMoa:       "t",
				TypeIdentifiantMoe:       "t",
				TypeIdentifiantValideur1: "t",
				TypeIdentifiantValideur2: "t",
			},
			Commentaire: "c",
			Destinataire: Destinataire{
				AdresseCodePays:         "a",
				AdresseCodePostal:       "a",
				AdresseComplement1:      "a",
				AdresseComplement2:      "a",
				AdresseDetail:           "a",
				AdresseId:               1,
				AdresseLibellePays:      "a",
				AdresseVille:            "a",
				Code:                    "c",
				CodeServiceExecutant:    "c",
				Etat:                    "d",
				Id:                      1,
				IdServiceExecutant:      1,
				Libelle:                 "l",
				LibelleServiceExecutant: "l",
			},
			Fournisseur: Fournisseur{
				AdresseCodePays:    "a",
				AdresseCodePostal:  "a",
				AdresseComplement1: "a",
				AdresseComplement2: "a",
				AdresseDetail:      "a",
				AdresseId:          1,
				AdresseLibellePays: "a",
				AdresseVille:       "a",
				Affactureur: Affactureur{
					Code:            "a",
					CodePays:        "a",
					Id:              1,
					RaisonSociale:   "a",
					TypeIdentifiant: "a",
				},
				Code:                         "c",
				CodeService:                  "c",
				CoordBancairesBicSwift:       "c",
				CoordBancairesCodeGuichet:    "c",
				CoordBancairesCleIban:        "c",
				CoordBancairesCleRib:         "c",
				CoordBancairesCodeBanque:     "c",
				CoordBancairesCodePays:       "c",
				CoordBancairesCompteBancaire: "c",
				CoordBancairesId:             1,
				CoordBancairesLibelle:        "c",
				CoordBancairesType:           "c",
				Designation:                  "d",
				Id:                           1,
				IdService:                    1,
				LibelleService:               "l",
				NumeroRcs:                    "n",
				TypeIdentifiant:              "t",
			},
			IdentifiantFactureCPP: 1,
			LignesDePoste: &ListeLignesDePoste{
				Lignes: []LignePoste{
					{
						Denomination:          "l",
						MontantHtApresRemise:  1,
						MontantRemiseHT:       1,
						MontantTtcApresRemise: 1,
						MontantTva:            1,
						MontantUnitaireHT:     1,
						Numero:                1,
						Quantite:              1,
						Reference:             "r",
						TauxTva:               "t",
						TauxTvaManuel:         1,
						UniteCode:             1,
						UniteLibelle:          "u",
					},
				},
				NbResultatsParPage: 1,
				PageCourante:       1,
				Pages:              1,
				Total:              1,
			},
			ListeDesPiecesJointes: &ListePiecesJointesFacture{
				PiecesJointes: []PieceJointeFacture{
					{
						Designation:        "p",
						Extension:          "e",
						Id:                 1,
						IdLiaison:          1,
						NumeroLigneFacture: 1,
						TypeCode:           "t",
						TypeLibelle:        "t",
					},
				},
				NbResultatsParPage: 1,
				PageCourante:       1,
				Pages:              1,
				Total:              1,
			},
			ModeDepot: "m",
			MontantTotal: MontantTotal{
				MontantAPayer:                  1,
				MontantHtTotal:                 1,
				MontantRemiseGlobaleTTC:        1,
				MontantTVA:                     1,
				MontantTtcAvantRemiseGlobalTTC: 1,
				MontantTtcTotal:                1,
				MotifRemiseGlobaleTTC:          "m",
			},
			NumeroFacture: "n",
			PieceJointePrincipale: &PieceJointePrincipale{
				IdLiaison: 1,
				Id:        1,
			},
			PiecePrecedente: &PiecePrecedente{
				CadreFacturation:          "c",
				CodeServiceExecutant:      "c",
				IdDestinataire:            1,
				Id:                        1,
				IdServiceExecutant:        1,
				IdentifiantDestinataire:   "i",
				NomServiceExecutant:       "n",
				Numero:                    "n",
				RaisonSocialeDestinataire: "r",
			},
			PieceSuivante: &PieceSuivante{
				CadreFacturation: "c",
				Id:               1,
				Numero:           "n",
			},
			RecapitulatifDeTva: &ListeLignesRecapTVA{
				Lignes: []LigneTva{
					{
						MontantBaseHtParTaux: 1,
						MontantTvaParTaux:    1,
						TauxManuel:           1,
						TauxRefCode:          "t",
						TauxRefId:            1,
						TauxRefLibelle:       "t",
						TauxRefValeur:        1,
					},
				},
				NbResultatsParPage: 1,
				PageCourante:       1,
				Pages:              1,
				Total:              1,
			},
			References: References{
				CodeDevise:                 "c",
				DateCreationFacture:        &Datetime{defaultDate},
				DateDepot:                  &Date{defaultDate},
				DateEcheancePaiement:       &Date{defaultDate},
				DateFacture:                &Date{defaultDate},
				LibelleDevise:              "l",
				LibelleMotifExonerationTva: "l",
				ModePaiement:               "m",
				MotifExonerationTva:        "m",
				NumeroBonCommande:          "n",
				NumeroDpMandat:             "n",
				NumeroFactureOrigine:       "n",
				NumeroMarche:               "n",
				TailleTotalePiecesJointes:  1,
				TypeFacture:                "t",
				TypeFactureTravaux:         "t",
				TypeTva:                    "t",
			},
			StatutFacture: "s",
		},
		Libelle: "l",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Factures.ConsulterFactureParRecipiendaire returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "ConsulterFactureParRecipiendaire", client, func() error {
		_, err := client.Factures.ConsulterFactureParRecipiendaire(ctx, opt)
		return err
	})
}

func TestFacturesService_ConsulterFactureParValideur(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/factures/v1/consulter/valideur", func(w http.ResponseWriter, r *http.Request) {
		v := new(ConsulterFactureOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"facture": {
				"cadreDeFacturation": {
					"codeCadreFacturation": "c",
					"codeServiceMoa": "c",
					"codeServiceMoe": "c",
					"codeValideur1": "c",
					"codeValideur2": "c",
					"dateValidation1": ` + defaultDateStr + `,
					"dateValidation2": ` + defaultDateStr + `,
					"designationMoa": "d",
					"designationMoe": "d",
					"designationValideur2": "d",
					"idMoa": 1,
					"idMoe": 1,
					"idServiceMoa": 1,
					"idServiceMoe": 1,
					"idValideur1": 1,
					"idValideur2": 1,
					"identifiantMoa": "i",
					"identifiantMoe": "i",
					"nomMoe": "n",
					"nomServiceMoa": "n",
					"nomServiceMoe": "n",
					"nomValideur1": "n",
					"nomValideur2": "n",
					"prenomMoe": "p",
					"prenomValideur1": "p",
					"prenomValideur2": "p",
					"raisonSocialeMoa": "r",
					"raisonSocialeMoe": "r",
					"raisonSocialeValideur1": "r",
					"raisonSocialeValideur2": "r",
					"typeIdentifiantMoa": "t",
					"typeIdentifiantMoe": "t",
					"typeIdentifiantValideur1": "t",
					"typeIdentifiantValideur2": "t"
				},
				"commentaire": "c",
				"destinataire": {
					"adresseDestinataireCodePays": "a",
					"adresseDestinataireCodePostal": "a",
					"adresseDestinataireComplement1": "a",
					"adresseDestinataireComplement2": "a",
					"adresseDestinataireDetail": "a",
					"adresseDestinataireId": 1,
					"adresseDestinataireLibellePays": "a",
					"adresseDestinataireVille": "a",
					"codeDestinataire": "c",
					"codeServiceExecutant": "c",
					"destinataireEtat": "d",
					"idDestinataire": 1,
					"idServiceExecutant": 1,
					"libelleDestinataire": "l",
					"libelleServiceExecutant": "l"
				},
				"fournisseur": {
					"adresseFournisseurCodePays": "a",
					"adresseFournisseurCodePostal": "a",
					"adresseFournisseurComplement1": "a",
					"adresseFournisseurComplement2": "a",
					"adresseFournisseurDetail": "a",
					"adresseFournisseurId": 1,
					"adresseFournisseurLibellePays": "a",
					"adresseFournisseurVille": "a",
					"affactureur": {
						"affactureurCode": "a",
						"affactureurCodePays": "a",
						"affactureurId": 1,
						"affactureurRaisonSociale": "a",
						"affactureurTypeIdentifiant": "a"
					},
					"codeFournisseur": "c",
					"codeServiceFournisseur": "c",
					"coordBancairesBicSwift": "c",
					"coordBancairesCodeGuichet": "c",
					"coordBancairesFournisseurCleIban": "c",
					"coordBancairesFournisseurCleRib": "c",
					"coordBancairesFournisseurCodeBanque": "c",
					"coordBancairesFournisseurCodePays": "c",
					"coordBancairesFournisseurCompteBancaire": "c",
					"coordBancairesFournisseurId": 1,
					"coordBancairesFournisseurLibelle": "c",
					"coordBancairesFournisseurType": "c",
					"designationFournisseur": "d",
					"idFournisseur": 1,
					"idServiceFournisseur": 1,
					"libelleServiceFournisseur": "l",
					"numeroRcsFournisseur": "n",
					"typeIdentifiantFournisseur": "t"
				},
				"identifiantFactureCPP": 1,
				"lignesDePoste": {
					"lignePoste": [
						{
							"lignePosteDenomination": "l",
							"lignePosteMontantHtApresRemise": 1,
							"lignePosteMontantRemiseHT": 1,
							"lignePosteMontantTtcApresRemise": 1,
							"lignePosteMontantTva": 1,
							"lignePosteMontantUnitaireHT": 1,
							"lignePosteNumero": 1,
							"lignePosteQuantite": 1,
							"lignePosteReference": "r",
							"lignePosteTauxTva": "t",
							"lignePosteTauxTvaManuel": 1,
							"lignePosteUniteCode": 1,
							"lignePosteUniteLibelle": "u"
						}
					],
					"nbResultatsParPageLignesPoste": 1,
					"pageCouranteLignesPoste": 1,
					"pagesLignesPoste": 1,
					"totalLignesPoste": 1
				},
				"listeDesPiecesJointes": {
					"pieceJointe": [
						{
							"pieceJointeDesignation": "p",
							"pieceJointeExtension": "e",
							"pieceJointeId": 1,
							"pieceJointeIdLiaison": 1,
							"pieceJointeNumeroLigneFacture": 1,
							"pieceJointeTypeCode": "t",
							"pieceJointeTypeLibelle": "t"
						}
					],
					"nbResultatsParPageListePiecesJointe": 1,
					"pageCouranteListePiecesJointe": 1,
					"pagesListePiecesJointe": 1,
					"totalListePiecesJointe": 1
				},
				"modeDepot": "m",
				"montantTotal": {
					"montantAPayer": 1,
					"montantHtTotal": 1,
					"montantRemiseGlobaleTTC": 1,
					"montantTVA": 1,
					"montantTtcAvantRemiseGlobalTTC": 1,
					"montantTtcTotal": 1,
					"motifRemiseGlobaleTTC": "m"
				},
				"numeroFacture": "n",
				"pieceJointePrincipale": {
					"idLiaisonPieceJointePrincipale": 1,
					"idPieceJointePrincipale": 1
				},
				"piecePrecedente": {
					"cadreFacturationPiecePrecedente": "c",
					"codeServiceExecutantPiecePrecedente": "c",
					"idDestinatairePiecePrecedente": 1,
					"idPiecePrecedente": 1,
					"idServiceExecutantPiecePrecedente": 1,
					"identifiantDestinatairePiecePrecedente": "i",
					"nomServiceExecutantPiecePrecedente": "n",
					"numeroPiecePrecedente": "n",
					"raisonSocialeDestinatairePiecePrecedente": "r"
				},
				"pieceSuivante": {
					"cadreFacturationPieceSuivante": "c",
					"idPieceSuivante": 1,
					"numeroPieceSuivante": "n"
				},
				"recapitulatifDeTva": {
					"ligneTva": [
						{
							"ligneTvaMontantBaseHtParTaux": 1,
							"ligneTvaMontantTvaParTaux": 1,
							"ligneTvaTauxManuel": 1,
							"ligneTvaTauxRefCode": "t",
							"ligneTvaTauxRefId": 1,
							"ligneTvaTauxRefLibelle": "t",
							"ligneTvaTauxRefValeur": 1
						}
					],
					"nbResultatsParPageLignesRecapitulativesTVA": 1,
					"pageCouranteLignesRecapitulativesTVA": 1,
					"pagesLignesRecapitulativesTVA": 1,
					"totalLignesRecapitulativesTVA": 1
				},
				"references": {
					"codeDeviseFacture": "c",
					"dateCreationFacture": ` + defaultDateTimeStr + `,
					"dateDepot": ` + defaultDateStr + `,
					"dateEcheancePaiement": ` + defaultDateStr + `,
					"dateFacture": ` + defaultDateStr + `,
					"libelleDeviseFacture": "l",
					"libelleMotifExonerationTva": "l",
					"modePaiement": "m",
					"motifExonerationTva": "m",
					"numeroBonCommande": "n",
					"numeroDpMandat": "n",
					"numeroFactureOrigine": "n",
					"numeroMarche": "n",
					"tailleTotalePiecesJointes": 1,
					"typeFacture": "t",
					"typeFactureTravaux": "t",
					"typeTva": "t"
				},
				"statutFacture": "s"
			},
			"libelle": "l"
		}
		`))
	})

	ctx := context.Background()
	opt := ConsulterFactureOptions{}
	got, err := client.Factures.ConsulterFactureParValideur(ctx, opt)

	if err != nil {
		t.Errorf("Factures.ConsulterFactureParValideur returned error : %v", err)
	}

	want := &ConsulterFactureResponse{
		CodeRetour: 1,
		Facture: Facture{
			CadreDeFacturation: CadreDeFacturation{
				Code:                     "c",
				CodeServiceMoa:           "c",
				CodeServiceMoe:           "c",
				CodeValideur1:            "c",
				CodeValideur2:            "c",
				DateValidation1:          &Date{defaultDate},
				DateValidation2:          &Date{defaultDate},
				DesignationMoa:           "d",
				DesignationMoe:           "d",
				DesignationValideur2:     "d",
				IdMoa:                    1,
				IdMoe:                    1,
				IdServiceMoa:             1,
				IdServiceMoe:             1,
				IdValideur1:              1,
				IdValideur2:              1,
				IdentifiantMoa:           "i",
				IdentifiantMoe:           "i",
				NomMoe:                   "n",
				NomServiceMoa:            "n",
				NomServiceMoe:            "n",
				NomValideur1:             "n",
				NomValideur2:             "n",
				PrenomMoe:                "p",
				PrenomValideur1:          "p",
				PrenomValideur2:          "p",
				RaisonSocialeMoa:         "r",
				RaisonSocialeMoe:         "r",
				RaisonSocialeValideur1:   "r",
				RaisonSocialeValideur2:   "r",
				TypeIdentifiantMoa:       "t",
				TypeIdentifiantMoe:       "t",
				TypeIdentifiantValideur1: "t",
				TypeIdentifiantValideur2: "t",
			},
			Commentaire: "c",
			Destinataire: Destinataire{
				AdresseCodePays:         "a",
				AdresseCodePostal:       "a",
				AdresseComplement1:      "a",
				AdresseComplement2:      "a",
				AdresseDetail:           "a",
				AdresseId:               1,
				AdresseLibellePays:      "a",
				AdresseVille:            "a",
				Code:                    "c",
				CodeServiceExecutant:    "c",
				Etat:                    "d",
				Id:                      1,
				IdServiceExecutant:      1,
				Libelle:                 "l",
				LibelleServiceExecutant: "l",
			},
			Fournisseur: Fournisseur{
				AdresseCodePays:    "a",
				AdresseCodePostal:  "a",
				AdresseComplement1: "a",
				AdresseComplement2: "a",
				AdresseDetail:      "a",
				AdresseId:          1,
				AdresseLibellePays: "a",
				AdresseVille:       "a",
				Affactureur: Affactureur{
					Code:            "a",
					CodePays:        "a",
					Id:              1,
					RaisonSociale:   "a",
					TypeIdentifiant: "a",
				},
				Code:                         "c",
				CodeService:                  "c",
				CoordBancairesBicSwift:       "c",
				CoordBancairesCodeGuichet:    "c",
				CoordBancairesCleIban:        "c",
				CoordBancairesCleRib:         "c",
				CoordBancairesCodeBanque:     "c",
				CoordBancairesCodePays:       "c",
				CoordBancairesCompteBancaire: "c",
				CoordBancairesId:             1,
				CoordBancairesLibelle:        "c",
				CoordBancairesType:           "c",
				Designation:                  "d",
				Id:                           1,
				IdService:                    1,
				LibelleService:               "l",
				NumeroRcs:                    "n",
				TypeIdentifiant:              "t",
			},
			IdentifiantFactureCPP: 1,
			LignesDePoste: &ListeLignesDePoste{
				Lignes: []LignePoste{
					{
						Denomination:          "l",
						MontantHtApresRemise:  1,
						MontantRemiseHT:       1,
						MontantTtcApresRemise: 1,
						MontantTva:            1,
						MontantUnitaireHT:     1,
						Numero:                1,
						Quantite:              1,
						Reference:             "r",
						TauxTva:               "t",
						TauxTvaManuel:         1,
						UniteCode:             1,
						UniteLibelle:          "u",
					},
				},
				NbResultatsParPage: 1,
				PageCourante:       1,
				Pages:              1,
				Total:              1,
			},
			ListeDesPiecesJointes: &ListePiecesJointesFacture{
				PiecesJointes: []PieceJointeFacture{
					{
						Designation:        "p",
						Extension:          "e",
						Id:                 1,
						IdLiaison:          1,
						NumeroLigneFacture: 1,
						TypeCode:           "t",
						TypeLibelle:        "t",
					},
				},
				NbResultatsParPage: 1,
				PageCourante:       1,
				Pages:              1,
				Total:              1,
			},
			ModeDepot: "m",
			MontantTotal: MontantTotal{
				MontantAPayer:                  1,
				MontantHtTotal:                 1,
				MontantRemiseGlobaleTTC:        1,
				MontantTVA:                     1,
				MontantTtcAvantRemiseGlobalTTC: 1,
				MontantTtcTotal:                1,
				MotifRemiseGlobaleTTC:          "m",
			},
			NumeroFacture: "n",
			PieceJointePrincipale: &PieceJointePrincipale{
				IdLiaison: 1,
				Id:        1,
			},
			PiecePrecedente: &PiecePrecedente{
				CadreFacturation:          "c",
				CodeServiceExecutant:      "c",
				IdDestinataire:            1,
				Id:                        1,
				IdServiceExecutant:        1,
				IdentifiantDestinataire:   "i",
				NomServiceExecutant:       "n",
				Numero:                    "n",
				RaisonSocialeDestinataire: "r",
			},
			PieceSuivante: &PieceSuivante{
				CadreFacturation: "c",
				Id:               1,
				Numero:           "n",
			},
			RecapitulatifDeTva: &ListeLignesRecapTVA{
				Lignes: []LigneTva{
					{
						MontantBaseHtParTaux: 1,
						MontantTvaParTaux:    1,
						TauxManuel:           1,
						TauxRefCode:          "t",
						TauxRefId:            1,
						TauxRefLibelle:       "t",
						TauxRefValeur:        1,
					},
				},
				NbResultatsParPage: 1,
				PageCourante:       1,
				Pages:              1,
				Total:              1,
			},
			References: References{
				CodeDevise:                 "c",
				DateCreationFacture:        &Datetime{defaultDate},
				DateDepot:                  &Date{defaultDate},
				DateEcheancePaiement:       &Date{defaultDate},
				DateFacture:                &Date{defaultDate},
				LibelleDevise:              "l",
				LibelleMotifExonerationTva: "l",
				ModePaiement:               "m",
				MotifExonerationTva:        "m",
				NumeroBonCommande:          "n",
				NumeroDpMandat:             "n",
				NumeroFactureOrigine:       "n",
				NumeroMarche:               "n",
				TailleTotalePiecesJointes:  1,
				TypeFacture:                "t",
				TypeFactureTravaux:         "t",
				TypeTva:                    "t",
			},
			StatutFacture: "s",
		},
		Libelle: "l",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Factures.ConsulterFactureParValideur returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "ConsulterFactureParValideur", client, func() error {
		_, err := client.Factures.ConsulterFactureParValideur(ctx, opt)
		return err
	})
}

func TestFacturesService_ConsulterFactureParFournisseur(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/factures/v1/consulter/fournisseur", func(w http.ResponseWriter, r *http.Request) {
		v := new(ConsulterFactureFournisseurOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"facture": {
				"cadreDeFacturation": {
					"codeCadreFacturation": "c",
					"codeServiceMoa": "c",
					"codeServiceMoe": "c",
					"codeValideur1": "c",
					"codeValideur2": "c",
					"dateValidation1": ` + defaultDateStr + `,
					"dateValidation2": ` + defaultDateStr + `,
					"designationMoa": "d",
					"designationMoe": "d",
					"designationValideur2": "d",
					"idMoa": 1,
					"idMoe": 1,
					"idServiceMoa": 1,
					"idServiceMoe": 1,
					"idValideur1": 1,
					"idValideur2": 1,
					"identifiantMoa": "i",
					"identifiantMoe": "i",
					"nomMoe": "n",
					"nomServiceMoa": "n",
					"nomServiceMoe": "n",
					"nomValideur1": "n",
					"nomValideur2": "n",
					"prenomMoe": "p",
					"prenomValideur1": "p",
					"prenomValideur2": "p",
					"raisonSocialeMoa": "r",
					"raisonSocialeMoe": "r",
					"raisonSocialeValideur1": "r",
					"raisonSocialeValideur2": "r",
					"typeIdentifiantMoa": "t",
					"typeIdentifiantMoe": "t",
					"typeIdentifiantValideur1": "t",
					"typeIdentifiantValideur2": "t"
				},
				"commentaire": "c",
				"destinataire": {
					"adresseDestinataireCodePays": "a",
					"adresseDestinataireCodePostal": "a",
					"adresseDestinataireComplement1": "a",
					"adresseDestinataireComplement2": "a",
					"adresseDestinataireDetail": "a",
					"adresseDestinataireId": 1,
					"adresseDestinataireLibellePays": "a",
					"adresseDestinataireVille": "a",
					"codeDestinataire": "c",
					"codeServiceExecutant": "c",
					"destinataireEtat": "d",
					"idDestinataire": 1,
					"idServiceExecutant": 1,
					"libelleDestinataire": "l",
					"libelleServiceExecutant": "l"
				},
				"fournisseur": {
					"adresseFournisseurCodePays": "a",
					"adresseFournisseurCodePostal": "a",
					"adresseFournisseurComplement1": "a",
					"adresseFournisseurComplement2": "a",
					"adresseFournisseurDetail": "a",
					"adresseFournisseurId": 1,
					"adresseFournisseurLibellePays": "a",
					"adresseFournisseurVille": "a",
					"affactureur": {
						"affactureurCode": "a",
						"affactureurCodePays": "a",
						"affactureurId": 1,
						"affactureurRaisonSociale": "a",
						"affactureurTypeIdentifiant": "a"
					},
					"codeFournisseur": "c",
					"codeServiceFournisseur": "c",
					"coordBancairesBicSwift": "c",
					"coordBancairesCodeGuichet": "c",
					"coordBancairesFournisseurCleIban": "c",
					"coordBancairesFournisseurCleRib": "c",
					"coordBancairesFournisseurCodeBanque": "c",
					"coordBancairesFournisseurCodePays": "c",
					"coordBancairesFournisseurCompteBancaire": "c",
					"coordBancairesFournisseurId": 1,
					"coordBancairesFournisseurLibelle": "c",
					"coordBancairesFournisseurType": "c",
					"designationFournisseur": "d",
					"idFournisseur": 1,
					"idServiceFournisseur": 1,
					"libelleServiceFournisseur": "l",
					"numeroRcsFournisseur": "n",
					"typeIdentifiantFournisseur": "t"
				},
				"identifiantFactureCPP": 1,
				"lignesDePoste": {
					"lignePoste": [
						{
							"lignePosteDenomination": "l",
							"lignePosteMontantHtApresRemise": 1,
							"lignePosteMontantRemiseHT": 1,
							"lignePosteMontantTtcApresRemise": 1,
							"lignePosteMontantTva": 1,
							"lignePosteMontantUnitaireHT": 1,
							"lignePosteNumero": 1,
							"lignePosteQuantite": 1,
							"lignePosteReference": "r",
							"lignePosteTauxTva": "t",
							"lignePosteTauxTvaManuel": 1,
							"lignePosteUniteCode": 1,
							"lignePosteUniteLibelle": "u"
						}
					],
					"nbResultatsParPageLignesPoste": 1,
					"pageCouranteLignesPoste": 1,
					"pagesLignesPoste": 1,
					"totalLignesPoste": 1
				},
				"listeDesPiecesJointes": {
					"pieceJointe": [
						{
							"pieceJointeDesignation": "p",
							"pieceJointeExtension": "e",
							"pieceJointeId": 1,
							"pieceJointeIdLiaison": 1,
							"pieceJointeNumeroLigneFacture": 1,
							"pieceJointeTypeCode": "t",
							"pieceJointeTypeLibelle": "t"
						}
					],
					"nbResultatsParPageListePiecesJointe": 1,
					"pageCouranteListePiecesJointe": 1,
					"pagesListePiecesJointe": 1,
					"totalListePiecesJointe": 1
				},
				"modeDepot": "m",
				"montantTotal": {
					"montantAPayer": 1,
					"montantHtTotal": 1,
					"montantRemiseGlobaleTTC": 1,
					"montantTVA": 1,
					"montantTtcAvantRemiseGlobalTTC": 1,
					"montantTtcTotal": 1,
					"motifRemiseGlobaleTTC": "m"
				},
				"numeroFacture": "n",
				"pieceJointePrincipale": {
					"idLiaisonPieceJointePrincipale": 1,
					"idPieceJointePrincipale": 1
				},
				"piecePrecedente": {
					"cadreFacturationPiecePrecedente": "c",
					"codeServiceExecutantPiecePrecedente": "c",
					"idDestinatairePiecePrecedente": 1,
					"idPiecePrecedente": 1,
					"idServiceExecutantPiecePrecedente": 1,
					"identifiantDestinatairePiecePrecedente": "i",
					"nomServiceExecutantPiecePrecedente": "n",
					"numeroPiecePrecedente": "n",
					"raisonSocialeDestinatairePiecePrecedente": "r"
				},
				"pieceSuivante": {
					"cadreFacturationPieceSuivante": "c",
					"idPieceSuivante": 1,
					"numeroPieceSuivante": "n"
				},
				"recapitulatifDeTva": {
					"ligneTva": [
						{
							"ligneTvaMontantBaseHtParTaux": 1,
							"ligneTvaMontantTvaParTaux": 1,
							"ligneTvaTauxManuel": 1,
							"ligneTvaTauxRefCode": "t",
							"ligneTvaTauxRefId": 1,
							"ligneTvaTauxRefLibelle": "t",
							"ligneTvaTauxRefValeur": 1
						}
					],
					"nbResultatsParPageLignesRecapitulativesTVA": 1,
					"pageCouranteLignesRecapitulativesTVA": 1,
					"pagesLignesRecapitulativesTVA": 1,
					"totalLignesRecapitulativesTVA": 1
				},
				"references": {
					"codeDeviseFacture": "c",
					"dateCreationFacture": ` + defaultDateTimeStr + `,
					"dateDepot": ` + defaultDateStr + `,
					"dateEcheancePaiement": ` + defaultDateStr + `,
					"dateFacture": ` + defaultDateStr + `,
					"libelleDeviseFacture": "l",
					"libelleMotifExonerationTva": "l",
					"modePaiement": "m",
					"motifExonerationTva": "m",
					"numeroBonCommande": "n",
					"numeroDpMandat": "n",
					"numeroFactureOrigine": "n",
					"numeroMarche": "n",
					"tailleTotalePiecesJointes": 1,
					"typeFacture": "t",
					"typeFactureTravaux": "t",
					"typeTva": "t"
				},
				"statutFacture": "s"
			},
			"libelle": "l"
		}
		`))
	})

	ctx := context.Background()
	opt := ConsulterFactureFournisseurOptions{}
	got, err := client.Factures.ConsulterFactureParFournisseur(ctx, opt)

	if err != nil {
		t.Errorf("Factures.ConsulterFactureParFournisseur returned error : %v", err)
	}

	want := &ConsulterFactureResponse{
		CodeRetour: 1,
		Facture: Facture{
			CadreDeFacturation: CadreDeFacturation{
				Code:                     "c",
				CodeServiceMoa:           "c",
				CodeServiceMoe:           "c",
				CodeValideur1:            "c",
				CodeValideur2:            "c",
				DateValidation1:          &Date{defaultDate},
				DateValidation2:          &Date{defaultDate},
				DesignationMoa:           "d",
				DesignationMoe:           "d",
				DesignationValideur2:     "d",
				IdMoa:                    1,
				IdMoe:                    1,
				IdServiceMoa:             1,
				IdServiceMoe:             1,
				IdValideur1:              1,
				IdValideur2:              1,
				IdentifiantMoa:           "i",
				IdentifiantMoe:           "i",
				NomMoe:                   "n",
				NomServiceMoa:            "n",
				NomServiceMoe:            "n",
				NomValideur1:             "n",
				NomValideur2:             "n",
				PrenomMoe:                "p",
				PrenomValideur1:          "p",
				PrenomValideur2:          "p",
				RaisonSocialeMoa:         "r",
				RaisonSocialeMoe:         "r",
				RaisonSocialeValideur1:   "r",
				RaisonSocialeValideur2:   "r",
				TypeIdentifiantMoa:       "t",
				TypeIdentifiantMoe:       "t",
				TypeIdentifiantValideur1: "t",
				TypeIdentifiantValideur2: "t",
			},
			Commentaire: "c",
			Destinataire: Destinataire{
				AdresseCodePays:         "a",
				AdresseCodePostal:       "a",
				AdresseComplement1:      "a",
				AdresseComplement2:      "a",
				AdresseDetail:           "a",
				AdresseId:               1,
				AdresseLibellePays:      "a",
				AdresseVille:            "a",
				Code:                    "c",
				CodeServiceExecutant:    "c",
				Etat:                    "d",
				Id:                      1,
				IdServiceExecutant:      1,
				Libelle:                 "l",
				LibelleServiceExecutant: "l",
			},
			Fournisseur: Fournisseur{
				AdresseCodePays:    "a",
				AdresseCodePostal:  "a",
				AdresseComplement1: "a",
				AdresseComplement2: "a",
				AdresseDetail:      "a",
				AdresseId:          1,
				AdresseLibellePays: "a",
				AdresseVille:       "a",
				Affactureur: Affactureur{
					Code:            "a",
					CodePays:        "a",
					Id:              1,
					RaisonSociale:   "a",
					TypeIdentifiant: "a",
				},
				Code:                         "c",
				CodeService:                  "c",
				CoordBancairesBicSwift:       "c",
				CoordBancairesCodeGuichet:    "c",
				CoordBancairesCleIban:        "c",
				CoordBancairesCleRib:         "c",
				CoordBancairesCodeBanque:     "c",
				CoordBancairesCodePays:       "c",
				CoordBancairesCompteBancaire: "c",
				CoordBancairesId:             1,
				CoordBancairesLibelle:        "c",
				CoordBancairesType:           "c",
				Designation:                  "d",
				Id:                           1,
				IdService:                    1,
				LibelleService:               "l",
				NumeroRcs:                    "n",
				TypeIdentifiant:              "t",
			},
			IdentifiantFactureCPP: 1,
			LignesDePoste: &ListeLignesDePoste{
				Lignes: []LignePoste{
					{
						Denomination:          "l",
						MontantHtApresRemise:  1,
						MontantRemiseHT:       1,
						MontantTtcApresRemise: 1,
						MontantTva:            1,
						MontantUnitaireHT:     1,
						Numero:                1,
						Quantite:              1,
						Reference:             "r",
						TauxTva:               "t",
						TauxTvaManuel:         1,
						UniteCode:             1,
						UniteLibelle:          "u",
					},
				},
				NbResultatsParPage: 1,
				PageCourante:       1,
				Pages:              1,
				Total:              1,
			},
			ListeDesPiecesJointes: &ListePiecesJointesFacture{
				PiecesJointes: []PieceJointeFacture{
					{
						Designation:        "p",
						Extension:          "e",
						Id:                 1,
						IdLiaison:          1,
						NumeroLigneFacture: 1,
						TypeCode:           "t",
						TypeLibelle:        "t",
					},
				},
				NbResultatsParPage: 1,
				PageCourante:       1,
				Pages:              1,
				Total:              1,
			},
			ModeDepot: "m",
			MontantTotal: MontantTotal{
				MontantAPayer:                  1,
				MontantHtTotal:                 1,
				MontantRemiseGlobaleTTC:        1,
				MontantTVA:                     1,
				MontantTtcAvantRemiseGlobalTTC: 1,
				MontantTtcTotal:                1,
				MotifRemiseGlobaleTTC:          "m",
			},
			NumeroFacture: "n",
			PieceJointePrincipale: &PieceJointePrincipale{
				IdLiaison: 1,
				Id:        1,
			},
			PiecePrecedente: &PiecePrecedente{
				CadreFacturation:          "c",
				CodeServiceExecutant:      "c",
				IdDestinataire:            1,
				Id:                        1,
				IdServiceExecutant:        1,
				IdentifiantDestinataire:   "i",
				NomServiceExecutant:       "n",
				Numero:                    "n",
				RaisonSocialeDestinataire: "r",
			},
			PieceSuivante: &PieceSuivante{
				CadreFacturation: "c",
				Id:               1,
				Numero:           "n",
			},
			RecapitulatifDeTva: &ListeLignesRecapTVA{
				Lignes: []LigneTva{
					{
						MontantBaseHtParTaux: 1,
						MontantTvaParTaux:    1,
						TauxManuel:           1,
						TauxRefCode:          "t",
						TauxRefId:            1,
						TauxRefLibelle:       "t",
						TauxRefValeur:        1,
					},
				},
				NbResultatsParPage: 1,
				PageCourante:       1,
				Pages:              1,
				Total:              1,
			},
			References: References{
				CodeDevise:                 "c",
				DateCreationFacture:        &Datetime{defaultDate},
				DateDepot:                  &Date{defaultDate},
				DateEcheancePaiement:       &Date{defaultDate},
				DateFacture:                &Date{defaultDate},
				LibelleDevise:              "l",
				LibelleMotifExonerationTva: "l",
				ModePaiement:               "m",
				MotifExonerationTva:        "m",
				NumeroBonCommande:          "n",
				NumeroDpMandat:             "n",
				NumeroFactureOrigine:       "n",
				NumeroMarche:               "n",
				TailleTotalePiecesJointes:  1,
				TypeFacture:                "t",
				TypeFactureTravaux:         "t",
				TypeTva:                    "t",
			},
			StatutFacture: "s",
		},
		Libelle: "l",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Factures.ConsulterFactureParFournisseur returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "ConsulterFactureParFournisseur", client, func() error {
		_, err := client.Factures.ConsulterFactureParFournisseur(ctx, opt)
		return err
	})
}

func TestFacturesService_ConsulterHistorique(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/factures/v1/consulter/historique", func(w http.ResponseWriter, r *http.Request) {
		v := new(ConsulterHistoriqueOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"derniereAction": {
				"derniereActionCode": "c",
				"derniereActionId": 1
			},
			"historiquesDesActionsUtilisateurs": {
				"histoAction": [
					{
						"histoActionCode": "c",
						"histoActionDateHeure": ` + defaultDateTimeStr + `,
						"histoActionId": 1,
						"histoActionUtilisateurEmail": "e",
						"histoActionUtilisateurNom": "n",
						"histoActionUtilisateurPrenom": "p",
						"histoActionUtilisateurTelephone": "t"
					}
				],
				"nbResultatsParPageHistoAction": 1,
				"pageCouranteHistoAction": 1,
				"pagesHistoAction": 1,
				"totalHistoAction": 1
			},
			"historiquesDesEvenementsComplementaires": {
				"evenement": [
					{
						"evenementDateHeure": ` + defaultDateTimeStr + `,
						"evenementId": 1,
						"evenementLibelle": "l",
						"evenementQui": "q"
					}
				],
				"nbResultatsParPageEvenement": 1,
				"pageCouranteEvenement": 1,
				"pagesEvenement": 1,
				"totalEvenement": 1
			},
			"historiquesDesStatuts": {
				"histoStatut": [
					{
						"histoStatutCode": "c",
						"histoStatutCommentaire": "c",
						"histoStatutDatePassage": ` + defaultDateTimeStr + `,
						"histoStatutId": 1,
						"histoStatutUtilisateurEmail": "e",
						"histoStatutUtilisateurNom": "n",
						"histoStatutUtilisateurPrenom": "p",
						"histoStatutUtilisateurTelephone": "t"
					}
				],
				"nbResultatsParPageHistoStatut": 1,
				"pageCouranteHistoStatut": 1,
				"pagesHistoStatut": 1,
				"totalHistoStatut": 1
			},
			"idFacture": 1,
			"libelle": "l",
			"modeDepot": "m",
			"numeroFacture": "n",
			"statutCourantCode": "s"
		}
		`))
	})

	ctx := context.Background()
	opt := ConsulterHistoriqueOptions{}
	got, err := client.Factures.ConsulterHistorique(ctx, opt)

	if err != nil {
		t.Errorf("Factures.ConsulterHistorique returned error : %v", err)
	}

	want := &ConsulterHistoriqueResponse{
		CodeRetour: 1,
		DerniereAction: &DerniereAction{
			Code: "c",
			Id:   1,
		},
		HistoriquesDesActionsUtilisateurs: &HistoriquesDesActionsUtilisateurs{
			HistoAction: []HistoAction{
				{
					Code:                 "c",
					DateHeure:            &Datetime{defaultDate},
					Id:                   1,
					UtilisateurEmail:     "e",
					UtilisateurNom:       "n",
					UtilisateurPrenom:    "p",
					UtilisateurTelephone: "t",
				},
			},
			NbResultatsParPageHistoAction: 1,
			PageCouranteHistoAction:       1,
			PagesHistoAction:              1,
			TotalHistoAction:              1,
		},
		HistoriquesDesEvenementsComplementaires: &HistoriquesDesEvenementsComplementaires{
			Evenement: []Evenement{
				{
					EvenementDateHeure: &Datetime{defaultDate},
					EvenementId:        1,
					EvenementLibelle:   "l",
					EvenementQui:       "q",
				},
			},

			NbResultatsParPageEvenement: 1,
			PageCouranteEvenement:       1,
			PagesEvenement:              1,
			TotalEvenement:              1,
		},
		HistoriquesDesStatuts: &HistoriquesDesStatuts{
			HistoStatut: []HistoStatut{
				{
					Code:                 "c",
					Commentaire:          "c",
					DatePassage:          &Datetime{defaultDate},
					Id:                   1,
					UtilisateurEmail:     "e",
					UtilisateurNom:       "n",
					UtilisateurPrenom:    "p",
					UtilisateurTelephone: "t",
				},
			},
			NbResultatsParPageHistoStatut: 1,
			PageCouranteHistoStatut:       1,
			PagesHistoStatut:              1,
			TotalHistoStatut:              1,
		},
		IdFacture:         1,
		Libelle:           "l",
		ModeDepot:         "m",
		NumeroFacture:     "n",
		StatutCourantCode: "s",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Factures.ConsulterHistorique returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "ConsulterHistorique", client, func() error {
		_, err := client.Factures.ConsulterHistorique(ctx, opt)
		return err
	})
}
