package apibrasil_test

import (
	"context"
	"testing"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
)

func TestRotasDeMessaging(t *testing.T) {
	RunRouteCases(t, []RouteCase{
		{
			Name: "whatsapp.Start",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.WhatsApp.Start(ctx, apibrasil.Json{"webhook_wh_message": "https://webhook.dev"})
			},
			Path: "/whatsapp/start",
			Body: core.Json{"webhook_wh_message": "https://webhook.dev"},
		},
		{
			Name: "whatsapp.QRCode",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.WhatsApp.QRCode(ctx, nil)
			},
			Path: "/whatsapp/qrcode",
		},
		{
			Name: "whatsapp.SendText",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.WhatsApp.SendText(ctx, apibrasil.Json{"number": "5511999999999", "text": "Olá!"})
			},
			Path: "/whatsapp/sendText",
			Body: core.Json{"number": "5511999999999", "text": "Olá!"},
		},
		{
			Name: "whatsapp.SendFile",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.WhatsApp.SendFile(ctx, apibrasil.Json{"number": "5511999999999", "path": "https://x.dev/a.pdf"})
			},
			Path: "/whatsapp/sendFile",
		},
		{
			Name: "whatsapp.Request (action livre)",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.WhatsApp.Request(ctx, "getAllChats", nil)
			},
			Path: "/whatsapp/getAllChats",
		},
		{
			Name: "whatsapp.Queue",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.WhatsApp.Queue(ctx, "sendText", apibrasil.Json{"number": "1", "text": "a"})
			},
			Path: "/whatsapp/sendText/queue",
		},
		{
			Name: "evolution.Request",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Evolution.Request(ctx, "message", "sendText", apibrasil.Json{"number": "1"})
			},
			Path: "/evolution/message/sendText",
		},
		{
			Name: "evolution.Call",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Evolution.Call(ctx, "instance/create", apibrasil.Json{"instanceName": "bot"})
			},
			Path: "/evolution/instance/create",
		},
		{
			Name: "evolution.Queue",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Evolution.Queue(ctx, "message", "sendText", nil)
			},
			Path: "/evolution/message/sendText/queue",
		},
		{
			Name: "whatsmeow.Request",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.WhatsMeow.Request(ctx, "send/text", apibrasil.Json{"number": "1"})
			},
			Path: "/whatsmeow/send/text",
		},
		{
			Name: "whatsmeow.SendText",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.WhatsMeow.SendText(ctx, apibrasil.Json{"number": "1"})
			},
			Path: "/whatsmeow/send/text",
		},
		{
			Name: "whatsmeow.Queue",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.WhatsMeow.Queue(ctx, "send/text", nil)
			},
			Path: "/whatsmeow/send/text/queue",
		},
		{
			Name: "sms.Send",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.SMS.Send(ctx, apibrasil.Json{"number": "5511999999999", "message": "Olá"})
			},
			Path: "/sms/send",
			Body: core.Json{"number": "5511999999999", "message": "Olá"},
		},
		{
			Name: "sms.SendWithCredits",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.SMS.SendWithCredits(ctx, apibrasil.Json{"number": "5511999999999", "message": "Olá"})
			},
			Path: "/sms/send/credits",
		},
	})
}
