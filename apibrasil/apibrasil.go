// Package apibrasil é o cliente oficial da plataforma APIBrasil em Go —
// WhatsApp, SMS, consultas de CPF/CNPJ, veículos, CEP, correios,
// pagamentos PIX/boleto e muito mais.
//
//	api := apibrasil.New(apibrasil.Config{
//		BearerToken: os.Getenv("APIBRASIL_BEARER_TOKEN"),
//		DeviceToken: os.Getenv("APIBRASIL_DEVICE_TOKEN"),
//	})
//
//	_, err := api.WhatsApp.SendText(ctx, apibrasil.Json{
//		"number": "5511999999999",
//		"text":   "Olá!",
//	})
//
// Credenciais não informadas são lidas das variáveis de ambiente
// APIBRASIL_BEARER_TOKEN, APIBRASIL_DEVICE_TOKEN, APIBRASIL_SECRET_KEY e
// APIBRASIL_BASE_URL.
package apibrasil

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services/data"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services/messaging"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services/platform"
)

// Client é o cliente da plataforma APIBrasil. Todos os serviços
// compartilham o mesmo *core.HTTPClient — um SetBearerToken vale para
// todos.
type Client struct {
	// HTTP é o cliente HTTP interno (headers, base URL, retry, hooks, erros).
	HTTP *core.HTTPClient

	// Auth cuida de login, 2FA, cadastro, senha e perfil.
	Auth *platform.AuthService
	// Devices gerencia os devices (criar, listar, atualizar, remover).
	Devices *platform.DevicesService

	// WhatsApp device-based (/whatsapp/{action}).
	WhatsApp *messaging.WhatsAppService
	// Evolution API (/evolution/{controller}/{action}).
	Evolution *messaging.EvolutionService
	// WhatsMeow (/whatsmeow/{action}).
	WhatsMeow *messaging.WhatsMeowService
	// SMS (/sms/{action} e /sms/send/credits).
	SMS *messaging.SmsService

	// Dados cadastrais device-based (/dados/cpf, /dados/cnpj...).
	Dados *data.DadosService
	// Vehicles consulta veículos por placa (/vehicles/dados, /vehicles/fipe).
	Vehicles *data.VehiclesService
	// Fipe é a tabela FIPE (/fipe/{action}).
	Fipe *data.FipeService
	// Correios (/correios/{action}).
	Correios *data.CorreiosService
	// Cep é o CEP + geolocalização (/cep/{action}).
	Cep *data.CepService
	// Geolocation (/geolocation/{action}).
	Geolocation *data.GeolocationService
	// Geomatrix é a matriz de distâncias (/geomatrix/{action}).
	Geomatrix *data.GeomatrixService
	// Recognize é o OCR / Google Vision (/recognize/{action}).
	Recognize *data.RecognizeService
	// DDD (/ddd/{action}).
	DDD *data.DddService
	// Holidays são os feriados (/holidays/{action}).
	Holidays *data.HolidaysService
	// Translate é a tradução (/translate/{action}).
	Translate *data.TranslateService
	// Weather é o clima (/weather/{action}).
	Weather *data.WeatherService
	// Loterias (/loterias/{sorteio}/*).
	Loterias *data.LoteriasService
	// DatabaseIP é o GeoIP (/database/ip).
	DatabaseIP *data.DatabaseIPService

	// Consulta são as consultas por crédito (/consulta/{servico}/credits).
	Consulta *data.ConsultaService
	// Ura é a URA reversa / ligações (/ura/call/*).
	Ura *data.UraService
	// ChipVirtual (/chip/virtual/*).
	ChipVirtual *data.ChipVirtualService
	// Bulk é a execução em lote (/bulk/*).
	Bulk *data.BulkService

	// Catalog é o catálogo de APIs, planos, docs e servidores.
	Catalog *platform.CatalogService
	// Account traz saldo, faturas, notificações e tickets.
	Account *platform.AccountService
	// Payments são as recargas e pagamentos (PIX, boleto, cartão).
	Payments *platform.PaymentsService
	// IPWhitelist é a whitelist de IPs da conta.
	IPWhitelist *platform.IPWhitelistService
	// BearerRateLimit é o rate limit por Bearer Token.
	BearerRateLimit *platform.BearerRateLimitService
	// Reports são os relatórios e o dashboard de consumo.
	Reports *platform.ReportsService
}

// ApiBrasil é um alias de Client, para leitura equivalente às demais
// SDKs da plataforma.
type ApiBrasil = Client

