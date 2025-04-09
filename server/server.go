package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "File_Syncer/proto"
	"File_Syncer/syncstate"

	"google.golang.org/grpc"
)

type syncServer struct {
	pb.UnimplementedSyncServiceServer
}

func (s *syncServer) SendChange(ctx context.Context, change *pb.FileChange) (*pb.Ack, error) {
	fmt.Println("====> Incoming file sync request")
	log.Printf("Received: %s (%s)", change.Filename, change.Action)

	// Mark this file with a timestamp, so subsequent events are skipped.
	syncstate.MarkAsRemoteUpdate(change.Filename)

	if change.Action == "delete" {
		if err := os.Remove(change.Filename); err != nil {
			log.Printf("Error deleting file %s: %v", change.Filename, err)
		}
	} else {
		if err := os.WriteFile(change.Filename, change.Content, 0644); err != nil {
			log.Printf("Error writing file %s: %v", change.Filename, err)
		}
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
