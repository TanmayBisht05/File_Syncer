package main

import "File_Syncer/server"

func main() {
	server.StartGRPCServer("50051")
}
