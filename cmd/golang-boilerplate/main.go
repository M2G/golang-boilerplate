package main

import (
	"context"
	"golang-boilerplate/internal/config"
	"golang-boilerplate/internal/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:    "boom",
		Usage:   "Golang boilerplate",
		Version: "v1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "port",
				Value: "8181",
				Usage: "HTTP server port",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			cfg := config.Load(cmd)

			server.RegisterRoutes()

			srv := &http.Server{
				Addr: ":" + cfg.Port,
			}

			go func() {
				log.Println("Server starting on", srv.Addr)
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatal(err)
				}
			}()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit

			shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			log.Println("Shutting down server...")
			return srv.Shutdown(shutdownCtx)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
