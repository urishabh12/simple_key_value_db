package main

import (
	"context"
	"log"

	proto "github.com/urishabh12/simple_key_value_db/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	dbClient := proto.NewDBServiceClient(conn)

	putReq := &proto.PutRequest{
		Key:   "a",
		Value: "b",
	}
	_, err = dbClient.Put(context.Background(), putReq)
	if err != nil {
		log.Fatal(err)
	}

	getReq := &proto.GetRequest{Key: "a"}
	getRes, err := dbClient.Get(context.Background(), getReq)
	if err != nil {
		log.Fatal(err)
	}

	if getRes.Value != putReq.Value {
		log.Fatalf("Values not equal %s != %s", getRes.Value, putReq.Value)
	}
}
