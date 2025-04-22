package observeme

import (
	"context"
	"errors"
	"math/rand"
	"strconv"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func Respond(ctx context.Context, b bool) string {
	tracer := otel.Tracer("obs")
	ctx, span := tracer.Start(ctx, "observeme.Respond")

	defer span.End()

	err := someInternalFunc(ctx)

	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return ""
	}

	if b {
		return "indeed"
	}
	return "not at all"
}

func someInternalFunc(ctx context.Context) error {
	tracer := otel.Tracer("obs")
	_, span := tracer.Start(ctx, "observeme.someInternalFunc")

	defer span.End()

	val := rand.Intn(10)
	span.SetAttributes(attribute.String("randomValue", strconv.Itoa(val)))
	if val > 5 {
		return errors.New("oh no")
	}

	return nil
}
