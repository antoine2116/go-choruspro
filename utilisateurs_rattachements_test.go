package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestUtilisateursService_RecupererRattachementsUtilisateur(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/utilisateurs/v1/monCompte/recuperer/rattachements", func(w http.ResponseWriter, r *http.Request) {
		v := new(ListeRattachementsUtilisateurOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"libelle": "l",
			"listeStructureRattachement": [
				{
					"designationStructure": "s",
					"idRattachementStructure": 1,
					"idStructure": 1,
					"identifiantStructure": "i",
					"listeServicesRattache": [
						{
							"codeService": "s",
							"dateDbtService": ` + defaultDateTimeStr + `,
							"dateFinService": ` + defaultDateTimeStr + `,
							"estActif": true,
							"idRattachementService": 1,
							"idService": 1,
							"libelleService": "l",
							"statutRattachementService": "a",
							"statutService": "e"
						}
					],
					"roleUtilisateur": "u",
					"statutRattachementStructure": "a",
					"statutUtilisateur": "a",
					"typeIdentifiantStructure": "t"
				}
			],
			"parametresRetour": {
				"pageCourante": 1,
				"pages": 1,
				"nbResultatsParPage": 1,
				"total": 1
			}
		}`))
	})

	ctx := context.Background()
	opt := ListeRattachementsUtilisateurOptions{}
	got, err := client.Utilisateurs.RecupererRattachementsUtilisateur(ctx, opt)

	if err != nil {
		t.Errorf("Utilisateurs.RecupererRattachementsUtilisateur returned error : %v", err)
	}

	want := &ListeRattachementsUtilisateurResponse{
		CodeRetour: 1,
		Libelle:    "l",
		Rattachements: []RattachementStructure{
			{
				DesignationStructure:    "s",
				IdRattachementStructure: 1,
				IdStructure:             1,
				IdentifiantStructure:    "i",
				ServicesRattaches: []RattachementService{
					{
						CodeService:               "s",
						DateDbtService:            &Datetime{defaultDate},
						DateFinService:            &Datetime{defaultDate},
						EstActif:                  true,
						IdRattachementService:     1,
						IdService:                 1,
						LibelleService:            "l",
						StatutRattachementService: "a",
						StatutService:             "e",
					},
				},

				RoleUtilisateur:             "u",
				StatutRattachementStructure: "a",
				StatutUtilisateur:           "a",
				TypeIdentifiantStructure:    "t",
			},
		},
		Pagination: &PaginationResponse{
			PageCourante:       1,
			Pages:              1,
			NbResultatsParPage: 1,
			Total:              1,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Utilisateurs.RecupererRattachementsUtilisateur returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererRattachementsUtilisateur", client, func() error {
		_, err := client.Utilisateurs.RecupererRattachementsUtilisateur(ctx, opt)
		return err
	})
}
