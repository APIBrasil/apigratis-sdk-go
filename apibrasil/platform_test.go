package apibrasil_test

import (
	"context"
	"testing"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
)

func TestRotasDePlatform(t *testing.T) {
	RunRouteCases(t, []RouteCase{
		{
			Name: "auth.Login",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Auth.Login(ctx, apibrasil.Json{"email": "a@b.c", "password": "x"})
			},
			Path: "/auth/login",
			Body: core.Json{"email": "a@b.c", "password": "x"},
		},
		{
			Name: "auth.Send2FA",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Auth.Send2FA(ctx, apibrasil.Json{"challenge": "abc", "method": "email"})
			},
			Path: "/auth/2fa/send",
		},
		{
			Name: "auth.Verify2FA",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Auth.Verify2FA(ctx, apibrasil.Json{"challenge": "abc", "code": "000000"})
			},
			Path: "/auth/login/verify-2fa",
		},
		{
			Name:   "auth.Me",
			Call:   func(ctx context.Context, api *apibrasil.Client) (any, error) { return api.Auth.Me(ctx) },
			Method: core.MethodGet,
			Path:   "/profile/me",
		},
		{
			Name: "auth.UpdateMe",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Auth.UpdateMe(ctx, apibrasil.Json{"first_name": "Jhon"})
			},
			Method: core.MethodPut,
			Path:   "/profile/me",
		},
		{
			Name:   "auth.Verify",
			Call:   func(ctx context.Context, api *apibrasil.Client) (any, error) { return api.Auth.Verify(ctx) },
			Method: core.MethodGet,
			Path:   "/auth/verify",
		},
		{
			Name:   "devices.List",
			Call:   func(ctx context.Context, api *apibrasil.Client) (any, error) { return api.Devices.List(ctx, nil) },
			Method: core.MethodGet,
			Path:   "/devices",
		},
		{
			Name: "devices.Store",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Devices.Store(ctx, apibrasil.Json{"device_name": "meu-bot", "type": "server"})
			},
			Path: "/devices/store",
		},
		{
			Name:   "devices.Show",
			Call:   func(ctx context.Context, api *apibrasil.Client) (any, error) { return api.Devices.Show(ctx, "") },
			Method: core.MethodGet,
			Path:   "/devices/show?search=dev",
		},
		{
			Name: "devices.Destroy",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Devices.Destroy(ctx, "device-x")
			},
			Method: core.MethodDelete,
			Path:   "/devices/destroy",
			Body:   core.Json{"search": "device-x"},
		},
		{
			Name:   "account.Balance",
			Call:   func(ctx context.Context, api *apibrasil.Client) (any, error) { return api.Account.Balance(ctx) },
			Method: core.MethodGet,
			Path:   "/balance",
		},
		{
			Name:   "account.Invoices",
			Call:   func(ctx context.Context, api *apibrasil.Client) (any, error) { return api.Account.Invoices(ctx) },
			Method: core.MethodGet,
			Path:   "/invoices",
		},
		{
			Name: "account.MarkNotificationRead",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Account.MarkNotificationRead(ctx, 42)
			},
			Method: core.MethodPatch,
			Path:   "/notifications/42/read",
		},
		{
			Name: "account.AddTicketMessage",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Account.AddTicketMessage(ctx, "TCK-1", apibrasil.Json{"message": "oi"})
			},
			Path: "/ticket/TCK-1/messages",
		},
		{
			Name: "payments.Recharge",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Payments.Recharge(ctx, apibrasil.Json{"amount": 50, "type": "pix"})
			},
			Path: "/recharge",
		},
		{
			Name: "payments.PixGenerate",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Payments.PixGenerate(ctx, "santander", apibrasil.Json{"amount": 50})
			},
			Path: "/santander/pix/generate",
		},
		{
			Name: "payments.PixStatus",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Payments.PixStatus(ctx, "inter", "tx-1")
			},
			Method: core.MethodGet,
			Path:   "/inter/pix/tx-1",
		},
		{
			Name: "payments.CardProcess",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Payments.CardProcess(ctx, apibrasil.Json{"token": "t"})
			},
			Path: "/mercadopago/card/process",
		},
		{
			Name:   "catalog.APIs com busca",
			Call:   func(ctx context.Context, api *apibrasil.Client) (any, error) { return api.Catalog.APIs(ctx, "cep") },
			Method: core.MethodGet,
			Path:   "/apis?search=cep",
		},
		{
			Name:   "catalog.Servers",
			Call:   func(ctx context.Context, api *apibrasil.Client) (any, error) { return api.Catalog.Servers(ctx) },
			Method: core.MethodGet,
			Path:   "/servers",
		},
		{
			Name:   "ipWhitelist.Get",
			Call:   func(ctx context.Context, api *apibrasil.Client) (any, error) { return api.IPWhitelist.Get(ctx) },
			Method: core.MethodGet,
			Path:   "/ip-whitelist",
		},
		{
			Name: "ipWhitelist.Set",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.IPWhitelist.Set(ctx, []string{"1.2.3.4"})
			},
			Method: core.MethodPut,
			Path:   "/ip-whitelist",
			Body:   core.Json{"ip_whitelist": []any{"1.2.3.4"}},
		},
		{
			Name: "ipWhitelist.Add",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.IPWhitelist.Add(ctx, "1.2.3.4")
			},
			Path: "/ip-whitelist/add",
			Body: core.Json{"entry": "1.2.3.4"},
		},
		{
			Name:   "bearerRateLimit.Get",
			Call:   func(ctx context.Context, api *apibrasil.Client) (any, error) { return api.BearerRateLimit.Get(ctx) },
			Method: core.MethodGet,
			Path:   "/bearer-rate-limit",
		},
		{
			Name:   "reports.QuickStats",
			Call:   func(ctx context.Context, api *apibrasil.Client) (any, error) { return api.Reports.QuickStats(ctx) },
			Method: core.MethodGet,
			Path:   "/reports/quick-stats",
		},
		{
			Name:   "reports.DashboardStats",
			Call:   func(ctx context.Context, api *apibrasil.Client) (any, error) { return api.Reports.DashboardStats(ctx) },
			Method: core.MethodGet,
			Path:   "/dashboard/stats",
		},
	})
}

