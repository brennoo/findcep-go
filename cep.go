// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/brennoo/findcep-go/internal/apijson"
	"github.com/brennoo/findcep-go/internal/requestconfig"
	"github.com/brennoo/findcep-go/option"
)

// CepService contains methods and other services that help with interacting with
// the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCepService] method instead.
type CepService struct {
	Options  []option.RequestOption
	Removido *CepRemovidoService
}

// NewCepService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCepService(opts ...option.RequestOption) (r *CepService) {
	r = &CepService{}
	r.Options = opts
	r.Removido = NewCepRemovidoService(opts...)
	return
}

// Returns a map with address. Lista de ceps para validar todos os campos de
// retorno. [01234000, 01010000, 14740000, 61603970, 60150901, 12918991, 99999999]
func (r *CepService) Get(ctx context.Context, cep string, opts ...option.RequestOption) (res *Cep, err error) {
	opts = append(r.Options[:], opts...)
	if cep == "" {
		err = errors.New("missing required cep parameter")
		return
	}
	path := fmt.Sprintf("v1/cep/%s.json", cep)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Cep struct {
	// Neighborhood
	Bairro string `json:"bairro"`
	// cep
	Cep string `json:"cep"`
	// City
	Cidade string `json:"cidade"`
	// IBGE Code
	CodigoIbge string `json:"codigo_ibge"`
	// Address complement only for directions
	Complemento string `json:"complemento"`
	// Address
	Logradouro string `json:"logradouro"`
	// Este campo só trará conteudo quando o campo "tipo" é igual a "cep_cod_especial e
	// cep_caixa_postal_comunitaria". Geralmente opcional para os formularios.
	Nome string `json:"nome"`
	// Status of CEP São denominados "completo" todos os ceps que não precisam de
	// interação manual do usuario para completar o endereço. Ver mais em "tipo". São
	// denominados "incompleto" todos os ceps que precisam ser completados pelo
	// usuario, pois esse tipo de cep geralmente não devolve os campos "logradouro/rua"
	// e "bairro". ver mais em "tipo".
	Status CepStatus `json:"status"`
	// cep_logradouro é um cep completo, com todos os campos do endereço fornecidas.
	// cep_localidade "14740-000" é um cep incompleto e regularmente conhecido como
	// "CEP ÚNICO", pois este cep engloba todo o municipio, pois a cidade ainda nao tem
	// populacao sufieciente para ter o cep catalogado por logradouro.
	// cep_agencia_correios "61603-970" é cep das agencias dos correios pelo Brasil.
	// cep_cod_especial "60150-901" é o cep grande usuario, e são os ceps destinados a
	// grandes empresas, universidades e outras denominacões que tenham um alto volume
	// de correspondencia, funcionando como uma mini-cidade, neste tipo de cep, o campo
	// "nome" trará a denominacao da instituição. cep_caixa_postal_comunitaria
	// "12918-991" é um cep incompleto e geralmente usado apenas para
	// "correspondencia", e regularmente "não é habilitado para entrega de encomendas"
	// por diversas transportadoras.
	Tipo CepTipo `json:"tipo"`
	// State
	Uf   string  `json:"uf"`
	JSON cepJSON `json:"-"`
}

// cepJSON contains the JSON metadata for the struct [Cep]
type cepJSON struct {
	Bairro      apijson.Field
	Cep         apijson.Field
	Cidade      apijson.Field
	CodigoIbge  apijson.Field
	Complemento apijson.Field
	Logradouro  apijson.Field
	Nome        apijson.Field
	Status      apijson.Field
	Tipo        apijson.Field
	Uf          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Cep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cepJSON) RawJSON() string {
	return r.raw
}

// Status of CEP São denominados "completo" todos os ceps que não precisam de
// interação manual do usuario para completar o endereço. Ver mais em "tipo". São
// denominados "incompleto" todos os ceps que precisam ser completados pelo
// usuario, pois esse tipo de cep geralmente não devolve os campos "logradouro/rua"
// e "bairro". ver mais em "tipo".
type CepStatus string

const (
	CepStatusCompleto   CepStatus = "completo"
	CepStatusIncompleto CepStatus = "incompleto"
)

func (r CepStatus) IsKnown() bool {
	switch r {
	case CepStatusCompleto, CepStatusIncompleto:
		return true
	}
	return false
}

// cep_logradouro é um cep completo, com todos os campos do endereço fornecidas.
// cep_localidade "14740-000" é um cep incompleto e regularmente conhecido como
// "CEP ÚNICO", pois este cep engloba todo o municipio, pois a cidade ainda nao tem
// populacao sufieciente para ter o cep catalogado por logradouro.
// cep_agencia_correios "61603-970" é cep das agencias dos correios pelo Brasil.
// cep_cod_especial "60150-901" é o cep grande usuario, e são os ceps destinados a
// grandes empresas, universidades e outras denominacões que tenham um alto volume
// de correspondencia, funcionando como uma mini-cidade, neste tipo de cep, o campo
// "nome" trará a denominacao da instituição. cep_caixa_postal_comunitaria
// "12918-991" é um cep incompleto e geralmente usado apenas para
// "correspondencia", e regularmente "não é habilitado para entrega de encomendas"
// por diversas transportadoras.
type CepTipo string

const (
	CepTipoCepLogradouro             CepTipo = "cep_logradouro"
	CepTipoCepLocalidade             CepTipo = "cep_localidade"
	CepTipoCepAgenciaCorreios        CepTipo = "cep_agencia_correios"
	CepTipoCepCodEspecial            CepTipo = "cep_cod_especial"
	CepTipoCepCaixaPostalComunitaria CepTipo = "cep_caixa_postal_comunitaria"
)

func (r CepTipo) IsKnown() bool {
	switch r {
	case CepTipoCepLogradouro, CepTipoCepLocalidade, CepTipoCepAgenciaCorreios, CepTipoCepCodEspecial, CepTipoCepCaixaPostalComunitaria:
		return true
	}
	return false
}
