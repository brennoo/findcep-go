// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep

import (
	"github.com/stainless-sdks/findcep-go/option"
)

// GeolocationDistanceService contains methods and other services that help with
// interacting with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGeolocationDistanceService] method instead.
type GeolocationDistanceService struct {
	Options    []option.RequestOption
	FromCep    *GeolocationDistanceFromCepService
	FromLatlon *GeolocationDistanceFromLatlonService
}

// NewGeolocationDistanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGeolocationDistanceService(opts ...option.RequestOption) (r *GeolocationDistanceService) {
	r = &GeolocationDistanceService{}
	r.Options = opts
	r.FromCep = NewGeolocationDistanceFromCepService(opts...)
	r.FromLatlon = NewGeolocationDistanceFromLatlonService(opts...)
	return
}
