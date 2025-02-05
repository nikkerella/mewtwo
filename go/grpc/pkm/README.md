# pkm

A gRPC Demo.
- The client inputs the ID of the Pokémon.
- The server returns the name of the Pokémon.

``` sh
protoc --go_out=. --go-grpc_out=. pkm.proto
```