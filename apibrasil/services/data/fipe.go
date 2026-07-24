package data

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// FipeService é a tabela FIPE (device-based): /fipe/{action}.
type FipeService struct {
	*services.DeviceProxyService
}

// NewFipeService cria o serviço da tabela FIPE.
func NewFipeService(client *core.HTTPClient) *FipeService {
	return &FipeService{DeviceProxyService: services.NewDeviceProxyService(client, "fipe")}
}

// ConsultarTabelaDeReferencia lista as tabelas de referência.
func (s *FipeService) ConsultarTabelaDeReferencia(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "ConsultarTabelaDeReferencia", body, opts...)
}

// ConsultarMarcas lista as marcas.
func (s *FipeService) ConsultarMarcas(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "ConsultarMarcas", body, opts...)
}

// ConsultarModelos lista os modelos de uma marca.
func (s *FipeService) ConsultarModelos(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "ConsultarModelos", body, opts...)
}

// ConsultarAnoModelo lista os anos de um modelo.
func (s *FipeService) ConsultarAnoModelo(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "ConsultarAnoModelo", body, opts...)
}

// ConsultarModelosAtravesDoAno lista os modelos de um ano.
func (s *FipeService) ConsultarModelosAtravesDoAno(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "ConsultarModelosAtravesDoAno", body, opts...)
}

// ConsultarValorComTodosParametros consulta o preço FIPE.
func (s *FipeService) ConsultarValorComTodosParametros(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "ConsultarValorComTodosParametros", body, opts...)
}
