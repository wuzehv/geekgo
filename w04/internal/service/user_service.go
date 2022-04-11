package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/wuzehv/geekgo/w04/api"
	"github.com/wuzehv/geekgo/w04/internal/biz"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedUserServer
}

func (s *server) GetUserList(ctx context.Context, in *pb.UserRequest) (*pb.UserList, error) {
	log.Printf("Received: %v", in)

	// DTO转换为DO，调用biz业务层
	d := biz.UserDo{
		Name: in.GetName(),
		Age:  int(in.Age),
	}

	userList := d.GetUserList()
	// DO转DTO
	res := new(pb.UserList)
	for _, v := range userList {
		res.Userinfo = append(res.Userinfo, &pb.UserInfo{
			Id:   int32(v.Id),
			Name: v.Name,
			Age:  int32(v.Age),
		})
	}

	return res, nil
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
