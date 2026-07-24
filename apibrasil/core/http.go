package core

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"
)

// Padrões do cliente HTTP.
const (
	// DefaultBaseURL é a base da API.
	DefaultBaseURL = "https://gateway.apibrasil.io/api/v2"
	// DefaultTimeout é o tempo limite das requisições.
	DefaultTimeout = 30 * time.Second
	// UserAgent identifica a SDK nas requisições.
	UserAgent = "APIBRASIL/SDK-GO"
)

// HTTPClient é o cliente HTTP interno da SDK. Injeta os headers de
// autenticação da plataforma (Authorization: Bearer, DeviceToken,
// SecretKey), aplica retry com backoff, dispara os hooks de
// observabilidade e converte falhas em *Error.
//
// É seguro para uso concorrente.
type HTTPClient struct {
	config    Config
	transport Transport
	retry     RetryConfig
	hooks     Hooks

	mu          sync.RWMutex
	bearerToken string
	deviceToken string
}

// NewHTTPClient cria o cliente HTTP resolvendo a configuração com as
// variáveis de ambiente (os campos informados em config têm prioridade).
func NewHTTPClient(config Config) *HTTPClient {
	resolved := ConfigFromEnv().Merge(config)

	transport := resolved.Transport
	if transport == nil {
		transport = NewHTTPTransport(nil)
	}

	hooks := Hooks{}
	if resolved.Hooks != nil {
		hooks = *resolved.Hooks
	}

	return &HTTPClient{
		config:      resolved,
		transport:   transport,
		retry:       ResolveRetry(resolved.Retry),
		hooks:       hooks,
		bearerToken: resolved.BearerToken,
		deviceToken: resolved.DeviceToken,
	}
}

// BaseURL devolve a base da API em uso.
func (c *HTTPClient) BaseURL() string {
	if c.config.BaseURL != "" {
		return c.config.BaseURL
	}
	return DefaultBaseURL
}

// Transport devolve o transporte HTTP em uso.
func (c *HTTPClient) Transport() Transport { return c.transport }

// BearerToken devolve o Bearer Token atual.
func (c *HTTPClient) BearerToken() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.bearerToken
}

// SetBearerToken define/atualiza o Bearer Token — string vazia remove a
// autenticação.
func (c *HTTPClient) SetBearerToken(token string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.bearerToken = token
}

// DeviceToken devolve o DeviceToken atual.
func (c *HTTPClient) DeviceToken() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.deviceToken
}

// SetDeviceToken define/atualiza o DeviceToken — string vazia remove o header.
func (c *HTTPClient) SetDeviceToken(token string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.deviceToken = token
}

// SecretKey devolve a SecretKey configurada (usada em devices.Store).
func (c *HTTPClient) SecretKey() string { return c.config.SecretKey }

// Config devolve a configuração atual do cliente, já resolvida com o
// ambiente e com os tokens em vigor.
func (c *HTTPClient) Config() Config {
	config := c.config
	config.BearerToken = c.BearerToken()
	config.DeviceToken = c.DeviceToken()
	return config
}

// Do executa uma requisição e devolve o corpo já decodificado (map,
// slice, string ou []byte, conforme a resposta e o ResponseType).
func (c *HTTPClient) Do(ctx context.Context, method Method, path string, body any, opts ...RequestOptions) (any, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	options := FirstOption(opts)

	headers := c.buildHeaders(options)
	target := JoinURL(c.BaseURL(), path) + BuildQueryString(options.Query)

	payload, err := serializeBody(body)
	if err != nil {
		return nil, &Error{
			Kind:    KindValidation,
			Message: fmt.Sprintf("Não foi possível serializar o corpo da requisição: %v", err),
			Err:     err,
		}
	}

	timeout := options.Timeout
	if timeout == 0 {
		timeout = c.config.Timeout
	}
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	maxAttempts := 1 + c.retry.Retries

	for attempt := 0; ; {
		if c.hooks.OnRequest != nil {
			c.hooks.OnRequest(ctx, RequestHookInfo{
				Method:  method,
				URL:     target,
				Headers: headers,
				Body:    body,
				Attempt: attempt,
			})
		}

		startedAt := time.Now()
		response, err := c.transport.Do(ctx, TransportRequest{
			Method:       method,
			URL:          target,
			Headers:      headers,
			Body:         payload,
			Timeout:      timeout,
			ResponseType: options.ResponseType,
		})

		if err != nil {
			apiErr := FromError(err)
			// Timeouts nunca são refeitos: a requisição pode ter sido
			// processada e o retry duplicaria cobranças/envios.
			retryable := apiErr.Kind == KindNetwork && ctx.Err() == nil
			if retryable && attempt+1 < maxAttempts {
				delay := BackoffDelay(attempt, c.retry)
				attempt++
				c.notifyRetry(ctx, method, target, attempt, delay, apiErr.Message)
				if sleepErr := Sleep(ctx, delay); sleepErr != nil {
					return nil, apiErr
				}
				continue
			}
			return nil, apiErr
		}

		if c.hooks.OnResponse != nil {
			c.hooks.OnResponse(ctx, ResponseHookInfo{
				Method:   method,
				URL:      target,
				Status:   response.Status,
				Duration: time.Since(startedAt),
				Attempt:  attempt,
			})
		}

		if response.Status >= 400 {
			apiErr := NewAPIError(response.Status, response.Data, response.Headers, nil)
			if c.retry.RetriesStatus(response.Status) && attempt+1 < maxAttempts && ctx.Err() == nil {
				delay := apiErr.RetryAfter
				if delay <= 0 {
					delay = BackoffDelay(attempt, c.retry)
				}
				attempt++
				c.notifyRetry(ctx, method, target, attempt, delay, fmt.Sprintf("HTTP %d", response.Status))
				if sleepErr := Sleep(ctx, delay); sleepErr != nil {
					return nil, apiErr
				}
				continue
			}
			return nil, apiErr
		}

		return response.Data, nil
	}
}

