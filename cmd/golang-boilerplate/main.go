package main

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli/v3"
)

func RunAPI(ctx context.Context, cmd *cli.Command) error {

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	http.HandleFunc("/bar/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Server listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
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
				Action:  RunAPI,
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatalf("An error occurred: %s", err)
	}
}