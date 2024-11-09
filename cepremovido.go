// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/brennoo/findcep-go/internal/requestconfig"
	"github.com/brennoo/findcep-go/option"
)

// CepRemovidoService contains methods and other services that help with
// interacting with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCepRemovidoService] method instead.
type CepRemovidoService struct {
	Options []option.RequestOption
}

// NewCepRemovidoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCepRemovidoService(opts ...option.RequestOption) (r *CepRemovidoService) {
	r = &CepRemovidoService{}
	r.Options = opts
	return
}

// Returns a map with address. Lista de ceps removidos. [11680000, 35570000,
// 62500000, 95180000, 13480637, 90040440, 77020902]
func (r *CepRemovidoService) Get(ctx context.Context, cepRemovido string, opts ...option.RequestOption) (res *Cep, err error) {
	opts = append(r.Options[:], opts...)
	if cepRemovido == "" {
		err = errors.New("missing required cep_removido parameter")
		return
	}
	path := fmt.Sprintf("v1/cep/removido/%s.json", cepRemovido)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
