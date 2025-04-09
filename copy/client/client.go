package client

import (
	"File_Syncer/proto"
	"File_Syncer/syncstate"
	"context"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

var stream proto.SyncService_ConnectClient

func StartClient(clientID, serverAddr string, applyChange func(*proto.FileChange)) error {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := proto.NewSyncServiceClient(conn)

	stream, err = client.Connect(context.Background())
	if err != nil {
		return err
	}

	// Send an initial no-op message to register with the server
	err = stream.Send(&proto.FileChange{
		ClientId: clientID,
		Action:   "noop",
	})
	if err != nil {
		log.Printf("Failed to send initial message: %v", err)
		return err
	}

	log.Println("Connected to server and registered stream.")

	go func() {
		for {
			change, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println("Error receiving:", err)
				break
			}
			log.Printf("Received change from peer: %s (%s)", change.Filename, change.Action)
			syncstate.MarkAsRemoteUpdate(change.Filename)
			applyChange(change)
		}
	}()
	return nil
}

func SendChange(clientID, filename, action string) {
	var content []byte
	if action != "delete" {
		var err error
		content, err = os.ReadFile(filename)
		if err != nil {
			log.Printf("Failed to read file %s: %v", filename, err)
			return
		}
	}

	err := stream.Send(&proto.FileChange{
		ClientId:  clientID,
		Filename:  filename,
		Action:    action,
		Timestamp: time.Now().Unix(),
		Content:   content,
	})
	if err != nil {
		log.Printf("Failed to send change: %v", err)
	}
}
