package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-envoy-docker/pb"
	"log"
	"net"
)

var (
	port = flag.Int("port", 9091, "The server port")
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	log.Printf("Received: %v", in)
	return &pb.User{UserId: uuid.New().String(), FirstName: in.FirstName, LastName: in.LastName}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
