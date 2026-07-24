package data

import (
	"context"
	"strconv"

	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/core"
	"github.com/jhowbhz/apigratis-sdk-go/apibrasil/services"
)

// LoteriasService são as loterias (device-based): /loterias/{sorteio}/*.
type LoteriasService struct {
	services.BaseService
}

// NewLoteriasService cria o serviço de loterias.
func NewLoteriasService(client *core.HTTPClient) *LoteriasService {
	return &LoteriasService{BaseService: services.NewBaseService(client)}
}

// Resultado consulta um concurso:
// POST /loterias/{sorteio}/{concurso}.
func (s *LoteriasService) Resultado(ctx context.Context, sorteio string, concurso int, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	path := "loterias/" + sorteio + "/" + strconv.Itoa(concurso)
	return services.Device(s.Post(ctx, path, body, opts...))
}

// Latest consulta o último resultado: POST /loterias/{sorteio}/latest.
func (s *LoteriasService) Latest(ctx context.Context, sorteio string, body core.Json, opts ...core.RequestOptions) (core.DeviceServiceResponse, error) {
	return services.Device(s.Post(ctx, "loterias/"+sorteio+"/latest", body, opts...))
}
