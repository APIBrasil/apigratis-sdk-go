// Package core reúne a infraestrutura da SDK: cliente HTTP, transporte
// plugável, política de retry, hooks de observabilidade, hierarquia de
// erros e os tipos compartilhados por todos os serviços.
package core

import (
	"context"
	"encoding/json"
	"time"
)

// Json é o objeto JSON decodificado — o formato de resposta da maioria
// das rotas do gateway.
type Json = map[string]any

// Method é um verbo HTTP usado pelo gateway.
type Method string

// Verbos HTTP suportados pela SDK.
const (
	MethodGet    Method = "GET"
	MethodPost   Method = "POST"
	MethodPut    Method = "PUT"
	MethodPatch  Method = "PATCH"
	MethodDelete Method = "DELETE"
)

// ResponseType define como o corpo da resposta deve ser decodificado.
type ResponseType int

const (
	// ResponseJSON decodifica JSON (padrão). Corpos não-JSON voltam como string.
	ResponseJSON ResponseType = iota
	// ResponseBytes devolve os bytes crus — use para PDFs, imagens etc.
	ResponseBytes
)

// RetryConfig é a política de retry do cliente. Por padrão a SDK tenta
// novamente apenas em HTTP 429 (rate limit) e em falhas de conexão —
// nunca em timeouts ou erros de negócio, para não duplicar cobranças/envios.
type RetryConfig struct {
	// Retries é o número de novas tentativas além da original. Padrão: 2.
	Retries int
	// MinDelay é o atraso base do backoff exponencial. Padrão: 300ms.
	MinDelay time.Duration
	// MaxDelay é o teto do atraso entre tentativas. Padrão: 5s.
	MaxDelay time.Duration
	// RetryOnStatuses são os status HTTP que disparam retry. Padrão: []int{429}.
	RetryOnStatuses []int
}

// RequestHookInfo são os dados da requisição entregues a Hooks.OnRequest.
type RequestHookInfo struct {
	Method  Method
	URL     string
	Headers map[string]string
	Body    any
	// Attempt é a tentativa atual (0 = primeira).
	Attempt int
}

// ResponseHookInfo são os dados da resposta entregues a Hooks.OnResponse.
type ResponseHookInfo struct {
	Method   Method
	URL      string
	Status   int
	Duration time.Duration
	Attempt  int
}

// RetryHookInfo são os dados do retry entregues a Hooks.OnRetry.
type RetryHookInfo struct {
	Method Method
	URL    string
	// Attempt é o número da próxima tentativa.
	Attempt int
	Delay   time.Duration
	Reason  string
}

// Hooks são os ganchos de observabilidade — logging, métricas, tracing.
type Hooks struct {
	OnRequest  func(ctx context.Context, info RequestHookInfo)
	OnResponse func(ctx context.Context, info ResponseHookInfo)
	OnRetry    func(ctx context.Context, info RetryHookInfo)
}

// Config é a configuração do cliente ApiBrasil.
//
// Campos não informados são lidos das variáveis de ambiente
// APIBRASIL_BEARER_TOKEN, APIBRASIL_DEVICE_TOKEN, APIBRASIL_SECRET_KEY
// e APIBRASIL_BASE_URL.
type Config struct {
	// BearerToken é o token JWT obtido no login (Authorization: Bearer <token>).
	BearerToken string
	// DeviceToken é o token do dispositivo, exigido pelos serviços
	// device-based (WhatsApp, SMS, veículos...).
	DeviceToken string
	// SecretKey é a SecretKey da API (usada apenas na criação de devices).
	SecretKey string
	// BaseURL é a base da API. Padrão: https://gateway.apibrasil.io/api/v2.
	BaseURL string
	// Timeout das requisições. Padrão: 30s.
	Timeout time.Duration
	// Headers adicionais enviados em todas as requisições.
	Headers map[string]string
	// Transport HTTP customizado. Padrão: HTTPTransport (net/http).
	Transport Transport
	// Retry é a política de retry. nil usa DefaultRetry();
	// NoRetry() desativa.
	Retry *RetryConfig
	// Hooks de observabilidade.
	Hooks *Hooks
}

// Merge devolve uma cópia desta configuração sobreposta por other — os
// campos preenchidos em other têm prioridade.
func (c Config) Merge(other Config) Config {
	merged := c

	if other.BearerToken != "" {
		merged.BearerToken = other.BearerToken
	}
	if other.DeviceToken != "" {
		merged.DeviceToken = other.DeviceToken
	}
	if other.SecretKey != "" {
		merged.SecretKey = other.SecretKey
	}
	if other.BaseURL != "" {
		merged.BaseURL = other.BaseURL
	}
	if other.Timeout != 0 {
		merged.Timeout = other.Timeout
	}
	if other.Headers != nil {
		merged.Headers = other.Headers
	}
	if other.Transport != nil {
		merged.Transport = other.Transport
	}
	if other.Retry != nil {
		merged.Retry = other.Retry
	}
	if other.Hooks != nil {
		merged.Hooks = other.Hooks
	}

	return merged
}

