package service

import (
	"context"
	"github.com/rbartolome/chatrooms/internal/entity"

	"github.com/go-kit/kit/endpoint"
)

type Service interface {
	GetGeneral(context.Context, *entity.Request) (interface{}, error)
	GetCharacter(context.Context, *entity.Request) (interface{}, error)
}

type service struct {
	client endpoint.Endpoint
}

func MakeService(endp endpoint.Endpoint) Service {
	return &service{
		client: endp,
	}
}

func (s *service) GetGeneral(ctx context.Context, request *entity.Request) (interface{}, error) {
	return s.client(ctx, nil)
}

func (s *service) GetCharacter(ctx context.Context, request *entity.Request) (interface{}, error) {
	return s.client(ctx, request)

}
