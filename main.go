package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/cors"
	"github.com/xpositivityx/obs-test/pkg/db"
	"github.com/xpositivityx/obs-test/pkg/tracing"
)

func main() {
	err := tracing.InitTracer()

	if err != nil {
		fmt.Printf("Error initializing tracer: %v\n", err)
		return
	}

	tp := tracing.GetTracer()

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			fmt.Printf("Error shutting down tracer provider: %v\n", err)
		}
	}()

	err = db.Init()
	if err != nil {
		fmt.Printf("Error initializing database: %v\n", err)
		return
	}

	defer db.Pool.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	// Don't do this in real life
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	})

	fmt.Println("Server starting on http://0.0.0.0:8000")

	if err := http.ListenAndServe("0.0.0.0:8000", c.Handler(mux)); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
