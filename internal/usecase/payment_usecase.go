package usecase

import (
	"context"

	"go.opentelemetry.io/otel"
)

func (p *usecase) IntegrateWithMandiri(ctx context.Context) (err error) {
	tracer := otel.Tracer("usecase-layer")
	ctx, span := tracer.Start(ctx, "usecase.IntegrateWithMandiri")
	defer span.End()

	return err
}
