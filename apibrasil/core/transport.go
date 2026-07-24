package core

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// TransportRequest é a requisição entregue à camada de transporte.
type TransportRequest struct {
	Method Method
	// URL absoluta, já com query string.
	URL     string
	Headers map[string]string
	// Body já serializado (JSON). nil quando não há corpo.
	Body         []byte
	Timeout      time.Duration
	ResponseType ResponseType
}

// TransportResponse é a resposta devolvida pela camada de transporte.
type TransportResponse struct {
	Status  int
	Headers http.Header
	// Data é o corpo já decodificado (JSON → map/slice; texto; bytes).
	Data any
	// Body são os bytes crus da resposta.
	Body []byte
}

// Transport é a camada de transporte HTTP da SDK. A implementação padrão
// é HTTPTransport (net/http); injete a sua para usar proxies
// corporativos, instrumentação ou mocks de teste.
//
// Contrato: devolve a resposta para QUALQUER status HTTP; devolve erro
// (KindNetwork/KindTimeout) apenas quando não houve resposta.
type Transport interface {
	Do(ctx context.Context, req TransportRequest) (*TransportResponse, error)
}

// HTTPTransport é o transporte padrão, baseado em net/http.
type HTTPTransport struct {
	// Client é o *http.Client usado nas requisições. Quando nil, usa um
	// cliente próprio com connection pooling.
	Client *http.Client
}

// NewHTTPTransport cria o transporte padrão. client opcional — informe o
// seu para configurar proxy, TLS, pool de conexões etc.
func NewHTTPTransport(client *http.Client) *HTTPTransport {
	return &HTTPTransport{Client: client}
}

var defaultHTTPClient = &http.Client{}

func (t *HTTPTransport) client() *http.Client {
	if t.Client != nil {
		return t.Client
	}
	return defaultHTTPClient
}

// Do executa a requisição HTTP.
func (t *HTTPTransport) Do(ctx context.Context, req TransportRequest) (*TransportResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if req.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, req.Timeout)
		defer cancel()
	}

	var payload io.Reader
	if len(req.Body) > 0 {
		payload = bytes.NewReader(req.Body)
	}

	httpReq, err := http.NewRequestWithContext(ctx, string(req.Method), req.URL, payload)
	if err != nil {
		return nil, NewNetworkError(
			fmt.Sprintf("Requisição inválida em %s %s: %v", req.Method, req.URL, err),
			err,
		)
	}
	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	httpResp, err := t.client().Do(httpReq)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, NewTimeoutError(
				fmt.Sprintf("Tempo limite de %s excedido em %s %s.", req.Timeout, req.Method, req.URL),
				err,
			)
		}
		return nil, NewNetworkError(
			fmt.Sprintf("Falha de rede em %s %s: %v", req.Method, req.URL, err),
			err,
		)
	}
	defer httpResp.Body.Close()

	raw, err := io.ReadAll(httpResp.Body)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, NewTimeoutError(
				fmt.Sprintf("Tempo limite de %s excedido ao ler a resposta de %s %s.", req.Timeout, req.Method, req.URL),
				err,
			)
		}
		return nil, NewNetworkError(
			fmt.Sprintf("Falha ao ler a resposta de %s %s: %v", req.Method, req.URL, err),
			err,
		)
	}

	return &TransportResponse{
		Status:  httpResp.StatusCode,
		Headers: httpResp.Header,
		Data:    DecodeBody(raw, req.ResponseType),
		Body:    raw,
	}, nil
}

// DecodeBody decodifica o corpo cru conforme o ResponseType: JSON vira
// map/slice, corpos não-JSON viram string e ResponseBytes devolve os
// próprios bytes.
func DecodeBody(raw []byte, responseType ResponseType) any {
	if responseType == ResponseBytes {
		return raw
	}
	if len(raw) == 0 {
		return nil
	}

	var decoded any
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return string(raw)
	}
	return decoded
}
