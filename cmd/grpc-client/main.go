package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/gildemberg-santos/webcrawlerurl_v2/grpc"
	grpc "google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Estabelecer conexão com o servidor gRPC
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	// Preparar a requisição
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	email := ""
	if len(os.Args) > 2 {
		email = os.Args[2]
	}

	phone := ""
	if len(os.Args) > 3 {
		phone = os.Args[3]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Enviar a requisição e receber a resposta
	r, err := client.CreateLead(ctx, &pb.LeadRequest{Name: name, Email: email, Phone: phone})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("ID:", r.GetId())
	log.Println("Status:", r.GetStatus())
	log.Println("Message:", r.GetMessage())
}
