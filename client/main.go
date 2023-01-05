package main

import (
	"context"
	"fmt"
	"log"

	pb "app/protos/user_service"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	// resp, err := c.GetUserById(context.Background(), &pb.UserPrimaryKey{Id: 1})
	// if err != nil {
	// 	log.Println("error whiling get user by id:", err.Error())
	// 	return
	// }
	// fmt.Println(resp)

	// resp2, err := c.GetSum(context.Background(), &pb.Variable{A: 2, B: 4})
	// if err != nil {
	// 	log.Println("error whiling get sum :", err.Error())
	// 	return
	// }
	// fmt.Println(resp2)

	// // var arr *pb.Array
	// // smth := []string{2, 3, 4, 5}

	// // var new := pb.Array{
	// // 	nums := []int32{2, 3, 4, 5}
	// // }
	// resp3, err := c.GetMax(context.Background(), &pb.Array{
	// 	Nums: []int32{2, 3, 4, 5, 55, 5555},
	// })
	// if err != nil {
	// 	log.Println("error whiling get min :", err.Error())
	// 	return
	// }
	// fmt.Println(resp3)

	// resp4, err := c.GetDiv(context.Background(), &pb.Variable{A: 25, B: 5})
	// if err != nil {
	// 	log.Println("error whiling get div :", err.Error())
	// 	return
	// }
	// fmt.Println(resp4)

	// resp5, err := c.GetMult(context.Background(), &pb.Variable{A: 20, B: 5})
	// if err != nil {
	// 	log.Println("error whiling get mult :", err.Error())
	// 	return
	// }
	// fmt.Println(resp5)

	// resp6, err := c.GetSub(context.Background(), &pb.Variable{A: 55, B: 34})
	// if err != nil {
	// 	log.Println("error whiling get sub :", err.Error())
	// 	return
	// }
	// fmt.Println(resp6)

	// resp7, err := c.GetSqrt(context.Background(), &pb.Sqrt{Var: 16})
	// if err != nil {
	// 	log.Println("error whiling get sqrt :", err.Error())
	// 	return
	// }
	// fmt.Println(resp7)

	// resp8, err := c.GetPow(context.Background(), &pb.Variable{A: 4, B: 3})
	// if err != nil {
	// 	log.Println("error whiling get pow :", err.Error())
	// 	return
	// }
	// fmt.Println(resp8)

	resp9, err := c.GetMin(context.Background(), &pb.Array{
		Nums: []int32{2, 3, 4, 5},
	})
	if err != nil {
		log.Println("error whiling get min :", err.Error())
		return
	}
	fmt.Println(resp9)

}
