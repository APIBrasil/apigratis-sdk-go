package apibrasil_test

import (
	"context"
	"errors"
	"testing"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
)

func TestClientInjetaHeadersDeAutenticacao(t *testing.T) {
	api, transport := BuildAPI()

	if _, err := api.WhatsApp.SendText(context.Background(), apibrasil.Json{"number": "5511999999999"}); err != nil {
		t.Fatalf("sendText falhou: %v", err)
	}

	headers := transport.Last(t).Headers
	want := map[string]string{
		"Authorization": "Bearer jwt",
		"DeviceToken":   "dev",
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"User-Agent":    apibrasil.UserAgent,
	}
	for key, value := range want {
		if headers[key] != value {
			t.Errorf("header %s = %q, esperado %q", key, headers[key], value)
		}
	}
}

func TestClientHeadersExtrasEPorRequisicao(t *testing.T) {
	api, transport := BuildAPI(apibrasil.Config{Headers: map[string]string{"X-Custom": "valor"}})

	_, err := api.WhatsApp.SendText(
		context.Background(),
		apibrasil.Json{"number": "1"},
		apibrasil.RequestOptions{
			BearerToken: "outro-jwt",
			DeviceToken: "outro-device",
			Headers:     map[string]string{"X-Request": "1"},
		},
	)
	if err != nil {
		t.Fatalf("sendText falhou: %v", err)
	}

	headers := transport.Last(t).Headers
	if headers["X-Custom"] != "valor" {
		t.Errorf("header extra do cliente não foi enviado: %v", headers)
	}
	if headers["X-Request"] != "1" {
		t.Errorf("header extra da requisição não foi enviado: %v", headers)
	}
	if headers["Authorization"] != "Bearer outro-jwt" {
		t.Errorf("bearer por requisição não sobrescreveu: %q", headers["Authorization"])
	}
	if headers["DeviceToken"] != "outro-device" {
		t.Errorf("device por requisição não sobrescreveu: %q", headers["DeviceToken"])
	}
}

func TestClientSetTokens(t *testing.T) {
	api, transport := BuildAPI()

	api.SetBearerToken("novo-jwt").SetDeviceToken("novo-device")

	if _, err := api.WhatsApp.QRCode(context.Background(), nil); err != nil {
		t.Fatalf("qrcode falhou: %v", err)
	}

	headers := transport.Last(t).Headers
	if headers["Authorization"] != "Bearer novo-jwt" {
		t.Errorf("Authorization = %q", headers["Authorization"])
	}
	if headers["DeviceToken"] != "novo-device" {
		t.Errorf("DeviceToken = %q", headers["DeviceToken"])
	}
}

func TestClientWithDeviceMantemCredenciais(t *testing.T) {
	api, _ := BuildAPI()

	outro := api.WithDevice("device-2")

	if outro.HTTP.DeviceToken() != "device-2" {
		t.Errorf("DeviceToken = %q, esperado device-2", outro.HTTP.DeviceToken())
	}
	if outro.HTTP.BearerToken() != "jwt" {
		t.Errorf("BearerToken = %q, esperado jwt", outro.HTTP.BearerToken())
	}
	if api.HTTP.DeviceToken() != "dev" {
		t.Errorf("cliente original foi alterado: %q", api.HTTP.DeviceToken())
	}
}

func TestClientRequestGenerico(t *testing.T) {
	api, transport := BuildAPI()

	_, err := api.Request(context.Background(), apibrasil.MethodPost, "/consulta/cpf/credits", apibrasil.Json{"cpf": "00000000000"})
	if err != nil {
		t.Fatalf("request falhou: %v", err)
	}

	last := transport.Last(t)
	if last.URL != BaseURL+"/consulta/cpf/credits" {
		t.Errorf("url = %s", last.URL)
	}
	if last.Method != core.MethodPost {
		t.Errorf("método = %s", last.Method)
	}
}

func TestClientLeConfiguracaoDoAmbiente(t *testing.T) {
	t.Setenv(apibrasil.EnvBearerToken, "jwt-do-ambiente")
	t.Setenv(apibrasil.EnvDeviceToken, "device-do-ambiente")
	t.Setenv(apibrasil.EnvBaseURL, "https://homolog.apibrasil.io/api/v2")

	api := apibrasil.New()

	if api.HTTP.BearerToken() != "jwt-do-ambiente" {
		t.Errorf("BearerToken = %q", api.HTTP.BearerToken())
	}
	if api.HTTP.DeviceToken() != "device-do-ambiente" {
		t.Errorf("DeviceToken = %q", api.HTTP.DeviceToken())
	}
	if api.HTTP.BaseURL() != "https://homolog.apibrasil.io/api/v2" {
		t.Errorf("BaseURL = %q", api.HTTP.BaseURL())
	}
}

