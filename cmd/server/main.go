package main

import (
	"os"
	"time"

	"github.com/go-kit/kit/log/level"

	jaegercfg "github.com/uber/jaeger-client-go/config"

	"github.com/opentracing/opentracing-go"
	commons "gitlab.falabella.com/fif/integracion/forthehorde/commons/go-microservices-commons"

	"github.com/rbartolome/chatrooms/cmd/config"
	"github.com/rbartolome/chatrooms/internal/client"
	"github.com/rbartolome/chatrooms/internal/endpoint"
	"github.com/rbartolome/chatrooms/internal/handler"
	"github.com/rbartolome/chatrooms/internal/service"
)

func main() {
	cfg := config.ReadConfiguration()

	port, _ := cfg["port"]
	loggingLevel, _ := cfg["logging_level"]
	tracingEnabled, _ := cfg["tracing_enabled"]
	metricsEnabled, _ := cfg["metrics_enabled"]

	url := cfg["url"].(string)
	timeout := time.Duration(cfg["timeout"].(int)) * time.Second

	logger := commons.ConfigureLogger(loggingLevel.(string))

	var metricsConf *commons.MetricsConfig
	if metricsEnabled.(bool) == true {
		metricsConf = commons.MakeDefaultEndpointMetrics("restructuraciones", "integracion")
	}

	var tracer opentracing.Tracer
	//Override de JAEGER_DISABLED en caso de tracing no habilitado
	if tracingEnabled.(bool) == true {
		// Instanciar tracer global Jaeger
		traceCfg, err := jaegercfg.FromEnv()
		if err != nil {
			level.Error(logger).Log("No se pudo parsear configuracion de Jaeger", err.Error())
			return
		}

		tracer, closer, err := traceCfg.NewTracer()
		defer closer.Close()
		if err != nil {
			level.Error(logger).Log("No se pudo inicializar Tracer Jaeger", err.Error())
			return
		}

		opentracing.SetGlobalTracer(tracer)
	} else {
		os.Setenv("JAEGER_DISABLED", "true")
	}

	//Implementación de un GET estatico
	endpoint_api := client.NewHTTPClientEndpoint(url, timeout, logger)
	endpoint_api = commons.EndpointLogMiddleware("message_service", "/messages", "GET", logger)(endpoint_api)

	//Implementación de un GET dinámico
	endpoint_character := client.MakeHTTPClientCharacterEndpoint(url, timeout, logger)
	endpoint_character = commons.EndpointLogMiddleware("message_service", "/messages", "POST", logger)(endpoint_character)

	//Llamado a GET estatico
	var (
		svc             = service.MakeService(endpoint_api)
		serviceEndpoint = endpoint.MakeServiceEndpoint(svc)
		httphandler     = handler.NewHTTPHandler(logger, serviceEndpoint, tracer, metricsConf)
	)

	g := commons.CreateServer(httphandler, port.(string), logger)

	logger.Log("exit", g.Run())
}
