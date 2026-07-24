# SDK GO - APIGratis by API BRASIL 🚀

SDK oficial Go da plataforma [APIBrasil](https://apibrasil.com.br) — WhatsApp, SMS, consultas de CPF/CNPJ, veículos, CEP, correios, pagamentos PIX/boleto e muito mais.

[![Go Reference](https://pkg.go.dev/badge/github.com/jhowbhz/apigratis-sdk-go.svg)](https://pkg.go.dev/github.com/jhowbhz/apigratis-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/jhowbhz/apigratis-sdk-go)](https://goreportcard.com/report/github.com/jhowbhz/apigratis-sdk-go)
<a href="https://github.com/APIBrasil/apigratis-sdk-go/issues" target="_blank"><img alt="GitHub issues" src="https://img.shields.io/github/issues/APIBrasil/apigratis-sdk-go"></a>
<a href="https://github.com/APIBrasil/apigratis-sdk-go/network" target="_blank"><img alt="GitHub forks" src="https://img.shields.io/github/forks/APIBrasil/apigratis-sdk-go"></a>
<a href="https://github.com/APIBrasil/apigratis-sdk-go/stargazers" target="_blank"><img alt="GitHub stars" src="https://img.shields.io/github/stars/APIBrasil/apigratis-sdk-go"></a>

## Canais de suporte (Comunidade)

[![WhatsApp Group](https://img.shields.io/badge/WhatsApp-Group-25D366?logo=whatsapp)](https://whatsapp.com/channel/0029VaMiaT6B4hdX3hrUcz3X)
[![Telegram Group](https://img.shields.io/badge/Telegram-Group-32AFED?logo=telegram)](https://t.me/apigratisoficial)

## Instalação

```bash
go get github.com/jhowbhz/apigratis-sdk-go
```

Requer **Go >= 1.22**. **Zero dependências** de runtime — só a biblioteca padrão.

Obtenha suas credenciais em https://apibrasil.com.br

## Começando

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil"
)

func main() {
	ctx := context.Background()

	api := apibrasil.New(apibrasil.Config{
		BearerToken: "SEU_BEARER_TOKEN", // JWT do login
		DeviceToken: "SEU_DEVICE_TOKEN", // device dos serviços device-based
	})

	// WhatsApp
	if _, err := api.WhatsApp.SendText(ctx, apibrasil.Json{
		"number": "5511999999999",
		"text":   "Olá! 👋",
	}); err != nil {
		log.Fatal(err)
	}

	// Consulta CNPJ (por créditos)
	empresa, err := api.Consulta.CNPJ(ctx, apibrasil.Json{"cnpj": "00000000000000"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(empresa.Data())
}
```

As credenciais também podem vir só do ambiente — `apibrasil.New()` lê automaticamente `APIBRASIL_BEARER_TOKEN`, `APIBRASIL_DEVICE_TOKEN`, `APIBRASIL_SECRET_KEY` e `APIBRASIL_BASE_URL`.

Também é possível autenticar por email/senha — o token retornado fica guardado no cliente:

```go
api, session, err := apibrasil.Login(ctx, apibrasil.Json{
	"email":    "voce@empresa.com.br",
	"password": "******",
})

// contas com 2FA:
api := apibrasil.New()
session, _ := api.Auth.Login(ctx, apibrasil.Json{"email": email, "password": senha})
if session["requires_2fa"] == true {
	challenge := session["challenge"]
	api.Auth.Send2FA(ctx, apibrasil.Json{"challenge": challenge, "method": "email"})
	api.Auth.Verify2FA(ctx, apibrasil.Json{"challenge": challenge, "code": "000000"})
}
```

## Como a plataforma funciona

A API Brasil tem duas famílias de serviços:

| Família          | Autenticação                                   | Exemplos                                                                    |
| ---------------- | ---------------------------------------------- | --------------------------------------------------------------------------- |
| **Device-based** | `Authorization: Bearer` + header `DeviceToken` | WhatsApp, SMS, veículos, CEP, correios, DDD, feriados, tradução, clima, OCR |
| **Por créditos** | apenas `Authorization: Bearer` (debita saldo)  | `Consulta.CPF`, `Consulta.CNPJ`, `Consulta.Veiculos`, Serasa, CNH, telefone |

Para os serviços device-based, crie um device com a `SecretKey` da API desejada (painel APIBrasil) e use o `device_token` retornado:

```go
device, err := api.Devices.Store(ctx,
	apibrasil.Json{"device_name": "meu-bot", "type": "server"},
	apibrasil.RequestOptions{SecretKey: "SUA_SECRET_KEY"},
)

api.SetDeviceToken("device_token_retornado")
```

## Serviços disponíveis

| Módulo                                                       | Descrição                                                                                             |
| ------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------- |
| `api.WhatsApp`                                               | WhatsApp: `Start`, `QRCode`, `SendText`, `SendFile`, `SendAudio`, `SendVideo`, fila (`Queue`)...      |
| `api.Evolution`                                              | Evolution API: `Request(ctx, controller, action, body)`                                               |
| `api.WhatsMeow`                                              | WhatsMeow: `Request(ctx, action, body)`                                                               |
| `api.SMS`                                                    | SMS device-based (`Send`) e por créditos (`SendWithCredits`)                                          |
| `api.Dados`                                                  | Dados cadastrais device-based (`CPF`, `CNPJ`, `ListaSocios`...)                                       |
| `api.Vehicles`                                               | Veículos por placa (`Dados`, `Fipe`, `ConsultaFipe`, `BaseDados`)                                     |
| `api.Fipe`                                                   | Tabela FIPE (`ConsultarMarcas`, `ConsultarModelos`...)                                                |
| `api.Correios`                                               | Correios (`Rastreio`, `Request`)                                                                      |
| `api.Cep`                                                    | CEP + geolocalização (`CEP`, `Cidades`, `Estados`, `CalcularDistancia`)                               |
| `api.Geolocation` / `api.Geomatrix`                          | Geolocalização e matriz de distâncias                                                                 |
| `api.Recognize`                                              | OCR / Google Vision (`Base64`, `URI`)                                                                 |
| `api.DDD` / `api.Holidays` / `api.Translate` / `api.Weather` | DDD, feriados, tradução, clima                                                                        |
| `api.Loterias`                                               | Loterias (`Latest`, `Resultado`)                                                                      |
| `api.DatabaseIP`                                             | GeoIP (`IP`)                                                                                          |
| `api.Consulta`                                               | Consultas por créditos: `CPF`, `CNPJ`, `CNH`, `CEP`, `Veiculos`, `Telefone`, `Generic(service, body)` |
| `api.Ura` / `api.ChipVirtual`                                | URA reversa e chip virtual                                                                            |
| `api.Bulk`                                                   | Execução em lote (`Direct`, `Queue`)                                                                  |
| `api.Auth`                                                   | Login, 2FA, cadastro, recuperação de senha, perfil                                                    |
| `api.Devices`                                                | CRUD de devices                                                                                       |
| `api.Catalog`                                                | Catálogo de APIs, planos, documentações, servidores                                                   |
| `api.Account`                                                | Saldo, faturas, notificações, tickets                                                                 |
| `api.Payments`                                               | Recargas e pagamentos PIX/boleto/cartão (Santander, Inter, Mercado Pago, Sicoob)                      |
| `api.IPWhitelist` / `api.BearerRateLimit`                    | Segurança da conta                                                                                    |
| `api.Reports`                                                | Relatórios e dashboard de consumo                                                                     |

Todo método recebe `context.Context` como primeiro parâmetro e aceita `apibrasil.RequestOptions` opcional no final.

### WhatsApp

```go
// iniciar sessão e obter QR Code
api.WhatsApp.Start(ctx, apibrasil.Json{
	"webhook_wh_message": "https://seu-webhook.com/mensagens",
})

qr, _ := api.WhatsApp.QRCode(ctx, nil)
fmt.Println(qr.Response()) // data URI base64

// envios
api.WhatsApp.SendText(ctx, apibrasil.Json{"number": "5511999999999", "text": "Olá!"})
api.WhatsApp.SendFile(ctx, apibrasil.Json{
	"number": "5511999999999",
	"path":   "https://exemplo.com/boleto.pdf",
})
api.WhatsApp.SendLocation(ctx, apibrasil.Json{"number": "5511999999999", "lat": -23.5, "lng": -46.6})

// qualquer action do catálogo
api.WhatsApp.Request(ctx, "getAllChats", nil)

// fila assíncrona
api.WhatsApp.Queue(ctx, "sendText", apibrasil.Json{"number": "5511999999999", "text": "vai por fila"})
```

O envelope device-based tem acessores tipados — e continua sendo um `map[string]any`:

```go
res, _ := api.WhatsApp.SendText(ctx, apibrasil.Json{"number": "...", "text": "..."})

res.IsError()   // bool
res.Message()   // string
res.Response()  // any
res.APILimit()  // any
res["response"] // acesso direto por chave

var enviado struct {
	ID string `json:"id"`
}
res.Decode(&enviado) // decodifica o campo response em uma struct
```

### Consultas por créditos

```go
cpf, _ := api.Consulta.CPF(ctx, apibrasil.Json{"cpf": "00000000000"})
fmt.Println(cpf.Balance(), cpf.Data())

// o campo `tipo` define o produto consultado
socios, _ := api.Consulta.CNPJ(ctx, apibrasil.Consulta{
	Tipo:   "lista-socios",
	Fields: apibrasil.Json{"cnpj": "00000000000000"},
}.Json())

// modo homologação (sandbox, sem cobrança)
api.Consulta.CNPJ(ctx, apibrasil.Consulta{
	Tipo:    "serasa-score-pj",
	Homolog: true,
	Fields:  apibrasil.Json{"cnpj": "00000000000000"},
}.Json())

// qualquer serviço do catálogo
api.Consulta.Generic(ctx, "cnh", apibrasil.Json{"cpf": "00000000000"})
```

### Veículos e FIPE (device-based)

```go
api.Vehicles.Dados(ctx, apibrasil.Json{"placa": "ABC1234"})
api.Vehicles.Fipe(ctx, apibrasil.Json{"placa": "ABC1234"})
api.Fipe.ConsultarMarcas(ctx, apibrasil.Json{"codigoTabelaReferencia": 300})
```

### SMS

```go
api.SMS.Send(ctx, apibrasil.Json{"number": "5511999999999", "message": "Olá!"})
api.SMS.SendWithCredits(ctx, apibrasil.Json{"number": "5511999999999", "message": "Olá!"})
```

### Pagamentos e recargas

```go
api.Payments.Recharge(ctx, apibrasil.Json{"amount": 50, "type": "pix"})
api.Payments.PixGenerate(ctx, "santander", apibrasil.Json{"amount": 50})
api.Payments.PixStatus(ctx, "santander", "TX_ID")

pdf, _ := api.Payments.BoletoPDF(ctx, "inter", "ID") // []byte
os.WriteFile("boleto.pdf", pdf, 0o644)
```

### Múltiplos devices

```go
bot1 := api.WithDevice("device_token_1")
bot2 := api.WithDevice("device_token_2")

bot1.WhatsApp.SendText(ctx, apibrasil.Json{"number": "5511999999999", "text": "do bot 1"})
bot2.WhatsApp.SendText(ctx, apibrasil.Json{"number": "5511999999999", "text": "do bot 2"})
```

## Tratamento de erros

Cada categoria de falha tem o seu sentinela — trate com `errors.Is`:

| Sentinela                          | Quando                                |
| ---------------------------------- | ------------------------------------- |
| `ErrValidation`                    | 400/422 — payload inválido            |
| `ErrAuthentication`                | 401 — token ausente/expirado          |
| `ErrInsufficientBalance`           | 402 — sem saldo/créditos              |
| `ErrPermission`                    | 403 — sem permissão (ex: exige PJ)    |
| `ErrNotFound`                      | 404/410 — sem dados / rota desativada |
| `ErrRateLimit`                     | 429 — limite atingido (`RetryAfter`)  |
| `ErrServer`                        | 5xx — erro do gateway/provedor        |
| `ErrNetwork` / `ErrTimeout`        | falha antes da resposta               |
| `ErrAPIBrasil`                     | qualquer erro da SDK                  |

```go
_, err := api.Consulta.CPF(ctx, apibrasil.Json{"cpf": "00000000000"})

switch {
case errors.Is(err, apibrasil.ErrInsufficientBalance):
	log.Println("Recarregue seus créditos")
case errors.Is(err, apibrasil.ErrRateLimit):
	if apiErr, ok := apibrasil.AsError(err); ok {
		log.Printf("Aguarde %s", apiErr.RetryAfter)
	}
}

// detalhes completos da falha
if apiErr, ok := apibrasil.AsError(err); ok {
	fmt.Println(apiErr.Status, apiErr.Code, apiErr.Message, apiErr.Response)
}
```

## Retry e observabilidade

Por padrão a SDK refaz a chamada em **HTTP 429** e em **falhas de conexão** (2 tentativas extras, backoff exponencial, respeitando `Retry-After`). Timeouts e erros de negócio nunca são refeitos — evita duplicar cobranças e envios.

```go
api := apibrasil.New(apibrasil.Config{
	Retry: &apibrasil.RetryConfig{ // ou apibrasil.NoRetry()
		Retries:         3,
		MinDelay:        500 * time.Millisecond,
		RetryOnStatuses: []int{429, 503},
	},
	Hooks: &apibrasil.Hooks{
		OnRequest: func(_ context.Context, info apibrasil.RequestHookInfo) {
			log.Printf("→ %s %s (#%d)", info.Method, info.URL, info.Attempt)
		},
		OnResponse: func(_ context.Context, info apibrasil.ResponseHookInfo) {
			log.Printf("← %d em %s", info.Status, info.Duration)
		},
		OnRetry: func(_ context.Context, info apibrasil.RetryHookInfo) {
			log.Printf("retry em %s: %s", info.Delay, info.Reason)
		},
	},
})
```

O `context.Context` também cancela e limita qualquer chamada:

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

api.WhatsApp.SendText(ctx, apibrasil.Json{"number": "...", "text": "..."})
```

## Transporte plugável

O HTTP é feito pelo `net/http`, mas a interface `Transport` permite trocar a camada inteira (proxy corporativo, instrumentação, mocks de teste):

```go
type MeuTransporte struct{}

func (MeuTransporte) Do(ctx context.Context, req apibrasil.TransportRequest) (*apibrasil.TransportResponse, error) {
	// use o cliente HTTP que quiser e devolva Status, Headers e Data
	return &apibrasil.TransportResponse{Status: 200, Data: apibrasil.Json{"ok": true}}, nil
}

api := apibrasil.New(apibrasil.Config{Transport: MeuTransporte{}})
```

Para apenas configurar proxy, TLS ou pool de conexões, reaproveite o transporte padrão:

```go
api := apibrasil.New(apibrasil.Config{
	Transport: apibrasil.NewHTTPTransport(&http.Client{Timeout: 60 * time.Second}),
})
```

## Pacotes por domínio

```go
import (
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil"                    // cliente + aliases do core
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"               // HTTP, transporte, erros, retry
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services/messaging" // WhatsApp, SMS, Evolution, WhatsMeow
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services/data"      // consultas, veículos, CEP, FIPE...
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services/platform"  // auth, devices, pagamentos, relatórios
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/generated"          // catálogo gerado
	"github.com/jhowbhz/apigratis-sdk-go/apigratis"                    // interface legada (deprecada)
)
```

Cada serviço também pode ser usado isoladamente sobre o mesmo cliente HTTP:

```go
client := apibrasil.NewHTTPClient(apibrasil.Config{BearerToken: "..."})
whatsapp := messaging.NewWhatsAppService(client)
```

## Catálogo gerado

As actions de WhatsApp/Evolution/WhatsMeow e os 210+ `tipo` de consulta são gerados do catálogo real da plataforma:

```bash
go generate ./...           # ou: go run ./cmd/codegen
```

```go
import "github.com/jhowbhz/apigratis-sdk-go/apibrasil/generated"

generated.WhatsAppActions              // []string com todas as actions
generated.ServiceActionsFor("cep")     // ["bairros", "cep", "cidades", ...]
generated.HasAction("whatsapp", "sendText")

meta, ok := generated.ConsultaTipo("acerta-essencial")
// meta.Service = "cpf", meta.Fields = ["cpf"]
```

## Endpoint sem método dedicado?

Todo o gateway fica acessível pela porta de saída genérica, já com seus headers de autenticação:

```go
api.Request(ctx, apibrasil.MethodPost, "/consulta/cpf/credits", apibrasil.Json{"cpf": "00000000000"})
api.Request(ctx, apibrasil.MethodGet, "/reports/quick-stats", nil)
```

Documentação completa dos endpoints: https://doc.apibrasil.io

## Configuração avançada

```go
api := apibrasil.New(apibrasil.Config{
	BearerToken: "...",                                  // ou APIBRASIL_BEARER_TOKEN
	DeviceToken: "...",                                  // ou APIBRASIL_DEVICE_TOKEN
	SecretKey:   "...",                                  // usada em Devices.Store (ou APIBRASIL_SECRET_KEY)
	BaseURL:     "https://gateway.apibrasil.io/api/v2",  // padrão (ou APIBRASIL_BASE_URL)
	Timeout:     30 * time.Second,
	Headers:     map[string]string{"X-Custom": "valor"}, // headers extras
	Retry:       &apibrasil.RetryConfig{Retries: 2},     // ou apibrasil.NoRetry()
	Hooks:       &apibrasil.Hooks{},
	Transport:   nil,                                    // Transport customizado (padrão: net/http)
})
```

## Interface legada (v0.1.x)

O pacote `apigratis` (`NewService`, `Whatsapp`, `Sms`, `Cpf`, `Cnpj`) continua funcionando exatamente como antes, mas está **deprecado** — prefira o cliente `apibrasil`.

```go
import "github.com/jhowbhz/apigratis-sdk-go/apigratis"

service := apigratis.NewService()

payload, _ := json.Marshal(map[string]any{
	"action": "sendText",
	"credentials": map[string]any{
		"DeviceToken": "YOUR_DEVICE_TOKEN",
		"BearerToken": "YOUR_BEARER_TOKEN",
	},
	"body": map[string]any{
		"number":      "5511999999999",
		"text":        "Hello World for Go",
		"time_typing": 1,
	},
})

response, err := service.Whatsapp(string(payload))
```

## Desenvolvimento

```bash
go build ./...
go test ./...
go vet ./...
gofmt -l .

APIBRASIL_CONTRACT=1 go test ./apibrasil -run Contrato -v  # testes de contrato (gateway real)
```

## Licença

MIT — veja [LICENSE](LICENSE).
