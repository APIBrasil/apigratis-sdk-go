package core

import "os"

// Variáveis de ambiente reconhecidas pela SDK.
const (
	EnvBearerToken = "APIBRASIL_BEARER_TOKEN"
	EnvDeviceToken = "APIBRASIL_DEVICE_TOKEN"
	EnvSecretKey   = "APIBRASIL_SECRET_KEY"
	EnvBaseURL     = "APIBRASIL_BASE_URL"
)

// EnvVars lista todas as variáveis de ambiente lidas pela SDK.
var EnvVars = []string{EnvBearerToken, EnvDeviceToken, EnvSecretKey, EnvBaseURL}

// ConfigFromEnv lê a configuração das variáveis de ambiente (quando
// disponíveis). Valores passados explicitamente na Config sempre têm
// prioridade — veja Config.Merge.
func ConfigFromEnv() Config {
	return Config{
		BearerToken: os.Getenv(EnvBearerToken),
		DeviceToken: os.Getenv(EnvDeviceToken),
		SecretKey:   os.Getenv(EnvSecretKey),
		BaseURL:     os.Getenv(EnvBaseURL),
	}
}
