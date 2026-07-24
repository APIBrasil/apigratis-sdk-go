package core

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Kind é a categoria de uma falha da SDK.
type Kind string

// Categorias de erro. Use os sentinelas Err* com errors.Is.
const (
	KindAPI                 Kind = "api"
	KindNetwork             Kind = "network"
	KindTimeout             Kind = "timeout"
	KindValidation          Kind = "validation"
	KindAuthentication      Kind = "authentication"
	KindInsufficientBalance Kind = "insufficient_balance"
	KindPermission          Kind = "permission"
	KindNotFound            Kind = "not_found"
	KindRateLimit           Kind = "rate_limit"
	KindServer              Kind = "server"
)

// Sentinelas para errors.Is — o equivalente Go às subclasses de erro das
// demais SDKs:
//
//	if errors.Is(err, core.ErrInsufficientBalance) { ... }
//
// ErrAPIBrasil casa com qualquer erro produzido pela SDK, e ErrNetwork
// também casa com timeouts (assim como TimeoutError estende NetworkError
// nas outras SDKs).
var (
	// ErrAPIBrasil casa com qualquer erro da SDK.
	ErrAPIBrasil = errors.New("apibrasil: erro da API")
	// ErrNetwork — falha de rede: a requisição pode não ter chegado ao servidor.
	ErrNetwork = errors.New("apibrasil: falha de rede")
	// ErrTimeout — a requisição pode ter sido processada; a SDK não faz
	// retry automático para não duplicar cobranças/envios.
	ErrTimeout = errors.New("apibrasil: tempo limite excedido")
	// ErrValidation — HTTP 400/422: payload inválido.
	ErrValidation = errors.New("apibrasil: payload inválido")
	// ErrAuthentication — HTTP 401: Bearer Token ausente, inválido ou expirado.
	ErrAuthentication = errors.New("apibrasil: falha de autenticação")
	// ErrInsufficientBalance — HTTP 402: saldo/créditos insuficientes.
	ErrInsufficientBalance = errors.New("apibrasil: saldo insuficiente")
	// ErrPermission — HTTP 403: sem permissão (ex: API exige conta PJ).
	ErrPermission = errors.New("apibrasil: sem permissão")
	// ErrNotFound — HTTP 404/410: recurso não encontrado ou desativado.
	ErrNotFound = errors.New("apibrasil: não encontrado")
	// ErrRateLimit — HTTP 429: rate limit atingido.
	ErrRateLimit = errors.New("apibrasil: limite de requisições atingido")
	// ErrServer — HTTP 5xx: erro interno do gateway/provedor.
	ErrServer = errors.New("apibrasil: erro interno do servidor")
)

var kindSentinels = map[Kind]error{
	KindAPI:                 ErrAPIBrasil,
	KindNetwork:             ErrNetwork,
	KindTimeout:             ErrTimeout,
	KindValidation:          ErrValidation,
	KindAuthentication:      ErrAuthentication,
	KindInsufficientBalance: ErrInsufficientBalance,
	KindPermission:          ErrPermission,
	KindNotFound:            ErrNotFound,
	KindRateLimit:           ErrRateLimit,
	KindServer:              ErrServer,
}

// Error é o erro devolvido por todas as chamadas da SDK.
//
// Cada categoria de falha tem o seu sentinela (ErrValidation,
// ErrAuthentication, ErrInsufficientBalance, ErrPermission, ErrNotFound,
// ErrRateLimit, ErrServer, ErrNetwork, ErrTimeout) — trate com
// errors.Is, e use errors.As para chegar aos detalhes:
//
//	var apiErr *core.Error
//	if errors.As(err, &apiErr) { fmt.Println(apiErr.Status, apiErr.Code) }
type Error struct {
	// Kind é a categoria da falha.
	Kind Kind
	// Message é a mensagem de erro — a da API quando disponível.
	Message string
	// Status é o status HTTP retornado pela API (ex: 401, 402, 404).
	Status int
	// Code é o código de erro retornado pela API (ex: NOT_FOUND).
	Code string
	// Response é o corpo completo da resposta de erro, quando existir.
	Response any
	// RetryAfter é a espera sugerida pelo servidor (header Retry-After),
	// preenchida em HTTP 429.
	RetryAfter time.Duration
	// Err é o erro original que causou esta falha, quando houver.
	Err error
}

