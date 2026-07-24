package apibrasil_test

import (
	"context"
	"testing"
	"time"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
)

func TestJoinURL(t *testing.T) {
	cases := []struct{ base, path, want string }{
		{"https://x.dev/api/v2", "whatsapp/sendText", "https://x.dev/api/v2/whatsapp/sendText"},
		{"https://x.dev/api/v2/", "/whatsapp/sendText", "https://x.dev/api/v2/whatsapp/sendText"},
		{"https://x.dev/api/v2//", "//balance", "https://x.dev/api/v2/balance"},
		{"https://x.dev/api/v2", "", "https://x.dev/api/v2"},
	}

	for _, testCase := range cases {
		if got := core.JoinURL(testCase.base, testCase.path); got != testCase.want {
			t.Errorf("JoinURL(%q, %q) = %q, esperado %q", testCase.base, testCase.path, got, testCase.want)
		}
	}
}

func TestBuildQueryString(t *testing.T) {
	if got := core.BuildQueryString(nil); got != "" {
		t.Errorf("query vazia = %q", got)
	}
	if got := core.BuildQueryString(core.Json{"search": nil}); got != "" {
		t.Errorf("valores nil deveriam ser ignorados, veio %q", got)
	}
	if got := core.BuildQueryString(core.Json{"a": 1, "b": "x y"}); got != "?a=1&b=x+y" {
		t.Errorf("query = %q", got)
	}
}

func TestBaseURLPadraoEPersonalizada(t *testing.T) {
	if got := apibrasil.New().HTTP.BaseURL(); got != apibrasil.DefaultBaseURL {
		t.Errorf("BaseURL padrão = %q", got)
	}

	custom := apibrasil.New(apibrasil.Config{BaseURL: "https://homolog.apibrasil.io/api/v2"})
	if got := custom.HTTP.BaseURL(); got != "https://homolog.apibrasil.io/api/v2" {
		t.Errorf("BaseURL = %q", got)
	}
}

func TestTimeoutDaConfiguracaoEDaRequisicao(t *testing.T) {
	api, transport := BuildAPI(apibrasil.Config{Timeout: 5 * time.Second})

	if _, err := api.Account.Balance(context.Background()); err != nil {
		t.Fatalf("balance falhou: %v", err)
	}
	if got := transport.Last(t).Timeout; got != 5*time.Second {
		t.Errorf("timeout = %v, esperado 5s", got)
	}

	if _, err := api.Account.Balance(context.Background(), apibrasil.RequestOptions{Timeout: time.Second}); err != nil {
		t.Fatalf("balance falhou: %v", err)
	}
	if got := transport.Last(t).Timeout; got != time.Second {
		t.Errorf("timeout por requisição = %v, esperado 1s", got)
	}
}

func TestTimeoutPadrao(t *testing.T) {
	api, transport := BuildAPI()

	if _, err := api.Account.Balance(context.Background()); err != nil {
		t.Fatalf("balance falhou: %v", err)
	}
	if got := transport.Last(t).Timeout; got != apibrasil.DefaultTimeout {
		t.Errorf("timeout = %v, esperado %v", got, apibrasil.DefaultTimeout)
	}
}

func TestRespostaNaoObjetoViraCampoData(t *testing.T) {
	api, transport := BuildAPI()
	transport.RespondWith(OK([]any{"a", "b"}))

	response, err := api.Account.Jobs(context.Background())
	if err != nil {
		t.Fatalf("jobs falhou: %v", err)
	}

	list, ok := response["data"].([]any)
	if !ok || len(list) != 2 {
		t.Errorf("resposta = %v", response)
	}
}

func TestRespostaVaziaViraMapaVazio(t *testing.T) {
	api, transport := BuildAPI()
	transport.RespondWith(OK(nil))

	response, err := api.Account.Balance(context.Background())
	if err != nil {
		t.Fatalf("balance falhou: %v", err)
	}
	if response == nil || len(response) != 0 {
		t.Errorf("resposta = %v, esperado mapa vazio", response)
	}
}

func TestDoDevolveOCorpoSemNormalizar(t *testing.T) {
	api, transport := BuildAPI()
	transport.RespondWith(OK([]any{1.0, 2.0}))

	data, err := api.Do(context.Background(), apibrasil.MethodGet, "/jobs", nil)
	if err != nil {
		t.Fatalf("do falhou: %v", err)
	}
	if list, ok := data.([]any); !ok || len(list) != 2 {
		t.Errorf("data = %v", data)
	}
}

func TestCorpoNaoEEnviadoQuandoNil(t *testing.T) {
	api, transport := BuildAPI()

	if _, err := api.WhatsApp.QRCode(context.Background(), nil); err != nil {
		t.Fatalf("qrcode falhou: %v", err)
	}
	if body := transport.Last(t).Body; body != nil {
		t.Errorf("body = %q, esperado nil", body)
	}
}

func TestVerbosDoClienteHTTP(t *testing.T) {
	client := core.NewHTTPClient(core.Config{BaseURL: BaseURL, Transport: NewFakeTransport(), Retry: core.NoRetry()})
	transport := client.Transport().(*FakeTransport)
	ctx := context.Background()

	cases := []struct {
		call   func() (core.Json, error)
		method core.Method
	}{
		{func() (core.Json, error) { return client.Get(ctx, "x") }, core.MethodGet},
		{func() (core.Json, error) { return client.Post(ctx, "x", nil) }, core.MethodPost},
		{func() (core.Json, error) { return client.Put(ctx, "x", nil) }, core.MethodPut},
		{func() (core.Json, error) { return client.Patch(ctx, "x", nil) }, core.MethodPatch},
		{func() (core.Json, error) { return client.Delete(ctx, "x", nil) }, core.MethodDelete},
	}

	for _, testCase := range cases {
		if _, err := testCase.call(); err != nil {
			t.Fatalf("%s falhou: %v", testCase.method, err)
		}
		if got := transport.Last(t).Method; got != testCase.method {
			t.Errorf("método = %s, esperado %s", got, testCase.method)
		}
	}
}

func TestIntoDecodificaEmStruct(t *testing.T) {
	client := core.NewHTTPClient(core.Config{BaseURL: BaseURL, Transport: NewFakeTransport(), Retry: core.NoRetry()})
	transport := client.Transport().(*FakeTransport)
	transport.RespondWith(OK(core.Json{"saldo": 12.5, "plano": "pro"}))

	var conta struct {
		Saldo float64 `json:"saldo"`
		Plano string  `json:"plano"`
	}
	if err := client.Into(context.Background(), core.MethodGet, "balance", nil, &conta); err != nil {
		t.Fatalf("into falhou: %v", err)
	}

	if conta.Saldo != 12.5 || conta.Plano != "pro" {
		t.Errorf("conta = %+v", conta)
	}
}

func TestConfigDoClienteRefleteOsTokensAtuais(t *testing.T) {
	api, _ := BuildAPI()
	api.SetBearerToken("jwt-2").SetDeviceToken("dev-2")

	config := api.Config()

	if config.BearerToken != "jwt-2" || config.DeviceToken != "dev-2" {
		t.Errorf("config = %+v", config)
	}
	if config.BaseURL != BaseURL {
		t.Errorf("BaseURL = %q", config.BaseURL)
	}
}
