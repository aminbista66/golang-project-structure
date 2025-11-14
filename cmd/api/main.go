package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"myapp/internal/app"
	"myapp/internal/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.Load()
	container, err := app.Initialize(cfg)

	if err != nil {
		log.Fatalf("failed to initialize container: %v", err)
	}

	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      container.Router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownErrCh := make(chan error, 1)
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// block until a signal is received
		s := <-quit
		container.Logger.PrintInfo(fmt.Sprintf("signal received; shutting down server: %s", s.String()), map[string]string{"signal": s.String()})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		shutdownErrCh <- srv.Shutdown(ctx)
	}()

	container.Logger.PrintInfo(fmt.Sprintf("starting server on port %s", cfg.Server.Port), map[string]string{"port": cfg.Server.Port})
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		container.Logger.PrintFatal("could not start server", map[string]string{"error": err.Error()})
	}

	if err := <-shutdownErrCh; err != nil {
		container.Logger.PrintFatal("shutdown error", map[string]string{"error": err.Error()})
	}

	container.Logger.PrintInfo(fmt.Sprintf("Server shutdown successfully on %s", srv.Addr), map[string]string{"address": srv.Addr})
}
