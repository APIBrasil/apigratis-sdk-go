package platform

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// Provedores de pagamento suportados pelo gateway.
const (
	ProviderSantander   = "santander"
	ProviderInter       = "inter"
	ProviderMercadoPago = "mercadopago"
	ProviderSicoob      = "sicoob"
)

// PaymentsService são os pagamentos e recargas (PIX, boleto e cartão):
// /recharge, /{provider}/pix/*, /{provider}/boleto/*,
// /mercadopago/card/*.
type PaymentsService struct {
	services.BaseService
}

// NewPaymentsService cria o serviço de pagamentos.
func NewPaymentsService(client *core.HTTPClient) *PaymentsService {
	return &PaymentsService{BaseService: services.NewBaseService(client)}
}

// Recharges lista as recargas: GET /recharges.
func (s *PaymentsService) Recharges(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "recharges", opts...)
}

// Recharge cria uma recarga (pix|boleto): POST /recharge.
func (s *PaymentsService) Recharge(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "recharge", body, opts...)
}

// RechargeShow detalha uma recarga: GET /recharge/{identifier}.
func (s *PaymentsService) RechargeShow(ctx context.Context, identifier string, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "recharge/"+identifier, opts...)
}

// PixGenerate gera uma cobrança PIX: POST /{provider}/pix/generate.
func (s *PaymentsService) PixGenerate(ctx context.Context, provider string, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, provider+"/pix/generate", body, opts...)
}

// PixStatus consulta uma cobrança PIX: GET /{provider}/pix/{txID}.
func (s *PaymentsService) PixStatus(ctx context.Context, provider, txID string, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, provider+"/pix/"+txID, opts...)
}

// BoletoGenerate gera um boleto: POST /{provider}/boleto/generate.
func (s *PaymentsService) BoletoGenerate(ctx context.Context, provider string, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, provider+"/boleto/generate", body, opts...)
}

// BoletoStatus consulta um boleto: GET /{provider}/boleto/{id}.
func (s *PaymentsService) BoletoStatus(ctx context.Context, provider, id string, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, provider+"/boleto/"+id, opts...)
}

// BoletoPDF baixa o PDF de um boleto:
// GET /{provider}/boleto/{id}/pdf.
func (s *PaymentsService) BoletoPDF(ctx context.Context, provider, id string, opts ...core.RequestOptions) ([]byte, error) {
	return s.Download(ctx, provider+"/boleto/"+id+"/pdf", opts...)
}

// CardProcess processa um pagamento com cartão (Mercado Pago):
// POST /mercadopago/card/process.
func (s *PaymentsService) CardProcess(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "mercadopago/card/process", body, opts...)
}

// CardInstallments consulta as parcelas do cartão:
// POST /mercadopago/card/installments.
func (s *PaymentsService) CardInstallments(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "mercadopago/card/installments", body, opts...)
}

// CardStatus consulta um pagamento de cartão:
// GET /mercadopago/card/{id}.
func (s *PaymentsService) CardStatus(ctx context.Context, id string, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "mercadopago/card/"+id, opts...)
}

// CheckoutPaymentMethods lista os métodos de pagamento do checkout:
// GET /checkout/payment-methods.
func (s *PaymentsService) CheckoutPaymentMethods(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "checkout/payment-methods", opts...)
}

// CheckoutPeriods lista os períodos do checkout: GET /checkout/periods.
func (s *PaymentsService) CheckoutPeriods(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "checkout/periods", opts...)
}

// ValidateCoupon valida um cupom: POST /checkout/validate-coupon.
func (s *PaymentsService) ValidateCoupon(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "checkout/validate-coupon", body, opts...)
}

// CheckoutFinalize finaliza o checkout: POST /checkout/finalize.
func (s *PaymentsService) CheckoutFinalize(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "checkout/finalize", body, opts...)
}
