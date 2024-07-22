// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep

import (
	"github.com/stainless-sdks/findcep-go/option"
)

// GeolocationDistanceFromCepService contains methods and other services that help
// with interacting with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGeolocationDistanceFromCepService] method instead.
type GeolocationDistanceFromCepService struct {
	Options  []option.RequestOption
	ToCep    *GeolocationDistanceFromCepToCepService
	ToLatlon *GeolocationDistanceFromCepToLatlonService
}

// NewGeolocationDistanceFromCepService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewGeolocationDistanceFromCepService(opts ...option.RequestOption) (r *GeolocationDistanceFromCepService) {
	r = &GeolocationDistanceFromCepService{}
	r.Options = opts
	r.ToCep = NewGeolocationDistanceFromCepToCepService(opts...)
	r.ToLatlon = NewGeolocationDistanceFromCepToLatlonService(opts...)
	return
}
