package main

import (
	"context"
	"log"
	"net"

	pb "../api"

	"google.golang.org/grpc"
)

const (
	port = ":50050"
)

// server is used to implement language.CountryLanguageServer.
type server struct {
	pb.UnimplementedCountryLanguageServer
}

func (s *server) GetCountryLanguage(ctx context.Context, in *pb.CountryLanguageRequest) (out *pb.CountryLanguageResponse, err error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.CountryLanguageResponse{Languages: nil}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCountryLanguageServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
