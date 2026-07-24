package platform

import (
	"context"
	"net/url"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// CatalogService é o catálogo público da plataforma: APIs, planos,
// documentações e servidores. A maioria das rotas é pública (não exige
// Bearer Token).
type CatalogService struct {
	services.BaseService
}

// NewCatalogService cria o serviço de catálogo.
func NewCatalogService(client *core.HTTPClient) *CatalogService {
	return &CatalogService{BaseService: services.NewBaseService(client)}
}

// APIs lista/busca as APIs disponíveis: GET /apis?search=.
func (s *CatalogService) APIs(ctx context.Context, search string, opts ...core.RequestOptions) (core.Json, error) {
	options := core.FirstOption(opts)
	if search != "" {
		options = options.WithQuery(core.Json{"search": search})
	}
	return s.Get(ctx, "apis", options)
}

// API detalha uma API por identificador: GET /apis/{identifier}.
func (s *CatalogService) API(ctx context.Context, identifier string, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "apis/"+identifier, opts...)
}

// APIByName detalha uma API pelo nome: GET /apis/name/{name}.
func (s *CatalogService) APIByName(ctx context.Context, name string, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "apis/name/"+url.PathEscape(name), opts...)
}

// APICategories lista as categorias de APIs: GET /apis/categories.
func (s *CatalogService) APICategories(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "apis/categories", opts...)
}

// MyAPIs lista as APIs contratadas pelo usuário (autenticado):
// GET /apis/list.
func (s *CatalogService) MyAPIs(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "apis/list", opts...)
}

// APIsByDevice lista as APIs vinculadas a um device:
// GET /apis/device/{device_token}.
func (s *CatalogService) APIsByDevice(ctx context.Context, deviceToken string, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "apis/device/"+deviceToken, opts...)
}

// Plans lista os planos disponíveis: GET /plans.
func (s *CatalogService) Plans(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "plans", opts...)
}

// Documentations lista as documentações: GET /documentations.
func (s *CatalogService) Documentations(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "documentations", opts...)
}

// DocumentationsByServer lista a documentação por servidor:
// GET /documentations/server/{server_search}.
func (s *CatalogService) DocumentationsByServer(ctx context.Context, serverSearch string, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "documentations/server/"+serverSearch, opts...)
}

// Servers lista os servidores disponíveis: GET /servers.
func (s *CatalogService) Servers(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "servers", opts...)
}

// EndpointURL resolve a URL de uma action (descoberta dinâmica de
// endpoints): POST /endpoint/url.
func (s *CatalogService) EndpointURL(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "endpoint/url", body, opts...)
}

// EndpointBody devolve o body esperado por uma action:
// POST /endpoint/body.
func (s *CatalogService) EndpointBody(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "endpoint/body", body, opts...)
}

// Status devolve o status do gateway: GET /status.
func (s *CatalogService) Status(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "status", opts...)
}
