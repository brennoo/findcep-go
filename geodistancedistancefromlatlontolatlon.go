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

// GeoDistanceDistanceFromLatlonToLatlonService contains methods and other services
// that help with interacting with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGeoDistanceDistanceFromLatlonToLatlonService] method instead.
type GeoDistanceDistanceFromLatlonToLatlonService struct {
	Options []option.RequestOption
}

// NewGeoDistanceDistanceFromLatlonToLatlonService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewGeoDistanceDistanceFromLatlonToLatlonService(opts ...option.RequestOption) (r *GeoDistanceDistanceFromLatlonToLatlonService) {
	r = &GeoDistanceDistanceFromLatlonToLatlonService{}
	r.Options = opts
	return
}

// Returns a distance between two cep. Lista de ceps para validar todos os campos
// de retorno. [01234000, 99945970, 01010000, 14740000, 61603970, 60150901,
// 12918991, 99999999]
func (r *GeoDistanceDistanceFromLatlonToLatlonService) List(ctx context.Context, latlonFrom string, latlonTo string, approximate float64, opts ...option.RequestOption) (res *GeolocationCep, err error) {
	opts = append(r.Options[:], opts...)
	if latlonFrom == "" {
		err = errors.New("missing required latlon_from parameter")
		return
	}
	if latlonTo == "" {
		err = errors.New("missing required latlon_to parameter")
		return
	}
	path := fmt.Sprintf("v1/geolocation/distance/from/latlon/%s/to/latlon/%s/approximate/%v", latlonFrom, latlonTo, approximate)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
