package apibrasil_test

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
)

// retryRapido evita esperas reais nos testes.
func retryRapido(retries int) *apibrasil.RetryConfig {
	return &apibrasil.RetryConfig{
		Retries:         retries,
		MinDelay:        time.Millisecond,
		MaxDelay:        5 * time.Millisecond,
		RetryOnStatuses: []int{429},
	}
}

func TestRetryEm429(t *testing.T) {
	api, transport := BuildAPI(apibrasil.Config{Retry: retryRapido(2)})
	transport.RespondWith(
		HTTPError(http.StatusTooManyRequests, core.Json{"message": "limite"}, nil),
		OK(core.Json{"ok": true}),
	)

	if _, err := api.WhatsApp.SendText(context.Background(), apibrasil.Json{"number": "1"}); err != nil {
		t.Fatalf("sendText falhou: %v", err)
	}
	if got := transport.Count(); got != 2 {
		t.Errorf("requisições = %d, esperado 2", got)
	}
}

func TestRetryRespeitaOLimiteDeTentativas(t *testing.T) {
	api, transport := BuildAPI(apibrasil.Config{Retry: retryRapido(2)})
	transport.SetFallback(HTTPError(http.StatusTooManyRequests, core.Json{"message": "limite"}, nil))

	_, err := api.WhatsApp.SendText(context.Background(), apibrasil.Json{"number": "1"})

	if !errors.Is(err, apibrasil.ErrRateLimit) {
		t.Fatalf("erro = %v, esperado ErrRateLimit", err)
	}
	if got := transport.Count(); got != 3 {
		t.Errorf("requisições = %d, esperado 3 (1 + 2 retries)", got)
	}
}

func TestRetryEmFalhaDeRede(t *testing.T) {
	api, transport := BuildAPI(apibrasil.Config{Retry: retryRapido(1)})
	transport.RespondWith(
		core.NewNetworkError("conexão recusada", nil),
		OK(core.Json{"ok": true}),
	)

	if _, err := api.Account.Balance(context.Background()); err != nil {
		t.Fatalf("balance falhou: %v", err)
	}
	if got := transport.Count(); got != 2 {
		t.Errorf("requisições = %d, esperado 2", got)
	}
}

func TestSemRetryEmTimeout(t *testing.T) {
	api, transport := BuildAPI(apibrasil.Config{Retry: retryRapido(2)})
	transport.RespondWith(core.NewTimeoutError("estourou", nil), core.NewTimeoutError("estourou", nil))

	_, err := api.WhatsApp.SendText(context.Background(), apibrasil.Json{"number": "1"})

	if !errors.Is(err, apibrasil.ErrTimeout) {
		t.Fatalf("erro = %v, esperado ErrTimeout", err)
	}
	if got := transport.Count(); got != 1 {
		t.Errorf("requisições = %d — timeout não pode ser refeito", got)
	}
}

func TestSemRetryEmErroDeNegocio(t *testing.T) {
	api, transport := BuildAPI(apibrasil.Config{Retry: retryRapido(2)})
	transport.SetFallback(HTTPError(http.StatusPaymentRequired, core.Json{"message": "sem saldo"}, nil))

	_, err := api.Consulta.CPF(context.Background(), apibrasil.Json{"cpf": "0"})

	if !errors.Is(err, apibrasil.ErrInsufficientBalance) {
		t.Fatalf("erro = %v", err)
	}
	if got := transport.Count(); got != 1 {
		t.Errorf("requisições = %d — erro de negócio não pode ser refeito", got)
	}
}

func TestRetryDesativado(t *testing.T) {
	api, transport := BuildAPI() // BuildAPI já usa NoRetry
	transport.SetFallback(HTTPError(http.StatusTooManyRequests, core.Json{"message": "limite"}, nil))

	if _, err := api.Account.Balance(context.Background()); !errors.Is(err, apibrasil.ErrRateLimit) {
		t.Fatalf("erro = %v", err)
	}
	if got := transport.Count(); got != 1 {
		t.Errorf("requisições = %d, esperado 1", got)
	}
}

