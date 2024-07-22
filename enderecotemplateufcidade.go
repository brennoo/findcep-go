// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep

import (
	"context"
	"net/http"

	"github.com/brennoo/findcep-go/internal/apijson"
	"github.com/brennoo/findcep-go/internal/param"
	"github.com/brennoo/findcep-go/internal/requestconfig"
	"github.com/brennoo/findcep-go/option"
)

// EnderecoTemplateUfCidadeService contains methods and other services that help
// with interacting with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnderecoTemplateUfCidadeService] method instead.
type EnderecoTemplateUfCidadeService struct {
	Options []option.RequestOption
}

// NewEnderecoTemplateUfCidadeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEnderecoTemplateUfCidadeService(opts ...option.RequestOption) (r *EnderecoTemplateUfCidadeService) {
	r = &EnderecoTemplateUfCidadeService{}
	r.Options = opts
	return
}

// Essa método aceita acentuação na busca e usa um modelo de aproximação das
// palavras para contornar erros de ortograficos de nomes e sobrenomes dos
// endereços.
func (r *EnderecoTemplateUfCidadeService) New(ctx context.Context, body EnderecoTemplateUfCidadeNewParams, opts ...option.RequestOption) (res *Endereco, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/endereco/template/uf/cidade"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EnderecoTemplateUfCidadeNewParams struct {
	// modelo do template de busca
	ID     param.Field[string]                                  `json:"id"`
	Params param.Field[EnderecoTemplateUfCidadeNewParamsParams] `json:"params"`
}

func (r EnderecoTemplateUfCidadeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnderecoTemplateUfCidadeNewParamsParams struct {
	// filtro para busca
	Cidade param.Field[string] `json:"cidade"`
	// paginação dos resultados
	From param.Field[int64] `json:"from"`
	// termo para busca no campo logradouro
	QueryString param.Field[string] `json:"query_string"`
	// quantidade de registros a serem exibidos
	Size param.Field[int64] `json:"size"`
	// filtro para busca
	Uf param.Field[string] `json:"uf"`
}

func (r EnderecoTemplateUfCidadeNewParamsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
