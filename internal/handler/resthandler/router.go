package resthandler

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/trace"
)

func SetupRouter(
	handler *handler,
	logger *slog.Logger,
) *gin.Engine {
	r := gin.Default()

	// Middleware OpenTelemetry
	r.Use(otelgin.Middleware("go-foldering-project-service"))

	r.Use(func(c *gin.Context) {
		start := time.Now()
		requestID := uuid.New().String()

		span := trace.SpanFromContext(c.Request.Context())

		requestLogger := logger.With(
			"request_id", requestID,
			"trace_id", span.SpanContext().TraceID().String(), // <-- Ambil Trace ID
		)

		c.Set("logger", requestLogger)

		c.Next()

		latency := time.Since(start)
		requestLogger.Info("request completed",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency", latency.String(),
		)
	})

	v1 := r.Group("/v1")
	{

		// Login User
		// v1.POST("/auth/login", handler.Login) // Endpoint Login
		v1.GET("/users/:id", handler.GetUsers)

		// v1.POST("/auth/refresh", handler.Refresh)

		// protected routes if you use RBAC to access specific endpoint
		// protected := v1.Group("/").Use(AuthMiddleware(cfg.JWTSecret))
		// {
		// 	protected.POST("/users/:userId/roles", handler.AssignRoleToUser)
		// }
	}

	return r
}
