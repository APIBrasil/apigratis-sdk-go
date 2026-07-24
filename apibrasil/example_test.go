package apibrasil_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/generated"
)

func ExampleNew() {
	api := apibrasil.New(apibrasil.Config{
		BearerToken: "SEU_BEARER_TOKEN",
		DeviceToken: "SEU_DEVICE_TOKEN",
	})

	_, err := api.WhatsApp.SendText(context.Background(), apibrasil.Json{
		"number": "5511999999999",
		"text":   "Olá!",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleNew_ambiente() {
	// Lê APIBRASIL_BEARER_TOKEN, APIBRASIL_DEVICE_TOKEN,
	// APIBRASIL_SECRET_KEY e APIBRASIL_BASE_URL.
	api := apibrasil.New()

	saldo, err := api.Account.Balance(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(saldo)
}

func ExampleLogin() {
	api, session, err := apibrasil.Login(context.Background(), apibrasil.Json{
		"email":    "voce@empresa.com.br",
		"password": "******",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(session["user"])
	_, _ = api.Account.Balance(context.Background())
}

func ExampleClient_Request() {
	api := apibrasil.New()

	// Porta de saída genérica para rotas sem método dedicado.
	response, err := api.Request(
		context.Background(),
		apibrasil.MethodPost,
		"/consulta/cpf/credits",
		apibrasil.Json{"cpf": "00000000000"},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}

func ExampleClient_WithDevice() {
	api := apibrasil.New()

	bot1 := api.WithDevice("device_token_1")
	bot2 := api.WithDevice("device_token_2")

	ctx := context.Background()
	_, _ = bot1.WhatsApp.SendText(ctx, apibrasil.Json{"number": "5511999999999", "text": "do bot 1"})
	_, _ = bot2.WhatsApp.SendText(ctx, apibrasil.Json{"number": "5511999999999", "text": "do bot 2"})
}

func ExampleConsulta() {
	api := apibrasil.New()

	socios, err := api.Consulta.CNPJ(context.Background(), apibrasil.Consulta{
		Tipo:    "serasa-score-pj",
		Homolog: true,
		Fields:  apibrasil.Json{"cnpj": "00000000000000"},
	}.Json())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(socios.Balance(), socios.Data())
}

func ExampleAsError() {
	api := apibrasil.New()

	_, err := api.Consulta.CPF(context.Background(), apibrasil.Json{"cpf": "00000000000"})

	switch {
	case errors.Is(err, apibrasil.ErrInsufficientBalance):
		fmt.Println("recarregue seus créditos")
	case errors.Is(err, apibrasil.ErrRateLimit):
		if apiErr, ok := apibrasil.AsError(err); ok {
			fmt.Printf("aguarde %s\n", apiErr.RetryAfter)
		}
	case err != nil:
		apiErr, _ := apibrasil.AsError(err)
		fmt.Println(apiErr.Status, apiErr.Code, apiErr.Message)
	}
}

func ExampleConfig_retry() {
	api := apibrasil.New(apibrasil.Config{
		Timeout: 45 * time.Second,
		Retry: &apibrasil.RetryConfig{
			Retries:         3,
			MinDelay:        500 * time.Millisecond,
			RetryOnStatuses: []int{429, 503},
		},
		Hooks: &apibrasil.Hooks{
			OnResponse: func(_ context.Context, info apibrasil.ResponseHookInfo) {
				log.Printf("← %d em %s", info.Status, info.Duration)
			},
		},
		Transport: apibrasil.NewHTTPTransport(&http.Client{}),
	})

	_, _ = api.Catalog.Status(context.Background())
}

func ExampleClient_generatedCatalog() {
	// Actions e tipos de consulta conhecidos, gerados do catálogo real.
	fmt.Println(generated.ServiceActionsFor("cep"))
	fmt.Println(generated.HasAction("whatsapp", "sendText"))

	if meta, ok := generated.ConsultaTipo("acerta-essencial"); ok {
		fmt.Println(meta.Service, meta.Fields)
	}
}
