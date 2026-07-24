package services

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
)

// CreditService é a base das consultas por crédito
// (POST /consulta/{servico}/credits).
//
// Não usam DeviceToken — debitam o saldo/créditos da conta.
type CreditService struct {
	BaseService
}

// NewCreditService cria a base das consultas por crédito.
func NewCreditService(client *core.HTTPClient) CreditService {
	return CreditService{BaseService: NewBaseService(client)}
}

// Request executa uma consulta: POST /consulta/{servico}/credits.
func (s CreditService) Request(ctx context.Context, service string, body core.Json, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return Credit(s.Post(ctx, "consulta/"+service+"/credits", body, opts...))
}

// Credits consulta os créditos disponíveis de um serviço:
// GET /consulta/{servico}/credits.
func (s CreditService) Credits(ctx context.Context, service string, opts ...core.RequestOptions) (core.CreditServiceResponse, error) {
	return Credit(s.Get(ctx, "consulta/"+service+"/credits", opts...))
}
