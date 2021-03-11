package main

import (
	"context"
	"log"
	"time"

	"github.com/almogori/muun/muun"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := muun.NewStockServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetStockPrice(ctx, &muun.GetStockPriceRequest{Ticker: "ABC"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Strike price: %f", r.GetStrike())
}
