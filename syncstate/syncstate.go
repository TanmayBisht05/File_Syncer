package syncstate

import (
	"sync"
	"time"
)

// skippedFiles tracks the time when a file was last updated remotely.
var skippedFiles sync.Map

// SkipWindow defines how long to ignore events following a remote update.
var SkipWindow = 500 * time.Millisecond

// MarkAsRemoteUpdate records the current time for the given filename.
func MarkAsRemoteUpdate(filename string) {
	skippedFiles.Store(filename, time.Now())
}

// ShouldSkip returns true if the file was updated remotely within the SkipWindow.
// If the SkipWindow has passed, the marker is removed.
func ShouldSkip(filename string) bool {
	if tsVal, ok := skippedFiles.Load(filename); ok {
		if ts, ok := tsVal.(time.Time); ok {
			if time.Since(ts) < SkipWindow {
				return true
			}
			// Expired – remove the marker.
			skippedFiles.Delete(filename)
		}
	}
	return false
}
