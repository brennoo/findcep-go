// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/brennoo/findcep-go"
	"github.com/brennoo/findcep-go/internal/testutil"
	"github.com/brennoo/findcep-go/option"
)

func TestEnderecoTemplateCodigoIbgeNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Enderecos.Templates.CodigoIbge.New(context.TODO(), findcep.EnderecoTemplateCodigoIbgeNewParams{
		ID: findcep.F("pesquisa_endereco_por_codigo_ibge"),
		Params: findcep.F(findcep.EnderecoTemplateCodigoIbgeNewParamsParams{
			CodigoIbge:  findcep.F("codigo_ibge"),
			From:        findcep.F(int64(0)),
			QueryString: findcep.F("avenida pacaembu"),
			Size:        findcep.F(int64(10)),
		}),
	})
	if err != nil {
		var apierr *findcep.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
