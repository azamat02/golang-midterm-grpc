package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	servicepb "task2/api/proto"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := servicepb.NewGreetServiceClient(conn)
	doLongGreet(c)
}

func doLongGreet(c servicepb.GreetServiceClient) {

	requests := []*servicepb.NumRequest{
		{
			Number: &servicepb.Number{
				Number: int64(1),
			},
		},
		{
			Number: &servicepb.Number{
				Number: int64(2),
			},
		},
		{
			Number: &servicepb.Number{
				Number: int64(3),
			},
		},
		{
			Number: &servicepb.Number{
				Number: int64(4),
			},
		},
	}

	ctx := context.Background()
	stream, err := c.LongGreet(ctx)
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res.Result)
}
