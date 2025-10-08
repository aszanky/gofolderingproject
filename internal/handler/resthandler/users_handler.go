package resthandler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func (h *handler) GetUsers(c *gin.Context) {
	logger, _ := c.Get("logger")
	requestLogger := logger.(*slog.Logger)

	// Dapatkan tracer dari OTel
	tracer := otel.Tracer("handler-layer")
	ctx, span := tracer.Start(c.Request.Context(), "handler.GetUsers")
	defer span.End()

	id := c.Param("id")
	if id == "" {
		requestLogger.Error("id parameter is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is required"})
		return
	}

	user, err := h.usecase.GetUsers(ctx, id)
	if err != nil {
		span.RecordError(err)
		requestLogger.Error("failed to get user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	span.SetAttributes(
		attribute.String("request.user.id", id),
	)

	c.JSON(http.StatusOK, gin.H{"user": user})
}
