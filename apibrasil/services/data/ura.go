package data

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// UraService é a URA reversa / ligações: /ura/call/*.
type UraService struct {
	services.BaseService
}

// NewUraService cria o serviço de URA.
func NewUraService(client *core.HTTPClient) *UraService {
	return &UraService{BaseService: services.NewBaseService(client)}
}

// Dialler dispara uma ligação: POST /ura/call/dialler.
func (s *UraService) Dialler(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "ura/call/dialler", body, opts...)
}

// Status consulta o status da ligação: POST /ura/call/status.
func (s *UraService) Status(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "ura/call/status", body, opts...)
}
