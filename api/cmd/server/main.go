package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/azeezlala/assessment/api/database"
	"github.com/azeezlala/assessment/api/database/seeder"
	r "github.com/azeezlala/assessment/api/pkg/transport/rest/router"
	"github.com/azeezlala/assessment/shared/queue/asyncmon"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if os.Getenv("ENV") != "PRODUCTION" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("loading env error: %v", err)
		}
	}

	err := database.AutoMigrate()
	if err != nil {
		log.Fatalf("Unable to migrate: %v", err)
		return
	}

	err = seeder.Seed()
	if err != nil {
		log.Fatalf("Unable to seed: %v", err)
		return
	}

	router := gin.Default()
	router.UseRawPath = true
	router.HandleMethodNotAllowed = true
	aqueue := asyncmon.NewManager(os.Getenv("REDIS_URL"))
	r.SetUpRoutes(router, aqueue)
	aqueue.Start()

	if err != nil {
		log.Fatal(err)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("error parsing port, must be numeric: %v", err)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	errChan := make(chan error, 1)

	// http
	go func() {
		fmt.Printf("Listening on port %d\n", port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("Shutdown signal received")
	case err := <-errChan:
		log.Printf("Server error: %v", err)
	}

	// Graceful shutdown
	stop()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

	log.Println("Server exiting...")

}
