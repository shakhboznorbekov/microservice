package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "app/calculatorpb"
)

func main() {

	conn, err := grpc.Dial("localhost:9002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	// doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	doBiDirectionalStream(c)
}

func doUnary(c pb.CalculatorServiceClient) {

	fmt.Println("Starting to do Sum Unary RPC...")

	req := pb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 35,
	}

	res, err := c.Sum(context.Background(), &req)
	if err != nil {
		log.Println("error while calling Sum RPC:", err)
	}

	fmt.Println("Reponse from Sum:", res.GetSumResult())
}

func doServerStreaming(c pb.CalculatorServiceClient) {

	fmt.Println("Starting to do PrimeNumberDecompition Server Streamin RPC...")

	stream, err := c.PrimeNumberDecompition(context.Background(), &pb.PrimeNumberDecompitionRequest{
		Number: 25,
	})
	if err != nil {
		log.Println("error while calling PrimeNumberDecompition RPC:", err)
	}

	for {

		resp, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println("error while calling PrimeNumberDecompition Stream Recv RPC:", err)
		}

		fmt.Println(resp)
	}
}

func doClientStreaming(c pb.CalculatorServiceClient) {

	stream, err := c.ComputeAvarage(context.Background())
	if err != nil {
		log.Println("error while calling ComputeAvarage RPC:", err)
	}

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range nums {
		err = stream.Send(&pb.ComputeAvarageRequest{
			Number: int64(num),
		})
		if err != nil {
			log.Println("error while calling ComputeAvarage Client Stream Send RPC:", err)
			return
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("error while calling ComputeAvarage Client Stream Send RPC:", err)
		return
	}

	fmt.Println(res)

}

func doBiDirectionalStream(c pb.CalculatorServiceClient) {

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Println("error while calling FindMaximum Send RPC:", err)
		return
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{4, 6, 2, 23, 13, 7, 8, 53, 1}

		for _, num := range numbers {
			stream.Send(&pb.FindMaximumRequest{Number: num})
			time.Sleep(time.Second * 1)
		}

		stream.CloseSend()
	}()

	go func() {
		for {

			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println("error while calling FindMaximum Recv RPC:", err)
				return
			}

			fmt.Println(res)
		}

		close(waitc)
	}()

	<-waitc
}
