package main

import (
	"context"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/h-celel/scaling-spoon/internal/config"

	"github.com/h-celel/scaling-spoon/proto/examples"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCService struct {
	examples.UnimplementedServiceServer
}

func (g GRPCService) Hello(_ context.Context, r *examples.HelloRequest) (*examples.HelloResponse, error) {
	if message := r.GetMessage(); message != nil {
		log.Println(message.GetValue())
	} else {
		log.Println("Nil")
	}

	return &examples.HelloResponse{Message: r.GetMessage()}, nil
}

func (g GRPCService) StreamingHello(_ *examples.Empty, client examples.Service_StreamingHelloServer) error {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-client.Context().Done():
			return nil
		case <-ticker.C:
			err := client.Send(&examples.HelloResponse{Message: &examples.Message{Value: "timer goes tick"}})
			if err != nil {
				return err
			}
		}
	}
}

func (g GRPCService) BidiStream(stream examples.Service_BidiStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		if message := req.GetMessage(); message != nil {
			log.Println(message.GetValue())
		} else {
			log.Println("Nil")
		}

		err = stream.Send(&examples.HelloResponse{Message: req.GetMessage()})
		if err != nil {
			return err
		}
	}
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
		service := &GRPCService{}
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
