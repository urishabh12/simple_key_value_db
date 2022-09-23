package main

import (
	"log"
	"net"

	proto "github.com/urishabh12/simple_key_value_db/proto"
	db_server "github.com/urishabh12/simple_key_value_db/server"
	"google.golang.org/grpc"
)

const (
	server_port = ":8000"
)

func main() {
	dbServer, err := db_server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", server_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterDBServiceServer(grpcServer, dbServer)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
