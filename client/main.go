package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "../api"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50050"
	defaultName = "Japan"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCountryLanguageClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetCountryLanguage(ctx, &pb.CountryLanguageRequest{Name: name})
	if err != nil {
		log.Fatalf("could not search: %v", err)
	}
	log.Printf("Result: %s", r)
}
