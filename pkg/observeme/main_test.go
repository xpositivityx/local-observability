package observeme_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/xpositivityx/local-observability/pkg/observeme"
	"github.com/xpositivityx/local-observability/pkg/tracing"
)

func TestMain(m *testing.M) {
	err := tracing.InitTracer()

	if err != nil {
		fmt.Println("failed to initialize tracer")
		os.Exit(1)
	}

	tp := tracing.GetTracer()

	code := m.Run()

	if err := tp.Shutdown(context.Background()); err != nil {
		fmt.Printf("Error shutting down tracer provider: %v\n", err)
	}

	os.Exit(code)
}

func TestObserveme(t *testing.T) {
	got := observeme.Respond(context.Background(), true)
	want := "indeed"
	if got != want {
		t.Errorf("expected %s, want %s", got, want)
	}
}
