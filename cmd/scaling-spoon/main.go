package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/h-celel/scaling-spoon/internal/config"

	"github.com/h-celel/scaling-spoon/proto/examples"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UserGRPCService struct {
	examples.UnimplementedServiceServer
}

func (u UserGRPCService) Hello(_ context.Context, r *examples.HelloRequest) (*examples.HelloResponse, error) {
	if message := r.GetMessage(); message != nil {
		log.Println(message.GetValue())
	} else {
		log.Println("Nil")
	}

	return &examples.HelloResponse{Message: r.GetMessage()}, nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	catchShutdown(cancel)

	env := config.NewEnvironment()

	lis, err := net.Listen("tcp", env.GRPCHost)
	if err != nil {
		log.Panicf("error with grpc api: %v", err)
	}

	go func() {
		defer cancel()
		server := grpc.NewServer()
		service := &UserGRPCService{}
		examples.RegisterServiceServer(server, service)
		reflection.Register(server)
		err := server.Serve(lis)
		if err != nil {
			log.Printf("error with grpc api: %v", err)
		}
	}()

	log.Println("running...")
	// awaits service context cancellation before closing the service
	<-ctx.Done()
}

// catchShutdown listens for SIGTERM and SIGINT, and cancels the service context when either signal is received.
func catchShutdown(cancel context.CancelFunc) {
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-cancelChan
		log.Printf("Caught SIGTERM %v", sig)
		cancel()
	}()
}
