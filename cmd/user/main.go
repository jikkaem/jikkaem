package main

import (
	"fmt"
	"log"
	"net"

	pb "jikkaem/internal/proto/user"
	"jikkaem/internal/services/user/server"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 6000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server.UserServer{})

	log.Printf("User server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
