package service

import (
	"context"

	"github.com/rbartolome/chatrooms/internal/entity"

	"github.com/go-kit/kit/endpoint"
)

type MessageService interface {
	CreateMessage(context.Context, *entity.Request) (interface{}, error)
	GetMessage(context.Context, *entity.Request) (interface{}, error)
}

func MakeMessageService(endp endpoint.Endpoint) MessageService {
	return &service{
		client: endp,
	}
}

func (s *service) CreateMessage(ctx context.Context, request *entity.Request) (interface{}, error) {
	return s.client(ctx, nil)
}

func (s *service) GetMessage(ctx context.Context, request *entity.Request) (interface{}, error) {
	return s.client(ctx, request)

}
