package main

import (
	"File_Syncer/client"
	"File_Syncer/server"
	"File_Syncer/syncstate"
	"File_Syncer/watcher"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

const watchDir = "./sync"
const peerAddr = "localhost:50052" // Address of the other machine

func main() {
	go server.StartGRPCServer("50051")

	// Ensure sync directory exists
	os.MkdirAll(watchDir, 0755)

	err := watcher.Watch(watchDir, func(event fsnotify.Event) {
		if syncstate.SkipNextEvent.Load() {
			log.Println("Skipping event from remote update")
			syncstate.SkipNextEvent.Store(false)
			return
		}

		log.Println("Detected change:", event)

		action := "update"
		if event.Op&fsnotify.Create != 0 {
			action = "create"
		}
		if event.Op&fsnotify.Remove != 0 {
			action = "delete"
		}

		filename := event.Name

		// Check if file is within watchDir
		relPath, _ := filepath.Rel(watchDir, filename)
		if !strings.HasPrefix(relPath, "..") {
			log.Println("Forwarding change to peer")
			client.SendChange(peerAddr, filename, action)
		} else {
			log.Println("Ignored change outside sync folder")
		}
	})

	if err != nil {
		log.Fatal(err)
	}

	select {} // Block forever
}
