package cmdutil

import (
	"google.golang.org/grpc"
	"log"
	"time"
)

func GetGrpcConn() (*grpc.ClientConn, error) {
	address := "localhost:8887"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithTimeout(5 * time.Second))
	if err != nil {
		log.Fatalf("can not connect: %v", err)
	}

	return conn, err
}