//需要在 protoc-gen-go目录执行go build和go install，生成 protoc-gen-go.exe
//protoc -I proto/ proto/gameservertest.proto --go_out=plugins=grpc:proto

package main

import (
	"context"
	"fmt"
	pb "gameservertest/proto"
	"google.golang.org/grpc"
	"net"
)

const (
	add = "150.109.145.126:10010"
)

type server struct {
	pb.UnimplementedTestServer
}

func(s *server) JustConnect(ctx context.Context, req *pb.ConnectRequest) (*pb.ConnectReply, error) {
	fmt.Printf("server.JustConnect|req = %#v", req)

	return &pb.ConnectReply{Who: req.GetAction() + "you"}, nil
}

func main() {
	lis, err := net.Listen("tcp", add)
	if err != nil {
		fmt.Printf("server.JustConnect|ERROR tcp to %s fail, err = %s", add, err)
	}

	s := grpc.NewServer()
	pb.RegisterTestServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Printf("server.JustConnect|ERROR server fail, err = %s", err)
	}
}