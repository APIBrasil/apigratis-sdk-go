package data

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// GeolocationService é a geolocalização (device-based):
// /geolocation/{action}.
type GeolocationService struct {
	*services.DeviceProxyService
}

// NewGeolocationService cria o serviço de geolocalização.
func NewGeolocationService(client *core.HTTPClient) *GeolocationService {
	return &GeolocationService{DeviceProxyService: services.NewDeviceProxyService(client, "geolocation")}
}

// Geocode converte endereço em coordenadas: POST /geolocation/geocode.
func (s *GeolocationService) Geocode(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "geocode", body, opts...)
}

// ForwardGeocoding executa o geocoding direto:
// POST /geolocation/forward-geocoding.
func (s *GeolocationService) ForwardGeocoding(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "forward-geocoding", body, opts...)
}

// GeomatrixService é a matriz de distâncias (device-based):
// /geomatrix/{action}.
type GeomatrixService struct {
	*services.DeviceProxyService
}

// NewGeomatrixService cria o serviço de matriz de distâncias.
func NewGeomatrixService(client *core.HTTPClient) *GeomatrixService {
	return &GeomatrixService{DeviceProxyService: services.NewDeviceProxyService(client, "geomatrix")}
}

// Distance calcula a distância entre pontos: POST /geomatrix/distance.
func (s *GeomatrixService) Distance(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "distance", body, opts...)
}

// RecognizeService é o OCR / Google Vision (device-based):
// /recognize/{action}.
type RecognizeService struct {
	*services.DeviceProxyService
}

// NewRecognizeService cria o serviço de OCR.
func NewRecognizeService(client *core.HTTPClient) *RecognizeService {
	return &RecognizeService{DeviceProxyService: services.NewDeviceProxyService(client, "recognize")}
}

// Base64 reconhece o conteúdo de uma imagem em base64:
// POST /recognize/base64.
func (s *RecognizeService) Base64(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "base64", body, opts...)
}

// URI reconhece o conteúdo de uma imagem por URL: POST /recognize/uri.
func (s *RecognizeService) URI(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "uri", body, opts...)
}

// DddService é a consulta de DDD (device-based): /ddd/{action}.
type DddService struct {
	*services.DeviceProxyService
}

// NewDddService cria o serviço de DDD.
func NewDddService(client *core.HTTPClient) *DddService {
	return &DddService{DeviceProxyService: services.NewDeviceProxyService(client, "ddd")}
}

// HolidaysService são os feriados (device-based): /holidays/{action}.
type HolidaysService struct {
	*services.DeviceProxyService
}

// NewHolidaysService cria o serviço de feriados.
func NewHolidaysService(client *core.HTTPClient) *HolidaysService {
	return &HolidaysService{DeviceProxyService: services.NewDeviceProxyService(client, "holidays")}
}

// Feriados lista os feriados: POST /holidays/feriados.
func (s *HolidaysService) Feriados(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "feriados", body, opts...)
}

// TranslateService é a tradução (device-based): /translate/{action}.
type TranslateService struct {
	*services.DeviceProxyService
}

// NewTranslateService cria o serviço de tradução.
func NewTranslateService(client *core.HTTPClient) *TranslateService {
	return &TranslateService{DeviceProxyService: services.NewDeviceProxyService(client, "translate")}
}

// Identify identifica o idioma de um texto: POST /translate/identify.
func (s *TranslateService) Identify(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "identify", body, opts...)
}

// Models lista os modelos de tradução: POST /translate/models.
func (s *TranslateService) Models(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "models", body, opts...)
}

// WeatherService é o clima (device-based): /weather/{action}.
type WeatherService struct {
	*services.DeviceProxyService
}

// NewWeatherService cria o serviço de clima.
func NewWeatherService(client *core.HTTPClient) *WeatherService {
	return &WeatherService{DeviceProxyService: services.NewDeviceProxyService(client, "weather")}
}

// City consulta o clima por cidade: POST /weather/city.
func (s *WeatherService) City(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "city", body, opts...)
}

// Coordenates consulta o clima por coordenadas:
// POST /weather/coordenates.
func (s *WeatherService) Coordenates(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Request(ctx, "coordenates", body, opts...)
}
