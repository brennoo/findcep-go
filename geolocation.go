// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/findcep-go/internal/apijson"
	"github.com/stainless-sdks/findcep-go/internal/requestconfig"
	"github.com/stainless-sdks/findcep-go/option"
)

// GeolocationService contains methods and other services that help with
// interacting with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGeolocationService] method instead.
type GeolocationService struct {
	Options   []option.RequestOption
	Distances *GeolocationDistanceService
}

// NewGeolocationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGeolocationService(opts ...option.RequestOption) (r *GeolocationService) {
	r = &GeolocationService{}
	r.Options = opts
	r.Distances = NewGeolocationDistanceService(opts...)
	return
}

// Returns a map with geo coordinates. Lista de ceps para validar todos os campos
// de retorno. [01234000, 99945970, 01010000, 14740000, 61603970, 60150901,
// 12918991, 99999999]
func (r *GeolocationService) Get(ctx context.Context, cep string, opts ...option.RequestOption) (res *GeolocationCep, err error) {
	opts = append(r.Options[:], opts...)
	if cep == "" {
		err = errors.New("missing required cep parameter")
		return
	}
	path := fmt.Sprintf("v1/geolocation/cep/%s", cep)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type GeolocationCep struct {
	// geo location
	Location GeolocationCepLocation `json:"location"`
	// cep
	PostalCode string `json:"postal_code"`
	// true quando existir geo coodenadas disponiveis para o cep
	Status bool               `json:"status"`
	JSON   geolocationCepJSON `json:"-"`
}

// geolocationCepJSON contains the JSON metadata for the struct [GeolocationCep]
type geolocationCepJSON struct {
	Location    apijson.Field
	PostalCode  apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeolocationCep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geolocationCepJSON) RawJSON() string {
	return r.raw
}

// geo location
type GeolocationCepLocation struct {
	Lat  float64                    `json:"lat"`
	Lon  float64                    `json:"lon"`
	JSON geolocationCepLocationJSON `json:"-"`
}

// geolocationCepLocationJSON contains the JSON metadata for the struct
// [GeolocationCepLocation]
type geolocationCepLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeolocationCepLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geolocationCepLocationJSON) RawJSON() string {
	return r.raw
}
