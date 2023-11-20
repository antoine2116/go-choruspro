package choruspro

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestUtilisateursService_RecupererRattachementsServicesUtilisateur(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/structures/v1/lister/rattachementsutilisateur", func(w http.ResponseWriter, r *http.Request) {
		v := new(ListeRattachementsServicesUtilisateurOptions)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))
		testMethod(t, r, http.MethodPost)
		w.Write([]byte(`{
			"codeRetour": 1,
			"libelle": "l",
			"designationStructure": "d",
			"idStructure": 1,
			"identifiantStructure": "i",
			"premierService": true,
			"listeServices": [
				{
					"codeService": "c",
					"commentaire": "c",
					"dateDebut": ` + defaultISODateTimeStr + `,
					"dateFin":` + defaultISODateTimeStr + `,
					"idRattachement": 1,
					"idService": 1,
					"nomService": "n",
					"statutRattachement": "s"
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
	opt := ListeRattachementsServicesUtilisateurOptions{}
	got, err := client.Structures.RecupererRattachementsServicesUtilisateur(ctx, opt)

	if err != nil {
		t.Errorf("Structures.RecupererRattachementsServicesUtilisateur returned error : %v", err)
	}

	want := &ListeRattachementsServicesUtilisateurResponse{
		CodeRetour:           1,
		Libelle:              "l",
		DesignationStructure: "d",
		IdStructure:          1,
		IdentifiantStructure: "i",
		PremierService:       true,
		Rattachements: []RattachementServiceUtilisateur{
			{
				CodeService:        "c",
				Commentaire:        "c",
				DateDebut:          &defaultDate,
				DateFin:            &defaultDate,
				IdRattachement:     1,
				IdService:          1,
				NomService:         "n",
				StatutRattachement: "s",
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
		t.Errorf("Structures.RecupererRattachementsServicesUtilisateur returned %+v, want %+v", got, want)
	}

	testNewRequestAndDoRequestFailure(t, "RecupererRattachementsServicesUtilisateur", client, func() error {
		_, err := client.Structures.RecupererRattachementsServicesUtilisateur(ctx, opt)
		return err
	})
}