// Error implementa a interface error.
func (e *Error) Error() string {
	var builder strings.Builder
	builder.WriteString(e.Message)
	if e.Status != 0 {
		fmt.Fprintf(&builder, " (HTTP %d)", e.Status)
	}
	if e.Code != "" {
		fmt.Fprintf(&builder, " [%s]", e.Code)
	}
	return builder.String()
}

// Unwrap devolve o erro original, permitindo errors.Is/As na causa.
func (e *Error) Unwrap() error { return e.Err }

// Is conecta o erro aos sentinelas Err* da SDK.
func (e *Error) Is(target error) bool {
	if target == ErrAPIBrasil {
		return true
	}
	if target == ErrNetwork && e.Kind == KindTimeout {
		return true
	}
	if sentinel, ok := kindSentinels[e.Kind]; ok {
		return sentinel == target
	}
	return false
}

// IsInsufficientBalance informa se a falha foi por saldo/créditos
// insuficientes (HTTP 402).
func (e *Error) IsInsufficientBalance() bool { return e.Status == http.StatusPaymentRequired }

// IsUnauthorized informa se a falha foi de autenticação (HTTP 401).
func (e *Error) IsUnauthorized() bool { return e.Status == http.StatusUnauthorized }

// AsError extrai o *Error de uma cadeia de erros.
func AsError(err error) (*Error, bool) {
	var apiErr *Error
	if errors.As(err, &apiErr) {
		return apiErr, true
	}
	return nil, false
}

// FromError converte qualquer erro em *Error, preservando a causa.
func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if apiErr, ok := AsError(err); ok {
		return apiErr
	}
	return &Error{Kind: KindAPI, Message: err.Error(), Err: err}
}

// NewNetworkError cria uma falha de rede (nenhuma resposta recebida).
func NewNetworkError(message string, cause error) *Error {
	return &Error{Kind: KindNetwork, Message: message, Err: cause}
}

// NewTimeoutError cria uma falha por tempo limite excedido.
func NewTimeoutError(message string, cause error) *Error {
	return &Error{Kind: KindTimeout, Message: message, Err: cause}
}

// NewAPIError mapeia um status HTTP + corpo de erro para o *Error
// adequado, replicando a hierarquia de erros das demais SDKs.
func NewAPIError(status int, data any, headers http.Header, cause error) *Error {
	err := &Error{
		Kind:     kindForStatus(status),
		Message:  extractMessage(status, data),
		Status:   status,
		Code:     extractCode(data),
		Response: data,
		Err:      cause,
	}
	if err.Kind == KindRateLimit {
		err.RetryAfter = ParseRetryAfter(headers, time.Now())
	}
	return err
}

func kindForStatus(status int) Kind {
	switch {
	case status == http.StatusBadRequest, status == http.StatusUnprocessableEntity:
		return KindValidation
	case status == http.StatusUnauthorized:
		return KindAuthentication
	case status == http.StatusPaymentRequired:
		return KindInsufficientBalance
	case status == http.StatusForbidden:
		return KindPermission
	case status == http.StatusNotFound, status == http.StatusGone:
		return KindNotFound
	case status == http.StatusTooManyRequests:
		return KindRateLimit
	case status >= 500:
		return KindServer
	default:
		return KindAPI
	}
}

func extractMessage(status int, data any) string {
	if body, ok := data.(map[string]any); ok {
		if message, ok := body["message"].(string); ok && message != "" {
			return message
		}
		if message, ok := body["error"].(string); ok && message != "" {
			return message
		}
	}
	return fmt.Sprintf("A API respondeu com HTTP %d.", status)
}

func extractCode(data any) string {
	if body, ok := data.(map[string]any); ok {
		if code, ok := body["code"].(string); ok {
			return code
		}
	}
	return ""
}

// ParseRetryAfter lê o header Retry-After (segundos ou data HTTP).
func ParseRetryAfter(headers http.Header, now time.Time) time.Duration {
	if headers == nil {
		return 0
	}
	raw := strings.TrimSpace(headers.Get("Retry-After"))
	if raw == "" {
		return 0
	}

	if seconds, err := strconv.ParseFloat(raw, 64); err == nil {
		if seconds <= 0 {
			return 0
		}
		return time.Duration(seconds * float64(time.Second))
	}

	if at, err := http.ParseTime(raw); err == nil {
		if delta := at.Sub(now); delta > 0 {
			return delta
		}
		return 0
	}

	return 0
}
