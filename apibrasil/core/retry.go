package core

import (
	"context"
	"math"
	"math/rand"
	"time"
)

// DefaultRetry devolve a política de retry padrão: 2 novas tentativas,
// backoff de 300ms a 5s, apenas em HTTP 429 e falhas de conexão.
func DefaultRetry() RetryConfig {
	return RetryConfig{
		Retries:         2,
		MinDelay:        300 * time.Millisecond,
		MaxDelay:        5 * time.Second,
		RetryOnStatuses: []int{429},
	}
}

// NoRetry devolve uma política que desativa o retry — passe em Config.Retry.
func NoRetry() *RetryConfig {
	return &RetryConfig{Retries: 0}
}

// ResolveRetry completa a política informada com os padrões. nil usa
// DefaultRetry().
func ResolveRetry(config *RetryConfig) RetryConfig {
	resolved := DefaultRetry()
	if config == nil {
		return resolved
	}

	resolved.Retries = config.Retries
	if config.MinDelay > 0 {
		resolved.MinDelay = config.MinDelay
	}
	if config.MaxDelay > 0 {
		resolved.MaxDelay = config.MaxDelay
	}
	if config.RetryOnStatuses != nil {
		resolved.RetryOnStatuses = config.RetryOnStatuses
	}
	return resolved
}

// RetriesStatus informa se um status HTTP dispara retry nesta política.
func (r RetryConfig) RetriesStatus(status int) bool {
	for _, candidate := range r.RetryOnStatuses {
		if candidate == status {
			return true
		}
	}
	return false
}

// BackoffDelay calcula o backoff exponencial com jitter:
// MinDelay * 2^attempt, limitado a MaxDelay.
func BackoffDelay(attempt int, retry RetryConfig) time.Duration {
	exponential := float64(retry.MinDelay) * math.Pow(2, float64(attempt))
	jitter := 0.5 + rand.Float64()*0.5 //nolint:gosec // jitter não precisa de CSPRNG
	delay := time.Duration(exponential * jitter)
	if delay > retry.MaxDelay {
		return retry.MaxDelay
	}
	return delay
}

// Sleep aguarda delay respeitando o cancelamento do contexto.
func Sleep(ctx context.Context, delay time.Duration) error {
	if delay <= 0 {
		return ctx.Err()
	}

	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
