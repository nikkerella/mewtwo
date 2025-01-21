package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "pkm/protobuf"

	"google.golang.org/grpc"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewQueryServiceClient(conn)

	// Prompt user for input
	fmt.Print("Enter the ID: ")

	// Create a scanner to read input from the command line
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Read the next line of input
	id := scanner.Text()

	// Convert the ID from string to an integer
	var queryId int32
	_, err = fmt.Sscanf(id, "%d", &queryId)
	if err != nil {
		log.Fatalf("Invalid input, please enter a valid integer ID")
	}

	// Call the Query method with the provided ID
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Query(ctx, &pb.QueryReq{Id: queryId})
	if err != nil {
		log.Fatalf("Could not query: %v", err)
	}

	// Display the response in the command prompt
	fmt.Printf("Name: %s\n", resp.Name)
}
