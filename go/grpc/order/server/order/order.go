package main

import (
	"context"
	"fmt"
	"log"
	"net"

	order_pb "order/proto/order"
	stock_pb "order/proto/stock"

	"google.golang.org/grpc"
)

type server struct {
	order_pb.UnimplementedOrderServiceServer
	stockClient stock_pb.StockServiceClient
}

func (s *server) PlaceOrder(ctx context.Context, req *order_pb.PlaceOrderRequest) (*order_pb.PlaceOrderResponse, error) {
	// Check stock availability by calling the StockService
	checkResp, err := s.stockClient.CheckStock(ctx, &stock_pb.CheckStockRequest{ProductId: req.ProductId})
	if err != nil {
		return nil, fmt.Errorf("failed to check stock: %v", err)
	}

	if !checkResp.Available || checkResp.CurrentStock < int32(req.Quantity) {
		fmt.Println("Order cancelled for not enough stock")
		return &order_pb.PlaceOrderResponse{
			Status:         "Order Cancelled - Not enough stock",
			RemainingStock: checkResp.CurrentStock,
		}, nil
	}

	// Deduct stock by calling StockService to update the stock
	_, err = s.stockClient.DeductStock(ctx, &stock_pb.DeductStockRequest{ProductId: req.ProductId})
	if err != nil {
		return nil, fmt.Errorf("failed to deduct stock: %v", err)
	}

	fmt.Println("Order placed successfully")
	return &order_pb.PlaceOrderResponse{
		Status:         "Order Placed Successfully",
		RemainingStock: checkResp.CurrentStock - 1,
	}, nil
}

func main() {
	// Connect to the StockService gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to StockService: %v", err)
	}
	defer conn.Close()

	// Create a StockService client
	stockClient := stock_pb.NewStockServiceClient(conn)

	// Set up the OrderService gRPC server
	s := grpc.NewServer()
	order_pb.RegisterOrderServiceServer(s, &server{stockClient: stockClient})

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Order service running on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
