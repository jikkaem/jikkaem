package main

import (
	"fmt"
	"log"
	"net"

	pb "jikkaem/internal/proto/fancam"
	"jikkaem/internal/services/fancam"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3333))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFancamServer(s, &fancam.FancamServer{})

	log.Printf("User server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
