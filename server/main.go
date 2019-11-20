package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "../api"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

const (
	databaseHost = "192.168.99.100:43306" // my docker-db
	port         = ":50050"
)

func searchCountryLanguage(db *sql.DB, name string) []*pb.Language {
	// send query
	rows, err := db.Query("SELECT cl.language, cl.percentage FROM countrylanguage AS cl JOIN country AS c ON cl.countrycode = c.code WHERE c.name = ? ORDER BY cl.percentage DESC", name)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	var result []*pb.Language
	for rows.Next() {
		var cl pb.Language
		err := rows.Scan(&cl.Name, &cl.Percentage)

		if err != nil {
			panic(err.Error())
		}
		fmt.Println(cl.Name, cl.Percentage)

		result = append(result, &cl)
	}
	return result
}

// server is used to implement language.CountryLanguageServer.
type server struct {
	pb.UnimplementedCountryLanguageServer
	db *sql.DB // gRPCサーバの実体となる server 構造体にdbフィールドを持たせる
}

func (s *server) GetCountryLanguage(ctx context.Context, in *pb.CountryLanguageRequest) (*pb.CountryLanguageResponse, error) {
	log.Printf("Received: %v", in.GetName())
	result := searchCountryLanguage(s.db, in.GetName())
	return &pb.CountryLanguageResponse{Languages: result}, nil
}

func main() {
	// connect DB
	db, err := sql.Open("mysql", "root:passwd@tcp("+databaseHost+")/world")
	if err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}
	log.Println("Connected to mysql.")
	defer db.Close()
	// run server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCountryLanguageServer(s, &server{db: db})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
