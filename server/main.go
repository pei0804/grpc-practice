package main

import (
	"flag"
	"log"
	"net"

	pb "github.com/pei0804/grpc-practice/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	addrFlag = flag.String("addr", ":5000", "Address host:post")
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("New Request: %v", in.String())
	return &pb.HelloReply{Message: "Hello, " + in.Name + "!"}, nil
}

func (s *server) Integer(ctx context.Context, in *pb.IntegerRequest) (*pb.IntegerResponse, error) {
	log.Printf("New Request: %v", in.String())
	return &pb.IntegerResponse{Val: in.Val}, nil
}

func main() {
	lis, err := net.Listen("tcp", *addrFlag)

	if err != nil {
		log.Fatalf("boo")
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)

}
