package data

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// CorreiosService são os Correios (device-based): /correios/{action}.
type CorreiosService struct {
	*services.DeviceProxyService
}

// NewCorreiosService cria o serviço dos Correios.
func NewCorreiosService(client *core.HTTPClient) *CorreiosService {
	return &CorreiosService{DeviceProxyService: services.NewDeviceProxyService(client, "correios")}
}

// Rastreio rastreia uma encomenda: POST /correios/rastreio.
func (s *CorreiosService) Rastreio(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "rastreio", body, opts...)
}
