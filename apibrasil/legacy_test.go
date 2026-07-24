package apibrasil_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jhowbhz/apigratis-sdk-go/apigratis"
)

func TestLegacyServiceMantemOContrato(t *testing.T) {
	var (
		gotPath   string
		gotAuth   string
		gotDevice string
		gotAgent  string
		gotBody   string
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		gotPath = r.URL.Path
		gotAuth = r.Header.Get("Authorization")
		gotDevice = r.Header.Get("DeviceToken")
		gotAgent = r.Header.Get("User-Agent")
		gotBody = string(body)

		_, _ = w.Write([]byte(`{"error":false,"response":"ok"}`))
	}))
	defer server.Close()

	service := apigratis.NewService()
	service.Server = server.URL + "/api/v2/"

	payload, _ := json.Marshal(map[string]any{
		"credentials": map[string]any{"BearerToken": "jwt", "DeviceToken": "dev"},
		"body":        map[string]any{"number": "5511999999999", "text": "Olá"},
		"action":      "sendText",
	})

	response, err := service.Whatsapp(string(payload))
	if err != nil {
		t.Fatalf("whatsapp falhou: %v", err)
	}

	if gotPath != "/api/v2/whatsapp/sendText" {
		t.Errorf("path = %s", gotPath)
	}
	if gotAuth != "Bearer jwt" {
		t.Errorf("Authorization = %q", gotAuth)
	}
	if gotDevice != "dev" {
		t.Errorf("DeviceToken = %q", gotDevice)
	}
	if gotAgent != "APIBRASIL/GOLANG-SDK" {
		t.Errorf("User-Agent = %q", gotAgent)
	}
	if gotBody != `{"number":"5511999999999","text":"Olá"}` {
		t.Errorf("body = %s", gotBody)
	}
	if response["response"] != "ok" {
		t.Errorf("resposta = %v", response)
	}
}

func TestLegacyServiceRotasPorServico(t *testing.T) {
	var gotPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	service := &apigratis.Service{Server: server.URL + "/"}
	payload := `{"credentials":{"BearerToken":"jwt","DeviceToken":"dev"},"body":{"cpf":"0"},"action":"dados"}`

	cases := []struct {
		name string
		call func(string) (map[string]interface{}, error)
		path string
	}{
		{"Whatsapp", service.Whatsapp, "/whatsapp/dados"},
		{"Sms", service.Sms, "/sms/dados"},
		{"Cpf", service.Cpf, "/cpf/dados/dados"},
		{"Cnpj", service.Cnpj, "/dados/dados"},
	}

	for _, testCase := range cases {
		if _, err := testCase.call(payload); err != nil {
			t.Fatalf("%s falhou: %v", testCase.name, err)
		}
		if gotPath != testCase.path {
			t.Errorf("%s: path = %s, esperado %s", testCase.name, gotPath, testCase.path)
		}
	}
}

func TestLegacyServiceDevolveErroDaAPIDecodificado(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusPaymentRequired)
		_, _ = w.Write([]byte(`{"error":true,"message":"Saldo insuficiente"}`))
	}))
	defer server.Close()

	service := &apigratis.Service{Server: server.URL + "/"}

	response, err := service.Cnpj(`{"credentials":{"BearerToken":"jwt","DeviceToken":"dev"},"body":{"cnpj":"0"},"action":"cnpj"}`)
	if err != nil {
		t.Fatalf("erros HTTP não devem virar erro na interface legada: %v", err)
	}
	if response["message"] != "Saldo insuficiente" {
		t.Errorf("resposta = %v", response)
	}
}

func TestLegacyServiceValidaOPayload(t *testing.T) {
	service := apigratis.NewService()

	if _, err := service.Whatsapp("{"); err == nil {
		t.Error("esperava erro com JSON inválido")
	}
	if _, err := service.Whatsapp(`{"body":{}}`); err == nil {
		t.Error("esperava erro sem credentials")
	}
	if _, err := service.Whatsapp(`{"credentials":{}}`); err == nil {
		t.Error("esperava erro sem body")
	}
}

func TestLegacyServiceUsaOGatewayPorPadrao(t *testing.T) {
	if got := apigratis.NewService().Server; got != "https://gateway.apibrasil.io/api/v2/" {
		t.Errorf("Server = %q", got)
	}
}
