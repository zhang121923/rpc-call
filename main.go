package main

import (
	"context"
	pb "github.com/zhang121923/rpc-proto-registry/protos/go"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCreateUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetUserName(ctx, &pb.User{Name: "llalalaa"})

	c1 := pb.NewStudentServiceClient(conn)

	_, err = c1.JoinSchool(ctx, &pb.Student{Name: "wangwu"})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting User: %s", r.GetName())
}
