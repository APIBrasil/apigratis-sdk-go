// Package services traz as bases compartilhadas pelos serviços da SDK:
// BaseService (atalhos HTTP), DeviceProxyService (rotas device-based) e
// CreditService (consultas por crédito).
package services

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
)

// BaseService é a base de todos os serviços — guarda o cliente HTTP e
// expõe os atalhos por verbo.
type BaseService struct {
	client *core.HTTPClient
}

// NewBaseService cria a base sobre um cliente HTTP já configurado.
func NewBaseService(client *core.HTTPClient) BaseService {
	return BaseService{client: client}
}

// HTTP devolve o cliente HTTP compartilhado.
func (s BaseService) HTTP() *core.HTTPClient { return s.client }

// URL monta a URL completa de um caminho.
func (s BaseService) URL(path string) string {
	return core.JoinURL(s.client.BaseURL(), path)
}

// Get executa GET path.
func (s BaseService) Get(ctx context.Context, path string, opts ...core.RequestOptions) (core.Json, error) {
	return s.client.Get(ctx, path, opts...)
}

// Post executa POST path.
func (s BaseService) Post(ctx context.Context, path string, body any, opts ...core.RequestOptions) (core.Json, error) {
	return s.client.Post(ctx, path, body, opts...)
}

// Put executa PUT path.
func (s BaseService) Put(ctx context.Context, path string, body any, opts ...core.RequestOptions) (core.Json, error) {
	return s.client.Put(ctx, path, body, opts...)
}

// Patch executa PATCH path.
func (s BaseService) Patch(ctx context.Context, path string, body any, opts ...core.RequestOptions) (core.Json, error) {
	return s.client.Patch(ctx, path, body, opts...)
}

// Delete executa DELETE path.
func (s BaseService) Delete(ctx context.Context, path string, body any, opts ...core.RequestOptions) (core.Json, error) {
	return s.client.Delete(ctx, path, body, opts...)
}

// Download baixa os bytes crus de uma rota (PDF, imagens...).
func (s BaseService) Download(ctx context.Context, path string, opts ...core.RequestOptions) ([]byte, error) {
	return s.client.Bytes(ctx, core.MethodGet, path, nil, opts...)
}
