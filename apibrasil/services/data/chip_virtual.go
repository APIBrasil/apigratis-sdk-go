package data

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// ChipVirtualService é o chip virtual: /chip/virtual/*.
type ChipVirtualService struct {
	services.BaseService
}

// NewChipVirtualService cria o serviço de chip virtual.
func NewChipVirtualService(client *core.HTTPClient) *ChipVirtualService {
	return &ChipVirtualService{BaseService: services.NewBaseService(client)}
}

// Operators lista as operadoras: POST /chip/virtual/operators.
func (s *ChipVirtualService) Operators(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "chip/virtual/operators", body, opts...)
}

// Buy compra um número: POST /chip/virtual/buy.
func (s *ChipVirtualService) Buy(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "chip/virtual/buy", body, opts...)
}

// Activation consulta a ativação: POST /chip/virtual/activation.
func (s *ChipVirtualService) Activation(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "chip/virtual/activation", body, opts...)
}

// Services lista os serviços disponíveis: POST /chip/virtual/services.
func (s *ChipVirtualService) Services(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "chip/virtual/services", body, opts...)
}
