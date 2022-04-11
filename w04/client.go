package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/wuzehv/geekgo/w04/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "test"
	defaultAge  = 10
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "username")
	age  = flag.Int("age", defaultAge, "user age")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUserList(ctx, &pb.UserRequest{Name: *name, Age: int32(*age)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("UserList: %d user return, %v", len(r.GetUserinfo()), r.GetUserinfo())
}
