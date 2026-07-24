package data

import (
	"context"
	"net/url"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// VehiclesService são os veículos por placa (device-based):
// /vehicles/{action}.
type VehiclesService struct {
	*services.DeviceProxyService
}

// NewVehiclesService cria o serviço de veículos.
func NewVehiclesService(client *core.HTTPClient) *VehiclesService {
	return &VehiclesService{DeviceProxyService: services.NewDeviceProxyService(client, "vehicles")}
}

// Dados consulta os dados do veículo pela placa:
// POST /vehicles/dados body {"placa": "..."}.
func (s *VehiclesService) Dados(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "dados", body, opts...)
}

// Fipe consulta a FIPE pela placa:
// POST /vehicles/fipe body {"placa": "..."}.
func (s *VehiclesService) Fipe(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "fipe", body, opts...)
}

// ConsultaFipe consulta a FIPE por placa na rota dedicada:
// POST /vehicles/consultafipe/{placa}.
func (s *VehiclesService) ConsultaFipe(ctx context.Context, placa string, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "consultafipe/"+url.PathEscape(placa), nil, opts...)
}

// BaseDados consulta a base nacional: POST /vehicles/base/000/dados.
func (s *VehiclesService) BaseDados(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "base/000/dados", body, opts...)
}

// BaseFipe consulta a FIPE na base nacional: POST /vehicles/base/000/fipe.
func (s *VehiclesService) BaseFipe(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "base/000/fipe", body, opts...)
}
