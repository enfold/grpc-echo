package main

import (
	"context"
	"log"
	"time"
// 	"io/ioutil"
 	"crypto/tls"
//	"crypto/x509"
// 	"fmt"
  "os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "google.golang.org/grpc/examples/features/proto/echo"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
    // Load certificate of the CA who signed server's certificate
//     pemServerCA, err := ioutil.ReadFile("rootCA.crt")
//     if err != nil {
//         return nil, err
//     }
//
//     certPool := x509.NewCertPool()
//     if !certPool.AppendCertsFromPEM(pemServerCA) {
//         return nil, fmt.Errorf("failed to add server CA's certificate")
//     }

    // Create the credentials and return it
    config := &tls.Config{
//        RootCAs:      certPool,
        InsecureSkipVerify: true,
    }

    return credentials.NewTLS(config), nil
}
func main() {
  tlsCredentials, err := loadTLSCredentials()
  if err != nil {
    log.Fatal("cannot load TLS credentials: ", err)
  }
	conn, err := grpc.Dial(os.Args[1], grpc.WithTransportCredentials(tlsCredentials))
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

