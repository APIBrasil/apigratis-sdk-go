package apibrasil_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/generated"
)

// Testes de contrato: batem no gateway real, então só rodam com
// APIBRASIL_CONTRACT=1 (as rotas usadas aqui são públicas).
//
//	APIBRASIL_CONTRACT=1 go test ./apibrasil -run Contrato -v
func skipSemContrato(t *testing.T) {
	t.Helper()
	if os.Getenv("APIBRASIL_CONTRACT") != "1" {
		t.Skip("defina APIBRASIL_CONTRACT=1 para rodar os testes de contrato")
	}
}

func TestContratoCatalogoRespondeDocumentacoes(t *testing.T) {
	skipSemContrato(t)

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	api := apibrasil.New(apibrasil.Config{Timeout: 120 * time.Second})

	response, err := api.Catalog.Documentations(ctx)
	if err != nil {
		t.Fatalf("documentations falhou: %v", err)
	}
	if len(response) == 0 {
		t.Fatal("o catálogo respondeu vazio")
	}
}

func TestContratoCatalogoDeServidores(t *testing.T) {
	skipSemContrato(t)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if _, err := apibrasil.New().Catalog.Servers(ctx); err != nil {
		t.Fatalf("servers falhou: %v", err)
	}
}

func TestCatalogoGeradoEstaPopulado(t *testing.T) {
	if len(generated.WhatsAppActions) == 0 {
		t.Error("WhatsAppActions vazio — rode go generate ./...")
	}
	if len(generated.ConsultaTipos) == 0 {
		t.Error("ConsultaTipos vazio — rode go generate ./...")
	}
	if len(generated.ServiceActions) == 0 {
		t.Error("ServiceActions vazio — rode go generate ./...")
	}

	if !generated.HasAction("whatsapp", "sendText") {
		t.Error("a action sendText deveria estar no catálogo de whatsapp")
	}
	if !generated.HasAction("dados", "lista-socios") {
		t.Error("a action lista-socios deveria estar no catálogo de dados")
	}
	if meta, ok := generated.ConsultaTipo("acerta-essencial"); !ok || meta.Service != "cpf" {
		t.Errorf("metadados de acerta-essencial = %+v", meta)
	}
}
