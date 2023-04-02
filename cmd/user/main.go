package main

import (
	"fmt"
	"log"
	"net"

	"jikkaem/internal/services/user/server"
	pb "jikkaem/internal/shared/proto/user"

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
