package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/gildemberg-santos/webcrawlerurl_v2/grpc"
	grpc "google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) CreateLead(ctx context.Context, in *pb.LeadRequest) (*pb.LeadReply, error) {
	log.Printf("Received: %v", in.GetName())

	return &pb.LeadReply{
		Id:      fmt.Sprint(rand.New(rand.NewSource(time.Now().UnixNano())).Int63()),
		Status:  "success",
		Message: fmt.Sprintf("Olá %s, seu e-mail %s foi cadastrado com sucesso.", in.GetName(), in.GetEmail()),
	}, nil
}

func main() {
	log.Println("Starting server...")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
