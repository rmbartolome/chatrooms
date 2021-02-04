package endpoint

import (
	"context"

	"github.com/rbartolome/chatrooms/internal/entity"
	"github.com/rbartolome/chatrooms/internal/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeCreateMessageEndpoint(svc service.MessageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.CreateMessage(ctx, nil)
	}
}

func MakeGetServiceEndpoint(svc service.MessageService) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		request := in.(*entity.Request)
		return svc.GetMessage(ctx, request)
	}
}
