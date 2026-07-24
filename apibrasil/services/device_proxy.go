package services

import (
	"context"
	"strings"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
)

// DeviceProxyService é a base dos serviços "device-based" do gateway
// (POST /api/v2/{servico}/{action}). Exigem Authorization: Bearer +
// header DeviceToken.
//
// Todos expõem Request(ctx, action, body) como porta de saída genérica —
// as actions são dinâmicas por provedor. As conhecidas estão em
// generated.ServiceActions; a documentação completa fica em
// https://doc.apibrasil.io
type DeviceProxyService struct {
	BaseService
	// Name é o serviço do gateway (ex: whatsapp, sms, cep).
	Name string
}

// NewDeviceProxyService cria um serviço device-based para name.
func NewDeviceProxyService(client *core.HTTPClient, name string) *DeviceProxyService {
	return &DeviceProxyService{BaseService: NewBaseService(client), Name: name}
}

// Request executa uma action do serviço: POST /{servico}/{action}.
func (s *DeviceProxyService) Request(ctx context.Context, action string, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.device(ctx, s.path(action), body, opts...)
}

// Queue executa a mesma action de forma assíncrona via fila:
// POST /{servico}/{action}/queue.
func (s *DeviceProxyService) Queue(ctx context.Context, action string, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.device(ctx, s.path(action)+"/queue", body, opts...)
}

func (s *DeviceProxyService) path(action string) string {
	action = strings.Trim(action, "/")
	if action == "" {
		return s.Name
	}
	return s.Name + "/" + action
}

func (s *DeviceProxyService) device(ctx context.Context, path string, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	response, err := s.Post(ctx, path, body, opts...)
	if err != nil {
		return nil, err
	}
	return core.DeviceServiceResponse(response), nil
}

// Device converte uma resposta JSON no envelope device-based — atalho
// para serviços que chamam rotas fora do padrão /{servico}/{action}.
func Device(response core.Json, err error) (core.DeviceServiceResponse, error) {
	if err != nil {
		return nil, err
	}
	return core.DeviceServiceResponse(response), nil
}

// Credit converte uma resposta JSON no envelope das consultas por crédito.
func Credit(response core.Json, err error) (core.CreditServiceResponse, error) {
	if err != nil {
		return nil, err
	}
	return core.CreditServiceResponse(response), nil
}
