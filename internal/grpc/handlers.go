package grpc

import (
	"context"

	"github.com/PepperoniOnzo/foReal-backend/internal/grpc/pb"
)

func (s *webRTCServer) JoinRoom(ctx context.Context, request *pb.JoinRoomRequest) (*pb.JoinRoomResponse, error) {
	return &pb.JoinRoomResponse{
		Message: "message",
	}, nil
}
