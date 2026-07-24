package apibrasil_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
)

func TestHTTPTransportEnviaMetodoHeadersECorpo(t *testing.T) {
	var (
		gotMethod string
		gotPath   string
		gotAuth   string
		gotBody   string
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		gotMethod, gotPath, gotAuth, gotBody = r.Method, r.URL.Path, r.Header.Get("Authorization"), string(body)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"error":false,"response":{"id":"1"}}`))
	}))
	defer server.Close()

	api := apibrasil.New(apibrasil.Config{
		BaseURL:     server.URL + "/api/v2",
		BearerToken: "jwt",
		DeviceToken: "dev",
		Retry:       apibrasil.NoRetry(),
	})

	response, err := api.WhatsApp.SendText(context.Background(), apibrasil.Json{"number": "5511999999999", "text": "Olá"})
	if err != nil {
		t.Fatalf("sendText falhou: %v", err)
	}

	if gotMethod != http.MethodPost {
		t.Errorf("método = %s", gotMethod)
	}
	if gotPath != "/api/v2/whatsapp/sendText" {
		t.Errorf("path = %s", gotPath)
	}
	if gotAuth != "Bearer jwt" {
		t.Errorf("Authorization = %q", gotAuth)
	}
	if gotBody != `{"number":"5511999999999","text":"Olá"}` {
		t.Errorf("body = %s", gotBody)
	}
	if response.IsError() {
		t.Error("IsError() = true")
	}
}

func TestHTTPTransportDevolveRespostaEmQualquerStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusPaymentRequired)
		_, _ = w.Write([]byte(`{"message":"Saldo insuficiente"}`))
	}))
	defer server.Close()

	transport := core.NewHTTPTransport(nil)

	response, err := transport.Do(context.Background(), core.TransportRequest{
		Method: core.MethodGet,
		URL:    server.URL,
	})
	if err != nil {
		t.Fatalf("o transporte não deve devolver erro para status HTTP: %v", err)
	}
	if response.Status != http.StatusPaymentRequired {
		t.Errorf("status = %d", response.Status)
	}
	if data, ok := response.Data.(map[string]any); !ok || data["message"] != "Saldo insuficiente" {
		t.Errorf("data = %v", response.Data)
	}
}

func TestHTTPTransportCorpoNaoJSONViraTexto(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("não é json"))
	}))
	defer server.Close()

	response, err := core.NewHTTPTransport(nil).Do(context.Background(), core.TransportRequest{
		Method: core.MethodGet,
		URL:    server.URL,
	})
	if err != nil {
		t.Fatalf("falhou: %v", err)
	}
	if response.Data != "não é json" {
		t.Errorf("data = %v", response.Data)
	}
}

func TestHTTPTransportResponseBytes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("%PDF-1.4"))
	}))
	defer server.Close()

	response, err := core.NewHTTPTransport(nil).Do(context.Background(), core.TransportRequest{
		Method:       core.MethodGet,
		URL:          server.URL,
		ResponseType: core.ResponseBytes,
	})
	if err != nil {
		t.Fatalf("falhou: %v", err)
	}
	if raw, ok := response.Data.([]byte); !ok || string(raw) != "%PDF-1.4" {
		t.Errorf("data = %v", response.Data)
	}
}

func TestHTTPTransportTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(150 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	_, err := core.NewHTTPTransport(nil).Do(context.Background(), core.TransportRequest{
		Method:  core.MethodGet,
		URL:     server.URL,
		Timeout: 20 * time.Millisecond,
	})

	if !errors.Is(err, apibrasil.ErrTimeout) {
		t.Fatalf("erro = %v, esperado ErrTimeout", err)
	}
}

func TestHTTPTransportFalhaDeConexao(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	url := server.URL
	server.Close() // derruba o servidor antes da chamada

	_, err := core.NewHTTPTransport(nil).Do(context.Background(), core.TransportRequest{
		Method: core.MethodGet,
		URL:    url,
	})

	if !errors.Is(err, apibrasil.ErrNetwork) {
		t.Fatalf("erro = %v, esperado ErrNetwork", err)
	}
}

func TestTransportCustomizadoNoCliente(t *testing.T) {
	transport := NewFakeTransport()
	api := apibrasil.New(apibrasil.Config{Transport: transport, BaseURL: BaseURL})

	if _, err := api.Account.Balance(context.Background()); err != nil {
		t.Fatalf("balance falhou: %v", err)
	}
	if transport.Count() != 1 {
		t.Errorf("o transporte customizado não foi usado")
	}
}

func TestHTTPTransportAceitaClientCustomizado(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	api := apibrasil.New(apibrasil.Config{
		BaseURL:   server.URL,
		Transport: apibrasil.NewHTTPTransport(&http.Client{Timeout: 5 * time.Second}),
	})

	response, err := api.Catalog.Status(context.Background())
	if err != nil {
		t.Fatalf("status falhou: %v", err)
	}
	if response["ok"] != true {
		t.Errorf("resposta = %v", response)
	}
}
