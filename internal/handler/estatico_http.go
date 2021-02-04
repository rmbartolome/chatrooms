package handler

import (
	"context"
	"net/http"

	"github.com/opentracing/opentracing-go"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	commons "gitlab.falabella.com/fif/integracion/forthehorde/commons/go-microservices-commons"
)

//NewHTTPHandler crea un nuevo http handler para un endpoint de service
func NewHTTPHandler(logger log.Logger, serviceEndpoint endpoint.Endpoint, tracer opentracing.Tracer, metricsConf *commons.MetricsConfig) http.Handler {

	endpointsCfg := []commons.EndpointConfig{
		commons.GET("/", "response_general", serviceEndpoint, DecodeRequest, nil),
	}

	builder := commons.MakeHTTPHandlerBuilder(logger, endpointsCfg).WithTracer(tracer).WithMetrics(metricsConf)

	return builder.Build()
}

// DecodeRequest request decode
func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
	/* var appRequest entity.RequestStruct

	if err := json.NewDecoder(r.Body).Decode(&appRequest); err != nil {
		return nil, err
	}

	return &appRequest, nil */
}
