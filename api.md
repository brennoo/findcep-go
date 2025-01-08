# Ceps

Response Types:

- <a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Cep">Cep</a>

Methods:

- <code title="get /v1/cep/{cep}.json">client.Ceps.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#CepService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cep <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Cep">Cep</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Removido

Methods:

- <code title="get /v1/cep/removido/{cep_removido}.json">client.Ceps.Removido.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#CepRemovidoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cepRemovido <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Cep">Cep</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Enderecos

Response Types:

- <a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Endereco">Endereco</a>

Methods:

- <code title="post /v1/endereco/pesquisa">client.Enderecos.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#EnderecoService.Pesquisa">Pesquisa</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#EnderecoPesquisaParams">EnderecoPesquisaParams</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Endereco">Endereco</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Templates

### CodigoIbge

Methods:

- <code title="post /v1/endereco/template/codigo_ibge">client.Enderecos.Templates.CodigoIbge.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#EnderecoTemplateCodigoIbgeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#EnderecoTemplateCodigoIbgeNewParams">EnderecoTemplateCodigoIbgeNewParams</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Endereco">Endereco</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Uf

Methods:

- <code title="post /v1/endereco/template/uf">client.Enderecos.Templates.Uf.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#EnderecoTemplateUfService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#EnderecoTemplateUfNewParams">EnderecoTemplateUfNewParams</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Endereco">Endereco</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Cidade

Methods:

- <code title="post /v1/endereco/template/uf/cidade">client.Enderecos.Templates.Uf.Cidade.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#EnderecoTemplateUfCidadeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#EnderecoTemplateUfCidadeNewParams">EnderecoTemplateUfCidadeNewParams</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Endereco">Endereco</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Geolocations

Response Types:

- <a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#GeolocationCep">GeolocationCep</a>

Methods:

- <code title="get /v1/geolocation/cep/{cep}">client.Geolocations.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#GeolocationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cep <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#GeolocationCep">GeolocationCep</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Distances

### FromCep

#### ToCep

Methods:

- <code title="get /v1/geolocation/distance/from/cep/{cep_from}/to/cep/{cep_to}">client.Geolocations.Distances.FromCep.ToCep.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#GeolocationDistanceFromCepToCepService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cepFrom <a href="https://pkg.go.dev/builtin#string">string</a>, cepTo <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#GeolocationCep">GeolocationCep</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### ToLatlon

Methods:

- <code title="get /v1/geolocation/distance/from/cep/{cep_from}/to/latlon/{latlon_to}">client.Geolocations.Distances.FromCep.ToLatlon.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#GeolocationDistanceFromCepToLatlonService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cepFrom <a href="https://pkg.go.dev/builtin#string">string</a>, latlonTo <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#GeolocationCep">GeolocationCep</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### FromLatlon

#### ToLatlon

Methods:

- <code title="get /v1/geolocation/distance/from/latlon/{latlon_from}/to/latlon/{latlon_to}">client.Geolocations.Distances.FromLatlon.ToLatlon.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#GeolocationDistanceFromLatlonToLatlonService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, latlonFrom <a href="https://pkg.go.dev/builtin#string">string</a>, latlonTo <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#GeolocationCep">GeolocationCep</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# GeoDistances

## DistanceFromCepToCep

Methods:

- <code title="get /v1/geolocation/distance/from/cep/{cep_from}/to/cep/{cep_to}/approximate/{approximate}">client.GeoDistances.DistanceFromCepToCep.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#GeoDistanceDistanceFromCepToCepService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cepFrom <a href="https://pkg.go.dev/builtin#string">string</a>, cepTo <a href="https://pkg.go.dev/builtin#string">string</a>, approximate <a href="https://pkg.go.dev/builtin#float64">float64</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#GeolocationCep">GeolocationCep</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## DistanceFromCepToLatlon

Methods:

- <code title="get /v1/geolocation/distance/from/cep/{cep_from}/to/latlon/{latlon_to}/approximate/{approximate}">client.GeoDistances.DistanceFromCepToLatlon.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#GeoDistanceDistanceFromCepToLatlonService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cepFrom <a href="https://pkg.go.dev/builtin#string">string</a>, latlonTo <a href="https://pkg.go.dev/builtin#string">string</a>, approximate <a href="https://pkg.go.dev/builtin#float64">float64</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#GeolocationCep">GeolocationCep</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## DistanceFromLatlonToLatlon

Methods:

- <code title="get /v1/geolocation/distance/from/latlon/{latlon_from}/to/latlon/{latlon_to}/approximate/{approximate}">client.GeoDistances.DistanceFromLatlonToLatlon.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#GeoDistanceDistanceFromLatlonToLatlonService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, latlonFrom <a href="https://pkg.go.dev/builtin#string">string</a>, latlonTo <a href="https://pkg.go.dev/builtin#string">string</a>, approximate <a href="https://pkg.go.dev/builtin#float64">float64</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#GeolocationCep">GeolocationCep</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Localidades

Response Types:

- <a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Bairros">Bairros</a>
- <a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Cidades">Cidades</a>
- <a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Estados">Estados</a>
- <a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Ibge">Ibge</a>

## Estados

Methods:

- <code title="get /v1/localidades/estados">client.Localidades.Estados.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#LocalidadeEstadoService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Estados">Estados</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Estado

### Cidades

Methods:

- <code title="get /v1/localidades/estado/{uf}/cidades">client.Localidades.Estado.Cidades.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#LocalidadeEstadoCidadeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, uf <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Cidades">Cidades</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Cidade

#### Bairros

Methods:

- <code title="get /v1/localidades/estado/{uf}/cidade/{cidade_key}/bairros">client.Localidades.Estado.Cidade.Bairros.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#LocalidadeEstadoCidadeBairroService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, uf <a href="https://pkg.go.dev/builtin#string">string</a>, cidadeKey <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Bairros">Bairros</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Ibge

Methods:

- <code title="get /v1/localidades/estado/{uf}/cidade/{cidade_key}/ibge">client.Localidades.Estado.Cidade.Ibge.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#LocalidadeEstadoCidadeIbgeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, uf <a href="https://pkg.go.dev/builtin#string">string</a>, cidadeKey <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Ibge">Ibge</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Cidade

### Bairros

Methods:

- <code title="get /v1/localidades/cidade/{cidade_hash}/bairros">client.Localidades.Cidade.Bairros.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#LocalidadeCidadeBairroService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cidadeHash <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Bairros">Bairros</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Ibge

Methods:

- <code title="get /v1/localidades/cidade/{cidade_hash}/ibge">client.Localidades.Cidade.Ibge.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#LocalidadeCidadeIbgeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cidadeHash <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Ibge">Ibge</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Codigo

### Municipio

Methods:

- <code title="get /v1/localidades/codigo/{codigo_ibge}/municipio">client.Localidades.Codigo.Municipio.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#LocalidadeCodigoMunicipioService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, codigoIbge <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/brennoo/findcep-go">findcep</a>.<a href="https://pkg.go.dev/github.com/brennoo/findcep-go#Ibge">Ibge</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
