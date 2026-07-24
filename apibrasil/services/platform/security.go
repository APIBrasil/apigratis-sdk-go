package platform

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// IPWhitelistService é a IP Whitelist da conta (/ip-whitelist/*).
// Restringe de quais IPs o Bearer Token pode ser usado.
type IPWhitelistService struct {
	services.BaseService
}

// NewIPWhitelistService cria o serviço de IP whitelist.
func NewIPWhitelistService(client *core.HTTPClient) *IPWhitelistService {
	return &IPWhitelistService{BaseService: services.NewBaseService(client)}
}

// Get devolve a configuração atual: GET /ip-whitelist.
func (s *IPWhitelistService) Get(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.BaseService.Get(ctx, "ip-whitelist", opts...)
}

// Set substitui a lista inteira: PUT /ip-whitelist.
func (s *IPWhitelistService) Set(ctx context.Context, entries []string, opts ...core.RequestOptions) (core.Json, error) {
	return s.Put(ctx, "ip-whitelist", core.Json{"ip_whitelist": entries}, opts...)
}

// Add adiciona uma entrada: POST /ip-whitelist/add.
func (s *IPWhitelistService) Add(ctx context.Context, entry string, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "ip-whitelist/add", core.Json{"entry": entry}, opts...)
}

// Remove remove uma entrada: DELETE /ip-whitelist/remove.
func (s *IPWhitelistService) Remove(ctx context.Context, entry string, opts ...core.RequestOptions) (core.Json, error) {
	return s.Delete(ctx, "ip-whitelist/remove", core.Json{"entry": entry}, opts...)
}

// AddCurrent adiciona o IP atual: POST /ip-whitelist/add-current.
func (s *IPWhitelistService) AddCurrent(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "ip-whitelist/add-current", nil, opts...)
}

// Reset libera todos os IPs (wildcard): POST /ip-whitelist/reset.
func (s *IPWhitelistService) Reset(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "ip-whitelist/reset", nil, opts...)
}

// Validate valida uma entrada: POST /ip-whitelist/validate.
func (s *IPWhitelistService) Validate(ctx context.Context, entry string, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "ip-whitelist/validate", core.Json{"entry": entry}, opts...)
}

// CurrentIP devolve o IP atual visto pelo gateway:
// GET /ip-whitelist/current-ip.
func (s *IPWhitelistService) CurrentIP(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.BaseService.Get(ctx, "ip-whitelist/current-ip", opts...)
}

// BearerRateLimitService é o rate limit por Bearer Token
// (/bearer-rate-limit).
type BearerRateLimitService struct {
	services.BaseService
}

// NewBearerRateLimitService cria o serviço de rate limit.
func NewBearerRateLimitService(client *core.HTTPClient) *BearerRateLimitService {
	return &BearerRateLimitService{BaseService: services.NewBaseService(client)}
}

// Get devolve o limite atual: GET /bearer-rate-limit.
func (s *BearerRateLimitService) Get(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.BaseService.Get(ctx, "bearer-rate-limit", opts...)
}

// Set define o limite por minuto: PUT /bearer-rate-limit.
func (s *BearerRateLimitService) Set(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Put(ctx, "bearer-rate-limit", body, opts...)
}
