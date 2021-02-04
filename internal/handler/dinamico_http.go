package handler

import (
	"context"
	"github.com/rbartolome/chatrooms/internal/entity"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	commons "gitlab.falabella.com/fif/integracion/forthehorde/commons/go-microservices-commons"
)

//NewHTTPHandler crea un nuevo http handler para un endpoint de service
func NewHTTPHandlerCharacter(logger log.Logger, serviceEndpoint endpoint.Endpoint, tracer opentracing.Tracer, metricsConf *commons.MetricsConfig) http.Handler {

	endpointsCfg := []commons.EndpointConfig{
		commons.GET("/character/{id}", "response_character", serviceEndpoint, DecodeRequestCharacters, nil),
	}

	builder := commons.MakeHTTPHandlerBuilder(logger, endpointsCfg).
		WithTracer(tracer).
		WithMetrics(metricsConf)

	return builder.Build()
}

//DecodeRequestCharacters obtiene el id del personaje que va a solicitar
func DecodeRequestCharacters(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	req := &entity.Request{
		Id_character: id,
	}

	return req, nil
}
