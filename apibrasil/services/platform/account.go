package platform

import (
	"context"
	"fmt"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// AccountService é a conta: saldo, faturas, notificações e tickets.
type AccountService struct {
	services.BaseService
}

// NewAccountService cria o serviço de conta.
func NewAccountService(client *core.HTTPClient) *AccountService {
	return &AccountService{BaseService: services.NewBaseService(client)}
}

// Balance devolve o saldo/créditos da conta: GET /balance.
func (s *AccountService) Balance(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "balance", opts...)
}

// Plan devolve o plano atual: GET /plan.
func (s *AccountService) Plan(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "plan", opts...)
}

// Invoices lista as faturas: GET /invoices.
func (s *AccountService) Invoices(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "invoices", opts...)
}

// InvoiceNotes lista as notas fiscais das faturas: GET /invoices/notes.
func (s *AccountService) InvoiceNotes(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "invoices/notes", opts...)
}

// PayInvoice paga uma fatura: POST /invoices/pay.
func (s *AccountService) PayInvoice(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "invoices/pay", body, opts...)
}

// Requests devolve o histórico de requisições da conta: POST /requests.
func (s *AccountService) Requests(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "requests", body, opts...)
}

// APIRequests devolve o histórico de requisições por API:
// POST /api/requests.
func (s *AccountService) APIRequests(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "api/requests", body, opts...)
}

// Jobs lista os jobs em fila do usuário: GET /jobs.
func (s *AccountService) Jobs(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "jobs", opts...)
}

// Credentials lista as credenciais do usuário: GET /credentials.
func (s *AccountService) Credentials(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "credentials", opts...)
}

// Indications lista as indicações do usuário: GET /indications.
func (s *AccountService) Indications(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "indications", opts...)
}

// Notifications lista as notificações: GET /notifications.
func (s *AccountService) Notifications(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "notifications", opts...)
}

// MarkNotificationRead marca uma notificação como lida:
// PATCH /notifications/{id}/read.
func (s *AccountService) MarkNotificationRead(ctx context.Context, id any, opts ...core.RequestOptions) (core.Json, error) {
	return s.Patch(ctx, fmt.Sprintf("notifications/%v/read", id), nil, opts...)
}

// MarkAllNotificationsRead marca todas as notificações como lidas:
// POST /notifications/mark-all-read.
func (s *AccountService) MarkAllNotificationsRead(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "notifications/mark-all-read", nil, opts...)
}

// Tickets lista os tickets de suporte: GET /tickets.
func (s *AccountService) Tickets(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "tickets", opts...)
}

// CreateTicket abre um ticket: POST /ticket.
func (s *AccountService) CreateTicket(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "ticket", body, opts...)
}

// UpdateTicket atualiza um ticket: PUT /ticket/{id}.
func (s *AccountService) UpdateTicket(ctx context.Context, id any, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Put(ctx, fmt.Sprintf("ticket/%v", id), body, opts...)
}

// TicketMessages lista as mensagens de um ticket:
// GET /ticket/{id}/messages.
func (s *AccountService) TicketMessages(ctx context.Context, id any, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, fmt.Sprintf("ticket/%v/messages", id), opts...)
}

// AddTicketMessage responde um ticket: POST /ticket/{id}/messages.
func (s *AccountService) AddTicketMessage(ctx context.Context, id any, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, fmt.Sprintf("ticket/%v/messages", id), body, opts...)
}
