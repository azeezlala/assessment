package main

import (
	"context"
	"fmt"
	"github.com/azeezlala/assessment/notification/pkg/transport/grpc/handler"
	"github.com/azeezlala/assessment/shared/queue/asyncmon"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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

	grpcServer := grpc.NewServer()
	grpc.NewServer()

	aqueue := asyncmon.NewManager(os.Getenv("REDIS_URL"))
	handler.NewNotificationServer(grpcServer, aqueue)
	aqueue.Start()

	grpcPort, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("error parsing port, must be numeric: %v", err)
	}
	// Error channel for goroutines.
	errChan := make(chan error, 1)

	go func() {
		log.Printf("Starting gRPC server on port %d", grpcPort)
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
		if err != nil {
			errChan <- err
			return
		}
		if err := grpcServer.Serve(listener); err != nil {
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("Received shutdown signal")
	case err := <-errChan:
		log.Printf("Error occurred: %v", err)
	}

	log.Println("Shutting down gracefully...")
	grpcServer.GracefulStop()
	log.Println("Server exiting...")
}
