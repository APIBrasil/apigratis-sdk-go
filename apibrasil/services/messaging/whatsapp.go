// Package messaging traz os serviços de mensageria da plataforma:
// WhatsApp, Evolution API, WhatsMeow e SMS.
package messaging

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// WhatsAppService é a API de WhatsApp (device-based):
// POST /whatsapp/{action}. Exige Authorization: Bearer + DeviceToken.
//
// Além dos métodos nomeados, Request(ctx, action, body) aceita qualquer
// action do catálogo (veja generated.WhatsAppActions).
type WhatsAppService struct {
	*services.DeviceProxyService
}

// NewWhatsAppService cria o serviço de WhatsApp.
func NewWhatsAppService(client *core.HTTPClient) *WhatsAppService {
	return &WhatsAppService{DeviceProxyService: services.NewDeviceProxyService(client, "whatsapp")}
}

// Start inicia a sessão do device (aceita webhooks opcionais):
// webhook_wh_status, webhook_wh_message, webhook_wh_connect, webhook_wh_qrcode.
func (s *WhatsAppService) Start(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "start", body, opts...)
}

// QRCode retorna o QR Code de pareamento (response.qrcode em base64).
func (s *WhatsAppService) QRCode(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "qrcode", body, opts...)
}

// Logout encerra a sessão.
func (s *WhatsAppService) Logout(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "logout", body, opts...)
}

// Close fecha o navegador/sessão no servidor.
func (s *WhatsAppService) Close(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "close", body, opts...)
}

// DeleteSession apaga a sessão no servidor.
func (s *WhatsAppService) DeleteSession(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "deleteSession", body, opts...)
}

// RestartSession reinicia a sessão do device.
func (s *WhatsAppService) RestartSession(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "restartSession", body, opts...)
}

// GetConnectionStatus consulta o status da conexão do device.
func (s *WhatsAppService) GetConnectionStatus(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "getConnectionStatus", body, opts...)
}

// SendText envia uma mensagem de texto:
// {"number": "5511999999999", "text": "Olá!"}.
func (s *WhatsAppService) SendText(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendText", body, opts...)
}

// SendFile envia um arquivo a partir de uma URL: {"number": ..., "path": ...}.
func (s *WhatsAppService) SendFile(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendFile", body, opts...)
}

// SendFile64 envia um arquivo em base64.
func (s *WhatsAppService) SendFile64(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendFile64", body, opts...)
}

// SendAudio envia áudio (URL; convertido para mp3 pelo gateway, máx. 6 min).
func (s *WhatsAppService) SendAudio(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendAudio", body, opts...)
}

// SendVideo envia um vídeo a partir de uma URL.
func (s *WhatsAppService) SendVideo(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendVideo", body, opts...)
}

// SendGif envia um GIF.
func (s *WhatsAppService) SendGif(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendGif", body, opts...)
}

// SendLink envia um link com preview.
func (s *WhatsAppService) SendLink(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendLink", body, opts...)
}

// SendLocation envia uma localização: {"number": ..., "lat": ..., "lng": ...}.
func (s *WhatsAppService) SendLocation(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendLocation", body, opts...)
}

// SendContact envia um contato.
func (s *WhatsAppService) SendContact(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendContact", body, opts...)
}

// SendSticker envia uma figurinha.
func (s *WhatsAppService) SendSticker(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendSticker", body, opts...)
}

// SendList envia uma mensagem de lista.
func (s *WhatsAppService) SendList(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendList", body, opts...)
}

// SendButtons envia uma mensagem com botões.
func (s *WhatsAppService) SendButtons(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendButtons", body, opts...)
}

// SendPollMessage envia uma enquete.
func (s *WhatsAppService) SendPollMessage(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendPollMessage", body, opts...)
}

// SendPixKey envia uma chave PIX.
func (s *WhatsAppService) SendPixKey(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendPixKey", body, opts...)
}

// SendTextToStorie publica um texto no status.
func (s *WhatsAppService) SendTextToStorie(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendTextToStorie", body, opts...)
}

// SendImageToStorie publica uma imagem no status.
func (s *WhatsAppService) SendImageToStorie(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendImageToStorie", body, opts...)
}

// SendVideoToStorie publica um vídeo no status.
func (s *WhatsAppService) SendVideoToStorie(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "sendVideoToStorie", body, opts...)
}

// CheckNumberStatus verifica se um número tem WhatsApp.
func (s *WhatsAppService) CheckNumberStatus(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "checkNumberStatus", body, opts...)
}

// GetAllChats lista as conversas do device.
func (s *WhatsAppService) GetAllChats(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "getAllChats", body, opts...)
}

// GetAllContacts lista os contatos do device.
func (s *WhatsAppService) GetAllContacts(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "getAllContacts", body, opts...)
}

// GetAllGroups lista os grupos do device.
func (s *WhatsAppService) GetAllGroups(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "getAllGroups", body, opts...)
}

// GetProfilePic retorna a foto de perfil de um número.
func (s *WhatsAppService) GetProfilePic(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "getProfilePic", body, opts...)
}
