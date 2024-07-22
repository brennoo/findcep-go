// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/findcep-go/internal/requestconfig"
	"github.com/stainless-sdks/findcep-go/option"
)

// LocalidadeCidadeIbgeService contains methods and other services that help with
// interacting with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLocalidadeCidadeIbgeService] method instead.
type LocalidadeCidadeIbgeService struct {
	Options []option.RequestOption
}

// NewLocalidadeCidadeIbgeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLocalidadeCidadeIbgeService(opts ...option.RequestOption) (r *LocalidadeCidadeIbgeService) {
	r = &LocalidadeCidadeIbgeService{}
	r.Options = opts
	return
}

// Returns a IBGE code id
func (r *LocalidadeCidadeIbgeService) Get(ctx context.Context, cidadeHash string, opts ...option.RequestOption) (res *Ibge, err error) {
	opts = append(r.Options[:], opts...)
	if cidadeHash == "" {
		err = errors.New("missing required cidade_hash parameter")
		return
	}
	path := fmt.Sprintf("v1/localidades/cidade/%s/ibge", cidadeHash)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
