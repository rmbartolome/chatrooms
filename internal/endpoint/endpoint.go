package endpoint

import (
	"context"
	"github.com/rbartolome/chatrooms/internal/entity"
	"github.com/rbartolome/chatrooms/internal/service"

	"github.com/go-kit/kit/endpoint"
)

//MakeServiceEndpoint crea el endpoint para la response general
func MakeServiceEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.GetGeneral(ctx, nil)
	}
}

//MakeServiceEndpoint crea el endpoint para un personaje
func MakeServiceCharacterEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		request := in.(*entity.Request)
		return svc.GetCharacter(ctx, request)
	}
}
