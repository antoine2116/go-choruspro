package choruspro

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cpro/transverses/v1/health-check", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"statusCode":"OK"}`)
	})

	ctx := context.Background()
	got, err := client.Transverses.HealthCheck(ctx)
	if err != nil {
		t.Errorf("Transverses.HealthCheck returned error : %v", err)
	}

	want := HealthCheck{StatusCode: "200"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transverses.HealthCheck returned %+v, want %+v", got, want)
	}
}
