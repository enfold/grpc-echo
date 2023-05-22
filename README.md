# gRPC Echo Server

This is a simple gRPC Echo Server implemented in Go. It provides unary and streaming echo services.

## Prerequisites

- Go programming language (version 1.16 or higher)
- Protocol Buffers compiler (`protoc`)
- Go gRPC plugin (`protoc-gen-go-grpc`)

## Installation

1. Clone the repository or download the source code.

2. Install the required dependencies:

```shell
go get google.golang.org/grpc
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## Protocol Buffers Compilation

1. Generate the Go code from the Protocol Buffers definition:

```shell
protoc -I=proto --go_out=. --go-grpc_out=. proto/echo.proto
```

2. This will generate the `echo.pb.go` and `echo_grpc.pb.go` files.

## Usage

1. Run the server:

```shell
go run main.go
```

2. The server will start listening on port 5050.

## API

The server provides the following gRPC API:

### UnaryEcho

UnaryEcho performs a unary RPC call to echo the message received.

```protobuf
rpc UnaryEcho(EchoRequest) returns (EchoResponse) {}
```

### ServerStreamingEcho

ServerStreamingEcho performs a server-streaming RPC call to repeatedly send the same message at intervals.

```protobuf
rpc ServerStreamingEcho(EchoRequest) returns (stream EchoResponse) {}
```

### ClientStreamingEcho

ClientStreamingEcho performs a client-streaming RPC call to receive multiple messages from the client.

```protobuf
rpc ClientStreamingEcho(stream EchoRequest) returns (EchoResponse) {}
```

### BidirectionalStreamingEcho

BidirectionalStreamingEcho performs a bidirectional-streaming RPC call to send and receive multiple messages.

```protobuf
rpc BidirectionalStreamingEcho(stream EchoRequest) returns (stream EchoResponse) {}
```

## Client Example

You can use the generated client code to interact with the server. Here's an example of a client calling the UnaryEcho service:

```go
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/features/proto/echo"
)

func main() {
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.UnaryEcho(ctx, &pb.EchoRequest{Message: "Hello, gRPC!"})
	if err != nil {
		log.Fatalf("failed to call UnaryEcho: %v", err)
	}

	log.Printf("Response: %s", response.Message)
}
```

## License

This project is licensed under the [MIT License](LICENSE).