package messaging

import (
	"context"
	"strings"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// EvolutionService é a Evolution API (device-based):
// POST /evolution/{controller}/{action}. Exige Authorization: Bearer +
// DeviceToken.
//
// Exemplos de controller/action: instance/create, message/sendText.
// Os caminhos documentados estão em generated.EvolutionPaths.
type EvolutionService struct {
	services.BaseService
}

// NewEvolutionService cria o serviço da Evolution API.
func NewEvolutionService(client *core.HTTPClient) *EvolutionService {
	return &EvolutionService{BaseService: services.NewBaseService(client)}
}

// Request executa POST /evolution/{controller}/{action}.
func (s *EvolutionService) Request(ctx context.Context, controller, action string, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Call(ctx, controller+"/"+action, body, opts...)
}

// Call executa um caminho completo do catálogo: POST /evolution/{path}.
func (s *EvolutionService) Call(ctx context.Context, path string, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return services.Device(s.Post(ctx, "evolution/"+strings.Trim(path, "/"), body, opts...))
}

// Queue executa a mesma chamada de forma assíncrona via fila.
func (s *EvolutionService) Queue(ctx context.Context, controller, action string, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return s.Call(ctx, controller+"/"+action+"/queue", body, opts...)
}
