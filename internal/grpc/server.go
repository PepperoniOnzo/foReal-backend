package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/PepperoniOnzo/foReal-backend/internal/grpc/pb"
	"google.golang.org/grpc"
)

type webRTCServer struct {
	pb.UnimplementedWebRTCServiceServer
}

func RunServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpServer := grpc.NewServer()
	pb.RegisterWebRTCServiceServer(grpServer, &webRTCServer{})

	log.Println("Starting gRPC server on port 8080")
	if err := grpServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