// RequestJSON executa a requisição e devolve o corpo como objeto JSON.
func (c *HTTPClient) RequestJSON(ctx context.Context, method Method, path string, body any, opts ...RequestOptions) (Json, error) {
	data, err := c.Do(ctx, method, path, body, opts...)
	if err != nil {
		return nil, err
	}
	return AsJSONMap(data), nil
}

// Get executa GET path.
func (c *HTTPClient) Get(ctx context.Context, path string, opts ...RequestOptions) (Json, error) {
	return c.RequestJSON(ctx, MethodGet, path, nil, opts...)
}

// Post executa POST path.
func (c *HTTPClient) Post(ctx context.Context, path string, body any, opts ...RequestOptions) (Json, error) {
	return c.RequestJSON(ctx, MethodPost, path, body, opts...)
}

// Put executa PUT path.
func (c *HTTPClient) Put(ctx context.Context, path string, body any, opts ...RequestOptions) (Json, error) {
	return c.RequestJSON(ctx, MethodPut, path, body, opts...)
}

// Patch executa PATCH path.
func (c *HTTPClient) Patch(ctx context.Context, path string, body any, opts ...RequestOptions) (Json, error) {
	return c.RequestJSON(ctx, MethodPatch, path, body, opts...)
}

// Delete executa DELETE path.
func (c *HTTPClient) Delete(ctx context.Context, path string, body any, opts ...RequestOptions) (Json, error) {
	return c.RequestJSON(ctx, MethodDelete, path, body, opts...)
}

// Bytes baixa o corpo cru (PDF de boleto, imagens...).
func (c *HTTPClient) Bytes(ctx context.Context, method Method, path string, body any, opts ...RequestOptions) ([]byte, error) {
	options := FirstOption(opts)
	options.ResponseType = ResponseBytes

	data, err := c.Do(ctx, method, path, body, options)
	if err != nil {
		return nil, err
	}

	switch value := data.(type) {
	case []byte:
		return value, nil
	case string:
		return []byte(value), nil
	case nil:
		return nil, nil
	default:
		return json.Marshal(value)
	}
}

// Into executa a requisição e decodifica a resposta em v.
func (c *HTTPClient) Into(ctx context.Context, method Method, path string, body any, v any, opts ...RequestOptions) error {
	data, err := c.Do(ctx, method, path, body, opts...)
	if err != nil {
		return err
	}
	return DecodeValue(data, v)
}

func (c *HTTPClient) buildHeaders(options RequestOptions) map[string]string {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
		"User-Agent":   UserAgent,
	}
	for key, value := range c.config.Headers {
		headers[key] = value
	}

	bearerToken := options.BearerToken
	if bearerToken == "" {
		bearerToken = c.BearerToken()
	}
	if bearerToken != "" {
		headers["Authorization"] = "Bearer " + bearerToken
	}

	deviceToken := options.DeviceToken
	if deviceToken == "" {
		deviceToken = c.DeviceToken()
	}
	if deviceToken != "" {
		headers["DeviceToken"] = deviceToken
	}

	if options.SecretKey != "" {
		headers["SecretKey"] = options.SecretKey
	}

	for key, value := range options.Headers {
		headers[key] = value
	}

	return headers
}

func (c *HTTPClient) notifyRetry(ctx context.Context, method Method, url string, attempt int, delay time.Duration, reason string) {
	if c.hooks.OnRetry == nil {
		return
	}
	c.hooks.OnRetry(ctx, RetryHookInfo{
		Method:  method,
		URL:     url,
		Attempt: attempt,
		Delay:   delay,
		Reason:  reason,
	})
}

func serializeBody(body any) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	if value, ok := body.(map[string]any); ok && value == nil {
		return nil, nil
	}
	if value, ok := body.([]byte); ok {
		return value, nil
	}
	return json.Marshal(body)
}

// AsJSONMap normaliza o corpo decodificado em um objeto JSON. Respostas
// vazias viram um mapa vazio; respostas que não são objetos JSON (listas,
// texto) são embrulhadas em {"data": ...}.
func AsJSONMap(data any) Json {
	switch value := data.(type) {
	case nil:
		return Json{}
	case map[string]any:
		return value
	default:
		return Json{"data": value}
	}
}

// JoinURL junta a base da API com o caminho, sem barras duplicadas.
func JoinURL(baseURL, path string) string {
	base := strings.TrimRight(baseURL, "/")
	suffix := strings.TrimLeft(path, "/")
	if suffix == "" {
		return base
	}
	return base + "/" + suffix
}

// BuildQueryString monta a query string a partir de um mapa, ignorando
// valores nil. As chaves saem ordenadas, o que torna a URL determinística.
func BuildQueryString(query map[string]any) string {
	if len(query) == 0 {
		return ""
	}

	values := url.Values{}
	for key, value := range query {
		if value == nil {
			continue
		}
		values.Set(key, fmt.Sprintf("%v", value))
	}

	encoded := values.Encode()
	if encoded == "" {
		return ""
	}
	return "?" + encoded
}
