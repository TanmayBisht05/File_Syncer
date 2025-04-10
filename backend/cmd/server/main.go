package main

import (
	"File_Syncer/auth/handlers"
	"File_Syncer/db"
	httpserver "File_Syncer/http"
	"File_Syncer/server"
	// "go.uber.org/zap"
)

func main() {
	db.Connect()
	handlers.InitAuthHandler() // Initialize after DB is connected
	go server.StartGRPCServer("50051")
	httpserver.StartHTTPServer("5000")
}
