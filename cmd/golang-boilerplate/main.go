package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/urfave/cli/v3"
)

func handleFoo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func handleBar(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "Hello, %s", id)
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func RunAPI(ctx context.Context, cmd *cli.Command) error {
	http.HandleFunc("GET /foo", handleFoo)
	http.HandleFunc("GET /bar/{id}", handleBar)
	http.HandleFunc("GET /healthz", handleHealthz)

	srv := &http.Server{
		Addr:         ":" + cmd.String("port"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		log.Printf("Server listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("server error: %w", err)
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		log.Println("Shutting down server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return srv.Shutdown(shutdownCtx)
	}
}

func main() {
	app := &cli.Command{
		Version: "v1.0.0",
		Usage:   "Test",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println(time.Now().String())
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "api",
				Aliases: []string{"a"},
				Usage:   "a cli application for the api crud.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "port",
						Value: "8080",
						Usage: "port d'écoute du serveur",
					},
				},
				Action: RunAPI,
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatalf("An error occurred: %s", err)
	}
}