// New cria o cliente. Sem config, as credenciais vêm do ambiente.
//
//	api := apibrasil.New()
//	api := apibrasil.New(apibrasil.Config{BearerToken: "...", DeviceToken: "..."})
func New(config ...Config) *Client {
	var resolved Config
	if len(config) > 0 {
		resolved = config[0]
	}
	return NewWithClient(core.NewHTTPClient(resolved))
}

// NewWithClient cria o cliente sobre um *core.HTTPClient já configurado.
func NewWithClient(client *core.HTTPClient) *Client {
	return &Client{
		HTTP: client,

		Auth:    platform.NewAuthService(client),
		Devices: platform.NewDevicesService(client),

		WhatsApp:  messaging.NewWhatsAppService(client),
		Evolution: messaging.NewEvolutionService(client),
		WhatsMeow: messaging.NewWhatsMeowService(client),
		SMS:       messaging.NewSmsService(client),

		Dados:       data.NewDadosService(client),
		Vehicles:    data.NewVehiclesService(client),
		Fipe:        data.NewFipeService(client),
		Correios:    data.NewCorreiosService(client),
		Cep:         data.NewCepService(client),
		Geolocation: data.NewGeolocationService(client),
		Geomatrix:   data.NewGeomatrixService(client),
		Recognize:   data.NewRecognizeService(client),
		DDD:         data.NewDddService(client),
		Holidays:    data.NewHolidaysService(client),
		Translate:   data.NewTranslateService(client),
		Weather:     data.NewWeatherService(client),
		Loterias:    data.NewLoteriasService(client),
		DatabaseIP:  data.NewDatabaseIPService(client),

		Consulta:    data.NewConsultaService(client),
		Ura:         data.NewUraService(client),
		ChipVirtual: data.NewChipVirtualService(client),
		Bulk:        data.NewBulkService(client),

		Catalog:         platform.NewCatalogService(client),
		Account:         platform.NewAccountService(client),
		Payments:        platform.NewPaymentsService(client),
		IPWhitelist:     platform.NewIPWhitelistService(client),
		BearerRateLimit: platform.NewBearerRateLimitService(client),
		Reports:         platform.NewReportsService(client),
	}
}

// Login autentica por email/senha e devolve um cliente já autenticado,
// junto da sessão retornada pela plataforma.
//
//	api, session, err := apibrasil.Login(ctx, apibrasil.Json{
//		"email":    "voce@empresa.com.br",
//		"password": "******",
//	})
//
// Devolve erro quando a conta exige 2FA — nesse caso crie o cliente com
// New e use Auth.Login + Auth.Send2FA + Auth.Verify2FA.
func Login(ctx context.Context, credentials Json, config ...Config) (*Client, Json, error) {
	client := New(config...)

	session, err := client.Auth.Login(ctx, credentials)
	if err != nil {
		return nil, nil, err
	}

	if platform.Requires2FA(session) {
		return nil, session, &core.Error{
			Kind:     core.KindAuthentication,
			Message:  "Esta conta exige autenticação em dois fatores. Use Auth.Login() + Auth.Verify2FA().",
			Response: session,
		}
	}

	return client, session, nil
}

// SetBearerToken define/atualiza o Bearer Token do cliente.
func (c *Client) SetBearerToken(token string) *Client {
	c.HTTP.SetBearerToken(token)
	return c
}

// SetDeviceToken define/atualiza o DeviceToken do cliente.
func (c *Client) SetDeviceToken(token string) *Client {
	c.HTTP.SetDeviceToken(token)
	return c
}

// WithDevice devolve um novo cliente com as mesmas credenciais, mas
// apontando para outro device — útil para gerenciar vários
// números/instâncias.
func (c *Client) WithDevice(deviceToken string) *Client {
	config := c.HTTP.Config()
	config.DeviceToken = deviceToken
	return New(config)
}

// Config devolve a configuração atual do cliente.
func (c *Client) Config() Config { return c.HTTP.Config() }

// Request é a porta de saída genérica: chama qualquer endpoint do
// gateway com os headers de autenticação já configurados. Use para
// rotas que ainda não têm método dedicado na SDK.
//
//	api.Request(ctx, apibrasil.MethodPost, "/consulta/cpf/credits", apibrasil.Json{"cpf": "..."})
func (c *Client) Request(ctx context.Context, method Method, path string, body any, opts ...RequestOptions) (Json, error) {
	return c.HTTP.RequestJSON(ctx, method, path, body, opts...)
}

// Do é como Request, mas devolve o corpo decodificado sem normalizar em
// objeto JSON (útil quando a rota responde uma lista ou texto).
func (c *Client) Do(ctx context.Context, method Method, path string, body any, opts ...RequestOptions) (any, error) {
	return c.HTTP.Do(ctx, method, path, body, opts...)
}
