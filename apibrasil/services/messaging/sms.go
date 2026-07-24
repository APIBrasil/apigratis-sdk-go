package messaging

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// SmsService é a API de SMS (device-based): POST /sms/{action}.
// Exige Authorization: Bearer + DeviceToken.
//
// Campos aceitos por Send/SendWithCredits: number, message, operator
// (padrão do gateway: claro), user_reply e webhook_url.
type SmsService struct {
	*services.DeviceProxyService
}

// NewSmsService cria o serviço de SMS.
func NewSmsService(client *core.HTTPClient) *SmsService {
	return &SmsService{DeviceProxyService: services.NewDeviceProxyService(client, "sms")}
}

// Send envia um SMS pelo device: POST /sms/send.
func (s *SmsService) Send(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "send", body, opts...)
}

// SendWithCredits envia um SMS debitando créditos da conta (sem
// DeviceToken): POST /sms/send/credits.
func (s *SmsService) SendWithCredits(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return services.Credit(s.Post(ctx, "sms/send/credits", body, opts...))
}
