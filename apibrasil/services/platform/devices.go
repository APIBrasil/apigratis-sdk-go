package platform

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// DevicesService é a gestão de devices (/devices/*).
//
// Devices são a credencial de consumo dos serviços device-based: crie um
// device com a SecretKey da API desejada e use o device_token retornado
// como header DeviceToken.
type DevicesService struct {
	services.BaseService
}

// NewDevicesService cria o serviço de devices.
func NewDevicesService(client *core.HTTPClient) *DevicesService {
	return &DevicesService{BaseService: services.NewBaseService(client)}
}

// List lista os devices do usuário: GET /devices.
func (s *DevicesService) List(ctx context.Context, query core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "devices", core.FirstOption(opts).WithQuery(query))
}

// Store cria um device: POST /devices/store.
//
// body aceita device_name, type (cellphone|tablet|computer|server),
// device_key, server_search e os webhooks. A SecretKey da API (painel
// APIBrasil) vai no header — informe em RequestOptions.SecretKey ou
// configure SecretKey no cliente.
func (s *DevicesService) Store(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	options := core.FirstOption(opts)
	if options.SecretKey == "" {
		options.SecretKey = s.HTTP().SecretKey()
	}
	return s.Post(ctx, "devices/store", body, options)
}

// Show detalha um device: GET /devices/show?search={device_token}.
// deviceToken vazio usa o device configurado no cliente.
func (s *DevicesService) Show(ctx context.Context, deviceToken string, opts ...core.RequestOptions) (core.Json, error) {
	if deviceToken == "" {
		deviceToken = s.HTTP().DeviceToken()
	}
	return s.Get(ctx, "devices/show", core.FirstOption(opts).WithQuery(core.Json{"search": deviceToken}))
}

// Update atualiza um device: POST /devices/update
// (body com device_token + campos).
func (s *DevicesService) Update(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "devices/update", body, opts...)
}

// Destroy remove um device: DELETE /devices/destroy.
// deviceToken vazio usa o device configurado no cliente.
func (s *DevicesService) Destroy(ctx context.Context, deviceToken string, opts ...core.RequestOptions) (core.Json, error) {
	if deviceToken == "" {
		deviceToken = s.HTTP().DeviceToken()
	}
	return s.Delete(ctx, "devices/destroy", core.Json{"search": deviceToken}, opts...)
}

// Requests devolve o histórico de requisições do device:
// POST /devices/requests.
func (s *DevicesService) Requests(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "devices/requests", body, opts...)
}
