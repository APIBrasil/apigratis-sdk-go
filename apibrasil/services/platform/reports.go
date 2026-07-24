package platform

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// ReportsService são os relatórios e o dashboard de consumo
// (/reports/*, /dashboard/stats).
type ReportsService struct {
	services.BaseService
}

// NewReportsService cria o serviço de relatórios.
func NewReportsService(client *core.HTTPClient) *ReportsService {
	return &ReportsService{BaseService: services.NewBaseService(client)}
}

// DashboardStats devolve as estatísticas do dashboard:
// GET /dashboard/stats.
func (s *ReportsService) DashboardStats(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "dashboard/stats", opts...)
}

// Consumption devolve o consumo: GET /reports/consumption.
func (s *ReportsService) Consumption(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "reports/consumption", opts...)
}

// GenerateConsumptionReport gera o relatório de consumo:
// POST /reports/generate-consumption-report.
func (s *ReportsService) GenerateConsumptionReport(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "reports/generate-consumption-report", body, opts...)
}

// Extract devolve o extrato: GET /reports/extract.
func (s *ReportsService) Extract(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "reports/extract", opts...)
}

// Dashboard devolve o dashboard de relatórios: GET /reports/dashboard.
func (s *ReportsService) Dashboard(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "reports/dashboard", opts...)
}

// Summary devolve o resumo: GET /reports/summary.
func (s *ReportsService) Summary(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "reports/summary", opts...)
}

// DailyUsage devolve o uso diário: GET /reports/daily-usage.
func (s *ReportsService) DailyUsage(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "reports/daily-usage", opts...)
}

// MonthlySummary devolve o resumo mensal:
// GET /reports/monthly-summary.
func (s *ReportsService) MonthlySummary(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "reports/monthly-summary", opts...)
}

// ErrorAnalysis devolve a análise de erros:
// GET /reports/error-analysis.
func (s *ReportsService) ErrorAnalysis(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "reports/error-analysis", opts...)
}

// DeviceAnalysis devolve a análise por device:
// GET /reports/device-analysis.
func (s *ReportsService) DeviceAnalysis(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "reports/device-analysis", opts...)
}

// RecentRequests devolve as requisições recentes:
// GET /reports/recent-requests.
func (s *ReportsService) RecentRequests(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "reports/recent-requests", opts...)
}

// QuickStats devolve as estatísticas rápidas:
// GET /reports/quick-stats.
func (s *ReportsService) QuickStats(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "reports/quick-stats", opts...)
}
