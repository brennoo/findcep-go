// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep

import (
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
	Bairros *LocalidadeEstadoCidadeBairroService
	Ibge    *LocalidadeEstadoCidadeIbgeService
}

// NewLocalidadeEstadoCidadeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLocalidadeEstadoCidadeService(opts ...option.RequestOption) (r *LocalidadeEstadoCidadeService) {
	r = &LocalidadeEstadoCidadeService{}
	r.Options = opts
	r.Bairros = NewLocalidadeEstadoCidadeBairroService(opts...)
	r.Ibge = NewLocalidadeEstadoCidadeIbgeService(opts...)
	return
}
