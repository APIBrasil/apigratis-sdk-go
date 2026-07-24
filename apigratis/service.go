// Package apigratis é a interface legada da SDK (v0.1.x), mantida por
// compatibilidade.
//
// Deprecated: prefira o cliente github.com/jhowbhz/apigratis-sdk-go/apibrasil,
// que cobre toda a plataforma com métodos dedicados, erros tipados,
// retry e hooks.
package apigratis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
)

// Service é o cliente legado da plataforma.
//
// Deprecated: prefira apibrasil.New(apibrasil.Config{...}).
type Service struct {
	// Server é a base da API, com barra ao final.
	Server string
}

// NewService cria o cliente legado apontando para o gateway padrão.
//
// Deprecated: prefira apibrasil.New(apibrasil.Config{...}).
func NewService() *Service {
	return &Service{
		Server: "https://gateway.apibrasil.io/api/v2/",
	}
}

// Request executa uma chamada no formato legado. dados é um JSON com
// credentials, body e action:
//
//	{"credentials": {"BearerToken": "...", "DeviceToken": "..."},
//	 "body": {"number": "5511999999999", "text": "Olá!"},
//	 "action": "sendText"}
//
// Assim como na versão anterior, respostas de erro da API são devolvidas
// decodificadas em vez de virarem erro.
//
// Deprecated: prefira os serviços do pacote apibrasil.
func (s *Service) Request(service string, dados string) (map[string]interface{}, error) {
	return s.RequestContext(context.Background(), service, dados)
}

// RequestContext é o Request com suporte a context.Context.
//
// Deprecated: prefira os serviços do pacote apibrasil.
func (s *Service) RequestContext(ctx context.Context, service string, dados string) (map[string]interface{}, error) {
	var payload struct {
		Credentials map[string]interface{} `json:"credentials"`
		Body        map[string]interface{} `json:"body"`
		Action      string                 `json:"action"`
	}
	if err := json.Unmarshal([]byte(dados), &payload); err != nil {
		return nil, err
	}
	if payload.Credentials == nil {
		return nil, fmt.Errorf("invalid request, missing credentials")
	}
	if payload.Body == nil {
		return nil, fmt.Errorf("invalid request, missing body")
	}

	client := core.NewHTTPClient(core.Config{
		BaseURL:     s.Server,
		BearerToken: stringField(payload.Credentials, "BearerToken"),
		DeviceToken: stringField(payload.Credentials, "DeviceToken"),
		Headers:     map[string]string{"User-Agent": "APIBRASIL/GOLANG-SDK"},
		Retry:       core.NoRetry(),
	})

	path := service
	if payload.Action != "" {
		path += "/" + payload.Action
	}

	response, err := client.Post(ctx, path, payload.Body)
	if err == nil {
		return response, nil
	}

	// Comportamento legado: erros HTTP voltam decodificados, não como erro.
	apiErr, ok := core.AsError(err)
	if !ok {
		return nil, err
	}
	if body, ok := apiErr.Response.(map[string]interface{}); ok {
		return body, nil
	}
	if apiErr.Status == 0 {
		return nil, err
	}
	return map[string]interface{}{"error": true, "message": apiErr.Message}, nil
}

// Whatsapp chama /whatsapp/{action}.
//
// Deprecated: prefira apibrasil.New(...).WhatsApp.
func (s *Service) Whatsapp(dados string) (map[string]interface{}, error) {
	return s.Request("whatsapp", dados)
}

// Sms chama /sms/{action}.
//
// Deprecated: prefira apibrasil.New(...).SMS.
func (s *Service) Sms(dados string) (map[string]interface{}, error) {
	return s.Request("sms", dados)
}

// Cpf chama /cpf/dados/{action}.
//
// Deprecated: prefira apibrasil.New(...).Dados.CPF ou .Consulta.CPF.
func (s *Service) Cpf(dados string) (map[string]interface{}, error) {
	return s.Request("cpf/dados", dados)
}

// Cnpj chama /dados/{action}.
//
// Deprecated: prefira apibrasil.New(...).Dados.CNPJ ou .Consulta.CNPJ.
func (s *Service) Cnpj(dados string) (map[string]interface{}, error) {
	return s.Request("dados", dados)
}

func stringField(source map[string]interface{}, key string) string {
	if value, ok := source[key].(string); ok {
		return value
	}
	return ""
}
