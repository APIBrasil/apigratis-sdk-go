package messaging

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// WhatsMeowService é a API WhatsMeow (device-based):
// POST /whatsmeow/{action}. Exige Authorization: Bearer + DeviceToken.
//
// As actions conhecidas estão em generated.WhatsMeowActions.
type WhatsMeowService struct {
	*services.DeviceProxyService
}

// NewWhatsMeowService cria o serviço WhatsMeow.
func NewWhatsMeowService(client *core.HTTPClient) *WhatsMeowService {
	return &WhatsMeowService{DeviceProxyService: services.NewDeviceProxyService(client, "whatsmeow")}
}

// SendText envia uma mensagem de texto: POST /whatsmeow/send/text.
func (s *WhatsMeowService) SendText(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "send/text", body, opts...)
}

// SendMedia envia uma mídia: POST /whatsmeow/send/media.
func (s *WhatsMeowService) SendMedia(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "send/media", body, opts...)
}

// InstanceCreate cria uma instância: POST /whatsmeow/instance/create.
func (s *WhatsMeowService) InstanceCreate(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "instance/create", body, opts...)
}

// InstanceConnect conecta a instância: POST /whatsmeow/instance/connect.
func (s *WhatsMeowService) InstanceConnect(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "instance/connect", body, opts...)
}

// InstanceQR retorna o QR Code da instância: POST /whatsmeow/instance/qr.
func (s *WhatsMeowService) InstanceQR(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "instance/qr", body, opts...)
}
