package main

import (
    "context"
    "sync"

    pb "distributed-kv-store/proto"
)

type server struct {
    pb.UnimplementedKVStoreServer
    mu   sync.RWMutex
    data map[string]string
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    value, ok := s.data[req.Key]
    if !ok {
        return &pb.GetResponse{Value: ""}, nil
    }
    return &pb.GetResponse{Value: value}, nil
}

func (s *server) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.data[req.Key] = req.Value
    return &pb.SetResponse{Success: true}, nil
}