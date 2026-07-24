// Package platform traz os serviços da plataforma: autenticação,
// devices, conta, pagamentos, catálogo, segurança e relatórios.
package platform

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// AuthService é a autenticação e conta (/auth/*, /profile*, /password/*).
//
// Login e Verify2FA guardam automaticamente o Bearer Token retornado no
// cliente, deixando as próximas chamadas autenticadas.
type AuthService struct {
	services.BaseService
}

// NewAuthService cria o serviço de autenticação.
func NewAuthService(client *core.HTTPClient) *AuthService {
	return &AuthService{BaseService: services.NewBaseService(client)}
}

// Login autentica com email/senha: POST /auth/login.
//
// body aceita email, password e turnstile_token. Se a conta tiver 2FA, a
// resposta traz requires_2fa e challenge — use Send2FA + Verify2FA para
// concluir.
func (s *AuthService) Login(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	response, err := s.Post(ctx, "auth/login", body, opts...)
	if err != nil {
		return nil, err
	}
	s.storeToken(response)
	return response, nil
}

// Send2FA envia o código 2FA pelo método escolhido (email, sms,
// whatsapp, call): POST /auth/2fa/send.
func (s *AuthService) Send2FA(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "auth/2fa/send", body, opts...)
}

// Verify2FA conclui o login com o código 2FA:
// POST /auth/login/verify-2fa.
func (s *AuthService) Verify2FA(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	response, err := s.Post(ctx, "auth/login/verify-2fa", body, opts...)
	if err != nil {
		return nil, err
	}
	s.storeToken(response)
	return response, nil
}

// TwoFactorMethods lista os métodos 2FA ativos da conta:
// GET /auth/2fa/methods.
func (s *AuthService) TwoFactorMethods(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "auth/2fa/methods", opts...)
}

// Register cria uma conta: POST /auth/register.
func (s *AuthService) Register(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "auth/register", body, opts...)
}

// RegisterSimple executa o cadastro simplificado:
// POST /auth/register/simple.
func (s *AuthService) RegisterSimple(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "auth/register/simple", body, opts...)
}

// VerificationSend dispara a verificação de email/celular:
// POST /auth/verification/send.
func (s *AuthService) VerificationSend(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "auth/verification/send", body, opts...)
}

// VerificationVerify confirma o código de verificação:
// POST /auth/verification/verify.
func (s *AuthService) VerificationVerify(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "auth/verification/verify", body, opts...)
}

// PasswordForgot inicia a recuperação de senha:
// POST /auth/password/forgot.
func (s *AuthService) PasswordForgot(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "auth/password/forgot", body, opts...)
}

// PasswordVerifyCode valida o código de recuperação:
// POST /auth/password/verify-code.
func (s *AuthService) PasswordVerifyCode(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "auth/password/verify-code", body, opts...)
}

// PasswordReset redefine a senha: POST /auth/password/reset.
func (s *AuthService) PasswordReset(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "auth/password/reset", body, opts...)
}

// PasswordResend reenvia o código de recuperação:
// POST /auth/password/resend.
func (s *AuthService) PasswordResend(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "auth/password/resend", body, opts...)
}

// ChangePassword troca a senha com a sessão ativa:
// POST /password/change.
func (s *AuthService) ChangePassword(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "password/change", body, opts...)
}

// Profile devolve o perfil completo (com estatísticas): POST /profile.
func (s *AuthService) Profile(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "profile", nil, opts...)
}

// Me devolve o perfil atual: GET /profile/me.
func (s *AuthService) Me(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "profile/me", opts...)
}

// UpdateMe atualiza o perfil: PUT /profile/me.
func (s *AuthService) UpdateMe(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.Json, error) {
	return s.Put(ctx, "profile/me", body, opts...)
}

// Verify valida o token atual: GET /auth/verify.
func (s *AuthService) Verify(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Get(ctx, "auth/verify", opts...)
}

// Refresh renova o JWT: POST /refresh.
func (s *AuthService) Refresh(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	response, err := s.Post(ctx, "refresh", nil, opts...)
	if err != nil {
		return nil, err
	}
	s.storeToken(response)
	return response, nil
}

// TokenRotate rotaciona o token: POST /auth/token/rotate.
func (s *AuthService) TokenRotate(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "auth/token/rotate", nil, opts...)
}

// TokenRevoke revoga o token atual: POST /auth/token/revoke.
func (s *AuthService) TokenRevoke(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	return s.Post(ctx, "auth/token/revoke", nil, opts...)
}

// Logout encerra a sessão e limpa o Bearer Token do cliente:
// POST /auth/logout.
func (s *AuthService) Logout(ctx context.Context, opts ...core.RequestOptions) (core.Json, error) {
	response, err := s.Post(ctx, "auth/logout", nil, opts...)
	s.HTTP().SetBearerToken("")
	if err != nil {
		return nil, err
	}
	return response, nil
}

// storeToken guarda no cliente o token devolvido pela plataforma
// (authorization.token ou token, conforme a rota).
func (s *AuthService) storeToken(response core.Json) {
	if token := ExtractToken(response); token != "" {
		s.HTTP().SetBearerToken(token)
	}
}

// ExtractToken lê o Bearer Token de uma resposta de autenticação
// (authorization.token ou token).
func ExtractToken(response core.Json) string {
	if response == nil {
		return ""
	}
	if authorization, ok := response["authorization"].(map[string]any); ok {
		if token, ok := authorization["token"].(string); ok && token != "" {
			return token
		}
	}
	if token, ok := response["token"].(string); ok {
		return token
	}
	return ""
}

// Requires2FA informa se a resposta de login exige segundo fator.
func Requires2FA(response core.Json) bool {
	return response != nil && response["requires_2fa"] == true
}