// RequestOptions são as opções por requisição — sobrescrevem a
// configuração do cliente. Os métodos dos serviços recebem esse tipo de
// forma variádica, então podem ser chamados sem nenhuma opção.
type RequestOptions struct {
	// Query string da requisição. Valores nil são ignorados.
	Query map[string]any
	// Headers extras desta requisição.
	Headers     map[string]string
	BearerToken string
	DeviceToken string
	SecretKey   string
	Timeout     time.Duration
	// ResponseType define como decodificar a resposta. Padrão: ResponseJSON.
	ResponseType ResponseType
}

// FirstOption devolve a primeira opção informada — atalho usado pelos
// serviços, que recebem RequestOptions de forma variádica.
func FirstOption(opts []RequestOptions) RequestOptions {
	if len(opts) == 0 {
		return RequestOptions{}
	}
	return opts[0]
}

// WithQuery devolve uma cópia com extra mesclado à query, mantendo as
// chaves já definidas nesta instância.
func (o RequestOptions) WithQuery(extra map[string]any) RequestOptions {
	query := make(map[string]any, len(extra)+len(o.Query))
	for key, value := range extra {
		query[key] = value
	}
	for key, value := range o.Query {
		query[key] = value
	}
	o.Query = query
	return o
}

// DeviceServiceResponse é o envelope de resposta dos serviços
// device-based ({ error, message, response, api_limit... }).
//
// É um map[string]any em tempo de execução — use os acessores tipados ou
// o acesso por chave (res["response"]), como preferir.
type DeviceServiceResponse Json

// IsError informa se o gateway sinalizou erro no envelope.
func (r DeviceServiceResponse) IsError() bool { return r["error"] == true }

// Message é a mensagem devolvida pelo gateway.
func (r DeviceServiceResponse) Message() string { return stringValue(r["message"]) }

// Response é o payload do provedor (campo response).
func (r DeviceServiceResponse) Response() any { return r["response"] }

// APILimit é o limite de requisições do plano.
func (r DeviceServiceResponse) APILimit() any { return r["api_limit"] }

// APILimitFor é a janela do limite (ex: day).
func (r DeviceServiceResponse) APILimitFor() any { return r["api_limit_for"] }

// APILimitUsed é quanto do limite já foi consumido.
func (r DeviceServiceResponse) APILimitUsed() any { return r["api_limit_used"] }

// Decode decodifica o campo response em v (struct, map, slice...).
func (r DeviceServiceResponse) Decode(v any) error { return DecodeValue(r["response"], v) }

// Json devolve o envelope como map genérico.
func (r DeviceServiceResponse) Json() Json { return Json(r) }

// CreditServiceResponse é o envelope de resposta das consultas por
// crédito ({ error, message, balance, tax, valor_consulta, data... }).
type CreditServiceResponse Json

// IsError informa se o gateway sinalizou erro no envelope.
func (r CreditServiceResponse) IsError() bool { return r["error"] == true }

// Message é a mensagem devolvida pelo gateway.
func (r CreditServiceResponse) Message() string { return stringValue(r["message"]) }

// Balance é o saldo restante após a consulta.
func (r CreditServiceResponse) Balance() any { return r["balance"] }

// Tax é a taxa aplicada.
func (r CreditServiceResponse) Tax() any { return r["tax"] }

// ValorConsulta é o valor cobrado pela consulta.
func (r CreditServiceResponse) ValorConsulta() any { return r["valor_consulta"] }

// Homolog informa se a resposta veio do modo homologação.
func (r CreditServiceResponse) Homolog() bool { return r["homolog"] == true }

// Data são os dados da consulta.
func (r CreditServiceResponse) Data() any { return r["data"] }

// Decode decodifica o campo data em v (struct, map, slice...).
func (r CreditServiceResponse) Decode(v any) error { return DecodeValue(r["data"], v) }

// Json devolve o envelope como map genérico.
func (r CreditServiceResponse) Json() Json { return Json(r) }

// Consulta são os campos comuns aceitos pelas consultas por crédito
// (/consulta/{servico}/credits). Fields aceita qualquer chave extra — o
// gateway define os campos por produto.
//
// Veja generated.ConsultaTipos para os tipos conhecidos.
type Consulta struct {
	// Tipo é o tipo/serviço da consulta (ex: lista-socios, serasa-score-pf).
	Tipo string
	// Homolog ativa o modo homologação (sandbox, sem cobrança) quando suportado.
	Homolog bool
	// Lite pede a versão reduzida da consulta, quando suportada.
	Lite bool
	// Agrupados são as consultas agrupadas.
	Agrupados []string
	// Extra são os serviços extras.
	Extra []string
	// Fields são os campos específicos do produto (ex: {"cnpj": "..."}).
	Fields Json
}

// Json monta o body da consulta.
func (c Consulta) Json() Json {
	body := Json{}
	if c.Tipo != "" {
		body["tipo"] = c.Tipo
	}
	if c.Homolog {
		body["homolog"] = true
	}
	if c.Lite {
		body["lite"] = true
	}
	if len(c.Agrupados) > 0 {
		body["agrupados"] = c.Agrupados
	}
	if len(c.Extra) > 0 {
		body["extra"] = c.Extra
	}
	for key, value := range c.Fields {
		body[key] = value
	}
	return body
}

// DecodeValue converte um valor já decodificado (map, slice, string...)
// em v, usando o encoding/json como intermediário.
func DecodeValue(value any, v any) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, v)
}

func stringValue(value any) string {
	if text, ok := value.(string); ok {
		return text
	}
	return ""
}
