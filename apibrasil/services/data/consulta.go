package data

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// ConsultaService são as consultas por crédito
// (/consulta/{servico}/credits e afins).
//
// Não usam DeviceToken — debitam o saldo/créditos da conta. O campo tipo
// do body define o produto consultado (ex: lista-socios,
// serasa-score-pf, spc-serasa); os tipos conhecidos estão em
// generated.ConsultaTipos. Use core.Consulta{Homolog: true} para sandbox.
type ConsultaService struct {
	services.CreditService
}

// NewConsultaService cria o serviço de consultas por crédito.
func NewConsultaService(client *core.HTTPClient) *ConsultaService {
	return &ConsultaService{CreditService: services.NewCreditService(client)}
}

// Generic executa uma consulta genérica: POST /consulta/{servico}/credits.
func (s *ConsultaService) Generic(ctx context.Context, service string, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, service, body, opts...)
}

// CPF consulta um CPF: POST /consulta/cpf/credits.
func (s *ConsultaService) CPF(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "cpf", body, opts...)
}

// CNPJ consulta um CNPJ: POST /consulta/cnpj/credits
// (ex: {"cnpj": "...", "tipo": "lista-socios"}).
func (s *ConsultaService) CNPJ(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "cnpj", body, opts...)
}

// CNH consulta uma CNH: POST /consulta/cnh/credits.
func (s *ConsultaService) CNH(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "cnh", body, opts...)
}

// CEP consulta um CEP: POST /consulta/cep/credits.
func (s *ConsultaService) CEP(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "cep", body, opts...)
}

// Veiculos executa a consulta veicular: POST /consulta/veiculos/credits.
func (s *ConsultaService) Veiculos(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "veiculos", body, opts...)
}

// Vehicles executa a consulta veicular na rota em inglês:
// POST /consulta/vehicles/credits.
func (s *ConsultaService) Vehicles(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "vehicles", body, opts...)
}

// Telefone consulta a operadora de um telefone:
// POST /consulta/telefone/credits.
func (s *ConsultaService) Telefone(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "telefone", body, opts...)
}

// DDDAnatel consulta o DDD na Anatel: POST /consulta/ddd-anatel/credits.
func (s *ConsultaService) DDDAnatel(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "ddd-anatel", body, opts...)
}

// GeoIP consulta a base de IPs: POST /consulta/geoip/credits.
func (s *ConsultaService) GeoIP(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "geoip", body, opts...)
}

// Rastreio rastreia uma encomenda: POST /consulta/rastreio/credits.
func (s *ConsultaService) Rastreio(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "rastreio", body, opts...)
}

// CRM consulta um registro no CRM: POST /consulta/crm/credits.
func (s *ConsultaService) CRM(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "crm", body, opts...)
}

// CRBM consulta um registro no CRBM: POST /consulta/crbm/credits.
func (s *ConsultaService) CRBM(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "crbm", body, opts...)
}

// CRO consulta um registro no CRO: POST /consulta/cro/credits.
func (s *ConsultaService) CRO(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "cro", body, opts...)
}

// Weather consulta a previsão do tempo:
// POST /consulta/weather-api/credits.
func (s *ConsultaService) Weather(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "weather-api", body, opts...)
}

// EmissaoNotas emite notas fiscais:
// POST /consulta/emissao-notas/credits.
func (s *ConsultaService) EmissaoNotas(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "emissao-notas", body, opts...)
}

// FreteAntt calcula o frete ANTT: POST /consulta/frete-antt/credits.
func (s *ConsultaService) FreteAntt(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "frete-antt", body, opts...)
}

// APIRntrc consulta o RNTRC: POST /consulta/api-rntrc/credits.
func (s *ConsultaService) APIRntrc(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return s.Request(ctx, "api-rntrc", body, opts...)
}

// Quod consulta o birô Quod: POST /quod/cnpj/credits.
func (s *ConsultaService) Quod(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return services.Credit(s.Post(ctx, "quod/cnpj/credits", body, opts...))
}

// VeiculosBase consulta as bases veiculares dedicadas:
// POST /vehicles/base/000/{base} (base: dados | fipe).
func (s *ConsultaService) VeiculosBase(ctx context.Context, base string, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return services.Credit(s.Post(ctx, "vehicles/base/000/"+base, body, opts...))
}

// CepDistancia calcula a distância entre CEPs:
// POST /cep/distancia/calcular.
func (s *ConsultaService) CepDistancia(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return services.Credit(s.Post(ctx, "cep/distancia/calcular", body, opts...))
}

// ProxySeller consulta o proxy seller: POST /proxy/seller/credits.
func (s *ConsultaService) ProxySeller(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return services.Credit(s.Post(ctx, "proxy/seller/credits", body, opts...))
}
