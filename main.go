package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/features/proto/echo"
)

type EchoServer struct {
	pb.UnimplementedEchoServer
}

func (s *EchoServer) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("received: %v", req.GetMessage())
	return &pb.EchoResponse{Message: req.GetMessage()}, nil
}

func (s *EchoServer) ServerStreamingEcho(req *pb.EchoRequest, streaming pb.Echo_ServerStreamingEchoServer) error {
	log.Printf("received streaming request: %v", req.GetMessage())
	resp := &pb.EchoResponse{Message: req.GetMessage()}
	i := 0

	for {
		if err := streaming.Send(resp); err != nil {
			return err
		}
		i++
		log.Printf("send the response [%v]: %v", i, resp.GetMessage())
		time.Sleep(time.Duration(2) * time.Second)
	}
}

func (s *EchoServer) ClientStreamingEcho(streaming pb.Echo_ClientStreamingEchoServer) error {
	for {
		req, err := streaming.Recv()
		if err != nil {
			return nil
		}

		log.Printf("received streaming message: %v", req.GetMessage())
	}
}

func (s *EchoServer) BidirectionalStreamingEcho(streaming pb.Echo_BidirectionalStreamingEchoServer) error {
	for {
		req, err := streaming.Recv()
		if err != nil {
			return nil
		}

		log.Printf("received streaming message: %v", req.Message)
		if err := streaming.SendMsg(&pb.EchoResponse{Message: req.Message}); err != nil {
			return nil
		}
	}
}

func main() {
	lis, err := net.Listen("tcp4", ":5050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	server := grpc.NewServer()
	pb.RegisterEchoServer(server, &EchoServer{})
	log.Printf("server listening at %v", lis.Addr().String())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Fail to server: %v", err)
	}
}
