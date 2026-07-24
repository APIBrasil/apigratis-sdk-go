package apibrasil

import (
	"net/http"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
)

// Tipos do core reexportados para que o uso do dia a dia precise de um
// único import.
type (
	// Json é o objeto JSON decodificado devolvido pelas rotas.
	Json = core.Json
	// Config é a configuração do cliente.
	Config = core.Config
	// RequestOptions são as opções por requisição.
	RequestOptions = core.RequestOptions
	// RetryConfig é a política de retry.
	RetryConfig = core.RetryConfig
	// Hooks são os ganchos de observabilidade.
	Hooks = core.Hooks
	// RequestHookInfo são os dados entregues a Hooks.OnRequest.
	RequestHookInfo = core.RequestHookInfo
	// ResponseHookInfo são os dados entregues a Hooks.OnResponse.
	ResponseHookInfo = core.ResponseHookInfo
	// RetryHookInfo são os dados entregues a Hooks.OnRetry.
	RetryHookInfo = core.RetryHookInfo
	// Method é um verbo HTTP.
	Method = core.Method
	// ResponseType define como decodificar o corpo da resposta.
	ResponseType = core.ResponseType
	// Transport é a camada de transporte HTTP plugável.
	Transport = core.Transport
	// TransportRequest é a requisição entregue ao transporte.
	TransportRequest = core.TransportRequest
	// TransportResponse é a resposta devolvida pelo transporte.
	TransportResponse = core.TransportResponse
	// HTTPTransport é o transporte padrão (net/http).
	HTTPTransport = core.HTTPTransport
	// HTTPClient é o cliente HTTP interno da SDK.
	HTTPClient = core.HTTPClient
	// Error é o erro devolvido por todas as chamadas da SDK.
	Error = core.Error
	// Kind é a categoria de uma falha.
	Kind = core.Kind
	// DeviceServiceResponse é o envelope dos serviços device-based.
	DeviceServiceResponse = core.DeviceServiceResponse
	// CreditServiceResponse é o envelope das consultas por crédito.
	CreditServiceResponse = core.CreditServiceResponse
	// Consulta monta o body das consultas por crédito.
	Consulta = core.Consulta
)

// Verbos HTTP.
const (
	MethodGet    = core.MethodGet
	MethodPost   = core.MethodPost
	MethodPut    = core.MethodPut
	MethodPatch  = core.MethodPatch
	MethodDelete = core.MethodDelete
)

// Tipos de decodificação da resposta.
const (
	ResponseJSON  = core.ResponseJSON
	ResponseBytes = core.ResponseBytes
)

// Padrões do cliente HTTP.
const (
	DefaultBaseURL = core.DefaultBaseURL
	DefaultTimeout = core.DefaultTimeout
	UserAgent      = core.UserAgent
)

// Variáveis de ambiente reconhecidas pela SDK.
const (
	EnvBearerToken = core.EnvBearerToken
	EnvDeviceToken = core.EnvDeviceToken
	EnvSecretKey   = core.EnvSecretKey
	EnvBaseURL     = core.EnvBaseURL
)

// Sentinelas de erro — use com errors.Is.
var (
	ErrAPIBrasil           = core.ErrAPIBrasil
	ErrNetwork             = core.ErrNetwork
	ErrTimeout             = core.ErrTimeout
	ErrValidation          = core.ErrValidation
	ErrAuthentication      = core.ErrAuthentication
	ErrInsufficientBalance = core.ErrInsufficientBalance
	ErrPermission          = core.ErrPermission
	ErrNotFound            = core.ErrNotFound
	ErrRateLimit           = core.ErrRateLimit
	ErrServer              = core.ErrServer
)

// AsError extrai o *Error de uma cadeia de erros.
func AsError(err error) (*Error, bool) { return core.AsError(err) }

// NoRetry devolve uma política que desativa o retry — passe em Config.Retry.
func NoRetry() *RetryConfig { return core.NoRetry() }

// DefaultRetry devolve a política de retry padrão.
func DefaultRetry() RetryConfig { return core.DefaultRetry() }

// ConfigFromEnv lê a configuração das variáveis de ambiente.
func ConfigFromEnv() Config { return core.ConfigFromEnv() }

// NewHTTPTransport cria o transporte padrão sobre um *http.Client.
func NewHTTPTransport(client *http.Client) *HTTPTransport {
	return core.NewHTTPTransport(client)
}

// NewHTTPClient cria o cliente HTTP interno da SDK.
func NewHTTPClient(config Config) *HTTPClient { return core.NewHTTPClient(config) }
