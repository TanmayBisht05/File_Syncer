package server

import (
	"File_Syncer/proto"
	"io"
	"log"
	"sync"

	"google.golang.org/grpc"
	"net"
)

type syncServer struct {
	proto.UnimplementedSyncServiceServer
	clients map[string]proto.SyncService_ConnectServer
	mu      sync.Mutex
}

func NewSyncServer() *syncServer {
	return &syncServer{
		clients: make(map[string]proto.SyncService_ConnectServer),
	}
}

func (s *syncServer) Connect(stream proto.SyncService_ConnectServer) error {
	var clientID string

	for {
		change, err := stream.Recv()
		if err == io.EOF {
			log.Println("Client disconnected:", clientID)
			s.mu.Lock()
			delete(s.clients, clientID)
			s.mu.Unlock()
			return nil
		}
		if err != nil {
			log.Println("Error receiving from stream:", err)
			return err
		}

		clientID = change.ClientId
		log.Printf("Received change from %s: %s (%s)", clientID, change.Filename, change.Action)

		// Store client stream if new
		s.mu.Lock()
		if _, exists := s.clients[clientID]; !exists {
			s.clients[clientID] = stream
		}
		s.mu.Unlock()

		// Broadcast to others
		s.broadcast(change, clientID)
	}
}

func (s *syncServer) broadcast(change *proto.FileChange, senderID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for id, clientStream := range s.clients {
		if id == senderID {
			continue
		}
		if err := clientStream.Send(change); err != nil {
			log.Printf("Failed to send to %s: %v", id, err)
		}
	}
}

func StartGRPCServer(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterSyncServiceServer(s, NewSyncServer())
	log.Println("Server listening on port", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
