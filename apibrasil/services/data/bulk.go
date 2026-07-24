package data

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// BulkService é a execução em lote: /bulk/direct/{action} e
// /bulk/queue/{action}.
type BulkService struct {
	services.BaseService
}

// NewBulkService cria o serviço de execução em lote.
func NewBulkService(client *core.HTTPClient) *BulkService {
	return &BulkService{BaseService: services.NewBaseService(client)}
}

// Direct executa um lote de forma síncrona: POST /bulk/direct/{action}.
func (s *BulkService) Direct(ctx context.Context, action string, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "bulk/direct/"+action, body, opts...)
}

// Queue enfileira um lote: POST /bulk/queue/{action}.
func (s *BulkService) Queue(ctx context.Context, action string, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "bulk/queue/"+action, body, opts...)
}
