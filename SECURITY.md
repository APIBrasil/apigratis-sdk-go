# Security Policy

## Versões suportadas

| Versão da SDK | Go        | Suportada          |
| ------------- | --------- | ------------------ |
| 1.x           | >= 1.22   | :white_check_mark: |
| 0.1.x         | >= 1.18   | :x:                |

## Reportando uma vulnerabilidade

Não abra issues públicas para falhas de segurança.

Envie os detalhes para **contato@apibrasil.com.br** com:

- descrição da falha e impacto;
- passos para reproduzir (de preferência com um caso mínimo);
- versão da SDK e do Go.

Você recebe uma confirmação em até 72 horas úteis e atualizações a cada
7 dias até a conclusão. Correções de segurança saem em uma release
dedicada, com crédito ao autor do report quando autorizado.

## Boas práticas ao usar a SDK

- Nunca versione `BearerToken`, `DeviceToken` ou `SecretKey` — use as
  variáveis de ambiente `APIBRASIL_BEARER_TOKEN`, `APIBRASIL_DEVICE_TOKEN`
  e `APIBRASIL_SECRET_KEY`.
- Restrinja de onde o token pode ser usado com `api.IPWhitelist`.
- Limite a taxa de uso do token com `api.BearerRateLimit`.
- Rotacione tokens comprometidos com `api.Auth.TokenRotate` e revogue os
  antigos com `api.Auth.TokenRevoke`.
