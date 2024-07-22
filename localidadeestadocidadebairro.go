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

// LocalidadeEstadoCidadeBairroService contains methods and other services that
// help with interacting with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLocalidadeEstadoCidadeBairroService] method instead.
type LocalidadeEstadoCidadeBairroService struct {
	Options []option.RequestOption
}

// NewLocalidadeEstadoCidadeBairroService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLocalidadeEstadoCidadeBairroService(opts ...option.RequestOption) (r *LocalidadeEstadoCidadeBairroService) {
	r = &LocalidadeEstadoCidadeBairroService{}
	r.Options = opts
	return
}

// Returns a map of bairros
func (r *LocalidadeEstadoCidadeBairroService) List(ctx context.Context, uf string, cidadeKey string, opts ...option.RequestOption) (res *Bairros, err error) {
	opts = append(r.Options[:], opts...)
	if uf == "" {
		err = errors.New("missing required uf parameter")
		return
	}
	if cidadeKey == "" {
		err = errors.New("missing required cidade_key parameter")
		return
	}
	path := fmt.Sprintf("v1/localidades/estado/%s/cidade/%s/bairros", uf, cidadeKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
