package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	servicepb "task2/api/proto"
)

//Server with embedded UnimplementedGreetServiceServer
type Server struct {
	servicepb.UnimplementedGreetServiceServer
}

//LongGreet is an example of stream from client side
func (s *Server) LongGreet(stream servicepb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet function was invoked with a streaming request\n")
	var result float32
	var counter float32=0
	var sum float32=0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			result = sum/counter
			// we have finished reading the client stream
			return stream.SendAndClose(&servicepb.NumResponse{
				Result: result,
			})

		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		sum += float32(req.GetNumber().Number)
		counter ++
	}
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	servicepb.RegisterGreetServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
