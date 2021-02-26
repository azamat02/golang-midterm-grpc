package main

import (
	servicepb "awesomeProject/api/primeFactorService"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := servicepb.NewPrimeCalcServiceClient(conn)
	doManyTimesFromServer(c)
}

func doManyTimesFromServer(c servicepb.PrimeCalcServiceClient) {
	ctx := context.Background()
	req := &servicepb.CalcRequest{Number: &servicepb.Number{
		Number: int64(120),
	}}

	stream, err := c.Calc(ctx, req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC %v", err)
	}
	defer stream.CloseSend()

LOOP:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// we've reached the end of the stream
				break LOOP
			}
			log.Fatalf("error while reciving from GreetManyTimes RPC %v", err)
		}
		log.Printf("response from GreetManyTimes:%v \n", res.GetResult())
	}

}

