package data

import (
	"context"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// DatabaseIPService é o GeoIP (device-based): POST /database/ip.
type DatabaseIPService struct {
	services.BaseService
}

// NewDatabaseIPService cria o serviço de GeoIP.
func NewDatabaseIPService(client *core.HTTPClient) *DatabaseIPService {
	return &DatabaseIPService{BaseService: services.NewBaseService(client)}
}

// IP consulta a base de IPs: POST /database/ip body {"ip": "..."}.
func (s *DatabaseIPService) IP(ctx context.Context, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return services.Device(s.Post(ctx, "database/ip", body, opts...))
}
