// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/findcep-go"
	"github.com/stainless-sdks/findcep-go/internal/testutil"
	"github.com/stainless-sdks/findcep-go/option"
)

func TestLocalidadeEstadoCidadeList(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := findcep.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Localidades.Estado.Cidades.List(context.TODO(), "sp")
	if err != nil {
		var apierr *findcep.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
