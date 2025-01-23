# order

A gRPC Demo.

Components:
- OrderService: The main service that handles order requests from clients.
- StockService: The service responsible for checking and updating the stock.
- Client: The client that will send the order request to the OrderService.

Flow:
1. Client sends an order request to the OrderService.
2. OrderService communicates with the StockService to check if enough stock is available.
  - If the stock is insufficient, the order is canceled.
  - If enough stock is available, OrderService updates the stock by sending a request to StockService to reduce the stock by 1.
3. OrderService then returns the order status to the client along with the updated stock level.

``` sh
protoc --go_out=. --go-grpc_out=. order.proto
protoc --go_out=. --go-grpc_out=. stock.proto
```