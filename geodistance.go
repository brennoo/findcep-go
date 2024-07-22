// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep

import (
	"github.com/brennoo/findcep-go/option"
)

// GeoDistanceService contains methods and other services that help with
// interacting with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGeoDistanceService] method instead.
type GeoDistanceService struct {
	Options                    []option.RequestOption
	DistanceFromCepToCep       *GeoDistanceDistanceFromCepToCepService
	DistanceFromCepToLatlon    *GeoDistanceDistanceFromCepToLatlonService
	DistanceFromLatlonToLatlon *GeoDistanceDistanceFromLatlonToLatlonService
}

// NewGeoDistanceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGeoDistanceService(opts ...option.RequestOption) (r *GeoDistanceService) {
	r = &GeoDistanceService{}
	r.Options = opts
	r.DistanceFromCepToCep = NewGeoDistanceDistanceFromCepToCepService(opts...)
	r.DistanceFromCepToLatlon = NewGeoDistanceDistanceFromCepToLatlonService(opts...)
	r.DistanceFromLatlonToLatlon = NewGeoDistanceDistanceFromLatlonToLatlonService(opts...)
	return
}
