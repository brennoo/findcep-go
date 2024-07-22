// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/findcep-go/internal/apijson"
	"github.com/stainless-sdks/findcep-go/internal/requestconfig"
	"github.com/stainless-sdks/findcep-go/option"
)

// LocalidadeEstadoService contains methods and other services that help with
// interacting with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLocalidadeEstadoService] method instead.
type LocalidadeEstadoService struct {
	Options []option.RequestOption
}

// NewLocalidadeEstadoService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLocalidadeEstadoService(opts ...option.RequestOption) (r *LocalidadeEstadoService) {
	r = &LocalidadeEstadoService{}
	r.Options = opts
	return
}

// Returns a list of states
func (r *LocalidadeEstadoService) List(ctx context.Context, opts ...option.RequestOption) (res *Estados, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/localidades/estados"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LocalidadeEstadoListResponse struct {
	Nome   string                             `json:"nome"`
	Region LocalidadeEstadoListResponseRegion `json:"region"`
	Sigla  string                             `json:"sigla"`
	JSON   localidadeEstadoListResponseJSON   `json:"-"`
}

// localidadeEstadoListResponseJSON contains the JSON metadata for the struct
// [LocalidadeEstadoListResponse]
type localidadeEstadoListResponseJSON struct {
	Nome        apijson.Field
	Region      apijson.Field
	Sigla       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocalidadeEstadoListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r localidadeEstadoListResponseJSON) RawJSON() string {
	return r.raw
}

type LocalidadeEstadoListResponseRegion struct {
	Nome  string                                 `json:"nome"`
	Sigla string                                 `json:"sigla"`
	JSON  localidadeEstadoListResponseRegionJSON `json:"-"`
}

// localidadeEstadoListResponseRegionJSON contains the JSON metadata for the struct
// [LocalidadeEstadoListResponseRegion]
type localidadeEstadoListResponseRegionJSON struct {
	Nome        apijson.Field
	Sigla       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocalidadeEstadoListResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r localidadeEstadoListResponseRegionJSON) RawJSON() string {
	return r.raw
}
