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

// LocalidadeCodigoMunicipioService contains methods and other services that help
// with interacting with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLocalidadeCodigoMunicipioService] method instead.
type LocalidadeCodigoMunicipioService struct {
	Options []option.RequestOption
}

// NewLocalidadeCodigoMunicipioService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLocalidadeCodigoMunicipioService(opts ...option.RequestOption) (r *LocalidadeCodigoMunicipioService) {
	r = &LocalidadeCodigoMunicipioService{}
	r.Options = opts
	return
}

// Returns a IBGE code id
func (r *LocalidadeCodigoMunicipioService) Get(ctx context.Context, codigoIbge string, opts ...option.RequestOption) (res *Ibge, err error) {
	opts = append(r.Options[:], opts...)
	if codigoIbge == "" {
		err = errors.New("missing required codigo_ibge parameter")
		return
	}
	path := fmt.Sprintf("v1/localidades/codigo/%s/municipio", codigoIbge)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
