package main

import (
	"flag"
	"log"

	pb "github.com/pei0804/grpc-practice/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	addrFlag = flag.String("addr", "localhost:5000", "server address host:post")
)

func main() {
	conn, err := grpc.Dial(*addrFlag, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "wdgk"})
	if err != nil {
		log.Fatalf("RPC error: %v", err)
	}
	log.Printf("Greeting: %s", resp.Message)
	resp2, err := c.Integer(context.Background(), &pb.IntegerRequest{Val: 1111})
	if err != nil {
		log.Fatalf("RPC error: %v", err)
	}
	log.Printf("ans : %s", resp2.Val)
}
