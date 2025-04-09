package client

import (
	"context"
	"log"
	"os"
	"time"

	pb "File_Syncer/proto"

	"google.golang.org/grpc"
)

func SendChange(addr, filename, action string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Failed to connect to peer at %s: %v", addr, err)
		return
	}
	defer conn.Close()

	c := pb.NewSyncServiceClient(conn)
	var content []byte
	if action != "delete" {
		content, err = os.ReadFile(filename)
		if err != nil {
			log.Printf("Failed to read file %s: %v", filename, err)
			return
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SendChange(ctx, &pb.FileChange{
		Filename:  filename,
		Action:    action,
		Timestamp: time.Now().Unix(),
		Content:   content,
	})
	if err != nil {
		log.Printf("Failed to send change: %v", err)
		return
	}

	log.Printf("Ack from peer: %s", r.Status)
}
