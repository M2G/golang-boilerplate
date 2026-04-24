package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v3"
)

func Run(context.Context, *cli.Command) error {

	fmt.Println("Hello World")

	http.HandleFunc("/bar/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		_, err := fmt.Fprintf(w, "Item ID: %s", id)
		if err != nil {
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}

func main() {
	cmd := &cli.Command{
		Name:    "boom",
		Version: "v1.0.0",
		Usage:   "Golang init",
		Action:  Run,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

	return
}
