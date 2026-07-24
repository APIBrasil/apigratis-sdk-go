package apibrasil_test

import (
	"context"
	"testing"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
)

func TestRotasDeData(t *testing.T) {
	RunRouteCases(t, []RouteCase{
		{
			Name: "dados.CPF",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Dados.CPF(ctx, apibrasil.Json{"cpf": "00000000000"})
			},
			Path: "/dados/cpf",
			Body: core.Json{"cpf": "00000000000"},
		},
		{
			Name: "dados.CNPJ",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Dados.CNPJ(ctx, apibrasil.Json{"cnpj": "00000000000000"})
			},
			Path: "/dados/cnpj",
		},
		{
			Name: "vehicles.Dados",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Vehicles.Dados(ctx, apibrasil.Json{"placa": "ABC1234"})
			},
			Path: "/vehicles/dados",
		},
		{
			Name: "vehicles.Fipe",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Vehicles.Fipe(ctx, apibrasil.Json{"placa": "ABC1234"})
			},
			Path: "/vehicles/fipe",
		},
		{
			Name: "vehicles.ConsultaFipe",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Vehicles.ConsultaFipe(ctx, "ABC1234")
			},
			Path: "/vehicles/consultafipe/ABC1234",
		},
		{
			Name: "vehicles.BaseDados",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Vehicles.BaseDados(ctx, apibrasil.Json{"placa": "ABC1234"})
			},
			Path: "/vehicles/base/000/dados",
		},
		{
			Name: "fipe.ConsultarMarcas",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Fipe.ConsultarMarcas(ctx, apibrasil.Json{"codigoTabelaReferencia": 300})
			},
			Path: "/fipe/ConsultarMarcas",
		},
		{
			Name: "correios.Rastreio",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Correios.Rastreio(ctx, apibrasil.Json{"code": "AA123456789BR"})
			},
			Path: "/correios/rastreio",
		},
		{
			Name: "cep.CEP",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Cep.CEP(ctx, apibrasil.Json{"cep": "01001000"})
			},
			Path: "/cep/cep",
		},
		{
			Name: "cep.CalcularDistancia",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Cep.CalcularDistancia(ctx, apibrasil.Json{"origem": "01001000", "destino": "30110000"})
			},
			Path: "/cep/distancia/calcular",
		},
		{
			Name: "databaseIP.IP",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.DatabaseIP.IP(ctx, apibrasil.Json{"ip": "8.8.8.8"})
			},
			Path: "/database/ip",
		},
		{
			Name: "geolocation.Geocode",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Geolocation.Geocode(ctx, apibrasil.Json{"address": "Av Paulista, 1000"})
			},
			Path: "/geolocation/geocode",
		},
		{
			Name: "geomatrix.Distance",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Geomatrix.Distance(ctx, apibrasil.Json{"origins": "a", "destinations": "b"})
			},
			Path: "/geomatrix/distance",
		},
		{
			Name: "recognize.Base64",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Recognize.Base64(ctx, apibrasil.Json{"content": "xxx"})
			},
			Path: "/recognize/base64",
		},
		{
			Name: "ddd.Request",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.DDD.Request(ctx, "consulta", apibrasil.Json{"ddd": "11"})
			},
			Path: "/ddd/consulta",
		},
		{
			Name: "holidays.Feriados",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Holidays.Feriados(ctx, apibrasil.Json{"ano": 2026})
			},
			Path: "/holidays/feriados",
		},
		{
			Name: "translate.Identify",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Translate.Identify(ctx, apibrasil.Json{"text": "olá"})
			},
			Path: "/translate/identify",
		},
		{
			Name: "weather.City",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Weather.City(ctx, apibrasil.Json{"city": "São Paulo"})
			},
			Path: "/weather/city",
		},
		{
			Name: "loterias.Latest",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Loterias.Latest(ctx, "megasena", nil)
			},
			Path: "/loterias/megasena/latest",
		},
		{
			Name: "loterias.Resultado",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Loterias.Resultado(ctx, "megasena", 2700, nil)
			},
			Path: "/loterias/megasena/2700",
		},
		{
			Name: "consulta.CPF",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Consulta.CPF(ctx, apibrasil.Json{"cpf": "00000000000"})
			},
			Path: "/consulta/cpf/credits",
			Body: core.Json{"cpf": "00000000000"},
		},
		{
			Name: "consulta.CNPJ com tipo",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Consulta.CNPJ(ctx, apibrasil.Consulta{
					Tipo:   "lista-socios",
					Fields: apibrasil.Json{"cnpj": "00000000000000"},
				}.Json())
			},
			Path: "/consulta/cnpj/credits",
			Body: core.Json{"tipo": "lista-socios", "cnpj": "00000000000000"},
		},
		{
			Name: "consulta.Generic",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Consulta.Generic(ctx, "serasa", apibrasil.Json{"cpf": "0"})
			},
			Path: "/consulta/serasa/credits",
		},
		{
			Name: "consulta.VeiculosBase",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Consulta.VeiculosBase(ctx, "dados", apibrasil.Json{"placa": "ABC1234"})
			},
			Path: "/vehicles/base/000/dados",
		},
		{
			Name: "consulta.ProxySeller",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Consulta.ProxySeller(ctx, nil)
			},
			Path: "/proxy/seller/credits",
		},
		{
			Name: "consulta.Quod",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Consulta.Quod(ctx, apibrasil.Json{"cnpj": "0"})
			},
			Path: "/quod/cnpj/credits",
		},
		{
			Name: "ura.Dialler",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Ura.Dialler(ctx, apibrasil.Json{"number": "5511999999999"})
			},
			Path: "/ura/call/dialler",
		},
		{
			Name: "chipVirtual.Buy",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.ChipVirtual.Buy(ctx, apibrasil.Json{"operator": "claro"})
			},
			Path: "/chip/virtual/buy",
		},
		{
			Name: "bulk.Direct",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Bulk.Direct(ctx, "cpf", apibrasil.Json{"items": []string{"1", "2"}})
			},
			Path: "/bulk/direct/cpf",
		},
		{
			Name: "bulk.Queue",
			Call: func(ctx context.Context, api *apibrasil.Client) (any, error) {
				return api.Bulk.Queue(ctx, "cpf", apibrasil.Json{"items": []string{"1"}})
			},
			Path: "/bulk/queue/cpf",
		},
	})
}

func TestConsultaCreditsUsaGet(t *testing.T) {
	api, transport := BuildAPI()

	if _, err := api.Consulta.Credits(context.Background(), "cpf"); err != nil {
		t.Fatalf("credits falhou: %v", err)
	}

	last := transport.Last(t)
	if last.Method != core.MethodGet {
		t.Errorf("método = %s, esperado GET", last.Method)
	}
	if last.URL != BaseURL+"/consulta/cpf/credits" {
		t.Errorf("url = %s", last.URL)
	}
}