func TestConfigExplicitaTemPrioridadeSobreAmbiente(t *testing.T) {
	t.Setenv(apibrasil.EnvBearerToken, "jwt-do-ambiente")

	api := apibrasil.New(apibrasil.Config{BearerToken: "jwt-explicito"})

	if api.HTTP.BearerToken() != "jwt-explicito" {
		t.Errorf("BearerToken = %q, esperado jwt-explicito", api.HTTP.BearerToken())
	}
}

func TestAuthLoginGuardaOToken(t *testing.T) {
	api, transport := BuildAPI()
	transport.RespondWith(OK(core.Json{
		"authorization": core.Json{"token": "jwt-do-login"},
	}))

	if _, err := api.Auth.Login(context.Background(), apibrasil.Json{"email": "a@b.c", "password": "x"}); err != nil {
		t.Fatalf("login falhou: %v", err)
	}

	if api.HTTP.BearerToken() != "jwt-do-login" {
		t.Errorf("BearerToken = %q, esperado jwt-do-login", api.HTTP.BearerToken())
	}
}

func TestAuthLogoutLimpaOToken(t *testing.T) {
	api, _ := BuildAPI()

	if _, err := api.Auth.Logout(context.Background()); err != nil {
		t.Fatalf("logout falhou: %v", err)
	}

	if api.HTTP.BearerToken() != "" {
		t.Errorf("BearerToken = %q, esperado vazio", api.HTTP.BearerToken())
	}
}

func TestLoginComDoisFatoresDevolveErro(t *testing.T) {
	transport := NewFakeTransport()
	transport.RespondWith(OK(core.Json{"requires_2fa": true, "challenge": "abc"}))

	_, session, err := apibrasil.Login(
		context.Background(),
		apibrasil.Json{"email": "a@b.c", "password": "x"},
		apibrasil.Config{Transport: transport, Retry: apibrasil.NoRetry(), BaseURL: BaseURL},
	)

	if err == nil {
		t.Fatal("esperava erro exigindo 2FA")
	}
	if !errors.Is(err, apibrasil.ErrAuthentication) {
		t.Errorf("erro = %v, esperado ErrAuthentication", err)
	}
	if session["challenge"] != "abc" {
		t.Errorf("sessão = %v", session)
	}
}

func TestEnvelopeDeviceServiceResponse(t *testing.T) {
	api, transport := BuildAPI()
	transport.RespondWith(OK(core.Json{
		"error":     false,
		"message":   "ok",
		"response":  core.Json{"qrcode": "data:image/png;base64,xxx"},
		"api_limit": 100,
	}))

	response, err := api.WhatsApp.QRCode(context.Background(), nil)
	if err != nil {
		t.Fatalf("qrcode falhou: %v", err)
	}

	if response.IsError() {
		t.Error("IsError() = true, esperado false")
	}
	if response.Message() != "ok" {
		t.Errorf("Message() = %q", response.Message())
	}
	if response.APILimit() != 100 {
		t.Errorf("APILimit() = %v", response.APILimit())
	}

	var payload struct {
		QRCode string `json:"qrcode"`
	}
	if err := response.Decode(&payload); err != nil {
		t.Fatalf("decode falhou: %v", err)
	}
	if payload.QRCode == "" {
		t.Error("decode não preencheu o qrcode")
	}

	// Continua sendo um mapa comum.
	if response["message"] != "ok" {
		t.Error("acesso por chave deixou de funcionar")
	}
}

func TestEnvelopeCreditServiceResponse(t *testing.T) {
	api, transport := BuildAPI()
	transport.RespondWith(OK(core.Json{
		"error":          false,
		"balance":        12.5,
		"valor_consulta": 0.5,
		"homolog":        true,
		"data":           core.Json{"nome": "Fulano"},
	}))

	response, err := api.Consulta.CPF(context.Background(), apibrasil.Json{"cpf": "00000000000"})
	if err != nil {
		t.Fatalf("consulta falhou: %v", err)
	}

	if response.Balance() != 12.5 {
		t.Errorf("Balance() = %v", response.Balance())
	}
	if !response.Homolog() {
		t.Error("Homolog() = false, esperado true")
	}

	var payload struct {
		Nome string `json:"nome"`
	}
	if err := response.Decode(&payload); err != nil {
		t.Fatalf("decode falhou: %v", err)
	}
	if payload.Nome != "Fulano" {
		t.Errorf("Nome = %q", payload.Nome)
	}
}

func TestConsultaMontaOBody(t *testing.T) {
	consulta := apibrasil.Consulta{
		Tipo:    "lista-socios",
		Homolog: true,
		Fields:  apibrasil.Json{"cnpj": "00000000000000"},
	}

	body := consulta.Json()

	if body["tipo"] != "lista-socios" || body["homolog"] != true || body["cnpj"] != "00000000000000" {
		t.Errorf("body = %v", body)
	}
	if _, ok := body["lite"]; ok {
		t.Error("campos não informados não deveriam aparecer no body")
	}
}
