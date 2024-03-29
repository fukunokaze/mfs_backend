package main

import (
	context "context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/fukunokaze/mfs_backend/LoginService/loginpb"
	"github.com/fukunokaze/mfs_backend/MFSRepository/ent"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	loginpb.LoginServiceServer
}

func main() {
	fmt.Println("Blog Service Started")

	client, err := ent.Open("postgres", "host=localhost port=5432 user=roshani dbname=entdemo password='roshani@123'")

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	loginpb.RegisterLoginServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch

}

func (s *server) VerifyLogin(ctx context.Context, in *loginpb.VerifyLoginRequest) (*loginpb.VerifyLoginResponse, error) {

	return &loginpb.VerifyLoginResponse{
		UserDetail: &loginpb.UserProto{
			Username: "Arga",
			Role:     "Admin",
		},
	}, nil
}
