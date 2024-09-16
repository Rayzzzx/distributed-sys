package main

import (
    "flag"
    "log"
    "net"

    pb "distributed-kv-store/proto"
    "google.golang.org/grpc"
)

func main() {
    isClient := flag.Bool("client", false, "Run as client")
    flag.Parse()

    if *isClient {
        runClient()
    } else {
        lis, err := net.Listen("tcp", ":50051")
        if err != nil {
            log.Fatalf("failed to listen: %v", err)
        }
        s := grpc.NewServer()
        pb.RegisterKVStoreServer(s, &server{})
        log.Printf("server listening at %v", lis.Addr())
        if err := s.Serve(lis); err != nil {
            log.Fatalf("failed to serve: %v", err)
        }
    }
}