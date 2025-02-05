package main

import (
	"context"
	"log"
	"net"
	pb "pkm/protobuf"

	"google.golang.org/grpc"
)

// The server struct
type server struct {
	pb.UnimplementedQueryServiceServer
}

// Implement the Query method defined in the protobuf file.
// This method handles incoming requests and returns the corresponding Pokémon name based on the ID.
func (s *server) Query(ctx context.Context, req *pb.QueryReq) (*pb.QueryResp, error) {
	switch req.GetId() { // Retrieve the ID from the request
	case 1:
		return &pb.QueryResp{Name: "Bulbasaur"}, nil
	case 2:
		return &pb.QueryResp{Name: "Ivysaur"}, nil
	case 3:
		return &pb.QueryResp{Name: "Venusaur"}, nil
	default:
		return &pb.QueryResp{Name: "Unknown"}, nil
	}
}

func main() {
	// Create a TCP listener on port 50051.
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err) // Log and exit if the listener cannot be created
	}

	// Create a new gRPC server instance.
	grpcServer := grpc.NewServer()

	// Register the QueryService server implementation with the gRPC server.
	pb.RegisterQueryServiceServer(grpcServer, &server{})

	log.Println("Server is running on port 50051")

	// Start the gRPC server and listen for client requests.
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
