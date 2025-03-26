// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep_test

import (
	"context"
	"os"
	"testing"

	"github.com/brennoo/findcep-go"
	"github.com/brennoo/findcep-go/internal/testutil"
	"github.com/brennoo/findcep-go/option"
)

func TestUsage(t *testing.T) {
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
	endereco, err := client.Enderecos.Pesquisa(context.TODO(), findcep.EnderecoPesquisaParams{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", endereco.Hits)
}
