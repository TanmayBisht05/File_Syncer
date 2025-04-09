package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "File_Syncer/proto"
	"File_Syncer/syncstate" // ✅ new import

	"google.golang.org/grpc"
)

type syncServer struct {
	pb.UnimplementedSyncServiceServer
}

func (s *syncServer) SendChange(ctx context.Context, change *pb.FileChange) (*pb.Ack, error) {
	fmt.Println("====> Incoming file sync request")
	log.Printf("Received: %s (%s)", change.Filename, change.Action)

	syncstate.SkipNextEvent.Store(true) // ✅ Prevent loop before applying change

	if change.Action == "delete" {
		os.Remove(change.Filename)
	} else {
		os.WriteFile(change.Filename, change.Content, 0644)
	}

	return &pb.Ack{Status: "OK"}, nil
}

func StartGRPCServer(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSyncServiceServer(s, &syncServer{})
	fmt.Println("Server listening on port", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
