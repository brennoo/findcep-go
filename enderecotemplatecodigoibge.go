// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/findcep-go/internal/apijson"
	"github.com/stainless-sdks/findcep-go/internal/param"
	"github.com/stainless-sdks/findcep-go/internal/requestconfig"
	"github.com/stainless-sdks/findcep-go/option"
)

// EnderecoTemplateCodigoIbgeService contains methods and other services that help
// with interacting with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnderecoTemplateCodigoIbgeService] method instead.
type EnderecoTemplateCodigoIbgeService struct {
	Options []option.RequestOption
}

// NewEnderecoTemplateCodigoIbgeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEnderecoTemplateCodigoIbgeService(opts ...option.RequestOption) (r *EnderecoTemplateCodigoIbgeService) {
	r = &EnderecoTemplateCodigoIbgeService{}
	r.Options = opts
	return
}

// Essa método aceita acentuação na busca e usa um modelo de aproximação das
// palavras para contornar erros de ortograficos de nomes e sobrenomes dos
// endereços.
func (r *EnderecoTemplateCodigoIbgeService) New(ctx context.Context, body EnderecoTemplateCodigoIbgeNewParams, opts ...option.RequestOption) (res *Endereco, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/endereco/template/codigo_ibge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EnderecoTemplateCodigoIbgeNewParams struct {
	// modelo do template de busca
	ID     param.Field[string]                                    `json:"id"`
	Params param.Field[EnderecoTemplateCodigoIbgeNewParamsParams] `json:"params"`
}

func (r EnderecoTemplateCodigoIbgeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnderecoTemplateCodigoIbgeNewParamsParams struct {
	// filtro para busca
	CodigoIbge param.Field[string] `json:"codigo_ibge"`
	// paginação dos resultados
	From param.Field[int64] `json:"from"`
	// termo para busca no campo logradouro
	QueryString param.Field[string] `json:"query_string"`
	// quantidade de registros a serem exibidos
	Size param.Field[int64] `json:"size"`
}

func (r EnderecoTemplateCodigoIbgeNewParamsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
