package apibrasil_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
)

func TestStatusHTTPViraSentinela(t *testing.T) {
	cases := []struct {
		status   int
		sentinel error
	}{
		{http.StatusBadRequest, apibrasil.ErrValidation},
		{http.StatusUnprocessableEntity, apibrasil.ErrValidation},
		{http.StatusUnauthorized, apibrasil.ErrAuthentication},
		{http.StatusPaymentRequired, apibrasil.ErrInsufficientBalance},
		{http.StatusForbidden, apibrasil.ErrPermission},
		{http.StatusNotFound, apibrasil.ErrNotFound},
		{http.StatusGone, apibrasil.ErrNotFound},
		{http.StatusTooManyRequests, apibrasil.ErrRateLimit},
		{http.StatusInternalServerError, apibrasil.ErrServer},
		{http.StatusBadGateway, apibrasil.ErrServer},
	}

	for _, testCase := range cases {
		api, transport := BuildAPI()
		transport.RespondWith(HTTPError(testCase.status, core.Json{"message": "falhou"}, nil))

		_, err := api.Consulta.CPF(context.Background(), apibrasil.Json{"cpf": "0"})
		if err == nil {
			t.Fatalf("HTTP %d: esperava erro", testCase.status)
		}
		if !errors.Is(err, testCase.sentinel) {
			t.Errorf("HTTP %d: erro = %v, esperado %v", testCase.status, err, testCase.sentinel)
		}
		if !errors.Is(err, apibrasil.ErrAPIBrasil) {
			t.Errorf("HTTP %d: todo erro da SDK deve casar com ErrAPIBrasil", testCase.status)
		}
	}
}

func TestErroExpoeStatusCodigoEResposta(t *testing.T) {
	api, transport := BuildAPI()
	transport.RespondWith(HTTPError(http.StatusPaymentRequired, core.Json{
		"message": "Saldo insuficiente",
		"code":    "NO_BALANCE",
	}, nil))

	_, err := api.Consulta.CNPJ(context.Background(), apibrasil.Json{"cnpj": "0"})

	apiErr, ok := apibrasil.AsError(err)
	if !ok {
		t.Fatalf("erro = %v, esperado *apibrasil.Error", err)
	}
	if apiErr.Status != http.StatusPaymentRequired {
		t.Errorf("Status = %d", apiErr.Status)
	}
	if apiErr.Code != "NO_BALANCE" {
		t.Errorf("Code = %q", apiErr.Code)
	}
	if apiErr.Message != "Saldo insuficiente" {
		t.Errorf("Message = %q", apiErr.Message)
	}
	if !apiErr.IsInsufficientBalance() {
		t.Error("IsInsufficientBalance() = false")
	}
	if response, ok := apiErr.Response.(core.Json); !ok || response["code"] != "NO_BALANCE" {
		t.Errorf("Response = %v", apiErr.Response)
	}

	var typed *apibrasil.Error
	if !errors.As(err, &typed) {
		t.Error("errors.As deveria encontrar *apibrasil.Error")
	}
}

func TestMensagemPadraoQuandoAAPINaoInforma(t *testing.T) {
	api, transport := BuildAPI()
	transport.RespondWith(HTTPError(http.StatusInternalServerError, "boom", nil))

	_, err := api.Account.Balance(context.Background())

	apiErr, _ := apibrasil.AsError(err)
	if apiErr == nil || apiErr.Message != "A API respondeu com HTTP 500." {
		t.Errorf("mensagem = %v", err)
	}
}

func TestErroUsaCampoErrorQuandoNaoHaMessage(t *testing.T) {
	api, transport := BuildAPI()
	transport.RespondWith(HTTPError(http.StatusBadRequest, core.Json{"error": "cpf inválido"}, nil))

	_, err := api.Consulta.CPF(context.Background(), apibrasil.Json{"cpf": "0"})

	apiErr, _ := apibrasil.AsError(err)
	if apiErr == nil || apiErr.Message != "cpf inválido" {
		t.Errorf("mensagem = %v", err)
	}
}

func TestRateLimitLeRetryAfter(t *testing.T) {
	headers := http.Header{}
	headers.Set("Retry-After", "3")

	api, transport := BuildAPI()
	transport.RespondWith(HTTPError(http.StatusTooManyRequests, core.Json{"message": "limite"}, headers))

	_, err := api.WhatsApp.SendText(context.Background(), apibrasil.Json{"number": "1"})

	apiErr, _ := apibrasil.AsError(err)
	if apiErr == nil || apiErr.RetryAfter != 3*time.Second {
		t.Errorf("RetryAfter = %v", err)
	}
}

func TestParseRetryAfterComData(t *testing.T) {
	now := time.Date(2026, 7, 23, 12, 0, 0, 0, time.UTC)

	headers := http.Header{}
	headers.Set("Retry-After", now.Add(90*time.Second).Format(http.TimeFormat))

	if got := core.ParseRetryAfter(headers, now); got != 90*time.Second {
		t.Errorf("ParseRetryAfter = %v, esperado 90s", got)
	}

	headers.Set("Retry-After", now.Add(-time.Minute).Format(http.TimeFormat))
	if got := core.ParseRetryAfter(headers, now); got != 0 {
		t.Errorf("data no passado deveria virar 0, veio %v", got)
	}

	if got := core.ParseRetryAfter(http.Header{}, now); got != 0 {
		t.Errorf("sem header deveria virar 0, veio %v", got)
	}
}

func TestFalhaDeRedeViraErrNetwork(t *testing.T) {
	api, transport := BuildAPI()
	transport.RespondWith(core.NewNetworkError("conexão recusada", errors.New("dial")))

	_, err := api.Account.Balance(context.Background())

	if !errors.Is(err, apibrasil.ErrNetwork) {
		t.Errorf("erro = %v, esperado ErrNetwork", err)
	}
	if errors.Is(err, apibrasil.ErrTimeout) {
		t.Error("falha de rede não deveria casar com ErrTimeout")
	}
}

func TestTimeoutTambemCasaComErrNetwork(t *testing.T) {
	api, transport := BuildAPI()
	transport.RespondWith(core.NewTimeoutError("estourou", nil))

	_, err := api.Account.Balance(context.Background())

	if !errors.Is(err, apibrasil.ErrTimeout) {
		t.Errorf("erro = %v, esperado ErrTimeout", err)
	}
	if !errors.Is(err, apibrasil.ErrNetwork) {
		t.Error("timeout deveria casar também com ErrNetwork")
	}
}

func TestErrorFormatado(t *testing.T) {
	err := &core.Error{
		Kind:    core.KindNotFound,
		Message: "Nada encontrado",
		Status:  404,
		Code:    "NOT_FOUND",
	}

	if got := err.Error(); got != "Nada encontrado (HTTP 404) [NOT_FOUND]" {
		t.Errorf("Error() = %q", got)
	}
}

func TestErroPreservaACausa(t *testing.T) {
	causa := errors.New("dial tcp: connection refused")
	err := core.NewNetworkError("falhou", causa)

	if !errors.Is(err, causa) {
		t.Error("a causa original deveria ser recuperável com errors.Is")
	}
}
