package main

import (
	"File_Syncer/client"
	"File_Syncer/proto"
	"File_Syncer/syncstate"
	"File_Syncer/watcher"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

const watchDir = "./sync"
const clientID = "client1"
const serverAddr = "172.20.55.253:50051"

func main() {
	log.Println("Client attempting to connect to server at", serverAddr)

	err := client.StartClient(clientID, serverAddr, func(change *proto.FileChange) {
		log.Println("Connected to server, stream established")

		localPath := filepath.Join(watchDir, filepath.Base(change.Filename))
		if change.Action == "delete" {
			os.Remove(localPath)
		} else {
			os.WriteFile(localPath, change.Content, 0644)
		}

		// if change.Action == "delete" {
		// 	os.Remove(change.Filename)
		// } else {
		// 	os.WriteFile(change.Filename, change.Content, 0644)
		// }
	})
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	os.MkdirAll(watchDir, 0755)

	var debounceMap = make(map[string]*time.Timer)
	var debounceMu sync.Mutex
	const debounceDuration = 100 * time.Millisecond

	scheduleSync := func(filename string, action string) {
		debounceMu.Lock()
		defer debounceMu.Unlock()
		if timer, ok := debounceMap[filename]; ok {
			timer.Stop()
		}
		debounceMap[filename] = time.AfterFunc(debounceDuration, func() {
			log.Println("Forwarding change to peer for file:", filename)
			client.SendChange(serverAddr, filename, action)
			debounceMu.Lock()
			delete(debounceMap, filename)
			debounceMu.Unlock()
		})
	}

	err = watcher.Watch(watchDir, func(event fsnotify.Event) {
		if syncstate.ShouldSkip(event.Name) {
			log.Println("Skipping event from remote update for file:", event.Name)
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

		relPath, _ := filepath.Rel(watchDir, event.Name)
		if strings.HasPrefix(relPath, "..") {
			log.Println("Ignored change outside sync folder:", event.Name)
			return
		}

		scheduleSync(event.Name, action)
	})

	if err != nil {
		log.Fatal(err)
	}

	select {}
}
