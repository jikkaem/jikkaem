package main

import (
	"fmt"
	"log"
	"net"

	pb "jikkaem/proto/fancam"
	"jikkaem/services/fancam"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 6001))
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
