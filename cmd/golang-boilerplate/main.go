package main

import (
	"fmt"
	"context"
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

	log.Info("Server listen at :8181")

	fmt.Println("Started server on - 127.0.0.1:8080")

	err = router.Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	app := &cli.Command{
		Version: "v1.0.0",
		Usage:   "Test",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Printf(time.Now().String())
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
		return
	}
}