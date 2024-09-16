package main

import (
	"context"
	"log"
	"time"

	pb "distributed-kv-store/proto"
	"google.golang.org/grpc"
)

func runClient() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewKVStoreClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Set a key-value pair
	_, err = c.Set(ctx, &pb.SetRequest{Key: "foo", Value: "bar"})
	if err != nil {
		log.Fatalf("could not set: %v", err)
	}

	// Get the value for the key
	r, err := c.Get(ctx, &pb.GetRequest{Key: "foo"})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
}