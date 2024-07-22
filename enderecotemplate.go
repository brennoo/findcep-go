// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep

import (
	"github.com/brennoo/findcep-go/option"
)

// EnderecoTemplateService contains methods and other services that help with
// interacting with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnderecoTemplateService] method instead.
type EnderecoTemplateService struct {
	Options    []option.RequestOption
	CodigoIbge *EnderecoTemplateCodigoIbgeService
	Uf         *EnderecoTemplateUfService
}

// NewEnderecoTemplateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEnderecoTemplateService(opts ...option.RequestOption) (r *EnderecoTemplateService) {
	r = &EnderecoTemplateService{}
	r.Options = opts
	r.CodigoIbge = NewEnderecoTemplateCodigoIbgeService(opts...)
	r.Uf = NewEnderecoTemplateUfService(opts...)
	return
}
