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

func TestEnderecoTemplateUfNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Enderecos.Templates.Uf.New(context.TODO(), findcep.EnderecoTemplateUfNewParams{
		ID: findcep.F("pesquisa_endereco_por_uf"),
		Params: findcep.F(findcep.EnderecoTemplateUfNewParamsParams{
			From:        findcep.F(int64(0)),
			QueryString: findcep.F("Av Prof Gioia Martins"),
			Size:        findcep.F(int64(10)),
			Uf:          findcep.F("sp"),
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