func TestDevicesStoreEnviaSecretKey(t *testing.T) {
	api, transport := BuildAPI(apibrasil.Config{SecretKey: "secret-do-cliente"})

	if _, err := api.Devices.Store(context.Background(), apibrasil.Json{"device_name": "bot"}); err != nil {
		t.Fatalf("store falhou: %v", err)
	}
	if got := transport.Last(t).Headers["SecretKey"]; got != "secret-do-cliente" {
		t.Errorf("SecretKey = %q", got)
	}

	if _, err := api.Devices.Store(
		context.Background(),
		apibrasil.Json{"device_name": "bot"},
		apibrasil.RequestOptions{SecretKey: "secret-da-requisicao"},
	); err != nil {
		t.Fatalf("store falhou: %v", err)
	}
	if got := transport.Last(t).Headers["SecretKey"]; got != "secret-da-requisicao" {
		t.Errorf("SecretKey = %q", got)
	}
}

func TestPaymentsBoletoPDFDevolveBytes(t *testing.T) {
	api, transport := BuildAPI()
	transport.RespondWith(&core.TransportResponse{Status: 200, Data: []byte("%PDF-1.4")})

	pdf, err := api.Payments.BoletoPDF(context.Background(), "inter", "123")
	if err != nil {
		t.Fatalf("boletoPDF falhou: %v", err)
	}
	if string(pdf) != "%PDF-1.4" {
		t.Errorf("pdf = %q", pdf)
	}
	if transport.Last(t).ResponseType != core.ResponseBytes {
		t.Error("responseType deveria ser ResponseBytes")
	}
}
