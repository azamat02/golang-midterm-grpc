package main

import (
	servicepb "awesomeProject/api/primeFactorService"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

//Server with embedded UnimplementedGreetServiceServer
type Server struct {
	servicepb.UnimplementedPrimeCalcServiceServer
}


// GreetManyTimes is an example of stream from server side
func (s *Server) Calc(req *servicepb.CalcRequest, stream servicepb.PrimeCalcService_CalcServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v \n", req)

	number := int(req.GetNumber().GetNumber())
	var result []int64

	for i := 2;number > i; i++  {
		for number%i == 0 {
			a := int64(i)
			result = append(result, a)
			number = number/i;
		}
	}
	if number >2 {
		num := int64(number)
		result = append(result, num)
	}

	for _, v:=range result {
		res := &servicepb.CalcResponse{Result: v}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error while sending greet many times responses: %v", err.Error())
		}
	}
	return nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	servicepb.RegisterPrimeCalcServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
