package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "app/calculatorpb"
)

type server struct {
	*pb.UnimplementedCalculatorServiceServer
}

func (s *server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {

	fmt.Println("Recieved: ", req)

	sum := req.FirstNumber + req.SecondNumber

	return &pb.SumResponse{
		SumResult: sum,
	}, nil
}

func (s *server) PrimeNumberDecompition(req *pb.PrimeNumberDecompitionRequest, stream pb.CalculatorService_PrimeNumberDecompitionServer) error {

	fmt.Println("Req:", req)

	number := req.GetNumber()
	divisor := int64(2)

	for number > 1 {

		if number%divisor == 0 {

			stream.Send(&pb.PrimeNumberDecompitionResponse{
				PrimeFactor: divisor,
			})

			number = number / divisor
		} else {
			divisor++
		}
	}

	return nil
}

func (s *server) ComputeAvarage(stream pb.CalculatorService_ComputeAvarageServer) error {

	var (
		sum, count int64
	)

	for {

		req, err := stream.Recv()

		if err == io.EOF {

			avarage := float64(sum) / float64(count)
			err = stream.SendAndClose(&pb.ComputeAvarageResponse{
				Avarage: avarage,
			})

			if err != nil {
				log.Println("error while ComputeAvarage Recv:", err)
				return err
			}

			return nil
		}

		if err != nil {
			log.Println("error while ComputeAvarage Recv:", err)
			return err
		}

		sum += req.Number
		count++
	}

	return nil
}

func (s *server) FindMaximum(stream pb.CalculatorService_FindMaximumServer) error {

	maximum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Println("error while FindMaximum Recv:", err)
			return err
		}

		fmt.Println(req)

		if req.Number > maximum {
			maximum = req.Number
			err = stream.Send(&pb.FindMaximumResponse{Maximum: maximum})
			if err != nil {
				log.Println("error while FindMaximum Send:", err)
				return err
			}
		}
	}

	return nil
}

func main() {

	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})

	fmt.Println("Listening :9002...")
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
