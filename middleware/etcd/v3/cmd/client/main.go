package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/grpc"

	pb "gowhole/middleware/etcd/v3/cmd/hello"

	api "gowhole/middleware/etcd/v3/api"
)

var (
	serv = flag.String("service", "hello_service", "service name")
	reg  = flag.String("reg", "127.0.0.1:2379,127.0.0.1:22379,127.0.0.1:32379", "register etcd address")
)

func main() {
	flag.Parse()
	r := api.NewResolver(*serv)
	b := grpc.RoundRobin(r)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	conn, err := grpc.DialContext(ctx, *reg, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock())
	cancel()
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(1000 * time.Millisecond)
	for t := range ticker.C {
		client := pb.NewGreeterClient(conn)
		resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world " + strconv.Itoa(t.Second())})
		if err == nil {
			fmt.Printf("%v: Reply is %s\n", t, resp.Message)
		}
	}
}
