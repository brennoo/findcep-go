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

// EnderecoService contains methods and other services that help with interacting
// with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnderecoService] method instead.
type EnderecoService struct {
	Options   []option.RequestOption
	Templates *EnderecoTemplateService
}

// NewEnderecoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEnderecoService(opts ...option.RequestOption) (r *EnderecoService) {
	r = &EnderecoService{}
	r.Options = opts
	r.Templates = NewEnderecoTemplateService(opts...)
	return
}

// Essa método aceita acentuação na busca e usa um modelo de aproximação das
// palavras para contornar erros de ortograficos de nomes e sobrenomes dos
// endereços.
func (r *EnderecoService) Pesquisa(ctx context.Context, body EnderecoPesquisaParams, opts ...option.RequestOption) (res *Endereco, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/endereco/pesquisa"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Endereco struct {
	Hits EnderecoHits `json:"hits"`
	JSON enderecoJSON `json:"-"`
}

// enderecoJSON contains the JSON metadata for the struct [Endereco]
type enderecoJSON struct {
	Hits        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Endereco) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enderecoJSON) RawJSON() string {
	return r.raw
}

type EnderecoHits struct {
	Hits  []EnderecoHitsHit `json:"hits"`
	Total int64             `json:"total"`
	JSON  enderecoHitsJSON  `json:"-"`
}

// enderecoHitsJSON contains the JSON metadata for the struct [EnderecoHits]
type enderecoHitsJSON struct {
	Hits        apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnderecoHits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enderecoHitsJSON) RawJSON() string {
	return r.raw
}

type EnderecoHitsHit struct {
	Source interface{}         `json:"_source"`
	JSON   enderecoHitsHitJSON `json:"-"`
}

// enderecoHitsHitJSON contains the JSON metadata for the struct [EnderecoHitsHit]
type enderecoHitsHitJSON struct {
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnderecoHitsHit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enderecoHitsHitJSON) RawJSON() string {
	return r.raw
}

type EnderecoPesquisaParams struct {
	// modelo do template de busca
	ID     param.Field[string]                       `json:"id"`
	Params param.Field[EnderecoPesquisaParamsParams] `json:"params"`
}

func (r EnderecoPesquisaParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EnderecoPesquisaParamsParams struct {
	// paginação dos resultados
	From param.Field[int64] `json:"from"`
	// termo para busca
	QueryString param.Field[string] `json:"query_string"`
	// quantidade de registros a serem exibidos
	Size param.Field[int64] `json:"size"`
}

func (r EnderecoPesquisaParamsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
