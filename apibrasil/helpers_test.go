package apibrasil_test

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"testing"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
)

// BaseURL usada em todos os testes de rota.
const BaseURL = "https://gateway.apibrasil.io/api/v2"

// FakeTransport é o transporte fake dos testes: grava todas as
// requisições e responde com uma fila programável (ou um fallback 200).
type FakeTransport struct {
	mu       sync.Mutex
	Calls    []core.TransportRequest
	queue    []any // *core.TransportResponse ou error
	fallback *core.TransportResponse
}

// NewFakeTransport cria o transporte fake com o fallback 200 {"ok": true}.
func NewFakeTransport() *FakeTransport {
	return &FakeTransport{fallback: OK(core.Json{"ok": true})}
}

// RespondWith enfileira respostas (ou erros a lançar), consumidas em ordem.
func (t *FakeTransport) RespondWith(responses ...any) *FakeTransport {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.queue = append(t.queue, responses...)
	return t
}

// SetFallback define a resposta padrão quando a fila está vazia.
func (t *FakeTransport) SetFallback(response *core.TransportResponse) *FakeTransport {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.fallback = response
	return t
}

// Do implementa core.Transport.
func (t *FakeTransport) Do(_ context.Context, req core.TransportRequest) (*core.TransportResponse, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.Calls = append(t.Calls, req)

	next := any(t.fallback)
	if len(t.queue) > 0 {
		next, t.queue = t.queue[0], t.queue[1:]
	}

	if err, ok := next.(error); ok {
		return nil, err
	}
	return next.(*core.TransportResponse), nil
}

// Last devolve a última requisição recebida.
func (t *FakeTransport) Last(tb testing.TB) core.TransportRequest {
	tb.Helper()
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(t.Calls) == 0 {
		tb.Fatal("nenhuma requisição foi feita")
	}
	return t.Calls[len(t.Calls)-1]
}

// LastBody devolve o body JSON decodificado da última requisição.
func (t *FakeTransport) LastBody(tb testing.TB) core.Json {
	tb.Helper()

	last := t.Last(tb)
	if len(last.Body) == 0 {
		return nil
	}

	var body core.Json
	if err := json.Unmarshal(last.Body, &body); err != nil {
		tb.Fatalf("body inválido: %v", err)
	}
	return body
}

// Count devolve quantas requisições foram feitas.
func (t *FakeTransport) Count() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	return len(t.Calls)
}

// OK monta uma resposta 200 com o corpo informado.
func OK(data any) *core.TransportResponse {
	return &core.TransportResponse{Status: http.StatusOK, Headers: http.Header{}, Data: data}
}

// HTTPError monta uma resposta de erro com status, corpo e headers.
func HTTPError(status int, data any, headers http.Header) *core.TransportResponse {
	if headers == nil {
		headers = http.Header{}
	}
	return &core.TransportResponse{Status: status, Headers: headers, Data: data}
}

// BuildAPI cria um cliente com transporte fake, retry desligado e
// credenciais fixas — base de todos os testes de rota.
func BuildAPI(overrides ...apibrasil.Config) (*apibrasil.Client, *FakeTransport) {
	transport := NewFakeTransport()

	config := apibrasil.Config{
		Transport:   transport,
		Retry:       apibrasil.NoRetry(),
		BaseURL:     BaseURL,
		BearerToken: "jwt",
		DeviceToken: "dev",
	}
	if len(overrides) > 0 {
		config = config.Merge(overrides[0])
	}

	return apibrasil.New(config), transport
}

// RouteCase é um caso de teste de rota: uma chamada da SDK e a
// requisição esperada.
type RouteCase struct {
	Name   string
	Call   func(ctx context.Context, api *apibrasil.Client) (any, error)
	Method core.Method
	Path   string
	Body   core.Json
}

// RunRouteCases executa uma tabela de casos de rota.
func RunRouteCases(t *testing.T, cases []RouteCase) {
	t.Helper()

	for _, testCase := range cases {
		t.Run(testCase.Name, func(t *testing.T) {
			api, transport := BuildAPI()

			if _, err := testCase.Call(context.Background(), api); err != nil {
				t.Fatalf("chamada falhou: %v", err)
			}

			last := transport.Last(t)

			expectedMethod := testCase.Method
			if expectedMethod == "" {
				expectedMethod = core.MethodPost
			}
			if last.Method != expectedMethod {
				t.Errorf("método = %s, esperado %s", last.Method, expectedMethod)
			}

			expectedURL := BaseURL + testCase.Path
			if last.URL != expectedURL {
				t.Errorf("url = %s, esperado %s", last.URL, expectedURL)
			}

			if testCase.Body != nil {
				got := transport.LastBody(t)
				if !jsonEqual(got, testCase.Body) {
					t.Errorf("body = %v, esperado %v", got, testCase.Body)
				}
			}
		})
	}
}

func jsonEqual(got, want any) bool {
	gotRaw, err := json.Marshal(got)
	if err != nil {
		return false
	}
	wantRaw, err := json.Marshal(want)
	if err != nil {
		return false
	}

	var gotValue, wantValue any
	if json.Unmarshal(gotRaw, &gotValue) != nil || json.Unmarshal(wantRaw, &wantValue) != nil {
		return false
	}
	return jsonDeepEqual(gotValue, wantValue)
}

func jsonDeepEqual(got, want any) bool {
	switch wantValue := want.(type) {
	case map[string]any:
		gotValue, ok := got.(map[string]any)
		if !ok || len(gotValue) != len(wantValue) {
			return false
		}
		for key, value := range wantValue {
			if !jsonDeepEqual(gotValue[key], value) {
				return false
			}
		}
		return true
	case []any:
		gotValue, ok := got.([]any)
		if !ok || len(gotValue) != len(wantValue) {
			return false
		}
		for index, value := range wantValue {
			if !jsonDeepEqual(gotValue[index], value) {
				return false
			}
		}
		return true
	default:
		return got == want
	}
}
