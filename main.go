package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"stockticker/internal/api/v1/stock"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {

	validateEnvVars()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Stock ticker live")

	s := stock.Server{}
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	stock.RegisterTickerServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	reflection.Register(grpcServer)
}

func validateEnvVars() {
	_, ok1 := os.LookupEnv("APIKEY")
	if !ok1 {
		log.Fatalf("APIKEY environment variable is not present")
	}
	_, ok2 := os.LookupEnv("SYMBOL")
	if !ok2 {
		log.Fatalf("SYMBOL environment variable is not present")
	}
	_, ok3 := os.LookupEnv("NDAYS")
	if !ok3 {
		log.Fatalf("NDAYS environment variable is not present")
	}
	days, err := strconv.Atoi(os.Getenv("NDAYS"))
	if err != nil {
		log.Fatalf("Cannot convert NDAYs environment variable to a number")
	}
	if days == 0 {
		log.Fatalf("a valid period of time is required for NDAYS environment variable")
	}
}
