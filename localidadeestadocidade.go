// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/brennoo/findcep-go/internal/apijson"
	"github.com/brennoo/findcep-go/internal/requestconfig"
	"github.com/brennoo/findcep-go/option"
)

// LocalidadeEstadoCidadeService contains methods and other services that help with
// interacting with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLocalidadeEstadoCidadeService] method instead.
type LocalidadeEstadoCidadeService struct {
	Options []option.RequestOption
}

// NewLocalidadeEstadoCidadeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLocalidadeEstadoCidadeService(opts ...option.RequestOption) (r *LocalidadeEstadoCidadeService) {
	r = &LocalidadeEstadoCidadeService{}
	r.Options = opts
	return
}

// Returns a list of Cities per State
func (r *LocalidadeEstadoCidadeService) List(ctx context.Context, uf string, opts ...option.RequestOption) (res *Cidades, err error) {
	opts = append(r.Options[:], opts...)
	if uf == "" {
		err = errors.New("missing required uf parameter")
		return
	}
	path := fmt.Sprintf("v1/localidades/estado/%s/cidades", uf)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LocalidadeEstadoCidadeListResponse struct {
	HashKey string                                 `json:"hash_key"`
	Nome    string                                 `json:"nome"`
	URLKey  string                                 `json:"url_key"`
	JSON    localidadeEstadoCidadeListResponseJSON `json:"-"`
}

// localidadeEstadoCidadeListResponseJSON contains the JSON metadata for the struct
// [LocalidadeEstadoCidadeListResponse]
type localidadeEstadoCidadeListResponseJSON struct {
	HashKey     apijson.Field
	Nome        apijson.Field
	URLKey      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocalidadeEstadoCidadeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r localidadeEstadoCidadeListResponseJSON) RawJSON() string {
	return r.raw
}
