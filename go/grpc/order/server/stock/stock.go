package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "order/proto/stock"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStockServiceServer
	stockMap map[string]int32
}

func (s *server) CheckStock(ctx context.Context, req *pb.CheckStockRequest) (*pb.CheckStockResponse, error) {
	stock, exists := s.stockMap[req.ProductId]
	if !exists {
		return &pb.CheckStockResponse{Available: false, CurrentStock: 0}, nil
	}
	return &pb.CheckStockResponse{Available: stock > 0, CurrentStock: stock}, nil
}

func (s *server) DeductStock(ctx context.Context, req *pb.DeductStockRequest) (*pb.DeductStockResponse, error) {
	stock, exists := s.stockMap[req.ProductId]
	if !exists || stock <= 0 {
		return &pb.DeductStockResponse{Success: false, UpdatedStock: stock}, nil
	}
	s.stockMap[req.ProductId] = stock - 1
	fmt.Println("stock -1, current", s.stockMap[req.ProductId])
	return &pb.DeductStockResponse{Success: true, UpdatedStock: s.stockMap[req.ProductId]}, nil
}

func main() {
	s := grpc.NewServer()
	pb.RegisterStockServiceServer(s, &server{
		stockMap: map[string]int32{
			"product": 3, // Initially, the stock for product is 3
		},
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Stock service running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