func TestRetryOnStatusesPersonalizado(t *testing.T) {
	retry := retryRapido(1)
	retry.RetryOnStatuses = []int{503}

	api, transport := BuildAPI(apibrasil.Config{Retry: retry})
	transport.RespondWith(
		HTTPError(http.StatusServiceUnavailable, core.Json{"message": "manutenção"}, nil),
		OK(core.Json{"ok": true}),
	)

	if _, err := api.Account.Balance(context.Background()); err != nil {
		t.Fatalf("balance falhou: %v", err)
	}
	if got := transport.Count(); got != 2 {
		t.Errorf("requisições = %d, esperado 2", got)
	}
}

func TestHooksDeObservabilidade(t *testing.T) {
	var mu sync.Mutex
	var requests, responses, retries int

	hooks := &apibrasil.Hooks{
		OnRequest: func(_ context.Context, info apibrasil.RequestHookInfo) {
			mu.Lock()
			defer mu.Unlock()
			requests++
			if info.URL == "" || info.Method == "" {
				t.Error("hook de request sem método/url")
			}
		},
		OnResponse: func(_ context.Context, info apibrasil.ResponseHookInfo) {
			mu.Lock()
			defer mu.Unlock()
			responses++
			if info.Status == 0 {
				t.Error("hook de response sem status")
			}
		},
		OnRetry: func(_ context.Context, info apibrasil.RetryHookInfo) {
			mu.Lock()
			defer mu.Unlock()
			retries++
			if info.Reason != "HTTP 429" {
				t.Errorf("motivo do retry = %q", info.Reason)
			}
		},
	}

	api, transport := BuildAPI(apibrasil.Config{Retry: retryRapido(1), Hooks: hooks})
	transport.RespondWith(
		HTTPError(http.StatusTooManyRequests, core.Json{"message": "limite"}, nil),
		OK(core.Json{"ok": true}),
	)

	if _, err := api.WhatsApp.SendText(context.Background(), apibrasil.Json{"number": "1"}); err != nil {
		t.Fatalf("sendText falhou: %v", err)
	}

	mu.Lock()
	defer mu.Unlock()
	if requests != 2 || responses != 2 || retries != 1 {
		t.Errorf("hooks: requests=%d responses=%d retries=%d", requests, responses, retries)
	}
}

func TestBackoffExponencialRespeitaOTeto(t *testing.T) {
	retry := core.RetryConfig{
		Retries:  5,
		MinDelay: 100 * time.Millisecond,
		MaxDelay: 300 * time.Millisecond,
	}

	for attempt := 0; attempt < 5; attempt++ {
		delay := core.BackoffDelay(attempt, retry)
		if delay > retry.MaxDelay {
			t.Errorf("tentativa %d: delay %v acima do teto %v", attempt, delay, retry.MaxDelay)
		}
		if delay <= 0 {
			t.Errorf("tentativa %d: delay %v deveria ser positivo", attempt, delay)
		}
	}
}

func TestResolveRetryUsaOsPadroes(t *testing.T) {
	padrao := core.ResolveRetry(nil)
	if padrao.Retries != 2 || padrao.MinDelay != 300*time.Millisecond || padrao.MaxDelay != 5*time.Second {
		t.Errorf("política padrão = %+v", padrao)
	}
	if !padrao.RetriesStatus(429) || padrao.RetriesStatus(500) {
		t.Error("por padrão só o 429 deve disparar retry")
	}

	desativado := core.ResolveRetry(core.NoRetry())
	if desativado.Retries != 0 {
		t.Errorf("NoRetry().Retries = %d", desativado.Retries)
	}
}

func TestContextoCanceladoInterrompeORetry(t *testing.T) {
	api, transport := BuildAPI(apibrasil.Config{Retry: retryRapido(3)})
	transport.SetFallback(HTTPError(http.StatusTooManyRequests, core.Json{"message": "limite"}, nil))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := api.Account.Balance(ctx); err == nil {
		t.Fatal("esperava erro com o contexto cancelado")
	}
	if got := transport.Count(); got != 1 {
		t.Errorf("requisições = %d, esperado 1", got)
	}
}
