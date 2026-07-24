// Package data traz os serviços de dados da plataforma: consultas
// cadastrais, veículos, FIPE, correios, CEP, geolocalização, clima,
// loterias, URA, chip virtual, lotes e as consultas por crédito.
package data

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// DadosService são os dados cadastrais device-based: /dados/{action}.
type DadosService struct {
	*services.DeviceProxyService
}

// NewDadosService cria o serviço de dados cadastrais.
func NewDadosService(client *core.HTTPClient) *DadosService {
	return &DadosService{DeviceProxyService: services.NewDeviceProxyService(client, "dados")}
}

// CNPJ consulta um CNPJ: POST /dados/cnpj body {"cnpj": "..."}.
func (s *DadosService) CNPJ(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "cnpj", body, opts...)
}

// CPF consulta um CPF: POST /dados/cpf body {"cpf": "..."}.
func (s *DadosService) CPF(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "cpf", body, opts...)
}

// CEP consulta um CEP: POST /dados/cep.
func (s *DadosService) CEP(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "cep", body, opts...)
}

// ListaSocios lista os sócios de uma empresa: POST /dados/lista-socios.
func (s *DadosService) ListaSocios(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "lista-socios", body, opts...)
}

// ListaCnaes lista os CNAEs: POST /dados/lista-cnaes.
func (s *DadosService) ListaCnaes(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "lista-cnaes", body, opts...)
}

// CapitalSocial consulta o capital social: POST /dados/capital-social.
func (s *DadosService) CapitalSocial(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "capital-social", body, opts...)
}

// ByQuery consulta por filtros livres: POST /dados/byquery.
func (s *DadosService) ByQuery(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "byquery", body, opts...)
}
