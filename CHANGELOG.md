# Changelog

## 1.0.0 — 2026-07-23

Reescrita completa da SDK, agora cobrindo toda a plataforma APIBrasil —
mesma arquitetura das SDKs de Node.js, PHP e Dart/Flutter.

### Novidades

- **Cliente central `apibrasil.Client`** com módulos por produto: `WhatsApp`, `Evolution`, `WhatsMeow`, `SMS`, `Dados`, `Vehicles`, `Fipe`, `Correios`, `Cep`, `Geolocation`, `Geomatrix`, `Recognize`, `DDD`, `Holidays`, `Translate`, `Weather`, `Loterias`, `DatabaseIP`, `Consulta` (créditos), `Ura`, `ChipVirtual`, `Bulk`, `Auth` (login/2FA), `Devices`, `Catalog`, `Account`, `Payments` (PIX/boleto/cartão), `IPWhitelist`, `BearerRateLimit`, `Reports`.
- **Zero dependências**: transporte padrão sobre `net/http`; interface `Transport` plugável para proxies, instrumentação e mocks.
- **`context.Context` em toda chamada**, com timeout e cancelamento propagados até o transporte.
- **Retry com backoff exponencial** (padrão: HTTP 429 e falhas de conexão; nunca timeouts nem erros de negócio) com suporte a `Retry-After`.
- **Hooks de observabilidade**: `OnRequest`, `OnResponse`, `OnRetry`.
- **Erros tipados** com sentinelas para `errors.Is`: `ErrValidation`, `ErrAuthentication`, `ErrInsufficientBalance`, `ErrPermission`, `ErrNotFound`, `ErrRateLimit`, `ErrServer`, `ErrNetwork`, `ErrTimeout` — todos casando também com `ErrAPIBrasil`, e `*apibrasil.Error` acessível via `errors.As`/`AsError`.
- **Envelopes com acessores**: `DeviceServiceResponse` e `CreditServiceResponse` seguem sendo `map[string]any`, com `IsError()`, `Message()`, `Response()`/`Data()`, `Balance()` e `Decode(&struct)`.
- **Variáveis de ambiente**: `APIBRASIL_BEARER_TOKEN`, `APIBRASIL_DEVICE_TOKEN`, `APIBRASIL_SECRET_KEY` e `APIBRASIL_BASE_URL` lidas automaticamente.
- **Catálogo gerado** (`go generate ./...`): actions de WhatsApp/Evolution/WhatsMeow, serviços de consulta e os 210+ `tipo` com seus campos.
- **Pacotes por domínio**: `apibrasil/core`, `apibrasil/services/messaging`, `/data`, `/platform`, `apibrasil/generated`.
- **Testes** unitários com transporte fake (rotas, headers, erros, retry, transporte, legado) e de contrato opcionais (`APIBRASIL_CONTRACT=1`).
- CI no GitHub Actions (Go 1.22/1.23/1.24) com `gofmt`, `go vet`, `go test -race` e build.

### Breaking changes

- Go **>= 1.22**.
- O pacote `apigratis` (`NewService`, `Whatsapp`, `Sms`, `Cpf`, `Cnpj`) segue funcionando com o mesmo contrato — inclusive devolvendo erros da API decodificados em vez de `error` — agora implementado sobre o novo núcleo e marcado como **deprecated**. Prefira `apibrasil.New(...)`.

## 0.1.x

Primeira versão da SDK: pacote `apigratis` com `Service.Request` sobre
`net/http` e os atalhos `Whatsapp`, `Sms`, `Cpf` e `Cnpj`.
