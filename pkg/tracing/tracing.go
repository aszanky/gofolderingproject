package tracing

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

const (
	samplerType  = "const"
	samplerParam = 1
)

func Init(serviceName string) (opentracing.Tracer, io.Closer, error) {
	cfg, err := config.FromEnv()
	if err != nil {
		log.Error().Msgf("cannot parse Jaeger env vars %v", err)
		return nil, nil, fmt.Errorf("jaeger init error: %v", err)
	}

	cfg.ServiceName = serviceName
	cfg.Sampler.Type = samplerType
	cfg.Sampler.Param = samplerParam

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Error().Msgf("cannot parse Jaeger env vars %v", err)
		return tracer, closer, fmt.Errorf("jaeger init error: %v", err)
	}

	return tracer, closer, nil
}

func GenerateTraceId(span opentracing.Span) string {
	if span != nil {
		sc, ok := span.Context().(jaeger.SpanContext)
		if ok {
			return sc.TraceID().String()
		}
	}
	return ""
}
