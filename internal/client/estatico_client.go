package client

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rbartolome/chatrooms/internal/entity"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	commons "gitlab.falabella.com/fif/integracion/forthehorde/commons/go-microservices-commons"
)

func NewHTTPClientEndpoint(url string, timeout time.Duration, logger log.Logger) endpoint.Endpoint {
	return commons.MakeHTTPClientBuilder("GET",
		url,
		timeout,
		commons.DefaultRequestEncode,
		DecodeResponse,
		logger).Build()
}

func DecodeResponse(_ context.Context, r *http.Response) (interface{}, error) {
	response := new(entity.ResponseGeneral)

	err := json.NewDecoder(r.Body).Decode(response)
	if err != nil {
		return nil, errors.New("Error en DecodeResponse de client")
	}

	return response, nil
}
