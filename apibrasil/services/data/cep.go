package data

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// CepService é o CEP + geolocalização (device-based): /cep/{action}.
type CepService struct {
	*services.DeviceProxyService
}

// NewCepService cria o serviço de CEP.
func NewCepService(client *core.HTTPClient) *CepService {
	return &CepService{DeviceProxyService: services.NewDeviceProxyService(client, "cep")}
}

// CEP consulta um CEP: POST /cep/cep body {"cep": "..."}.
func (s *CepService) CEP(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "cep", body, opts...)
}

// Bairros lista os bairros: POST /cep/bairros.
func (s *CepService) Bairros(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "bairros", body, opts...)
}

// Cidades lista as cidades: POST /cep/cidades.
func (s *CepService) Cidades(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "cidades", body, opts...)
}

// CidadesPorDDD lista as cidades de um DDD: POST /cep/cidadesPorDDD.
func (s *CepService) CidadesPorDDD(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "cidadesPorDDD", body, opts...)
}

// Estados lista os estados: POST /cep/estados.
func (s *CepService) Estados(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "estados", body, opts...)
}

// CalcularDistancia calcula a distância entre CEPs:
// POST /cep/distancia/calcular.
func (s *CepService) CalcularDistancia(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "distancia/calcular", body, opts...)
}
