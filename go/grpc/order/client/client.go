package main

import (
	"context"
	"fmt"
	"log"
	pb "order/proto/order"

	"google.golang.org/grpc"
)

func main() {
	// Connect to the OrderService gRPC server
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	// Make an order request
	resp, err := client.PlaceOrder(context.Background(), &pb.PlaceOrderRequest{
		ProductId: "product",
		Quantity:  1,
	})
	if err != nil {
		log.Fatalf("could not place order: %v", err)
	}

	fmt.Printf("Order Status: %s\n", resp.Status)
	fmt.Printf("Remaining Stock: %d\n", resp.RemainingStock)
}
