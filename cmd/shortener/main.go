package main

import "IDontKnowWhatIsThis/cmd/server"

func main() {
	var protocol = "tcp"
	var port = "127.0.0.1:8080"
	server.Server(protocol,port)
}
