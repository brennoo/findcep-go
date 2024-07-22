// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep

import (
	"github.com/stainless-sdks/findcep-go/option"
)

// LocalidadeCidadeService contains methods and other services that help with
// interacting with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLocalidadeCidadeService] method instead.
type LocalidadeCidadeService struct {
	Options []option.RequestOption
	Bairros *LocalidadeCidadeBairroService
	Ibge    *LocalidadeCidadeIbgeService
}

// NewLocalidadeCidadeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLocalidadeCidadeService(opts ...option.RequestOption) (r *LocalidadeCidadeService) {
	r = &LocalidadeCidadeService{}
	r.Options = opts
	r.Bairros = NewLocalidadeCidadeBairroService(opts...)
	r.Ibge = NewLocalidadeCidadeIbgeService(opts...)
	return
}
