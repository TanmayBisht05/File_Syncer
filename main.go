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
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

const watchDir = "./sync"
const peerAddr = "localhost:50052" // Address of the other machine

func main() {
	go server.StartGRPCServer("50051")

	// Ensure sync directory exists
	os.MkdirAll(watchDir, 0755)

	// Map and mutex for debouncing
	var debounceMap = make(map[string]*time.Timer)
	var debounceMu sync.Mutex
	const debounceDuration = 100 * time.Millisecond

	// Helper to schedule sending changes for a file.
	scheduleSync := func(filename string, action string) {
		debounceMu.Lock()
		defer debounceMu.Unlock()
		// If there's an existing timer, stop it.
		if timer, ok := debounceMap[filename]; ok {
			timer.Stop()
		}
		// Create a new timer.
		debounceMap[filename] = time.AfterFunc(debounceDuration, func() {
			log.Println("Forwarding change to peer for file:", filename)
			client.SendChange(peerAddr, filename, action)
			// Remove after triggering.
			debounceMu.Lock()
			delete(debounceMap, filename)
			debounceMu.Unlock()
		})
	}

	err := watcher.Watch(watchDir, func(event fsnotify.Event) {
		// If this file change was caused by a recent remote update, skip it.
		if syncstate.ShouldSkip(event.Name) {
			log.Println("Skipping event from remote update for file:", event.Name)
			return
		}

		log.Println("Detected change:", event)

		// Determine the action based on event type.
		action := "update"
		if event.Op&fsnotify.Create != 0 {
			action = "create"
		}
		if event.Op&fsnotify.Remove != 0 {
			action = "delete"
		}

		// Only process events inside the watch directory.
		relPath, _ := filepath.Rel(watchDir, event.Name)
		if strings.HasPrefix(relPath, "..") {
			log.Println("Ignored change outside sync folder:", event.Name)
			return
		}

		// Debounce the event so multiple rapid events lead to a single sync.
		scheduleSync(event.Name, action)
	})

	if err != nil {
		log.Fatal(err)
	}

	select {} // Block forever
}
