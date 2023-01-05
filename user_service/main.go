package main

import (
	pb "app/protos/user_service"
	"context"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type user_service struct {
	*pb.UnimplementedUserServiceServer
}

func (u *user_service) GetMin(ctx context.Context, req *pb.Array) (*pb.Arr, error) {

	min := pb.Arr{}
	min.Num = req.Nums[0]
	for _, num := range req.Nums {
		fmt.Println(num)
		if min.Num > num {
			min.Num = num
		}
	}
	fmt.Println(min.Num)
	return &min, nil
}

func (u *user_service) GetSub(ctx context.Context, req *pb.Variable) (*pb.Var, error) {
	sub := pb.Var{
		C: req.A - req.B,
	}
	return &sub, nil
}

func (u *user_service) GetSqrt(ctx context.Context, req *pb.Sqrt) (*pb.Var2, error) {
	javob := pb.Var2{}
	var sr float32 = req.Var / 2
	var temp float32
	for {
		temp = sr
		sr = (temp + (float32(req.Var) / temp)) / 2
		if (temp - sr) == 0 {
			break
		}
	}
	javob.Var2 = sr
	return &javob, nil
}

func (u *user_service) GetPow(ctx context.Context, req *pb.Variable) (*pb.Var, error) {

	pow := int32(1)
	for i := int32(0); i < req.B; i++ {
		pow *= req.A
	}

	return &pb.Var{
		C: pow,
	}, nil
}

func (u *user_service) GetMult(ctx context.Context, req *pb.Variable) (*pb.Var, error) {
	mult := pb.Var{
		C: req.A * req.B,
	}
	return &mult, nil
}

func (u *user_service) GetDiv(ctx context.Context, req *pb.Variable) (*pb.Var, error) {
	div := pb.Var{
		C: req.A / req.B,
	}
	return &div, nil
}

func (u *user_service) GetSum(ctx context.Context, req *pb.Variable) (*pb.Var, error) {
	sum := pb.Var{
		C: req.A + req.B,
	}
	return &sum, nil
}

func (u *user_service) GetMax(ctx context.Context, req *pb.Array) (*pb.Arr, error) {

	max := int32(0)
	for _, num := range req.Nums {

		if max < num {
			max = num
		}
	}

	return &pb.Arr{
		Num: max,
	}, nil
}

func (u *user_service) GetUserById(ctx context.Context, req *pb.UserPrimaryKey) (*pb.User, error) {

	users := []pb.User{
		{
			Id:       1,
			FullName: "Shaxboz Norbekov",
			Age:      24,
		},
		{
			Id:       2,
			FullName: "Jahongir Normurodov",
			Age:      25,
		},
		{
			Id:       3,
			FullName: "Samandar Foziljonov",
			Age:      20,
		},
		{
			Id:       4,
			FullName: "Moxirbek Abduvaliyev",
			Age:      21,
		},
	}

	for _, user := range users {
		if req.Id == user.Id {
			return &user, nil
		}
	}
	return nil, errors.New("Not found user")
}

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &user_service{})

	fmt.Println("Listening :9001...")
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
