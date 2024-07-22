// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package findcep

import (
	"github.com/stainless-sdks/findcep-go/internal/apijson"
	"github.com/stainless-sdks/findcep-go/option"
)

// LocalidadeService contains methods and other services that help with interacting
// with the findcep API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLocalidadeService] method instead.
type LocalidadeService struct {
	Options []option.RequestOption
	Estados *LocalidadeEstadoService
	Estado  *LocalidadeEstadoService
	Cidade  *LocalidadeCidadeService
	Codigo  *LocalidadeCodigoService
}

// NewLocalidadeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLocalidadeService(opts ...option.RequestOption) (r *LocalidadeService) {
	r = &LocalidadeService{}
	r.Options = opts
	r.Estados = NewLocalidadeEstadoService(opts...)
	r.Estado = NewLocalidadeEstadoService(opts...)
	r.Cidade = NewLocalidadeCidadeService(opts...)
	r.Codigo = NewLocalidadeCodigoService(opts...)
	return
}

type Bairros struct {
	// Lista com todos os bairros
	Bairros []string `json:"bairros"`
	// Nome da Cidade
	Cidade string `json:"cidade"`
	// reservado para o futuro
	Extras []interface{} `json:"extras"`
	// é a chave unica da localidade resultada da junçao dos campos "uf|cidade" da base
	// de Cep usando "shake_128" com Output Lenght setado em 5 bytes em hexadecimal.
	// "python3.8" "import hashlib;
	// print(hashlib.shake_128(b'SP|Bebedouro').hexdigest(5))"
	HashKey string `json:"hash_key"`
	// todos os nomes dos bairros
	Uf string `json:"uf"`
	// campo cidade normalizado com lower case e replace dos espaços por "-" para ser
	// usado como parte da url
	URLKey string `json:"url_key"`
	// timestamp com ultima atualizacao do dado
	Version int64       `json:"version"`
	JSON    bairrosJSON `json:"-"`
}

// bairrosJSON contains the JSON metadata for the struct [Bairros]
type bairrosJSON struct {
	Bairros     apijson.Field
	Cidade      apijson.Field
	Extras      apijson.Field
	HashKey     apijson.Field
	Uf          apijson.Field
	URLKey      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Bairros) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bairrosJSON) RawJSON() string {
	return r.raw
}

type Cidades []CidadesItem

type CidadesItem struct {
	HashKey string          `json:"hash_key"`
	Nome    string          `json:"nome"`
	URLKey  string          `json:"url_key"`
	JSON    cidadesItemJSON `json:"-"`
}

// cidadesItemJSON contains the JSON metadata for the struct [CidadesItem]
type cidadesItemJSON struct {
	HashKey     apijson.Field
	Nome        apijson.Field
	URLKey      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CidadesItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cidadesItemJSON) RawJSON() string {
	return r.raw
}

type Estados []EstadosItem

type EstadosItem struct {
	Nome   string            `json:"nome"`
	Region EstadosItemRegion `json:"region"`
	Sigla  string            `json:"sigla"`
	JSON   estadosItemJSON   `json:"-"`
}

// estadosItemJSON contains the JSON metadata for the struct [EstadosItem]
type estadosItemJSON struct {
	Nome        apijson.Field
	Region      apijson.Field
	Sigla       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EstadosItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r estadosItemJSON) RawJSON() string {
	return r.raw
}

type EstadosItemRegion struct {
	Nome  string                `json:"nome"`
	Sigla string                `json:"sigla"`
	JSON  estadosItemRegionJSON `json:"-"`
}

// estadosItemRegionJSON contains the JSON metadata for the struct
// [EstadosItemRegion]
type estadosItemRegionJSON struct {
	Nome        apijson.Field
	Sigla       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EstadosItemRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r estadosItemRegionJSON) RawJSON() string {
	return r.raw
}

type Ibge struct {
	// Nome da Cidade
	Cidade string `json:"cidade"`
	// Texto com codigo ibge
	CodigoIbge string `json:"codigo_ibge"`
	// reservado para o futuro
	Extras []interface{} `json:"extras"`
	// é a chave unica da localidade resultada da junçao dos campos "uf|cidade" da base
	// de Cep usando "shake_128" com Output Lenght setado em 5 bytes em hexadecimal.
	// python3 -c "import hashlib;
	// print(hashlib.shake_128(b'SP|Bebedouro').hexdigest(5))"
	HashKey string `json:"hash_key"`
	// todos os nomes dos bairros
	Uf string `json:"uf"`
	// campo cidade normalizado com lower case e replace dos espaços por "-" para ser
	// usado como parte da url
	URLKey string   `json:"url_key"`
	JSON   ibgeJSON `json:"-"`
}

// ibgeJSON contains the JSON metadata for the struct [Ibge]
type ibgeJSON struct {
	Cidade      apijson.Field
	CodigoIbge  apijson.Field
	Extras      apijson.Field
	HashKey     apijson.Field
	Uf          apijson.Field
	URLKey      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Ibge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ibgeJSON) RawJSON() string {
	return r.raw
}
