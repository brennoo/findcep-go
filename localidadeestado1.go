// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep

import (
	"github.com/brennoo/findcep-go/option"
)

// LocalidadeEstadoService contains methods and other services that help with
// interacting with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLocalidadeEstadoService] method instead.
type LocalidadeEstadoService struct {
	Options []option.RequestOption
	Cidades *LocalidadeEstadoCidadeService
	Cidade  *LocalidadeEstadoCidadeService
}

// NewLocalidadeEstadoService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLocalidadeEstadoService(opts ...option.RequestOption) (r *LocalidadeEstadoService) {
	r = &LocalidadeEstadoService{}
	r.Options = opts
	r.Cidades = NewLocalidadeEstadoCidadeService(opts...)
	r.Cidade = NewLocalidadeEstadoCidadeService(opts...)
	return
}
