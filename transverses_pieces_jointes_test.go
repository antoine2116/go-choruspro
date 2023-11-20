package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestTransversesService_RecupererTypesPieceJointe(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/recuperer/typespj", func(w http.ResponseWriter, r *http.Request) {
		v := new(ListeTypesPieceJointeOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "TRA_MSG_00.000",
			"listeTypePieceJointe": [
				{
					"codeTypePieceJointe": "c1",
					"libelleTypePieceJointe": "l1",
					"idTypePieceJointe": 1,
					"typeObjet": "t1"
				},
				{
					"codeTypePieceJointe": "c2",
					"libelleTypePieceJointe": "l2",
					"idTypePieceJointe": 2,
					"typeObjet": "t2"
				}
			]
		}`))
	})

	ctx := context.Background()
	opt := ListeTypesPieceJointeOptions{}
	got, err := client.Transverses.RecupererTypesPieceJointe(ctx, opt)

	if err != nil {
		t.Errorf("Transverses.RecupererTypesPieceJointe returned error : %v", err)
	}

	want := &ListeTypesPieceJointeResponse{
		CodeRetour: 0,
		Libelle:    "TRA_MSG_00.000",
		Types: []TypePieceJointe{
			{
				Code:      "c1",
				Libelle:   "l1",
				Id:        1,
				TypeObjet: "t1",
			},
			{
				Code:      "c2",
				Libelle:   "l2",
				Id:        2,
				TypeObjet: "t2",
			},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RecupererTypesPieceJointe returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererTypesPieceJointe", client, func() error {
		_, err := client.Transverses.RecupererTypesPieceJointe(ctx, opt)
		return err
	})
}

func TestTransversesService_AjouterPieceJointe(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/ajouter/fichier", func(w http.ResponseWriter, r *http.Request) {
		v := new(AjouterPieceOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour" : 0,
			"libelle" : "GCU_MSG_01_000",
			"pieceJointeId" : 1
		}`))
	})

	ctx := context.Background()
	opt := AjouterPieceOptions{
		Extension: "e",
		Fichier:   "f",
		Nom:       "n",
		TypeMime:  "t",
	}
	got, err := client.Transverses.AjouterPieceJointe(ctx, opt)

	if err != nil {
		t.Errorf("Transverses.AjouterPieceJointe returned error : %v", err)
	}

	want := &AjouterPieceResponse{
		CodeRetour:    0,
		Libelle:       "GCU_MSG_01_000",
		IdPieceJointe: 1,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.AjouterPieceJointe returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "AjouterPieceJointe", client, func() error {
		_, err := client.Transverses.AjouterPieceJointe(ctx, opt)
		return err
	})
}

func TestTransversesService_RechercherPiecesJointesStructure(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/rechercher/pj/structure", func(w http.ResponseWriter, r *http.Request) {
		v := new(ListePiecesJointesStructureOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "GCU_MSG_01_000",
			"listePiecesJointes": [
				{
					"pieceJointeDesignation": "d",
					"pieceJointeExtension": "e",
					"pieceJointeId": 1,
					"pieceJointeObjetId": 1,
					"pieceJointeStructureCode": "sc",
					"pieceJointeStructureId": 1,
					"pieceJointeStructureRaisonSociale": "srs",
					"pieceJointeTypeCode": "c",
					"pieceJointeTypeLibelle": "c",
					"pieceJointeUtilisateurId": 0
				}
			],
			"parametresRetour": {
				"nbResultatsParPage": 10,
				"pageCourante": 1,
				"pages": 1,
				"total": 1
			}
		}`))
	})

	ctx := context.Background()
	opt := ListePiecesJointesStructureOptions{IdStructureCPP: 1}
	got, err := client.Transverses.RechercherPiecesJointesStructure(ctx, opt)

	if err != nil {
		t.Errorf("Transverses.RechercherPiecesJointesStructure returned error : %v", err)
	}

	want := &ListePiecesJointesStructureResponse{
		CodeRetour: 0,
		Libelle:    "GCU_MSG_01_000",
		PiecesJointes: []PieceJointeStructure{
			{
				Designation:            "d",
				Extension:              "e",
				Id:                     1,
				ObjetId:                1,
				StructureCode:          "sc",
				StructureId:            1,
				StructureRaisonSociale: "srs",
				TypeCode:               "c",
				TypeLibelle:            "c",
				IdUtilisateur:          0,
			},
		},
		Pagination: &PaginationResponse{
			NbResultatsParPage: 10,
			PageCourante:       1,
			Pages:              1,
			Total:              1,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RechercherPiecesJointesStructure returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RechercherPiecesJointesStructure", client, func() error {
		_, err := client.Transverses.RechercherPiecesJointesStructure(ctx, opt)
		return err
	})
}

func TestTransversesService_RechercherPiecesJointesMonCompte(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/rechercher/pj/moncompte", func(w http.ResponseWriter, r *http.Request) {
		v := new(ListePiecesJointesStructureOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 0,
			"libelle": "GCU_MSG_01_000",
			"listePiecesJointes": [
				{
					"pieceJointeDesignation": "d",
					"pieceJointeExtension": "e",
					"pieceJointeId": 1,
					"pieceJointeObjetId": 1,
					"pieceJointeTypeCode": "c",
					"pieceJointeTypeLibelle": "c",
					"pieceJointeUtilisateurId": 1
				}
			],
			"parametresRetour": {
				"nbResultatsParPage": 10,
				"pageCourante": 1,
				"pages": 1,
				"total": 1
			}
		}`))
	})

	ctx := context.Background()
	opt := ListePiecesJointesMonCompteOptions{PieceJointeTypeId: 1}
	got, err := client.Transverses.RechercherPiecesJointesMonCompte(ctx, opt)

	if err != nil {
		t.Errorf("Transverses.RechercherPiecesJointesMonCompte returned error : %v", err)
	}

	want := &ListePiecesJointesMonCompteResponse{
		CodeRetour: 0,
		Libelle:    "GCU_MSG_01_000",
		PiecesJointes: []PieceJointeMonCompte{
			{
				Designation:   "d",
				Extension:     "e",
				Id:            1,
				ObjetId:       1,
				TypeCode:      "c",
				TypeLibelle:   "c",
				IdUtilisateur: 1,
			},
		},
		Pagination: &PaginationResponse{
			NbResultatsParPage: 10,
			PageCourante:       1,
			Pages:              1,
			Total:              1,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.RechercherPiecesJointesMonCompte returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RechercherPiecesJointesMonCompte", client, func() error {
		_, err := client.Transverses.RechercherPiecesJointesMonCompte(ctx, opt)
		return err
	})
}

func TestTransversesService_TelechargerPieceJointe(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/telecharger/pieceJointe", func(w http.ResponseWriter, r *http.Request) {
		v := new(TelechargerPieceJointeOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"pieceJointe": "dGVzdA=="
		}`))
	})

	ctx := context.Background()
	opt := TelechargerPieceJointeOptions{
		Objet:               ObjetPJEngagementJuridique,
		EngagementJuridique: &TelechargerPieceJointeEngagementJuridiqueOptions{"n"},
	}
	got, err := client.Transverses.TelechargerPieceJointe(ctx, opt)

	if err != nil {
		t.Errorf("Transverses.TelechargerPieceJointe returned error : %v", err)
	}

	want := &PieceJointe{"dGVzdA=="}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.TelechargerPieceJointe returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "TelechargerPieceJointe", client, func() error {
		_, err := client.Transverses.TelechargerPieceJointe(ctx, opt)
		return err
	})
}
