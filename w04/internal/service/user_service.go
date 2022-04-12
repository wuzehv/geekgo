package service

import (
	"context"
	pb "github.com/wuzehv/geekgo/w04/api"
	"github.com/wuzehv/geekgo/w04/internal/biz"
	"log"
)

type Server struct {
	pb.UnimplementedUserServer
}

func (s *Server) GetUserList(ctx context.Context, in *pb.UserRequest) (*pb.UserList, error) {
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
